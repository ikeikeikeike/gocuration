package ormapper

import (
	"database/sql"
	"time"
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

	Icon   Image
	IconId sql.NullInt64

	VideosCount int
	Videos      []Video
}
