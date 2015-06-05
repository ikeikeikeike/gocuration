package main

import (
	"testing"

	"github.com/astaxie/beego/orm"
	"github.com/jinzhu/now"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	"bitbucket.org/ikeikeikeike/antenna/models"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestQueryOutput(t *testing.T) {
	day := now.BeginningOfDay()

	cond := orm.NewCondition()
	dayly := cond.And("begin_time", day).And("begin_name", "dayly")

	var list []*models.VideoRanking
	o := orm.NewOrm()
	_, _ = o.QueryTable("video_ranking").SetCond(dayly).Filter("video__entry__id", 41402).All(&list)
}
