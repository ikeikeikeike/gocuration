package ormapper

import (
	"database/sql"
	"time"
)

type Picture struct {
	Id int64

	Created time.Time
	Updated time.Time

	Entry   *Entry
	EntryId sql.NullInt64

	Anime   *Anime
	AnimeId sql.NullInt64

	Images     []Image
	Characters []Character `gorm:"many2many:picture_character;"`
}
