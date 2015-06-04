package public

import (
	"bitbucket.org/ikeikeikeike/antenna/models"

	"github.com/astaxie/beego/utils/pagination"
	// "github.com/k0kubun/pp"
)

type RankingsController struct {
	BaseController
}

// func (c *RankingsController) NestPrepare() {}

func (c *RankingsController) NestFinish() {
	c.PushInAccessLog()
}

func (c *RankingsController) Index() {
	c.TplNames = "public/rankings/index.tpl"

	pers := c.DefaultPers
	qs := models.EntryRankings().RelatedSel()

	cnt, _ := models.CountObjects(qs)
	pager := pagination.SetPaginator(c.Ctx, pers, cnt)

	qs = qs.Limit(pers, pager.Offset())

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["QURL"] = ""
	c.Data["Rankings"] = rankings
}
