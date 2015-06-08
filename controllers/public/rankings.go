package public

import (
	"time"

	"bitbucket.org/ikeikeikeike/antenna/models"
	"github.com/jinzhu/now"
)

// "github.com/k0kubun/pp"

type RankingsController struct {
	BaseController
}

// func (c *RankingsController) NestPrepare() {}

func (c *RankingsController) NestFinish() {
	c.PushInAccessLog()
}

func (c *RankingsController) Dayly() {
	c.TplNames = "public/rankings/dayly.tpl"

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("rank__gt", 0)
	qs = qs.Filter("begin_name", "dayly")
	qs = qs.Filter("begin_time", c.GetParamatedNow().BeginningOfDay())
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Weekly() {
	c.TplNames = "public/rankings/weekly.tpl"

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("rank__gt", 0)
	qs = qs.Filter("begin_name", "weekly")
	qs = qs.Filter("begin_time", c.GetParamatedNow().BeginningOfWeek())
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Monthly() {
	c.TplNames = "public/rankings/monthly.tpl"

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("rank__gt", 0)
	qs = qs.Filter("begin_name", "monthly")
	qs = qs.Filter("begin_time", c.GetParamatedNow().BeginningOfMonth())
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Yearly() {
	c.TplNames = "public/rankings/yearly.tpl"

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("rank__gt", 0)
	qs = qs.Filter("begin_name", "yearly")
	qs = qs.Filter("begin_time", c.GetParamatedNow().BeginningOfYear())
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["Rankings"] = rankings
}
