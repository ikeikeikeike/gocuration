package main

import (
	"testing"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	"bitbucket.org/ikeikeikeike/antenna/ormapper"
	"bitbucket.org/ikeikeikeike/antenna/ormapper/anime"
	"bitbucket.org/ikeikeikeike/antenna/ormapper/diva"
	"bitbucket.org/ikeikeikeike/antenna/ormapper/entry"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestPictureEntries(t *testing.T) {
	// normal struct test
	var list []ormapper.Entry
	entry.PictureEntries().
		Limit(20).
		Find(&list)

	for _, m := range list {
		m.RelLoader()
		if m.Blog.User.Id < 1 {
			t.Errorf("m.Blog.User.Id %d", m.Blog.User.Id)
		}
	}

	// pointer struct test
	list = make([]ormapper.Entry, 0)
	entry.PictureEntries().
		Limit(20).
		Find(&list)

	for _, m := range list {
		entry := m.Picture.Entry
		m.RelLoader()
		if entry == m.Picture.Entry {
			t.Errorf("m.Picture.Entry %v == %v", entry, m.Picture.Entry)
		}
	}
}

func TestVideoGoddess(t *testing.T) {
	var list []ormapper.Diva
	diva.VideoGoddess().
		Limit(20).
		Find(&list)

	for _, m := range list {
		if m.Id < 1 {
			t.Errorf("Diva.Id %d", m.Id)
		}
	}

}

func TestPictureAnimations(t *testing.T) {
	var list []ormapper.Anime
	anime.PictureAnimations().
		Limit(20).
		Find(&list)

	for _, m := range list {
		if m.Id < 1 {
			t.Errorf("Anime.Id %d", m.Id)
		}
	}
}
