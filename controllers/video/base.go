package video

import ctl "bitbucket.org/ikeikeikeike/antenna/controllers"

type BaseController struct {
	ctl.PublicController
}

func (c *BaseController) Prepare() {
	c.PublicController.Prepare()

	c.Layout = "base/video/base.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["BaseHeader"] = "base/video/header.tpl"
	c.LayoutSections["BaseFooter"] = "base/video/footer.tpl"
	// c.LayoutSections["Sidebar"] = ""
}
