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

	Tags   []Tag   `gorm:"many2many:entry_tag;"`
	Images []Image `gorm:"many2many:entry_image;"`

	Scores []Score
}
