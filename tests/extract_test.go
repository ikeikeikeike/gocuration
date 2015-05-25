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
		"http://okazukyuubin.blog.fc2.com/blog-entry-1140.html",
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

		// ext.Urls()
		// ext.Codes()

		pp.Println(ext.Urls(), ext.Codes())
	}

}
