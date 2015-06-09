package ormapper

import (
	"database/sql"
	"time"

	"github.com/jinzhu/now"
)

type Picture struct {
	Id         int64
	PageView   int64
	ImageCount int

	Created time.Time
	Updated time.Time

	Entry   *Entry
	EntryId sql.NullInt64

	Anime   *Anime
	AnimeId sql.NullInt64

	Images     []*Image
	Characters []*Character `gorm:"many2many:picture_character;"`
}

func (m *Picture) TodayRanking(name string) *PictureRanking {

	db := DB.
		Preload("Entry").Preload("Anime").Preload("Images").
		Select("picture_ranking.*").
		Where("picture_id = ?", m.Id)

	n := now.New(time.Now().UTC())

	switch name {
	case "w", "weekly":
		db = db.Where("begin_time = ?", n.BeginningOfWeek()).Where("begin_name = ?", "weekly")
	case "m", "monthly":
		db = db.Where("begin_time = ?", n.BeginningOfMonth()).Where("begin_name = ?", "monthly")
	case "y", "yearly":
		db = db.Where("begin_time = ?", n.BeginningOfYear()).Where("begin_name = ?", "yearly")
	case "d", "dayly":
		fallthrough
	default:
		db = db.Where("begin_time = ?", n.BeginningOfDay()).Where("begin_name = ?", "dayly")
	}

	var ranking PictureRanking
	if db.First(&ranking).Error != nil {
		return nil
	}

	return &ranking
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

func (m *Picture) RankingsLoader() {
	m.Entry.NewsLoader()
}

func (m *Picture) ShowLoader() {
	m.NewsLoader()
}
