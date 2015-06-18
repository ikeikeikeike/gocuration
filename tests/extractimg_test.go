package main

import (
	"testing"

	"github.com/k0kubun/pp"

	libentry "bitbucket.org/ikeikeikeike/antenna/lib/models/entry"
	"bitbucket.org/ikeikeikeike/antenna/models"

	// _ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	// _ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestExtract(t *testing.T) {

	var urlsForTest = []string{
		"http://doujinsieromangasouko.net/blog-entry-2150.html",
		"http://2jiero.dip.jp/archives/58172",
	}

	for _, url := range urlsForTest {
		println("")
		println("___________")
		println(url)
		entry := &models.Entry{Url: url}

		ext := libentry.NewExtractor()
		if err := ext.Do(entry); err != nil {
			t.Error(err)
		}

		pp.Println(ext.Imgs())
	}

}
