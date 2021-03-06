package book

import (
	"bitbucket.org/ikeikeikeike/antenna/models"

	// "github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/ikeikeikeike/gopkg/convert"
	// "github.com/k0kubun/pp"

	// "github.com/k0kubun/pp"
)

type BlogsController struct {
	BaseController
}

func (c *BlogsController) NestFinish() {
	c.PushInAccessLog()
}

func (c *BlogsController) Show() {
	c.TplNames = "public/blogs/show.tpl"

	id := c.Ctx.Input.Param(":id")
	if id == "" {
		c.Ctx.WriteString("entry does not exists")
		return
	}

	uid, _ := convert.StrTo(id).Int64()
	blog := &models.Blog{Id: uid}
	blog.Read()
	blog.LoadRelated()
	c.Data["Blog"] = blog

	pers := c.DefaultPers
	qs := models.Entries().RelatedSel()
	qs = c.SetQ(qs.Filter("blog_id", uid), "")
	qs = qs.Filter("blog__mediatype", "image")

	cnt, _ := models.CountObjects(qs)
	pager := pagination.SetPaginator(c.Ctx, pers, cnt)

	qs = qs.Limit(pers, pager.Offset())

	var entries []*models.Entry
	models.ListObjects(qs, &entries)

	for _, s := range entries {
		s.LoadRelated()
	}

	c.Data["Entries"] = entries
}
