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

func DivaImageByGoogleimages() (err error) {
	var divas []*models.Diva
	diva.StarringDivas().RelatedSel().Limit(1000000).All(&divas)

	err = FillupFromGoogleimages(divas, "AV女優")
	return
}

func DivaInfoByWikipedia() (err error) {
	c := divaextractor.NewWikipedia()
	c.Header("User-Agent", beego.AppConfig.String("UserAgent"))

	var divas []*models.Diva
	diva.StarringDivas().RelatedSel().Limit(1000000).All(&divas)

	var msg string
	var errs []error
	t := shuffler.Shuffler(divas).(reflect.Value)
	for i := 0; i < t.Len(); i++ {
		time.Sleep(41 * time.Second)

		d := t.Index(i).Interface().(*models.Diva)

		// TODO: Birthday
		if d.Bracup != "" {
			continue // File ok
		}

		err := c.Do(d.Name)
		if err != nil {
			msg = fmt.Sprintf(
				"Wikipedia fetch error: %s", d.Name)
			beego.Warning(msg)
			errs = append(errs, errors.New(msg))
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
			"Max errors by google image update: %d length.",
			len(errs),
		)
		beego.Error(msg, errs)
		err = errors.New(msg)
	}

	return
}
