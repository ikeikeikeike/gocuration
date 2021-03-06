package ormapper

import (
	"database/sql"
	"time"

	"bitbucket.org/ikeikeikeike/antenna/ormapper/blog"
	"bitbucket.org/ikeikeikeike/antenna/ormapper/picture"
	"bitbucket.org/ikeikeikeike/antenna/ormapper/video"
)

type Blog struct {
	Id int64

	Rss         string
	Url         string
	Name        string
	Mediatype   string
	Adsensetype string

	VerifyParts     int
	VerifyRss       int
	VerifyLink      int
	VerifyBookRss   int
	VerifyBookLink  int
	VerifyVideoRss  int
	VerifyVideoLink int

	IsBan     string
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

func (m *Blog) PictureShowLoader() {
	PictureEntries().
		Scopes(blog.FilterMediatype("image")).
		Scopes(picture.ImageCountMoreThanOnce).
		Where("blog.id = ?", m.Id).
		Limit("20").
		Order("entry.id DESC").
		Find(&m.Entries)

	for _, e := range m.Entries {
		e.NewsLoader()
	}
}

func (m *Blog) VideoShowLoader() {
	VideoEntries().
		Scopes(blog.FilterMediatype("movie")).
		Scopes(video.HasVideo).
		Where("blog.id = ?", m.Id).
		Limit("20").
		Order("entry.id DESC").
		Find(&m.Entries)

	for _, e := range m.Entries {
		e.NewsLoader()
	}
}
