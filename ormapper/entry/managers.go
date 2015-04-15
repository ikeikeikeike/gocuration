package entry

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func FilterNameKana(words []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, word := range words {
			if word != "" {
				w := fmt.Sprintf("%%%s%%", word)
				q := "diva.name  like ? OR diva.kana  like ? OR "
				q += "anime.name like ? OR anime.kana like ?"
				db = db.Where(q, w, w, w, w)
			}
		}
		return db
	}
}

func FilterQ(words []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, word := range words {
			if word != "" {
				db = db.Where("entry.q like ?", fmt.Sprintf("%%%s%%", word))
			}
		}
		return db
	}
}
