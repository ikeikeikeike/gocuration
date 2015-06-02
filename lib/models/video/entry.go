package video

import (
	"strings"

	libentry "bitbucket.org/ikeikeikeike/antenna/lib/models/entry"
	"bitbucket.org/ikeikeikeike/antenna/lib/models/site"
	"bitbucket.org/ikeikeikeike/antenna/models"
	"bitbucket.org/ikeikeikeike/antenna/models/diva"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/*
	Added models below ::

		- Video
		- Site
		- Diva

	.. note:: only create

*/
func AddsByEntries(entries []*models.Entry) (errs []error) {
	var err error
	o := orm.NewOrm()

	for _, e := range entries {
		ext := libentry.NewExtractor()
		ext.Header("User-Agent", beego.AppConfig.String("UserAgent"))

		if err := ext.Do(e); err != nil {
			beego.Warning("Error Entry{id=", e.Id, "}:", err)
			errs = append(errs, err)
			continue
		}

		// Save Video: last once value in slice...
		v := new(models.Video)
		v.Entry = e

		urls := ext.Urls()
		for _, val := range urls {
			v.Url = val
		}
		codes := ext.Codes()
		v.Code = strings.Join(codes, "")

		if v.Site, err = site.ReadOrCreateByVideo(v); err != nil {
			beego.Notice("Add site Entry{id=", e.Id, "}", err)
		}

		// XXX: Ignore ero-video, asg, redtube
		if len(urls) <= 0 && len(codes) <= 0 {
			beego.Notice("Not extract Entry{id=", e.Id, "}")
		}

		// TODO: Duration
		// ext.Duration()

		// Need video's primary id
		v.Insert()

		// Save video metas
		for _, val := range urls {
			vm := &models.VideoUrl{Name: val, Video: v}
			vm.Insert()
		}
		for _, val := range codes {
			vm := &models.VideoCode{Name: val, Video: v}
			vm.Insert()
		}

		// Save Diva
		var divas []*models.Diva
		for _, name := range ext.ByNames(diva.DivasName()) {
			d := &models.Diva{Name: name}
			d.Read("Name")
			if d.Id > 0 {
				divas = append(divas, d)
			}
		}
		if len(divas) > 0 {
			_, err := o.QueryM2M(v, "Divas").Add(divas)
			if err != nil {
				beego.Warn("Warn add divas:", err)
			}
		}
	}
	return
}
