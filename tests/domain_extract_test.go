package main

import (
	"fmt"
	"strings"
	"testing"

	urlparse "net/url"

	gq "github.com/PuerkitoBio/goquery"
	behavior "github.com/ikeikeikeike/gopkg/net/http"
	"github.com/k0kubun/pp"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestRssFeed(t *testing.T) {
	// url := "http://futomomofetigazo.blog.fc2.com/blog-entry-1721.html"
	// url := "http://djinem.blog.fc2.com/blog-entry-2009.html"
	// url := "http://dmm-free-movie.com/archives/22665340.html"
	// url := "http://moroahedoujin.com/blog-entry-4988.html"
	url := "http://blog.livedoor.jp/eresoku/archives/43383453.html"
	u, err := urlparse.Parse(strings.TrimSpace(url))
	if err != nil || u.Host == "" {
		return
	}

	var src string
	domain := strings.Split(u.Host, ":")[0]

	if strings.Contains(domain, "livedoor.jp") {
		src = strings.Split(u.Path, "/")[1]
	} else if strings.Contains(domain, "fc2.com") {
		src = strings.Split(domain, ".")[0]
	} else {
		src = strings.Split(domain, ".")[0]
	}

	doc, err := func(urlStr string) (*gq.Document, error) {
		e := behavior.NewUserBehavior()
		resp, err := e.Behave(urlStr)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return gq.NewDocumentFromResponse(resp)
	}(url)
	if err != nil {
		t.Fatal(err)
	}

	selector := fmt.Sprintf(`img[src*='%s']`, src)
	doc.Find(selector).Each(func(i int, s *gq.Selection) {
		src, ok := s.Attr("src")
		if !ok {
			return
		}
		println(src)
	})

	pp.Println(selector)
}
