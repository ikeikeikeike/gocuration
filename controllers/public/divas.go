package public

import (
	"time"

	"bitbucket.org/ikeikeikeike/antenna/models"

	"github.com/astaxie/beego/utils/pagination"
	// "github.com/k0kubun/pp"
)

type DivasController struct {
	BaseController
}

func (c *DivasController) NestFinish() {
	c.PushInAccessLog()
}

func (c *DivasController) Index() {
	c.TplNames = "public/divas/index.tpl"

	pers := c.DefaultPers
	qs := models.Divas().RelatedSel()
	qs = c.SetBracupLines(c.SetPrefixLines(qs.Filter("videos_count__gt", 0), ""))

	cnt, _ := models.CountObjects(qs)
	pager := pagination.SetPaginator(c.Ctx, pers, cnt)

	qs = qs.OrderBy("-videos_count").Limit(pers, pager.Offset())

	var divas []*models.Diva
	models.ListObjects(qs, &divas)

	c.Data["Divas"] = divas
}

func (c *DivasController) Show() {
	c.TplNames = "public/divas/show.tpl"

	name := c.Ctx.Input.Param(":name")
	if name == "" {
		c.Ctx.Abort(404, "404 NotFound")
		return
	}

	s := &models.Diva{Name: name}
	s.Read("Name")

	if s.Id <= 0 {
		c.Ctx.Abort(404, "404 NotFound")
		return
	}

	// Update raw html
	if c.IsAjax() && c.Ctx.Input.IsPost() {
		s.Html = c.GetString("data")
		s.Update("Html", "Updated")
		c.ServeJson()
		return
	}
	// Raw html update expire: 5 day
	if (s.Updated.Unix() + (60 * 60 * 24 * 5)) < time.Now().Unix() {
		s.Html = ""
	}
	c.Data["xsrftoken"] = c.XsrfToken()

	s.LoadRelated()
	c.Data["Diva"] = s

	pers := c.DefaultPers
	qs := models.Videos()
	qs = qs.Filter("divas__diva__name", name)
	qs = qs.Filter("entry__isnull", false)

	cnt, _ := models.CountObjects(qs)
	pager := pagination.SetPaginator(c.Ctx, pers, cnt)

	qs = qs.Limit(pers, pager.Offset()).RelatedSel()

	var videos []*models.Video
	models.ListObjects(qs, &videos)

	c.Data["Videos"] = videos
}
