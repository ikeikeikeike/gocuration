package diva

import (
	"fmt"

	"bitbucket.org/ikeikeikeike/antenna/ormapper"
	"github.com/jinzhu/gorm"
)

// Scope

func VideoCountMoreThanZero(db *gorm.DB) *gorm.DB {
	return db.Where("diva.videos_count > ?", 0)
}

// Where(args) scope

func FilterMediatype(mtype string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("blog.mediatype = ?", mtype)
	}
}

func FilterBracupLines(cup string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if cup != "" {
			db = db.Where("diva.bracup in (?)", ormapper.BracupLines[cup])
		}
		return db
	}
}

func FilterBracup(cups []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(cups) > 0 {
			db = db.Where("diva.bracup in (?)", cups)
		}
		return db
	}
}

func FilterBlood(blood string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if blood != "" {
			db = db.Where("diva.blood = ?", blood)
		}
		return db
	}
}

func FilterNameKana(words []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, word := range words {
			if word != "" {
				w := fmt.Sprintf("%%%s%%", word)
				db = db.Where("diva.name like ? OR diva.kana like ?", w, w)
			}
		}
		return db
	}
}

func FilterPrefixLines(line string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if line != "" {
			db = db.Where("diva.gyou in (?)", ormapper.PrefixLines[line])
		}
		return db
	}
}

// Has video record list
func VideoGoddess() *gorm.DB {
	return ormapper.DB.Table("diva").
		Preload("Videos").Preload("Icon").
		Select("diva.*").
		Joins(`
		INNER JOIN video_diva ON video_diva.diva_id = diva.id 
		INNER JOIN video ON video.id = video_diva.video_id
		INNER JOIN entry ON entry.id = video.entry_id
		INNER JOIN blog ON blog.id = entry.blog_id 
		LEFT OUTER JOIN image ON image.id = diva.icon_id
		`).
		Group("diva.id").
		Order("diva.videos_count DESC")
}
