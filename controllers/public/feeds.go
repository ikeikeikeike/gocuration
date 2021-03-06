package public

import (
	"fmt"
	"net/url"
	"time"

	"bitbucket.org/ikeikeikeike/antenna/models"
	"bitbucket.org/ikeikeikeike/antenna/models/image"

	"github.com/gorilla/feeds"
	"github.com/ikeikeikeike/gopkg/convert"
)

type FeedsController struct {
	BaseController
}

func (c *FeedsController) Rdf() {}

func (c *FeedsController) Atom() {}

func (c *FeedsController) Rss() {
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	c.Ctx.Output.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS")

	now := time.Now()
	meta := c.Meta
	feed := &feeds.Feed{
		Title:       meta.Title,
		Link:        &feeds.Link{Href: meta.Host},
		Description: meta.Description,
		Author:      &feeds.Author{meta.Author, meta.Email},
		Created:     now,
	}

	var (
		src       string
		summaries []*models.Summary
	)
	qs := models.Summaries().RelatedSel().Limit(10)
	// qs = qs.Filter("entry__scores__name", "twitter").OrderBy("entry__scores__count")
	models.ListObjects(qs, &summaries)

	for _, summary := range summaries {
		e := summary.Entry
		e.LoadRelated()

		path := c.UrlFor("EntriesController.Show",
			":id", convert.ToStr(e.Id),
			":title", url.QueryEscape(e.SeoTitle),
		)
		href := c.BuildRequestUrl(path)

		src = image.CachedRandomSrc("large")
		if len(e.Images) > 0 {
			src = e.Images[0].Src
		}

		feed.Items = append(feed.Items, &feeds.Item{
			Title: e.Title,
			Link:  &feeds.Link{Href: href},
			Description: fmt.Sprintf(`<![CDATA[
				<div>
					<a href="%s">
						<img src="%s" alt="%s" 
							 style="max-height:400px; max-width:300px" />
						<div>
							<h6>%s</h6>
							<p>%s</p>
							<p>続きを読む</p>
						</div>
					</a>
				</div>
				]]>`,
				href, src, e.Blog.Name, e.Blog.Name, e.Title,
			),
			Author:  &feeds.Author{e.Creator, ""},
			Created: e.PublishedAt,
		})
	}

	rss, _ := feed.ToRss()
	c.Ctx.Output.Body([]byte(rss))
	c.ServeXml()
}
