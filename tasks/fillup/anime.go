package fillup

import (
	"errors"
	"fmt"
	"time"

	"bitbucket.org/ikeikeikeike/antenna/lib/models/picture"
	"bitbucket.org/ikeikeikeike/antenna/models"
	"bitbucket.org/ikeikeikeike/antenna/models/anime"
	"github.com/astaxie/beego"
	"github.com/ikeikeikeike/charneoapo"
	"github.com/ikeikeikeike/gopkg/convert"
)

func AnimeFromEntries() (err error) {
	var entries []*models.Entry

	qs := models.Entries().RelatedSel()
	qs.Limit(10000000).OrderBy("Id").All(&entries)

	errs := picture.AddsByEntries(entries)
	if len(errs) > 0 {
		beego.Error(errs)
	}
	return
}

func AnimeByCharacterModel() (err error) {
	var (
		anime *models.Anime
		clist []*models.Character
	)
	models.Characters().Limit(10000000).All(&clist)

	for _, c := range clist {
		c.RelLoader()

		anime = &models.Anime{Name: c.Product}
		anime.ReadOrCreate("Name")
		anime.RelLoader()

		c.Anime = anime
		c.Update("Anime")

		for _, p := range c.Pictures {
			p.Anime = anime
			p.Update("Anime")
		}
	}

	return
}

func AnimeByNeoapoByPrefix(prefix string) error {
	c := charneoapo.NewNeoapo()
	c.Header("User-Agent", beego.AppConfig.String("UserAgent"))

	err := c.Do("animes", prefix)
	if err != nil {
		beego.Error(err)
		return err
	} else if c.AnimeName() == "" {
		msg := fmt.Sprintf(
			"Name does not exists: %s", prefix)
		beego.Warning(msg)
		return errors.New(msg)
	}

	_, m, err := anime.ReadOrCreateByNeoapo(c)
	if m != nil && err == nil {
		m.UpdateByNeoapo(c)
	}
	return err
}

func AnimeByNeoapo() error {

	cnt, err := models.Animes().Count()
	if err != nil {
		return err
	}

	if cnt <= 0 {
		err := AnimeByNeoapoByPrefix("1")
		if err != nil {
			return err
		}
	}

	var c models.Anime
	err = models.Animes().Limit(1).One(&c)
	if err != nil {
		return err
	}
	prefix := c.Id

	var errs []error
	for true {

		prefix++

		err = AnimeByNeoapoByPrefix(convert.ToStr(prefix))
		if err != nil {
			errs = append(errs, err)
		}

		if len(errs) >= 100 {
			msg := fmt.Sprintf(
				"Max errors in neoapo update: %s length",
				len(errs),
			)
			beego.Error(msg, errs)
			return errors.New(msg)
		}

		time.Sleep(41 * time.Second)
	}
	return nil
}

func AnimeImageByGoogleimages() (err error) {
	var animes []*models.Anime
	// models.Animes().RelatedSel().Limit(1000000).All(&animes)
	anime.StarringAnimes().RelatedSel().Limit(1000000).All(&animes)

	err = FillupFromGoogleimages(animes, "エロアニメ")
	return
}
