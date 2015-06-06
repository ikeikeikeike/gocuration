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
	qs = qs.Filter("begin_name", "dayly")
	qs = qs.Filter("begin_time", c.getParamatedNow().BeginningOfDay())
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Weekly() {
	c.TplNames = "public/rankings/weekly.tpl"

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("begin_name", "weekly")
	qs = qs.Filter("begin_time", c.getParamatedNow().BeginningOfWeek())
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Monthly() {
	c.TplNames = "public/rankings/monthly.tpl"

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("begin_name", "monthly")
	qs = qs.Filter("begin_time", c.getParamatedNow().BeginningOfMonth())
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Yearly() {
	c.TplNames = "public/rankings/yearly.tpl"

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("begin_name", "yearly")
	qs = qs.Filter("begin_time", c.getParamatedNow().BeginningOfYear())
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["Rankings"] = rankings
}

func (c *RankingsController) getParamatedNow() *now.Now {
	t := time.Now().UTC()
	date := now.New(t)

	if t, err := date.Parse(c.GetString("date")); err == nil {
		date = now.New(t)
	} else {
		c.Data["Params"].(map[string]string)["date"] = date.Format("2006-01-02")
	}
	return date
}
