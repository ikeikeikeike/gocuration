package admin

type LicensesController struct {
	BaseController
}

// func (c *LicensesController) NestPrepare() {}
// func (c *LicensesController) NestFinish() { }

func (c *LicensesController) Terms() {
	c.TplNames = "admin/licenses/terms.tpl"
}

func (c *LicensesController) Privacy() {
	c.TplNames = "admin/licenses/privacy.tpl"
}

func (c *LicensesController) Guideline() {
	c.TplNames = "admin/licenses/guideline.tpl"
}
