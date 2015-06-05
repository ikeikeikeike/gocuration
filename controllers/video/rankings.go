package video

// "github.com/k0kubun/pp"

type RankingsController struct {
	BaseController
}

func (c *RankingsController) NestFinish() {
	c.PushInAccessLog()
}

func (c *RankingsController) Index() {
	c.TplNames = "video/rankings/index.tpl"
}
