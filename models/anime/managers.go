package anime

import (
	"fmt"

	"bitbucket.org/ikeikeikeike/antenna/models"
	"github.com/astaxie/beego/orm"
	"github.com/ikeikeikeike/charneoapo"
)

func ReadOrCreateByNeoapo(n *charneoapo.Neoapo) (bool, *models.Anime, error) {
	m := &models.Anime{Name: n.AnimeName()}

	created, _, err := m.ReadOrCreate("Name")
	if created {
		m.UpdateByNeoapo(n)
	}
	return created, m, err
}

// Search more than 0, showon video.
func StarringAnimes() orm.QuerySeter {
	return orm.NewOrm().QueryTable("anime").Filter("pictures_count__gt", 0).OrderBy("-pictures_count")
}

func MediatypedAnimes(mtype string, limit int, animes interface{}) {
	q := fmt.Sprintf(`
	SELECT a.* 
	FROM anime a 
	LEFT OUTER JOIN image i ON i.id = a.icon_id 
	LEFT OUTER JOIN picture p ON p.id = i.picture_id 
	INNER JOIN picture p2 ON p2.anime_id = a.id 
	LEFT OUTER JOIN entry e ON e.id = p2.entry_id
	LEFT OUTER JOIN blog b ON b.id = e.blog_id 
	WHERE a.pictures_count > 0 AND b.mediatype = '%s'
	GROUP BY a.id
	ORDER BY a.pictures_count DESC 
	LIMIT %d
	`, mtype, limit)

	orm.NewOrm().Raw(q).QueryRows(animes)
}
