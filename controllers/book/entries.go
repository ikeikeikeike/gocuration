package book

import (
	"bitbucket.org/ikeikeikeike/antenna/models"
	"bitbucket.org/ikeikeikeike/antenna/models/summary"
	"bitbucket.org/ikeikeikeike/antenna/ormapper"
	ani "bitbucket.org/ikeikeikeike/antenna/ormapper/anime"
	en "bitbucket.org/ikeikeikeike/antenna/ormapper/entry"
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

	// var pers = c.DefaultPers

	// var entries []*models.Entry
	// c.SetImage(c.SetAdvancedSearch(models.PictureEntries().RelatedSel(), ""), "").Limit(pers).All(&entries)
	// c.Data["Entries"] = entries

	var summaries []*models.Summary

	// c.SetImage(c.SetAdvancedSearch(models.PictureSummaries().RelatedSel(), "entry__").RelatedSel(), "entry__").Limit(pers).All(&summaries)
	c.Data["Summaries"] = summaries

	var list []*ormapper.Entry
	en.PictureEntries().
		Scopes(ani.FilterMediatype("image")).
		Limit(c.DefaultPers).
		Find(&list)

	c.Data["Entries"] = list

	var divas []*ormapper.Diva
	// di.VideoGoddess().
	// Scopes(di.VideoCountMoreThanZero).
	// Scopes(di.FilterMediatype("image")).
	// Scopes(di.FilterBlood(c.GetString("blood"))).
	// Scopes(di.FilterBracup(c.GetStrings("cup"))).
	// Scopes(di.FilterPrefixLines(c.GetString("line"))).
	// Scopes(di.FilterNameKana(convert.StrTo(c.GetString("q")).MultiWord())).
	// Limit(4).
	// Find(&divas)
	c.Data["Divas"] = divas

	var animes []*ormapper.Anime
	// ani.PictureAnimations().
	// Scopes(ani.PictureCountMoreThanZero).
	// Scopes(ani.FilterMediatype("image")).
	// Scopes(ani.FilterPrefixLines(c.GetString("line"))).
	// Scopes(ani.FilterNameKana(convert.StrTo(c.GetString("q")).MultiWord())).
	// Limit(4).
	// Find(&animes)
	c.Data["Animes"] = animes
}

func (c *EntriesController) News() {
	c.TplNames = "public/entries/news.tpl"

	pers := c.DefaultPers
	qs := c.SetAdvancedSearch(models.Entries().RelatedSel(), "")

	cnt, _ := models.CountObjects(qs)
	pager := pagination.SetPaginator(c.Ctx, pers, cnt)

	qs = qs.Limit(pers, pager.Offset())

	var entries []*models.Entry
	models.ListObjects(qs, &entries)

	c.Data["QURL"] = ""
	c.Data["Entries"] = entries
}

func (c *EntriesController) Hots() {
	c.TplNames = "book/entries/hots.tpl"

	pers := c.DefaultPers
	qs := c.SetAdvancedSearch(models.Summaries().RelatedSel(), "entry__")

	cnt, _ := models.CountObjects(qs)
	pager := pagination.SetPaginator(c.Ctx, pers, cnt)

	qs = qs.Limit(pers, pager.Offset())

	var summaries []*models.Summary
	models.ListObjects(qs, &summaries)

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
	s := &models.Entry{Id: uid}
	s.Read()

	if !s.IsLiving() {
		c.Ctx.Abort(404, "404")
		return
	}

	var (
		in         []string
		divas      []*models.Diva
		animes     []*models.Anime
		summaries  []*models.Summary
		characters []*models.Character
	)

	s.Blog.LoadRelated()
	if s.Video != nil {
		s.Video.LoadRelated()
		divas = s.Video.Divas
	}
	if s.Picture != nil {
		s.Picture.LoadRelated()
		if s.Picture.Anime != nil {
			animes = []*models.Anime{s.Picture.Anime}
		}
		characters = s.Picture.Characters
	}

	for _, t := range s.Tags {
		if t.Name != "" {
			in = append(in, t.Name)
		}
	}
	summary.RelatedSummaries(s.Id, in, &summaries)

	c.Data["Entry"] = s
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
	s := &models.Entry{Id: uid}
	s.Read()

	if !s.IsLiving() {
		c.Ctx.Abort(404, "404")
		return
	}

	s.Blog.LoadRelated()
	s.Picture.LoadRelated()

	c.Data["Entry"] = s
}

func (c *EntriesController) Search() {
	c.TplNames = "public/entries/search.tpl"
}
