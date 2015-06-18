package main

import (
	"fmt"
	"testing"

	rss "github.com/ikeikeikeike/go-pkg-rss"
	"github.com/ikeikeikeike/gopkg/extract/image"
	"github.com/k0kubun/pp"
	sani "github.com/kennygrant/sanitize"
)

func TestFeeds(t *testing.T) {

	urls := []string{
		// "http://blog.livedoor.jp/pakoncho/index.rdf",
		// "http://erotic.agirls.jp/?xml",
		// "http://newerosanpo.blog.fc2.com/?xml",
		// "http://nametaisyou.blog.fc2.com/?xml",
		// "http://eigarape.blog.fc2.com/?xml",
		// "http://wonderfulerosworld.blog.fc2.com/?xml",
		// "http://erosoul.com/?xml",
		// "http://chichishirimomo.com/?xml",
		// "http://onaniedouga.net/?xml",
		// "http://funkymovie.blog.fc2.com/?xml",
		// "http://okazukyuubin.blog.fc2.com/?xml",
		// "http://abzyoyusaikou.blog.fc2.com/?xml",
		"http://2jiero.dip.jp/feed",
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

	var Encoded, Content string

	for idx, it := range ch.Items {

		pp.Println(idx, it)
		// it.Links[0].Href

		Encoded = it.Encoded
		Content = it.Description
		if Content == "" {
			Content = Encoded
		}
		Content, _ = sani.HTMLAllowing(Content)
		Content = sani.HTML(Content)

		paths, _ := image.ExtractImagePaths(it.Encoded)
		if len(paths) <= 0 {
			paths, _ = image.ExtractImagePaths(it.Description)
		}

		pp.Println(paths)

		// s.PublishedAt, _ = time.Parse(tf, it.Date)
		// if it.PubDate != "" {
		// s.PublishedAt, _ = time.Parse(tf, it.PubDate)
		// }

	}

}
