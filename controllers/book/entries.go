package book

import (
	"fmt"

	"bitbucket.org/ikeikeikeike/antenna/ormapper"
	"bitbucket.org/ikeikeikeike/antenna/ormapper/anime"
	"bitbucket.org/ikeikeikeike/antenna/ormapper/blog"
	"bitbucket.org/ikeikeikeike/antenna/ormapper/entry"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/ikeikeikeike/gopkg/convert"
	// "github.com/k0kubun/pp"
)

type EntriesController struct {
	BaseController
}

// func (c *EntriesController) NestPrepare() {}

func (c *EntriesController) NestFinish() {
	c.PushInAccessLog()
}

func (c *EntriesController) Home() {
	c.TplNames = "book/entries/home.tpl"

	var summaries []*ormapper.Summary
	ormapper.PictureSummaries().
		Scopes(blog.FilterMediatype("image")).
		Scopes(blog.FilterAdsensetype(c.GetString("at"))).
		Scopes(anime.FilterPrefixLines(c.GetString("line"))).
		Scopes(entry.FilterQ(convert.StrTo(c.GetString("q")).MultiWord())).
		Limit(c.DefaultPers).
		Order("summary.sort DESC").
		Find(&summaries)
	for _, s := range summaries {
		s.NewsLoader()
	}
	c.Data["Summaries"] = summaries

	var entries []*ormapper.Entry
	ormapper.PictureEntries().
		Scopes(blog.FilterMediatype("image")).
		Scopes(blog.FilterAdsensetype(c.GetString("at"))).
		Scopes(anime.FilterPrefixLines(c.GetString("line"))).
		Scopes(entry.FilterQ(convert.StrTo(c.GetString("q")).MultiWord())).
		Limit(c.DefaultPers).
		Order("entry.id DESC").
		Find(&entries)
	for _, e := range entries {
		e.NewsLoader()
	}
	c.Data["Entries"] = entries

	var divas []*ormapper.Diva
	c.Data["Divas"] = divas

	var animes []*ormapper.Anime
	ormapper.PictureAnimations().
		Scopes(blog.FilterMediatype("image")).
		Scopes(anime.PictureCountMoreThanZero).
		Scopes(anime.FilterPrefixLines(c.GetString("line"))).
		Scopes(anime.FilterNameKana(convert.StrTo(c.GetString("q")).MultiWord())).
		Limit(4).
		Order("anime.pictures_count DESC").
		Find(&animes)
	c.Data["Animes"] = animes
}

func (c *EntriesController) News() {
	c.TplNames = "book/entries/news.tpl"

	db := ormapper.PictureEntries().
		Scopes(blog.FilterMediatype("image")).
		Scopes(anime.FilterPrefixLines(c.GetString("line"))).
		Scopes(anime.FilterNameKana(convert.StrTo(c.GetString("q")).MultiWord()))

	var count int64
	db.Count(&count)

	pager := pagination.SetPaginator(c.Ctx, c.DefaultPers, count)
	db = db.Limit(c.DefaultPers).Offset(pager.Offset())

	var entries []*ormapper.Entry
	db.Order("entry.id DESC").Find(&entries)

	for _, e := range entries {
		e.NewsLoader()
	}

	c.Data["QURL"] = ""
	c.Data["Entries"] = entries
}

func (c *EntriesController) Hots() {
	c.TplNames = "book/entries/hots.tpl"

	db := ormapper.PictureSummaries().
		Scopes(blog.FilterMediatype("image")).
		Scopes(anime.FilterPrefixLines(c.GetString("line"))).
		Scopes(anime.FilterNameKana(convert.StrTo(c.GetString("q")).MultiWord()))

	var count int64
	db.Count(&count)

	pager := pagination.SetPaginator(c.Ctx, c.DefaultPers, count)
	db = db.Limit(c.DefaultPers).Offset(pager.Offset())

	var summaries []*ormapper.Summary
	db.Order("summary.sort DESC").Find(&summaries)

	for _, s := range summaries {
		s.NewsLoader()
	}

	c.Data["QURL"] = ""
	c.Data["Summaries"] = summaries
}

func (c *EntriesController) Show() {
	c.TplNames = "book/entries/show.tpl"

	id := c.Ctx.Input.Param(":id")
	if id == "" {
		c.Ctx.Abort(404, "404")
		return
	}

	uid, _ := convert.StrTo(id).Int64()

	m := &ormapper.Entry{Id: uid}
	ormapper.DB.
		Preload("Picture").Preload("Video").Preload("Blog").
		First(m)

	if !m.IsLiving() {
		c.Ctx.Abort(404, "404")
		return
	}

	m.PictureShowLoader()

	var (
		divas      []*ormapper.Diva
		animes     []*ormapper.Anime
		summaries  []*ormapper.Summary
		characters []*ormapper.Character
	)

	if m.Video != nil {
		divas = m.Video.Divas
	}
	if m.Picture != nil {
		characters = m.Picture.Characters
		if m.Picture.Anime != nil {
			animes = []*ormapper.Anime{m.Picture.Anime}
		}
	}

	var in []string
	for _, t := range m.Tags {
		if t.Name != "" {
			in = append(in, t.Name)
		}
	}
	if len(in) <= 0 {
		in = append(in, "同人")
	}
	ormapper.PictureShowSummaries().
		Scopes(blog.FilterMediatype("image")).
		Where("entry.id != ?", m.Id).
		Where("tag.name IN (?) OR entry.q like ?", in, fmt.Sprintf("%%%s%%", in[0])).
		Group("summary.id").
		Order("summary.sort DESC").
		Limit(3).
		Find(&summaries)
	for _, s := range summaries {
		s.ShowLoader()
	}

	c.Data["Entry"] = m
	c.Data["Divas"] = divas
	c.Data["Animes"] = animes
	c.Data["Summaries"] = summaries
	c.Data["Characters"] = characters
}

func (c *EntriesController) Viewer() {
	c.Layout = "base/none/base.tpl"
	c.TplNames = "book/entries/viewer.tpl"

	id := c.Ctx.Input.Param(":id")
	if id == "" {
		c.Ctx.Abort(404, "404")
		return
	}

	uid, _ := convert.StrTo(id).Int64()

	m := &ormapper.Entry{Id: uid}
	ormapper.DB.
		Preload("Picture").Preload("Video").Preload("Blog").
		First(m)

	if !m.IsLiving() {
		c.Ctx.Abort(404, "404")
		return
	}

	m.PictureShowLoader()

	c.Data["Entry"] = m
}

func (c *EntriesController) Search() {
	c.TplNames = "public/entries/search.tpl"
}
