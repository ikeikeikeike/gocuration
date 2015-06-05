package public

import (
	"time"

	"bitbucket.org/ikeikeikeike/antenna/models"
	"github.com/astaxie/beego"
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

	var day time.Time
	if beego.AppConfig.String("runmode") == "prod" {
		day = now.BeginningOfDay()
	} else {
		day = now.BeginningOfDay().Add(-time.Duration(15) * time.Hour)
	}

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("begin_name", "dayly")
	qs = qs.Filter("begin_time", day)
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["QURL"] = ""
	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Dayly() {
	c.TplNames = "public/rankings/index.tpl"

	var day time.Time
	if beego.AppConfig.String("runmode") == "prod" {
		day = now.BeginningOfDay()
	} else {
		day = now.BeginningOfDay().Add(-time.Duration(15) * time.Hour)
	}

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("begin_name", "dayly")
	qs = qs.Filter("begin_time", day)
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["QURL"] = ""
	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Weekly() {
	c.TplNames = "public/rankings/index.tpl"

	var day time.Time
	if beego.AppConfig.String("runmode") == "prod" {
		day = now.BeginningOfDay()
	} else {
		day = now.BeginningOfDay().Add(-time.Duration(15) * time.Hour)
	}

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("begin_name", "dayly")
	qs = qs.Filter("begin_time", day)
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["QURL"] = ""
	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Monthly() {
	c.TplNames = "public/rankings/index.tpl"

	var day time.Time
	if beego.AppConfig.String("runmode") == "prod" {
		day = now.BeginningOfDay()
	} else {
		day = now.BeginningOfDay().Add(-time.Duration(15) * time.Hour)
	}

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("begin_name", "dayly")
	qs = qs.Filter("begin_time", day)
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["QURL"] = ""
	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Yearly() {
	c.TplNames = "public/rankings/index.tpl"

	var day time.Time
	if beego.AppConfig.String("runmode") == "prod" {
		day = now.BeginningOfDay()
	} else {
		day = now.BeginningOfDay().Add(-time.Duration(15) * time.Hour)
	}

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("begin_name", "dayly")
	qs = qs.Filter("begin_time", day)
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["QURL"] = ""
	c.Data["Rankings"] = rankings
}
