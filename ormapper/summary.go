package ormapper

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

type Summary struct {
	Id int64

	Sort int64

	Created time.Time
	Updated time.Time

	Entry   *Entry
	EntryId sql.NullInt64

	Scores []*Score
}

func (m *Summary) RelLoader() {
	if m.Entry != nil {
		DB.Model(&m).
			Preload("Picture").
			// Preload("Video").
			// Preload("Blog").
			// Preload("Summary").
			// Preload("Scores").
			// Preload("Tags").Preload("Images")
			Related(&m.Entry)
	}

	// DB.Model(&m).Preload("Entry").Preload("Blog").Preload("Summary").
	// Related(&m.Scores)
}

func (m *Summary) NewsLoader() {
	if m.Entry != nil {
		DB.Model(&m).
			Preload("Picture").Preload("Video").Preload("Blog").
			Related(&m.Entry)
		m.Entry.NewsLoader()
	}
}

func (m *Summary) ShowLoader() {
	m.NewsLoader()
}

func VideoSummaries() *gorm.DB {
	return DB.Table("summary").
		Preload("Entry").Preload("Scores").
		Select("summary.*").
		Joins(`
		INNER JOIN entry ON entry.id = summary.entry_id 
		INNER JOIN blog ON blog.id = entry.blog_id 
		INNER JOIN video ON entry.id = video.entry_id
		`)
}

func PictureSummaries() *gorm.DB {
	return DB.Table("summary").
		Preload("Entry").Preload("Scores").
		Select("summary.*").
		Joins(`
		INNER JOIN entry ON entry.id = summary.entry_id 
		INNER JOIN blog ON blog.id = entry.blog_id 
		INNER JOIN picture ON entry.id = picture.entry_id
		LEFT OUTER JOIN anime ON anime.id = picture.anime_id
		`)
}

func PictureShowSummaries() *gorm.DB {
	return DB.Table("summary").
		Preload("Entry").Preload("Scores").
		Select("summary.*").
		Joins(`
		INNER JOIN entry ON entry.id = summary.entry_id 
		INNER JOIN blog ON blog.id = entry.blog_id 
		INNER JOIN picture ON entry.id = picture.entry_id
		LEFT OUTER JOIN anime ON anime.id = picture.anime_id
		LEFT OUTER JOIN entry_tag et ON et.entry_id = entry.id
		LEFT OUTER JOIN tag ON tag.id = et.tag_id
		`)
}

func VideoShowSummaries() *gorm.DB {
	return DB.Table("summary").
		Preload("Entry").Preload("Scores").
		Select("summary.*").
		Joins(`
		INNER JOIN entry ON entry.id = summary.entry_id 
		INNER JOIN blog ON blog.id = entry.blog_id 
		INNER JOIN video ON entry.id = video.entry_id
		LEFT OUTER JOIN entry_tag et ON et.entry_id = entry.id
		LEFT OUTER JOIN tag ON tag.id = et.tag_id
		`)
}
