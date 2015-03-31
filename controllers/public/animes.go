package public

import (
	"time"

	"bitbucket.org/ikeikeikeike/antenna/models"
	"bitbucket.org/ikeikeikeike/antenna/models/character"
	"github.com/astaxie/beego/utils/pagination"
)

type AnimesController struct {
	BaseController
}

func (c *AnimesController) NestFinish() {
	c.PushInAccessLog()
}

func (c *AnimesController) Index() {
	c.TplNames = "public/animes/index.tpl"

	pers := c.DefaultPers
	qs := models.Animes().RelatedSel()
	qs = c.SetPrefixLines(qs.Filter("pictures_count__gt", 0), "")

	cnt, _ := models.CountObjects(qs)
	pager := pagination.SetPaginator(c.Ctx, pers, cnt)

	qs = qs.OrderBy("-pictures_count").Limit(pers, pager.Offset())

	var animes []*models.Anime
	models.ListObjects(qs, &animes)

	c.Data["Animes"] = animes
}

func (c *AnimesController) Show() {
	c.TplNames = "public/animes/show.tpl"

	name := c.Ctx.Input.Param(":name")
	if name == "" {
		c.Ctx.Abort(404, "404 NotFound")
		return
	}

	s := &models.Anime{Name: name}
	s.Read("Name")

	if s.Id <= 0 {
		c.Ctx.Abort(404, "404 NotFound")
		return
	}

	// Update raw html
	if c.IsAjax() && c.Ctx.Input.IsPost() {
		s.Html = c.GetString("data")
		s.HtmlExpire = time.Now()
		s.Update("Html", "HtmlExpire", "Updated")
		c.ServeJson()
		return
	}

	s.LoadRelated()

	pers := c.DefaultPers
	qs := models.Pictures()
	qs = qs.Filter("anime__name", name)
	qs = qs.Filter("entry__isnull", false)

	cnt, _ := models.CountObjects(qs)
	pager := pagination.SetPaginator(c.Ctx, pers, cnt)
	qs = qs.Limit(pers, pager.Offset()).RelatedSel()

	var pictures []*models.Picture
	models.ListObjects(qs, &pictures)
	c.Data["Pictures"] = pictures

	qs = character.StarringCharacters().RelatedSel()
	qs = qs.Filter("anime__name", name)

	var clist []*models.Character
	models.ListObjects(qs, &clist)

	c.Data["Anime"] = s
	c.Data["Characters"] = clist
	c.Data["xsrftoken"] = c.XsrfToken()
	expire := s.HtmlExpire.Unix() + (60 * 60 * 24 * 30) // Raw html update expire: 30 day
	c.Data["doCache"] = expire < time.Now().Unix()
}
