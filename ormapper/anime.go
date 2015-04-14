package ormapper

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

type Anime struct {
	Id int64

	Name  string
	Alias string

	Kana   string
	Romaji string
	Gyou   string

	Url         string
	Author      string
	Works       string
	ReleaseDate time.Time

	Outline string

	Created time.Time
	Updated time.Time

	Icon   *Image
	IconId sql.NullInt64

	Characters []*Character

	Pictures      []*Picture
	PicturesCount int
}

// Choiced mediatyped animation list
func PictureAnimations() *gorm.DB {
	return DB.Table("anime").
		Preload("Pictures").Preload("Characters").Preload("Icon").
		Select("anime.*, count(*) as pictures_count").
		Joins(`
		INNER JOIN picture ON picture.anime_id = anime.id 
		INNER JOIN entry ON entry.id = picture.entry_id
		INNER JOIN blog ON blog.id = entry.blog_id 
		LEFT OUTER JOIN image ON image.id = anime.icon_id
		`).
		Group("anime.id")
}
