package anime

import (
	"fmt"

	"bitbucket.org/ikeikeikeike/antenna/ormapper"
	"github.com/jinzhu/gorm"
)

// Scope

func PictureCountMoreThanZero(db *gorm.DB) *gorm.DB {
	return db.Where("anime.pictures_count > ?", 0)
}

// Where(args) scope

func FilterMediatype(mtype string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("blog.mediatype = ?", mtype)
	}
}

func FilterNameKana(words []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var conds []interface{}

		for _, w := range words {
			conds = append(conds, fmt.Sprintf("%%%s%%", w))
		}
		return db.Where("anime.name like ? OR anime.kana like ?", conds...)
	}
}

func FilterPrefixLines(line string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if line != "" {
			db = db.Where("anime.gyou in ?", ormapper.PrefixLines[line])
		}
		return db
	}
}

// Choiced mediatyped animation list
func PictureAnimations() *gorm.DB {
	return ormapper.DB.Table("anime").
		Preload("Pictures").Preload("Characters").Preload("Icon").
		Select("anime.*").
		Joins(`
		JOIN picture ON picture.anime_id = anime.id 
		JOIN entry ON entry.id = picture.entry_id
		JOIN blog ON blog.id = entry.blog_id 
		LEFT OUTER JOIN image ON image.id = anime.icon_id
		`).
		Group("anime.id").
		Order("anime.pictures_count DESC")
}
