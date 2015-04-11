package ormapper

import (
	"database/sql"
	"time"
)

type Character struct {
	Id int64

	Name string

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

	Product string
	Anime   *Anime
	AnimeId sql.NullInt64

	PicturesCount int
	Pictures      []*Picture
}
