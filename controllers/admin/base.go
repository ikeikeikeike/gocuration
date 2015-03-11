package admin

import (
	"bitbucket.org/ikeikeikeike/antenna/controllers"
	"bitbucket.org/ikeikeikeike/antenna/models"
)

type BaseController struct {
	controllers.BaseController

	Userinfo *models.User
	IsLogin  bool
}

func (c *BaseController) Prepare() {

	c.IsLogin = c.GetSession("userinfo") != nil
	if c.IsLogin {
		c.Userinfo = c.GetLogin()
	}

	c.Data["IsLogin"] = c.IsLogin
	c.Data["Userinfo"] = c.Userinfo

	c.Layout = "base/admin/base.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["BaseHeader"] = "base/admin/header.tpl"
	c.LayoutSections["BaseFooter"] = "base/admin/footer.tpl"

	c.BaseController.Prepare()
}

func (c *BaseController) GetLogin() *models.User {
	u := &models.User{Id: c.GetSession("userinfo").(int64)}
	u.Read()
	u.LoadRelated()
	return u
}

func (c *BaseController) DelLogin() {
	c.DelSession("userinfo")
}

func (c *BaseController) SetLogin(user *models.User) {
	c.SetSession("userinfo", user.Id)
}

func (c *BaseController) LoginPath() string {
	return c.UrlFor("LoginController.Login")
}
