package fillup

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"bitbucket.org/ikeikeikeike/antenna/models"
	"bitbucket.org/ikeikeikeike/antenna/models/diva"
	"github.com/astaxie/beego"
	"github.com/ikeikeikeike/divaextractor"
	apiactress "github.com/ikeikeikeike/go-apiactress"
	"github.com/ikeikeikeike/shuffler"
)

// act.Id, act.Gyou, act.Thumb, act.Yomi, act.Oto
func DivaByApiActress(prefix string) error {
	c := apiactress.NewClient()
	c.Header("User-Agent", beego.AppConfig.String("UserAgent"))

	r, err := c.Fetch(prefix)
	if err != nil {
		beego.Error(err)
	}

	for _, act := range r.Actresses {
		diva.ReadOrCreateByActress(act)
	}

	return nil
}

func DivaByApiActresses() error {
	c := apiactress.NewClient()
	c.Header("User-Agent", beego.AppConfig.String("UserAgent"))

	all, errs := c.FetchAll()
	if len(errs) > 0 {
		beego.Error(errs)
	}

	for _, r := range all {
		for _, act := range r.Actresses {
			diva.ReadOrCreateByActress(act)
		}
	}

	return nil
}

func AllDivaImageByGoogleimages() (err error) {
	var divas []*models.Diva
	models.Divas().RelatedSel().Limit(1000000).OrderBy("Updated").All(&divas)

	err = FillupFromGoogleimages(divas, "AV女優")
	return
}

func StarringDivaImageByGoogleimages() (err error) {
	var divas []*models.Diva
	diva.StarringDivas().RelatedSel().Limit(1000000).OrderBy("Updated").All(&divas)

	err = FillupFromGoogleimages(divas, "AV女優")
	return
}

func AllDivaInfoByWikipedia() (err error) {
	var divas []*models.Diva
	models.Divas().RelatedSel().Limit(1000000).OrderBy("Updated").All(&divas)
	return updateDivaInfoByWikipedia(divas)
}

func StarringDivaInfoByWikipedia() (err error) {
	var divas []*models.Diva
	diva.StarringDivas().RelatedSel().Limit(1000000).OrderBy("Updated").All(&divas)
	return updateDivaInfoByWikipedia(divas)
}

func updateDivaInfoByWikipedia(divas []*models.Diva) (err error) {
	c := divaextractor.NewWikipedia()
	c.Header("User-Agent", beego.AppConfig.String("UserAgent"))

	var msg string
	var errs []error
	t := shuffler.Shuffler(divas).(reflect.Value)
	for i := 0; i < t.Len(); i++ {
		time.Sleep(73 * time.Second)

		d := t.Index(i).Interface().(*models.Diva)

		// TODO: Birthday
		if d.Bracup != "" {
			continue // File ok
		}
		if len(errs) >= 100 {
			msg := fmt.Sprintf(
				"Max errors in wiki update: %d length",
				len(errs),
			)
			beego.Error(msg, errs)
			break
		}

		err := c.Do(d.Name)
		if err != nil {
			msg = fmt.Sprintf(
				"Wikipedia fetch error: %s", d.Name)
			beego.Warning(msg)
			errs = append(errs, errors.New(msg))

			// refresh client
			c = divaextractor.NewWikipedia()
			c.Header("User-Agent", beego.AppConfig.String("UserAgent"))
			continue
		}

		err = d.UpdateByWikipedia(c)
		if err != nil {
			msg = fmt.Sprintf(
				"Update diva by wikipedia: %s", d.Name)
			beego.Warning(msg)
			errs = append(errs, errors.New(msg))
		}
	}

	if len(errs) > 0 {
		msg := fmt.Sprintf(
			"Max errors by wiki update: %d length.",
			len(errs),
		)
		beego.Error(msg, errs)
		err = errors.New(msg)
	}

	return
}
