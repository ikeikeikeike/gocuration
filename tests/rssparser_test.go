package main

import (
	"testing"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"

	rss "github.com/ikeikeikeike/go-pkg-rss"
	"github.com/k0kubun/pp"
)

func TestParser(t *testing.T) {
	var feed *rss.Feed

	feed = rss.New(20, true, chHandler, itemHandler)
	if err := feed.Fetch("http://erosha39.blog.fc2.com/?xml", nil); err != nil {
		t.Error(err)
		return
	}

	pp.Println(feed)

	for _, ch := range feed.Channels {
		for _, it := range ch.Items {
			pp.Println(it.Date)
		}
	}
}

func chHandler(f *rss.Feed, chs []*rss.Channel)                   {} // fmt.Printf("%d new channel(s) in %s\n", len(chs), f.Url)
func itemHandler(f *rss.Feed, ch *rss.Channel, items []*rss.Item) {} // fmt.Printf("%d new item(s) in %s\n", len(items), f.Url)
