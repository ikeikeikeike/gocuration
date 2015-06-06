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

func (c *RankingsController) Index() {
	c.TplNames = "public/rankings/index.tpl"

	t := time.Now().UTC()
	t = now.New(t).BeginningOfDay()

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("begin_name", "dayly")
	qs = qs.Filter("begin_time", t)
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Dayly() {
	c.TplNames = "public/rankings/index.tpl"

	t := time.Now().UTC()
	t = now.New(t).BeginningOfDay()

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("begin_name", "dayly")
	qs = qs.Filter("begin_time", t)
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Weekly() {
	c.TplNames = "public/rankings/index.tpl"

	t := time.Now().UTC()
	t = now.New(t).BeginningOfWeek()

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("begin_name", "weekly")
	qs = qs.Filter("begin_time", t)
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Monthly() {
	c.TplNames = "public/rankings/index.tpl"

	t := time.Now().UTC()
	t = now.New(t).BeginningOfMonth()

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("begin_name", "monthly")
	qs = qs.Filter("begin_time", t)
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Yearly() {
	c.TplNames = "public/rankings/index.tpl"

	t := time.Now().UTC()
	t = now.New(t).BeginningOfYear()

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("begin_name", "yearly")
	qs = qs.Filter("begin_time", t)
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["Rankings"] = rankings
}
