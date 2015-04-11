package entry

import (
	"bitbucket.org/ikeikeikeike/antenna/ormapper"
	"github.com/jinzhu/gorm"
)

func PictureEntries() *gorm.DB {
	return ormapper.DB.Table("entry").
		Preload("Picture").Preload("Video").Preload("Summary").Preload("Blog").
		Preload("Scores").
		// Preload("Tags").Preload("Images"). XXX: not supported relation
		Select("entry.*").
		Joins(`
		INNER JOIN blog ON blog.id = entry.blog_id 
		INNER JOIN picture ON entry.id = picture.entry_id
		`).
		Order("entry.id DESC")
}
