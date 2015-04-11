package ormapper

import (
	"database/sql"
	"time"
)

type Image struct {
	Id int64

	Name string
	Src  string

	Ext    string
	Mime   string
	Width  int
	Height int

	Created time.Time
	Updated time.Time

	Picture   *Picture
	PictureId sql.NullInt64 // `gorm:"column:picture_id"`
}
