package ranking

import (
	"time"

	"github.com/jinzhu/gorm"
)

func RankMoreThanZero(db *gorm.DB) *gorm.DB {
	return db.Where("entry_ranking.rank > ?", 0)
}

func HasDayly(db *gorm.DB) *gorm.DB {
	return db.Where("entry_ranking.begin_name = ?", "dayly")
}

func HasWeekly(db *gorm.DB) *gorm.DB {
	return db.Where("entry_ranking.begin_name = ?", "weekly")
}

func HasMonthly(db *gorm.DB) *gorm.DB {
	return db.Where("entry_ranking.begin_name = ?", "monthly")
}

func HasYearly(db *gorm.DB) *gorm.DB {
	return db.Where("entry_ranking.begin_name = ?", "yearly")
}

// Where(args) scope

func FilterBeginTime(t time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if t.Unix() > 0 {
			db = db.Where("entry_ranking.begin_time = ?", t)
		}
		return db
	}
}
