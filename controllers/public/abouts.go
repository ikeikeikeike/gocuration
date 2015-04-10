package public

type AboutsController struct {
	BaseController
}

func (c *AboutsController) Index() {
	c.TplNames = "public/abouts/index.tpl"
}
