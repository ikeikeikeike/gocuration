package main

import (
	"fmt"
	"testing"

	rss "github.com/ikeikeikeike/go-pkg-rss"
)

func TestSQL(t *testing.T) {

	urls  := []string{
		"http://blog.livedoor.jp/pakoncho/index.rdf",
		"http://erotic.agirls.jp/?xml",
		"http://newerosanpo.blog.fc2.com/?xml",
		"http://nametaisyou.blog.fc2.com/?xml",
		"http://eigarape.blog.fc2.com/?xml",
		"http://wonderfulerosworld.blog.fc2.com/?xml",
		"http://erosoul.com/?xml",
		"http://chichishirimomo.com/?xml",
		"http://onaniedouga.net/?xml",
		"http://funkymovie.blog.fc2.com/?xml",
		"http://okazukyuubin.blog.fc2.com/?xml",
		"http://abzyoyusaikou.blog.fc2.com/?xml",
	}

	for _, url := range urls {
		feed := rss.New(20, true, chHandler, itemHandler)
		if err := feed.Fetch(url, nil); err != nil {
			t.Fatalf("Rss fetch error: %s", err)
		}
	}

}

func chHandler(f *rss.Feed, chs []*rss.Channel) {
	fmt.Printf("%d new channel(s) in %s\n", len(chs), f.Url)
}

func itemHandler(f *rss.Feed, ch *rss.Channel, items []*rss.Item) {
	fmt.Printf("%d new item(s) in %s\n", len(items), f.Url)
}
