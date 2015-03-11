package admin

import (
	"html/template"

	"bitbucket.org/ikeikeikeike/antenna/lib/models/user/changemail"
	"bitbucket.org/ikeikeikeike/antenna/lib/verify"
	"bitbucket.org/ikeikeikeike/antenna/models"
	"github.com/astaxie/beego"
	"github.com/ikeikeikeike/gopkg/convert"
	"github.com/ikeikeikeike/gopkg/str"
)

type UsersController struct {
	BaseController
}

func (c *UsersController) NestPrepare() {
	if !c.IsLogin {
		c.Ctx.Redirect(302, c.LoginPath())
		return
	}
}

// func (c *UsersController) NestFinish() {}

func (c *UsersController) Index() {
	beego.ReadFromRequest(&c.Controller)

	c.Data["Path"] = c.UrlFor("UsersController.Index")
	c.TplNames = "admin/users/dashboard.tpl"
}

func (c *UsersController) ChangeMail() {
	beego.ReadFromRequest(&c.Controller)

	c.Data["Path"] = c.UrlFor("UsersController.ChangeMail")
	c.Data["xsrfdata"] = template.HTML(c.XsrfFormHtml())
	c.TplNames = "admin/users/changemail.tpl"

	if !c.Ctx.Input.IsPost() {
		return
	}

	flash := beego.NewFlash()
	email := c.GetString("Email")

	tu, err := changemail.CreateTmpuser(email)
	if err != nil || tu.Id < 1 {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		return
	}

	// Send activate mail
	err = changemail.SendChangeMail(tu, c.mailConfirmUrl(tu.Token))
	if err != nil {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		return
	}

	flash.Success(`新しいメールアドレスの確認のため、
	メッセージを送信しました。`)
	flash.Store(&c.Controller)
}

func (c *UsersController) MailConfirm() {
	token := c.GetString("token")

	tu, err := changemail.ReceiveChangeMail(token)
	if err != nil {
		c.Ctx.Abort(403, err.Error())
		return
	}

	user, err := changemail.RegisterMail(tu, c.Userinfo)
	if err != nil || user.Id < 1 {
		c.Ctx.Abort(403, err.Error())
		return
	}

	flash := beego.NewFlash()
	flash.Success("メールアドレスを更新しました。")
	flash.Store(&c.Controller)

	c.Ctx.Redirect(302, c.UrlFor("UsersController.ChangeMail"))
}

func (c *UsersController) Blogs() {
	beego.ReadFromRequest(&c.Controller)

	c.Data["Path"] = c.UrlFor("UsersController.Blogs")
	c.TplNames = "admin/users/blogs.tpl"
}

func (c *UsersController) Blog() {
	id := c.Ctx.Input.Param(":id")
	if id == "" {
		c.Ctx.Abort(403, "403 Forbidden")
		return
	}
	uid, _ := convert.StrTo(id).Int64()
	for _, b := range c.Userinfo.Blogs {
		if b.Id == uid {
			c.Data["Blog"] = b
			break
		}
	}
	b, ok := c.Data["Blog"]
	if !ok {
		c.Ctx.Abort(403, "403 Forbidden")
		return
	}

	beego.ReadFromRequest(&c.Controller)

	c.Data["Path"] = c.UrlFor("UsersController.Blog")
	c.Data["xsrfdata"] = template.HTML(c.XsrfFormHtml())
	c.TplNames = "admin/users/blog.tpl"

	if !c.Ctx.Input.IsPost() {
		return
	}

	var err error
	flash := beego.NewFlash()

	// validation
	if err = c.ParseForm(b); err != nil {
		flash.Error("Blog update invalid!")
		flash.Store(&c.Controller)
		return
	}
	if err = models.IsValid(b); err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		return
	}

	blog := b.(*models.Blog)

	if !verify.IsRss(blog.Rss) {
		flash.Error("RSSの形式が正しくありません。")
		flash.Store(&c.Controller)
		return
	}

	// update
	if err = blog.Update(); err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		return
	}

	flash.Success("サイトを更新しました")
	flash.Store(&c.Controller)
}

func (c *UsersController) BlogRegister() {
	beego.ReadFromRequest(&c.Controller)
	b := new(models.Blog)

	c.Data["Blog"] = b
	c.Data["Path"] = c.UrlFor("UsersController.BlogRegister")
	c.Data["xsrfdata"] = template.HTML(c.XsrfFormHtml())
	c.TplNames = "admin/users/blog.tpl"

	if !c.Ctx.Input.IsPost() {
		return
	}

	var err error
	flash := beego.NewFlash()

	// validation
	if err = c.ParseForm(b); err != nil {
		flash.Error("Blog register invalid!")
		flash.Store(&c.Controller)
		return
	}
	if err = models.IsValid(b); err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		return
	}

	rss := str.Clean(b.Rss)

	feed, ok := verify.GetFeed(rss)
	if !ok {
		flash.Error("RSSの形式が正しくありません。")
		flash.Store(&c.Controller)
		return
	}

	// Insert
	if len(feed.Channels) > 0 {
		b.Url = str.Clean(feed.Channels[0].Links[0].Href)
	}
	b.Rss = rss
	b.User = c.Userinfo

	if err = b.Insert(); err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		return
	}

	flash.Success("新しくサイトを登録しました")
	flash.Store(&c.Controller)

	c.Ctx.Redirect(302, c.UrlFor("UsersController.Blog", ":id", convert.ToStr(b.Id)))
}

func (c *UsersController) AccountRemove() {
	for _, b := range c.Userinfo.Blogs {
		b.LogicalDelete()
		// b.Delete()
	}
	c.Userinfo.Delete()
	c.DelLogin()
	c.Ctx.Redirect(302, c.UrlFor("LoginController.Login"))
}

func (c *UsersController) BlogRemove() {

	id := c.GetString("id")
	if id == "" {
		c.Ctx.Abort(403, "403 Forbidden")
		return
	}

	var err error

	uid, _ := convert.StrTo(id).Int64()
	for _, b := range c.Userinfo.Blogs {
		if b.Id == uid {
			flash := beego.NewFlash()

			// b.Delete()
			if err = b.LogicalDelete(); err != nil {
				flash.Error(err.Error())
				flash.Store(&c.Controller)
			} else {
				flash.Success("サイトを削除しました")
				flash.Store(&c.Controller)
			}

			c.Ctx.Redirect(302, c.UrlFor("UsersController.Blogs"))
			return
		}
	}

	c.Ctx.Abort(403, "403 Forbidden")
}

func (c *UsersController) mailConfirmUrl(token string) string {
	return c.BuildRequestUrl(c.UrlFor("UsersController.MailConfirm", "token", token))
}
