package book

import (
	"bitbucket.org/ikeikeikeike/antenna/models"

	"github.com/astaxie/beego/utils/pagination"
	// "github.com/k0kubun/pp"
)

type TagsController struct {
	BaseController
}

func (c *TagsController) Index() {}

func (c *TagsController) Show() {
	c.TplNames = "book/tags/show.tpl"

	name := c.Ctx.Input.Param(":name")
	if name == "" {
		c.Abort("404")
		return
	}

	s := &models.Tag{Name: name}
	s.Read("Name")

	if s == nil || s.Id <= 0 {
		c.Abort("404")
		return
	}

	s.LoadRelated()
	c.Data["Tag"] = s

	pers := c.DefaultPers
	qs := models.Entries()
	qs = c.SetQ(qs.Filter("tags__tag__name", name), "")

	cnt, _ := models.CountObjects(qs)
	pager := pagination.SetPaginator(c.Ctx, pers, cnt)

	qs = qs.Limit(pers, pager.Offset()).RelatedSel()

	var entries []*models.Entry
	models.ListObjects(qs, &entries)

	for _, s := range entries {
		s.LoadRelated()
	}

	c.Data["Entries"] = entries
}
