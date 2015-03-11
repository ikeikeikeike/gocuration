package public

import (
	"bitbucket.org/ikeikeikeike/antenna/models"
	"github.com/astaxie/beego/utils/pagination"
)

type CharactersController struct {
	BaseController
}

func (c *CharactersController) NestFinish() {
	c.PushInAccessLog()
}

func (c *CharactersController) Index() {
	c.TplNames = "public/characters/index.tpl"

	pers := c.DefaultPers
	qs := models.Characters().RelatedSel()
	// qs = c.SetBracupLines(c.SetPrefixLines(qs.Filter("pictures_count__gt", 0), ""))

	cnt, _ := models.CountObjects(qs)
	pager := pagination.SetPaginator(c.Ctx, pers, cnt)

	qs = qs.OrderBy("-pictures_count").Limit(pers, pager.Offset())

	var characters []*models.Character
	models.ListObjects(qs, &characters)

	c.Data["Characters"] = characters
}

func (c *CharactersController) Show() {
	c.TplNames = "public/characters/show.tpl"

	name := c.Ctx.Input.Param(":name")
	if name == "" {
		c.Ctx.Abort(404, "404 NotFound")
		return
	}

	s := &models.Character{Name: name}
	s.Read("Name")

	if s.Id <= 0 {
		c.Ctx.Abort(404, "404 NotFound")
		return
	}

	s.LoadRelated()
	c.Data["Character"] = s

	pers := c.DefaultPers
	qs := models.Pictures()
	qs = qs.Filter("characters__character__name", name)
	qs = qs.Filter("entry__isnull", false)

	cnt, _ := models.CountObjects(qs)
	pager := pagination.SetPaginator(c.Ctx, pers, cnt)

	qs = qs.Limit(pers, pager.Offset()).RelatedSel()

	var pictures []*models.Picture
	models.ListObjects(qs, &pictures)

	c.Data["Pictures"] = pictures
	c.Data["Animes"] = []*models.Anime{s.Anime}
}
