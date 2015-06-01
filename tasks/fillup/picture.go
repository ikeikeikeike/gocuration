package fillup

import (
	"bitbucket.org/ikeikeikeike/antenna/models"
	"github.com/astaxie/beego/orm"
)

func PictureImageCount() error {
	o := orm.NewOrm()

	var all []*models.Picture
	models.Pictures().Limit(100000000000000).All(&all)

	for _, p := range all {
		p.RelLoader()

		_, _ = o.LoadRelated(p, "Images", 2, -1, 0, "-id")

		p.ImageCount = len(p.Images)
		p.Update("ImageCount")
	}

	return nil
}
