package main

import (
	"strings"
	"testing"

	"github.com/k0kubun/pp"

	libentry "bitbucket.org/ikeikeikeike/antenna/lib/models/entry"
	"bitbucket.org/ikeikeikeike/antenna/models"

	// _ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	// _ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestExtract(t *testing.T) {

	var urlsForTest = []string{
		// "http://xvideo0matome.blog.fc2.com/blog-entry-6671.html",
		// "http://xvideocollection.blog.fc2.com/blog-entry-3974.html",
		// "http://xvideocollection.blog.fc2.com/blog-entry-3976.html",
		// "http://kyonyuufera.blog.fc2.com/blog-entry-806.html",
		// "http://hinapavloe.blog.fc2.com/blog-entry-2035.html",
		// "http://chirari2ch.com/blog-entry-19433.html",
		// "http://erodougacollections.blog.fc2.com/blog-entry-3006.html",
		// "http://muu.be/v/6240pji",
		// "http://mumouai.blog.fc2.com/blog-entry-2328.html",
		// "http://proyou.blog.fc2.com/blog-entry-227.html",
		// "http://aiaikurumi.blog.fc2.com/blog-entry-2245.html",
		// "http://eromatometayo.blog.fc2.com/blog-entry-2374.html",
		// "http://erogoo2525.blog.fc2.com/blog-entry-1009.html",
		"http://erolog.info/watch/ero139978",
		"http://erolog.info/watch/ero140719",
		"http://erolog.info/watch/ero136313",
		"http://erolog.info/watch/ero140717",
		"http://erolog.info/watch/ero140192",
		"http://erolog.info/watch/ero140515",
		"http://erolog.info/watch/ero138002",
		"http://erolog.info/watch/ero137084",
		"http://erolog.info/watch/ero140075",
		"http://erolog.info/watch/ero140281",
		"http://erolog.info/watch/ero140438",
		"http://erolog.info/watch/ero140521",
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

		pp.Println(ext.Urls(), strings.Join(ext.Codes(), ""))
	}

}
