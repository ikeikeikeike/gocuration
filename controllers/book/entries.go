package book

import (
	"bitbucket.org/ikeikeikeike/antenna/models"
	"bitbucket.org/ikeikeikeike/antenna/models/anime"
	"bitbucket.org/ikeikeikeike/antenna/models/diva"
	"bitbucket.org/ikeikeikeike/antenna/models/summary"

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

	var (
		divas     []*models.Diva
		animes    []*models.Anime
		entries   []*models.Entry
		summaries []*models.Summary
		pers      = c.DefaultPers
	)

	qs := c.SetNameKana(c.SetBracup(c.SetBlood(c.SetPrefixLines(diva.StarringDivas().RelatedSel(), ""), ""), ""))
	c.SetImage(qs, "videos__video__entry__").Limit(4).All(&divas)
	c.SetImage(c.SetNameKana(c.SetPrefixLines(anime.StarringAnimes().RelatedSel(), "")), "pictures__entry__").Limit(4).All(&animes)

	c.SetImage(c.SetAdvancedSearch(models.Entries().RelatedSel(), ""), "").Limit(pers).All(&entries)
	c.SetImage(c.SetAdvancedSearch(models.Summaries().RelatedSel(), "entry__").RelatedSel(), "entry__").Limit(pers).All(&summaries)

	c.Data["Divas"] = divas
	c.Data["Animes"] = animes
	c.Data["Entries"] = entries
	c.Data["Summaries"] = summaries
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
		c.Ctx.Abort(404, "404 NotFound")
		return
	}

	uid, _ := convert.StrTo(id).Int64()
	s := &models.Entry{Id: uid}
	s.Read()

	if !s.IsLiving() {
		c.Ctx.Abort(404, "404 NotFound")
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

func (c *EntriesController) Search() {
	c.TplNames = "public/entries/search.tpl"
}
