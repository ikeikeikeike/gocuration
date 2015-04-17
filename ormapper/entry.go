package ormapper

import (
	"database/sql"
	"strings"
	"time"

	"bitbucket.org/ikeikeikeike/antenna/ormapper/blog"

	"github.com/jinzhu/gorm"
)

type Entry struct {
	Id          int64
	Url         string
	Title       string
	Content     string
	SeoTitle    string
	SeoContent  string
	Encoded     string
	Creator     string
	Publisher   string
	PublishedAt time.Time

	Q string

	Created time.Time
	Updated time.Time

	Blog   *Blog
	BlogId sql.NullInt64

	Video   *Video
	Picture *Picture
	Summary *Summary

	Tags   []*Tag   `gorm:"many2many:entry_tag;"`
	Images []*Image `gorm:"many2many:entry_image;"`

	Scores []*Score
}

// Is liveing entry?
func (m *Entry) IsLiving() bool {
	if m.Id <= 0 || m.Blog == nil {
		return false
	}
	return true
}

func (m *Entry) GetByBlog() (*gorm.DB, error) {
	if !m.IsLiving() {
		return nil, nil
	}

	db := DB.
		Preload("Picture").Preload("Video").Preload("Blog").
		Select("entry.*").
		Joins(`INNER JOIN blog ON blog.id = entry.blog_id`).
		Scopes(blog.FilterMediatype(m.Blog.Mediatype))

	return db, nil
}

func (m *Entry) PreviousByBlog() (*Entry, error) {
	db, err := m.GetByBlog()
	if err != nil {
		return nil, err
	}

	var list []*Entry
	db = db.
		Where("entry.blog_id = ?", m.Blog.Id).
		Where("entry.id < ?", m.Id).
		Last(&list)

	if db.RecordNotFound() {
		return nil, db.Error
	}

	e := list[0]
	e.NewsLoader()

	return e, nil
}

func (m *Entry) NextByBlog() (*Entry, error) {
	db, err := m.GetByBlog()
	if err != nil {
		return nil, err
	}

	var list []*Entry
	db = db.
		Where("entry.blog_id = ?", m.Blog.Id).
		Where("entry.id > ?", m.Id).
		First(&list)

	if db.RecordNotFound() {
		return nil, db.Error
	}

	e := list[0]
	e.NewsLoader()

	return e, nil
}

func (m *Entry) CommaTags() string {
	tags := []string{}
	if m.Tags != nil {
		for _, tag := range m.Tags {
			if tag.Name != "" {
				tags = append(tags, tag.Name)
			}
		}
	}
	return strings.Join(tags, ",")
}

func (m *Entry) RelLoader() {
	// if m.Summary != nil {
	// DB.Model(&m).Preload("Entry").Preload("Scores").
	// Related(&m.Summary)
	// }
	if m.Blog != nil {
		DB.Model(&m).Preload("User").Preload("Icon").
			// Preload("Scores").
			// Preload("Entries").
			Related(&m.Blog)
	}
	if m.Video != nil {
		DB.Model(&m).Preload("Entry").Preload("Site").
			//Preload("Divas").
			Related(&m.Video)
	}
	if m.Picture != nil {
		DB.Model(&m).Preload("Entry").Preload("Anime").Preload("Images").
			//Preload("Characters").
			Related(&m.Picture)

		m.Picture.NewsLoader()
	}

	// DB.Model(&m).Preload("Entry").Preload("Blog").Preload("Summary").
	// Related(&m.Scores)

	DB.Model(&m).Preload("Image").Preload("Entries").Association("Tags").
		Find(&m.Tags)

	DB.Model(&m).Preload("Picture").Association("Images").
		Find(&m.Images)

}

func (m *Entry) NewsLoader() {
	if m.Blog != nil {
		DB.Model(&m).Preload("Icon").Related(&m.Blog)
	}
	if m.Video != nil {
		DB.Model(&m).Preload("Site").Related(&m.Video)
	}
	if m.Picture != nil {
		DB.Model(&m).Preload("Anime").Preload("Images").
			Related(&m.Picture)
		m.Picture.NewsLoader()
	}
	DB.Model(&m).Association("Tags").Find(&m.Tags)
	DB.Model(&m).Association("Images").Find(&m.Images)
}

func (m *Entry) PictureShowLoader() {
	if m.Blog != nil {
		DB.Model(&m).Preload("Icon").Related(&m.Blog)
		m.Blog.PictureShowLoader()
	}

	m.NewsLoader()
}

func PictureEntries() *gorm.DB {
	return DB.Table("entry").
		Preload("Picture").Preload("Video").Preload("Summary").Preload("Blog").
		Preload("Scores").
		// Preload("Tags").Preload("Images"). XXX: not supported relation
		Select("entry.*").
		Joins(`
		INNER JOIN blog ON blog.id = entry.blog_id 
		INNER JOIN picture ON entry.id = picture.entry_id
		LEFT OUTER JOIN anime ON anime.id = picture.anime_id
		`)
}

func VideoEntries() *gorm.DB {
	return DB.Table("entry").
		Preload("Picture").Preload("Video").Preload("Summary").Preload("Blog").
		Preload("Scores").
		// Preload("Tags").Preload("Images"). XXX: not supported relation
		Select("entry.*").
		Joins(`
		INNER JOIN blog ON blog.id = entry.blog_id 
		INNER JOIN video ON entry.id = video.entry_id
		`)
}
