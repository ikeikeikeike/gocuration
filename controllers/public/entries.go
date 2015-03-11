package public

import (
	"fmt"
	"strings"

	"bitbucket.org/ikeikeikeike/antenna/models"
	"bitbucket.org/ikeikeikeike/antenna/models/anime"
	"bitbucket.org/ikeikeikeike/antenna/models/diva"

	"github.com/astaxie/beego/orm"
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
	c.TplNames = "public/entries/home.tpl"

	var (
		divas     []*models.Diva
		animes    []*models.Anime
		entries   []*models.Entry
		summaries []*models.Summary
		pers      = c.DefaultPers
	)

	dqs := c.SetBracup(c.SetBlood(c.SetPrefixLines(diva.StarringDivas().RelatedSel(), ""), ""), "")
	aqs := c.SetPrefixLines(anime.StarringAnimes().RelatedSel(), "")

	v := c.GetString("q")
	if v != "" {
		for _, word := range convert.StrTo(v).MultiWord() {
			c := orm.NewCondition()
			c = c.Or("name__icontains", word)
			c = c.Or("kana__icontains", word)

			dqs = dqs.SetCond(c)
			aqs = aqs.SetCond(c)
		}
	}
	dqs.Limit(4).All(&divas)
	aqs.Limit(4).All(&animes)

	c.SetAdvancedSearch(models.Entries().RelatedSel(), "").Limit(pers).All(&entries)
	c.SetAdvancedSearch(models.Summaries().RelatedSel(), "entry__").RelatedSel().Limit(pers).All(&summaries)

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
	c.TplNames = "public/entries/hots.tpl"

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
	c.TplNames = "public/entries/show.tpl"

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

	s.Blog.LoadRelated()
	if s.Video != nil {
		s.Video.LoadRelated()
	}
	if s.Picture != nil {
		s.Picture.LoadRelated()
	}

	var (
		in        []string
		summaries []*models.Summary
	)

	for _, t := range s.Tags {
		if t.Name != "" {
			in = append(in, t.Name)
		}
	}
	if len(in) <= 0 {
		in = append(in, "巨乳", "エロ")
	}

	// models.Summaries().RelatedSel().
	// Filter("entry__tags__tag__name__in", in).
	// Limit(15).All(&summaries)
	//
	// 上記を `DISTINCT` 付きでやっている
	names := fmt.Sprintf("'%s'", strings.Join(in, "','"))
	q := fmt.Sprintf(`
	SELECT DISTINCT s.* FROM summary as s 
	LEFT OUTER JOIN entry e ON e.id = s.entry_id 
	LEFT OUTER JOIN blog b ON b.id = e.blog_id 
	LEFT OUTER JOIN entry_tag et ON et.entry_id = e.id 
	LEFT OUTER JOIN tag tag ON tag.id = et.tag_id 
	WHERE (tag.name IN (%s) OR e.q like '%%%s%%') AND e.id != '%d'
	ORDER BY s.sort DESC LIMIT 15`, names, names[0], s.Id)
	orm.NewOrm().Raw(q).QueryRows(&summaries)

	c.Data["Summaries"] = summaries
	c.Data["Entry"] = s
}

func (c *EntriesController) Search() {
	c.TplNames = "public/entries/search.tpl"
}
