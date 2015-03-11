package book

import ctl "bitbucket.org/ikeikeikeike/antenna/controllers"

type BaseController struct {
	ctl.PublicController
}

func (c *BaseController) Prepare() {
	c.PublicController.Prepare()

	c.Layout = "base/book/base.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["BaseHeader"] = "base/book/header.tpl"
	c.LayoutSections["BaseFooter"] = "base/book/footer.tpl"
	// c.LayoutSections["Sidebar"] = ""
}
