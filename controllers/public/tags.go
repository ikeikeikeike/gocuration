package public

import (
	"bitbucket.org/ikeikeikeike/antenna/models"

	"github.com/astaxie/beego/utils/pagination"
	// "github.com/k0kubun/pp"
)

type TagsController struct {
	BaseController
}

func (c *TagsController) Index() {
	c.TplNames = "public/tags/index.tpl"

	pers := c.DefaultPers
	qs := models.Tags()

	cnt, _ := models.CountObjects(qs)
	pager := pagination.SetPaginator(c.Ctx, pers, cnt)

	qs = qs.OrderBy("-Created").Limit(pers, pager.Offset()).RelatedSel()
	//   qs := orm.QueryRelated(post,"Tag")
	//  qs.All(&[]*Tag{})

	var tags []*models.Tag
	models.ListObjects(qs, &tags)

	for _, s := range tags {
		s.LoadRelated()
	}

	c.Data["Tags"] = tags
}

func (c *TagsController) Show() {
	c.TplNames = "public/tags/show.tpl"

	name := c.Ctx.Input.Param(":name")
	if name == "" {
		c.Ctx.WriteString("tags does not exists")
		return
	}

	s := &models.Tag{Name: name}
	s.Read("Name")
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
