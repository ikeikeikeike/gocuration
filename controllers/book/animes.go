package book

import (
	"time"

	"bitbucket.org/ikeikeikeike/antenna/models"
	"bitbucket.org/ikeikeikeike/antenna/models/character"
	"bitbucket.org/ikeikeikeike/antenna/ormapper"
	"bitbucket.org/ikeikeikeike/antenna/ormapper/anime"
	"bitbucket.org/ikeikeikeike/antenna/ormapper/blog"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/ikeikeikeike/gopkg/convert"
)

type AnimesController struct {
	BaseController
}

func (c *AnimesController) NestFinish() {
	c.PushInAccessLog()
}

func (c *AnimesController) Index() {
	c.TplNames = "book/animes/index.tpl"

	db := ormapper.PictureAnimations().
		Scopes(blog.FilterMediatype("image")).
		Scopes(anime.PictureCountMoreThanZero).
		Scopes(anime.FilterPrefixLines(c.GetString("line"))).
		Scopes(anime.FilterNameKana(convert.StrTo(c.GetString("q")).MultiWord())).
		Limit(c.DefaultPers)

	var count int64
	db.Count(&count)

	pager := pagination.SetPaginator(c.Ctx, c.DefaultPers, count)
	db = db.Limit(c.DefaultPers).Offset(pager.Offset())

	var animes []*ormapper.Anime
	db.Order("anime.pictures_count DESC").Find(&animes)

	c.Data["Animes"] = animes
}

func (c *AnimesController) Show() {
	c.TplNames = "book/animes/show.tpl"

	name := c.Ctx.Input.Param(":name")
	if name == "" {
		c.Ctx.Abort(404, "404")
		return
	}

	s := &models.Anime{Name: name}
	s.Read("Name")

	if s.Id <= 0 {
		c.Ctx.Abort(404, "404")
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
