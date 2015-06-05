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

	var t time.Time
	if beego.AppConfig.String("runmode") == "prod" {
		t = now.BeginningOfDay()
	} else {
		t = now.BeginningOfDay().Add(-time.Duration(15) * time.Hour)
	}

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("begin_name", "dayly")
	qs = qs.Filter("begin_time", t)
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["QURL"] = ""
	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Dayly() {
	c.TplNames = "public/rankings/index.tpl"

	var t time.Time
	if beego.AppConfig.String("runmode") == "prod" {
		t = now.BeginningOfDay()
	} else {
		t = now.BeginningOfDay().Add(-time.Duration(15) * time.Hour)
	}

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("begin_name", "dayly")
	qs = qs.Filter("begin_time", t)
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["QURL"] = ""
	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Weekly() {
	c.TplNames = "public/rankings/index.tpl"

	var t time.Time
	if beego.AppConfig.String("runmode") == "prod" {
		t = now.BeginningOfWeek()
	} else {
		t = now.BeginningOfWeek().Add(time.Duration(9) * time.Hour)
	}

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("begin_name", "weekly")
	qs = qs.Filter("begin_time", t)
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["QURL"] = ""
	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Monthly() {
	c.TplNames = "public/rankings/index.tpl"

	var t time.Time
	if beego.AppConfig.String("runmode") == "prod" {
		t = now.BeginningOfMonth()
	} else {
		t = now.BeginningOfMonth().Add(time.Duration(9) * time.Hour)
	}

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("begin_name", "monthly")
	qs = qs.Filter("begin_time", t)
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["QURL"] = ""
	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Yearly() {
	c.TplNames = "public/rankings/index.tpl"

	var t time.Time
	if beego.AppConfig.String("runmode") == "prod" {
		t = now.BeginningOfYear()
	} else {
		t = now.BeginningOfYear().Add(time.Duration(9) * time.Hour)
	}

	qs := models.EntryRankings().RelatedSel()
	qs = qs.Filter("begin_name", "yearly")
	qs = qs.Filter("begin_time", t)
	qs = qs.Limit(100, 0)

	var rankings []*models.EntryRanking
	models.ListObjects(qs, &rankings)

	c.Data["QURL"] = ""
	c.Data["Rankings"] = rankings
}
