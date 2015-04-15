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

func FilterNameKana(words []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, word := range words {
			if word != "" {
				w := fmt.Sprintf("%%%s%%", word)
				db = db.Where("anime.name like ? OR anime.kana like ?", w, w)
			}
		}
		return db
	}
}

func FilterPrefixLines(line string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if line != "" {
			in, ok := ormapper.PrefixLines[line]
			if ok {
				db = db.Where("anime.gyou in (?)", in)
			}
		}
		return db
	}
}
