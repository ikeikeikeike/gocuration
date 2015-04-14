package blog

import "github.com/jinzhu/gorm"

// Where(args) scope

func FilterMediatype(mtype string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("blog.mediatype = ?", mtype)
	}
}
