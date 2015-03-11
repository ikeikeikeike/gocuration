package fillup

import (
	"errors"
	"fmt"
	"time"

	"bitbucket.org/ikeikeikeike/antenna/models"
	"bitbucket.org/ikeikeikeike/antenna/models/character"
	"github.com/astaxie/beego"
	"github.com/ikeikeikeike/charneoapo"
	"github.com/ikeikeikeike/gopkg/convert"
)

func CharacterByNeoapoByPrefix(prefix string) error {
	c := charneoapo.NewNeoapo()
	c.Header("User-Agent", beego.AppConfig.String("UserAgent"))

	err := c.Do("characters", prefix)
	if err != nil {
		beego.Error(err)
		return err
	} else if c.Name() == "" {
		msg := fmt.Sprintf(
			"Name does not exists: %s", prefix)
		beego.Warning(msg)
		return errors.New(msg)
	} else if c.Product() == "" {
		msg := fmt.Sprintf(
			"Product does not exists: %s", prefix)
		beego.Warning(msg)
		return errors.New(msg)
	}

	_, _, err = character.ReadOrCreateByNeoapo(c)
	return err
}

func CharacterByNeoapo() error {

	cnt, err := models.Characters().Count()
	if err != nil {
		return err
	}

	if cnt <= 0 {
		err := CharacterByNeoapoByPrefix("1")
		if err != nil {
			return err
		}
	}

	var c models.Character
	err = models.Characters().Limit(1).One(&c)
	if err != nil {
		return err
	}
	prefix := c.Id

	var errs []error
	for true {

		prefix++

		err = CharacterByNeoapoByPrefix(convert.ToStr(prefix))
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

func CharacterImageByGoogleimages() (err error) {
	var clist []*models.Character
	//models.Characters().RelatedSel().Limit(1000000).All(&clist)
	character.StarringCharacters().RelatedSel().Limit(1000000).All(&clist)

	err = FillupFromGoogleimages(clist, "エロアニメ")
	return
}
