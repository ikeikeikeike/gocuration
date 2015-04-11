package ormapper

import (
	"database/sql"
	"time"
)

type Tag struct {
	Id int64

	Name   string
	Kana   string
	Romaji string

	Created time.Time
	Updated time.Time

	Image   *Image
	ImageId sql.NullInt64

	Entries []*Entry
}

func (m Tag) TableName() string {
	return "tag"
}
