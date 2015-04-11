package ormapper

import (
	"database/sql"
	"time"
)

type Entry struct {
	Id          int64
	Url         string
	Title       string
	Content     string
	SeoTitle    string
	SeoContent  string
	Encoded     string
	Creator     string
	Publisher   string
	PublishedAt time.Time

	Q string

	Created time.Time
	Updated time.Time

	Blog   *Blog
	BlogId sql.NullInt64

	Video   *Video
	Picture *Picture
	Summary *Summary

	Tags   []*Tag   `gorm:"many2many:entry_tag;"`
	Images []*Image `gorm:"many2many:entry_image;"`

	Scores []*Score
}

func (m *Entry) RelLoader() {
	if m.Summary != nil {
		DB.Model(&m).Preload("Entry").Preload("Scores").
			Related(&m.Summary)
	}
	if m.Blog != nil {
		DB.Model(&m).Preload("User").Preload("Icon").Preload("Scores").Preload("Entries").
			Related(&m.Blog)
	}
	if m.Video != nil {
		DB.Model(&m).Preload("Entry").Preload("Site").
			//Preload("Divas").
			Related(&m.Video)
	}
	if m.Picture != nil {
		DB.Model(&m).Preload("Entry").Preload("Anime").Preload("Images").
			//Preload("Characters").
			Related(&m.Picture)
	}

	DB.Model(&m).Preload("Entry").Preload("Blog").Preload("Summary").
		Related(&m.Scores)

	DB.Model(&m).Preload("Image").Preload("Entries").Association("Tags").
		Find(&m.Tags)

	DB.Model(&m).Preload("Picture").Association("Images").
		Find(&m.Images)

}
