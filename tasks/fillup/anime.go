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
		c.Update("Anime", "Updated")

		for _, p := range c.Pictures {
			p.Anime = anime
			p.Update("Anime", "Updated")
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

// func AnimeKanaByMecab(option string) (err error) {
// option = fmt.Sprintf("-d %s", option)

// var animes []*models.Anime
// models.Animes().RelatedSel().Limit(1000000).All(&animes)

// for _, m := range animes {

// // "-d /usr/local/Cellar/mecab/0.996/lib/mecab/dic/mecab-ipadic-neologd"
// result, err := mecab.Parse(m.Name, option)
// if err != nil {
// beego.Error(err)
// continue
// }

// var name, kana, hira, roma []string
// for _, res := range result {
// name = append(name, res.Surface)
// roma = append(roma, res.Romaji)
// kana = append(kana, res.Read)
// hira = append(hira, res.Hiragana)
// }

// // namae := strings.Join(name, "")

// if m.Kana == "" {
// m.Kana = strings.Join(kana, "")
// }
// if m.Romaji == "" {
// m.Romaji = strings.Join(roma, " ")
// }
// if m.Gyou == "" {
// hiragana := strings.Join(hira, "")
// if hiragana != "" {
// result, _ = mecab.Parse(string([]rune(hiragana)[0]), option)
// for _, r := range result {
// m.Gyou = r.Kunrei
// }
// }
// }

// m.Update("Kana", "Romaji", "Gyou", "Updated")
// }

// return
// }
