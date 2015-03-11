package anime

import (
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
