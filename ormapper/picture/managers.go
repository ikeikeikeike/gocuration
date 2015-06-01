package picture

import "github.com/jinzhu/gorm"

func ImageCountMoreThanOnce(db *gorm.DB) *gorm.DB {
	return db.Where("picture.image_count > ?", 1)
}
