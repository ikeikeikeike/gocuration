package admin

import (
	"html/template"

	"bitbucket.org/ikeikeikeike/antenna/lib/models/user/authc"
	"bitbucket.org/ikeikeikeike/antenna/lib/models/user/reset"
	"bitbucket.org/ikeikeikeike/antenna/lib/models/user/signup"
	"bitbucket.org/ikeikeikeike/antenna/lib/verify"
	"bitbucket.org/ikeikeikeike/antenna/models"

	"github.com/astaxie/beego"
)

type LoginController struct {
	BaseController
}

// func (c *LoginController) NestPrepare() {}
// func (c *LoginController) NestFinish() { }

func (c *LoginController) Login() {
	success := c.UrlFor("UsersController.Index")

	if c.IsLogin {
		c.Ctx.Redirect(302, success)
		return
	}

	c.TplNames = "admin/login/login.tpl"
	c.Data["xsrfdata"] = template.HTML(c.XsrfFormHtml())

	if !c.Ctx.Input.IsPost() {
		return
	}

	flash := beego.NewFlash()
	email := c.GetString("Email")
	password := c.GetString("Password")

	// Get authenticated user
	user, err := authc.Authenticate(email, password)
	if err != nil || user.Id < 1 {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		return
	}

	flash.Success("Success logged in")
	flash.Store(&c.Controller)

	// set user session
	c.SetLogin(user)

	c.Redirect(success, 303)
}

func (c *LoginController) Logout() {
	c.DelLogin()
	c.Ctx.Redirect(302, c.UrlFor("LoginController.Login"))
}

func (c *LoginController) PasswordReset() {
	c.TplNames = "admin/login/reset.tpl"
	c.Data["xsrfdata"] = template.HTML(c.XsrfFormHtml())

	if !c.Ctx.Input.IsPost() {
		return
	}

	flash := beego.NewFlash()
	email := c.GetString("Email")

	tu, err := reset.ResetPassword(email)
	if err != nil || tu.Id < 1 {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		return
	}

	// Send activate mail
	err = reset.SendResetMail(tu, c.resetUrl(tu.Token))
	if err != nil {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		return
	}

	msg := "パスワード再発行のためにメールアドレスをご確認ください。"
	msg += "%sに確認メールを送信しました。"
	flash.Success(msg, tu.Email)
	flash.Store(&c.Controller)
}

func (c *LoginController) ResetConfirm() {
	token := c.GetString("token")

	tu, err := reset.ReceiveResetMail(token)
	if err != nil {
		c.Ctx.Abort(403, err.Error())
		return
	}

	c.TplNames = "admin/login/resetconfirm.tpl"
	c.Data["xsrfdata"] = template.HTML(c.XsrfFormHtml())
	c.Data["token"] = token

	if !c.Ctx.Input.IsPost() {
		return
	}

	flash := beego.NewFlash()
	r := &reset.ResetForm{}

	if err = c.ParseForm(r); err != nil {
		flash.Error("Reset invalid!")
		flash.Store(&c.Controller)
		return
	}
	if err = models.IsValid(r); err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		return
	}

	tu.Password = r.Password
	user, err := reset.RegisterPassword(tu)
	if err != nil || user.Id < 1 {
		c.Ctx.Abort(403, err.Error())
		return
	}

	flash.Success("Success register password")
	flash.Store(&c.Controller)

	// set user session
	c.SetLogin(user)

	c.Ctx.Redirect(302, c.UrlFor("UsersController.Index"))
}

func (c *LoginController) Signup() {
	c.TplNames = "admin/login/signup.tpl"
	c.Data["xsrfdata"] = template.HTML(c.XsrfFormHtml())

	if !c.Ctx.Input.IsPost() {
		return
	}

	var err error
	flash := beego.NewFlash()

	tu := &models.Tmpuser{}
	if err = c.ParseForm(tu); err != nil {
		flash.Error("Signup invalid!")
		flash.Store(&c.Controller)
		return
	}
	if err = models.IsValid(tu); err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		return
	}
	if !verify.IsRss(tu.Rss) {
		flash.Error("RSSの形式が正しくありません。ご確認ください。")
		flash.Store(&c.Controller)
		return
	}

	// issue one timed user
	id, err := signup.CreateTmpuser(tu)
	if err != nil || id < 1 {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		return
	}

	// Send activate mail
	err = signup.SendSignupMail(tu, c.confirmUrl(tu.Token))
	if err != nil {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		return
	}

	msg := "アカウント作成のためにメールアドレスをご確認ください。"
	msg += "%sに確認メールを送信しました。"
	flash.Success(msg, tu.Email)
	flash.Store(&c.Controller)
}

func (c *LoginController) SignupConfirm() {
	token := c.GetString("token")

	tu, err := signup.ReceiveSignupMail(token)
	if err != nil {
		c.Ctx.Abort(403, err.Error())
		return
	}

	user, err := signup.RegisterUser(tu)
	if err != nil || user.Id < 1 {
		c.Ctx.Abort(403, err.Error())
		return
	}

	flash := beego.NewFlash()
	flash.Success(`ご登録ありがとうございます。
	ただいま御サイトの確認作業をおこなっております。今しばらくお待ち下さい。
	`)
	flash.Store(&c.Controller)

	// set user session
	c.SetLogin(user)

	c.Ctx.Redirect(302, c.UrlFor("UsersController.Index"))
}

func (c *LoginController) resetUrl(token string) string {
	return c.BuildRequestUrl(c.UrlFor("LoginController.ResetConfirm", "token", token))
}

func (c *LoginController) confirmUrl(token string) string {
	return c.BuildRequestUrl(c.UrlFor("LoginController.SignupConfirm", "token", token))
}
