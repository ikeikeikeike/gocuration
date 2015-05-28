package ormapper

import (
	"database/sql"
	"time"
)

type Video struct {
	Id int64

	Url      string
	Code     string
	Duration int
	PageView int64

	Created time.Time
	Updated time.Time

	Entry   *Entry
	EntryId sql.NullInt64

	Site   *Site
	SiteId sql.NullInt64

	Divas []*Diva `gorm:"many2many:video_diva;"`
}

func (m *Video) NewsLoader() {
	if len(m.Divas) <= 0 {
		// XXX: 以下のクエリは order by diva.id asc なので注意が必要
		DB.Model(&m).Preload("Icon").Association("Divas").
			Find(&m.Divas)
	}
}

func (m *Video) ShowLoader() {
	m.NewsLoader()
}
