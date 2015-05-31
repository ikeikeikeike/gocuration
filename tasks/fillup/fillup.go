package fillup

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"bitbucket.org/ikeikeikeike/antenna/lib/verify"
	"bitbucket.org/ikeikeikeike/antenna/models"
	"github.com/astaxie/beego"
	"github.com/ikeikeikeike/go-googleimages"
	"github.com/ikeikeikeike/gopkg/extract/image"
	"github.com/ikeikeikeike/shuffler"
)

type GoogleimagesUpdater interface {
	UpdateIconByFileInfo(*image.FileInfo, string) error
}

func newImageClient() *googleimages.Client {
	c := googleimages.NewClient()
	c.Header("User-Agent", beego.AppConfig.String("UserAgent"))
	return c
}

func FillupFromGoogleimages(records interface{}, keyword string) (err error) {
	var (
		errs               []error
		src, name, product string
		fileinfo           *image.FileInfo
	)

	c := newImageClient()

	checker := verify.AvailableImageChecker()
	t := shuffler.Shuffler(records).(reflect.Value)

	for i := 0; i < t.Len(); i++ {
		time.Sleep(73 * time.Second)

		ref := t.Index(i).Interface()

		// TODO: To Interface logic.
		switch ref.(type) {
		case *models.Diva:
			obj := ref.(*models.Diva)
			obj.LoadRelated()
			src, name, product = obj.Icon.Src, obj.Name, ""
		case *models.Character:
			obj := ref.(*models.Character)
			obj.LoadRelated()
			if obj.Icon == nil {
				obj.Icon = &models.Image{}
				obj.Icon.Insert()
			}
			src, name, product = obj.Icon.Src, obj.Name, obj.Product
		case *models.Anime:
			obj := ref.(*models.Anime)
			obj.LoadRelated()
			if obj.Icon == nil {
				obj.Icon = &models.Image{}
				obj.Icon.Insert()
			}
			src, name, product = obj.Icon.Src, obj.Name, ""
		default:
			msg := "assetion type eror, give Diva or Character object"
			return errors.New(msg)
		}

		fileinfo, err = checker.Fetch(src)
		if err == nil && fileinfo != nil {
			if fileinfo.Width >= 300 && fileinfo.Height >= 200 {
				continue // File ok
			}
		}
		// File ng

		result, err := c.Fetch(fmt.Sprintf("%s %s %s", name, keyword, product))

		if err != nil {
			c = newImageClient()
			beego.Warning(err)
			errs = append(errs, err)
			continue
		} else if len(result.ResponseData.Results) <= 0 {
			c = newImageClient()
			continue
		} else if len(errs) >= 10 {
			msg := fmt.Sprintf(
				"Max errors by google image update: %s length.",
				len(errs),
			)
			beego.Error(msg, errs)
			return errors.New(msg)
		}

		obj := ref.(GoogleimagesUpdater)

		for _, r := range result.ResponseData.Results {

			fileinfo, err = checker.Fetch(r.Url)
			if err == nil && fileinfo != nil {

				if fileinfo.Width >= 300 && fileinfo.Height >= 200 {
					err := obj.UpdateIconByFileInfo(fileinfo, r.TitleNoFormatting)
					if err != nil {
						beego.Warning(err)
					} else {
						break
					}
				}
			}
		}

	}

	return
}
