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
			db = db.Where("diva.bracup in ?", ormapper.BracupLines[cup])
		}
		return db
	}
}

func FilterBracup(cups []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(cups) > 0 {
			db = db.Where("diva.bracup in ?", cups)
		}
		return db
	}
}

func FilterBlood(blood string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if blood != "" {
			db = db.Where("diva.blood", blood)
		}
		return db
	}
}

func FilterNameKana(words []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var conds []interface{}

		for _, w := range words {
			conds = append(conds, fmt.Sprintf("%%%s%%", w))
		}
		return db.Where("diva.name like ? OR diva.kana like ?", conds...)
	}
}

func FilterPrefixLines(line string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if line != "" {
			db = db.Where("diva.gyou in ?", ormapper.PrefixLines[line])
		}
		return db
	}
}
