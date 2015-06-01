package video

import "github.com/jinzhu/gorm"

func HasVideo(db *gorm.DB) *gorm.DB {
	return db.Where("video.url != ? OR video.code != ?", "", "")
}
