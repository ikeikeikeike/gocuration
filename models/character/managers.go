package character

import (
	"bitbucket.org/ikeikeikeike/antenna/models"
	"github.com/astaxie/beego/orm"
	"github.com/ikeikeikeike/charneoapo"
)

func ReadOrCreateByNeoapo(n *charneoapo.Neoapo) (bool, *models.Character, error) {
	c := &models.Character{Name: n.Name(), Product: n.Product()}

	created, _, err := c.ReadOrCreate("Name", "Product")
	if created {
		anime := &models.Anime{Name: c.Product}
		anime.ReadOrCreate("Name")

		c.Anime = anime
		c.UpdateByNeoapo(n)
	}

	return created, c, err
}

// Search more than 0, showon video.
func StarringCharacters() orm.QuerySeter {
	return orm.NewOrm().QueryTable("character").Filter("pictures_count__gt", 0).OrderBy("-pictures_count")
}
