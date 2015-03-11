package main

import (
	"fmt"
	"testing"

	rss "github.com/ikeikeikeike/go-pkg-rss"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestSQL(t *testing.T) {
	var (
		feed *rss.Feed
		url  = "http://blog.livedoor.jp/pakoncho/index.rdf"
	)

	feed = rss.New(20, true, chHandler, itemHandler)
	if err := feed.Fetch(url, nil); err != nil {
		t.Fatalf("Rss fetch error: %s", err)
	}

}

func chHandler(f *rss.Feed, chs []*rss.Channel) {
	fmt.Printf("%d new channel(s) in %s\n", len(chs), f.Url)
}

func itemHandler(f *rss.Feed, ch *rss.Channel, items []*rss.Item) {
	fmt.Printf("%d new item(s) in %s\n", len(items), f.Url)
}
