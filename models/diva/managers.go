package diva

import (
	"bitbucket.org/ikeikeikeike/antenna/models"
	"github.com/astaxie/beego/orm"
	apiactress "github.com/ikeikeikeike/go-apiactress"
	googleimages "github.com/ikeikeikeike/go-googleimages"
)

func ReadOrCreateByActress(act *apiactress.Actress) *models.Diva {
	d := &models.Diva{Name: act.Name}

	created, _, _ := d.ReadOrCreate("Name")
	if created {
		d.Icon = &models.Image{
			Name:   act.Name,
			Src:    act.Thumb,
			Ext:    ".jpg",
			Mime:   "image/jpeg",
			Width:  65,
			Height: 65,
		}
		d.Icon.Insert()

		d.Kana = act.Yomi
		d.Romaji = act.Oto
		d.Gyou = act.Gyou

		d.Update()
	}

	return d
}

func UpdateByGoogleimage(r *googleimages.Result) error {
	return nil
}

// Search more than 0, showon video.
func StarringDivas() orm.QuerySeter {
	return orm.NewOrm().QueryTable("diva").Filter("videos_count__gt", 0).OrderBy("-videos_count")
}
