package public

type RegistersController struct {
	BaseController
}

func (c *RegistersController) Index() {
	c.TplNames = "public/registers/index.tpl"
}

func (c *RegistersController) Parts() {
	c.TplNames = "public/registers/parts.tpl"
}

func (c *RegistersController) Links() {
	c.TplNames = "public/registers/links.tpl"
}

func (c *RegistersController) Rss() {
	c.TplNames = "public/registers/rss.tpl"
}
