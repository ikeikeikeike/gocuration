package verify

import (
	rss "github.com/ikeikeikeike/go-pkg-rss"
)

func IsRss(url string) bool {
	feed := rss.New(20, true, chHandler, itemHandler)
	err := feed.Fetch(url, nil)
	if err == nil {
		return true
	} else {
		return false
	}
}

func GetFeed(url string) (*rss.Feed, bool) {
	feed := rss.New(20, true, chHandler, itemHandler)
	err := feed.Fetch(url, nil)
	if err == nil {
		return feed, true
	} else {
		return feed, false
	}
}

func chHandler(f *rss.Feed, chs []*rss.Channel)                   {} // fmt.Printf("%d new channel(s) in %s\n", len(chs), f.Url)
func itemHandler(f *rss.Feed, ch *rss.Channel, items []*rss.Item) {} // fmt.Printf("%d new item(s) in %s\n", len(items), f.Url)
