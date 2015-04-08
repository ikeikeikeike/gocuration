package diva

import (
	"fmt"

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

func MediatypedDivas(mtype string, limit int, divas interface{}) {
	q := fmt.Sprintf(`
	SELECT d.*
	FROM diva d 
	LEFT OUTER JOIN image i ON i.id = d.icon_id 
	LEFT OUTER JOIN picture p ON p.id = i.picture_id 
	INNER JOIN video_diva vd ON vd.diva_id = d.id 
	INNER JOIN video v ON v.id = vd.video_id 
	LEFT OUTER JOIN entry e ON e.id = v.entry_id 
	LEFT OUTER JOIN blog b ON b.id = e.blog_id 
	WHERE d.videos_count > 0 AND b.mediatype = '%s'
	GROUP BY d.id 
	ORDER BY d.videos_count 
	DESC LIMIT %d
	`, mtype, limit)

	orm.NewOrm().Raw(q).QueryRows(divas)
}
