package public

import ctl "bitbucket.org/ikeikeikeike/antenna/controllers"

type BaseController struct {
	ctl.PublicController
}

func (c *BaseController) Prepare() {
	c.PublicController.Prepare()

	c.Layout = "base/public/base.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["BaseHeader"] = "base/public/header.tpl"
	c.LayoutSections["BaseFooter"] = "base/public/footer.tpl"
	// c.LayoutSections["Sidebar"] = ""
}
