package ormapper

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

type Diva struct {
	Id int64

	Name   string
	Kana   string
	Romaji string
	Gyou   string

	Birthday time.Time

	Blood string

	Height int
	Weight int

	Bust   int
	Waste  int
	Hip    int
	Bracup string

	Outline string

	Created time.Time
	Updated time.Time

	Icon   *Image
	IconId sql.NullInt64

	VideosCount int
	Videos      []*Video
}

func (m Diva) TableName() string {
	return "diva"
}

// Has video record list
func VideoGoddess() *gorm.DB {
	return DB.Table("diva").
		Preload("Videos").Preload("Icon").
		Select("diva.*").
		Joins(`
		INNER JOIN video_diva ON video_diva.diva_id = diva.id 
		INNER JOIN video ON video.id = video_diva.video_id
		INNER JOIN entry ON entry.id = video.entry_id
		INNER JOIN blog ON blog.id = entry.blog_id 
		LEFT OUTER JOIN image ON image.id = diva.icon_id
		`).
		Group("diva.id")
}
