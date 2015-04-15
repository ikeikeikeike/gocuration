package blog

import "github.com/jinzhu/gorm"

// Where(args) scope

func FilterMediatype(mtype string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if mtype != "" {
			db = db.Where("blog.mediatype = ?", mtype)
		}
		return db
	}
}

func FilterAdsensetype(adtype string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if adtype != "" {
			db = db.Where("blog.adsensetype = ?", adtype)
		}
		return db
	}
}
