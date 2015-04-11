package ormapper

import (
	"database/sql"
	"time"
)

type Blog struct {
	Id int64

	Rss         string
	Url         string
	Name        string
	Mediatype   string
	Adsensetype string

	VerifyLink  int
	VerifyRss   int
	VerifyParts int

	IsPenalty bool

	LastModified time.Time

	Created time.Time
	Updated time.Time

	User   *User
	UserId sql.NullInt64

	Icon   *Image
	IconId sql.NullInt64

	Scores  []*Score
	Entries []*Entry
}
