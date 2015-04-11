package ormapper

import (
	"database/sql"
	"time"
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
