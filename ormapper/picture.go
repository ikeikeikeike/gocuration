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

	Images     []*Image
	Characters []*Character `gorm:"many2many:picture_character;"`
}

func (m *Picture) RelLoader() {
	// if m.Entry != nil {
	// DB.Model(&m).
	// Preload("Picture").Preload("Video").Preload("Summary").Preload("Blog").Preload("Scores").
	// Preload("Tags").Preload("Images")
	// Related(&m.Entry)
	// }
	if m.Anime != nil {
		DB.Model(&m).Preload("Icon").
			// Preload("Characters").
			// Preload("Pictures").
			Related(&m.Anime)
	}

	// DB.Model(&m).
	// Preload("Picture").
	// Related(&m.Images)

	// DB.Model(&m).
	// Related(&Characters)
}

func (m *Picture) NewsLoader() {
	if m.Anime != nil {
		DB.Model(&m).
			Preload("Icon").
			// Preload("Characters"). // XXX: full scan
			Related(&m.Anime)
	}

	// DB.Model(&m).Find(&m.Characters)  // XXX: full scan
}

func (m *Picture) ShowLoader() {
	m.NewsLoader()
}
