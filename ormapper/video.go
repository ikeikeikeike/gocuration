package ormapper

import (
	"database/sql"
	"time"
)

type VideoUrl struct {
	Id int64

	Name string

	Created time.Time
	Updated time.Time

	Video   *Video
	VideoId sql.NullInt64
}

type VideoCode struct {
	Id int64

	Name string

	Created time.Time
	Updated time.Time

	Video   *Video
	VideoId sql.NullInt64
}

type Video struct {
	Id int64

	PageView int64

	Url      string
	Code     string
	Duration int

	Urls  []*VideoUrl
	Codes []*VideoCode

	Created time.Time
	Updated time.Time

	Entry   *Entry
	EntryId sql.NullInt64

	Site   *Site
	SiteId sql.NullInt64

	Divas []*Diva `gorm:"many2many:video_diva;"`
}

func (m *Video) NewsLoader() {

	if len(m.Codes) <= 0 {
		DB.Model(&m).Related(&m.Codes, "Codes")
	}
	if len(m.Urls) <= 0 {
		DB.Model(&m).Related(&m.Urls, "Urls")
	}

	if len(m.Divas) <= 0 {
		// XXX: 以下のクエリは order by diva.id asc なので注意が必要
		DB.Model(&m).Preload("Icon").Association("Divas").
			Find(&m.Divas)
	}
}

func (m *Video) RankingsLoader() {
	m.Entry.NewsLoader()
}

func (m *Video) ShowLoader() {
	m.NewsLoader()
}
