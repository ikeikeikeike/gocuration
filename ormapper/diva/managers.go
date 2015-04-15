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

func FilterBracupLines(cup string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if cup != "" {
			in, ok := ormapper.BracupLines[cup]
			if ok {
				db = db.Where("diva.bracup in (?)", in)
			}
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
			in, ok := ormapper.PrefixLines[line]
			if ok {
				db = db.Where("diva.gyou in (?)", in)
			}
		}
		return db
	}
}
