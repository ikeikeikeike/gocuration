package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/jinzhu/now"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	"bitbucket.org/ikeikeikeike/antenna/models/ranking"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestRanking(t *testing.T) {
	o := orm.NewOrm()
	cond := orm.NewCondition()

	jst := time.Duration(9) * time.Hour
	day := now.BeginningOfDay().Add(jst)
	week := now.BeginningOfWeek().Add(jst)
	month := now.BeginningOfMonth().Add(jst)
	year := now.BeginningOfYear().Add(jst)

	dayly := cond.And("begin_time", day).And("begin_name", "dayly")
	weekly := cond.And("begin_time", week).And("begin_name", "weekly")
	monthly := cond.And("begin_time", month).And("begin_name", "monthly")
	yearly := cond.And("begin_time", year).And("begin_name", "yearly")

	var id int64 = 47649
	path := "/unko"
	docount := orm.Params{"page_view": orm.ColValue(orm.Col_Add, 1)}

	var err error

	_, _, err = ranking.ReadOrCreateEntry(id, "dayly", day)
	check(t, err, "dayly create", path)
	_, err = o.QueryTable("entry_ranking").Filter("entry", id).SetCond(dayly).Update(docount)
	check(t, err, "dayly", path)

	_, _, err = ranking.ReadOrCreateEntry(id, "weekly", week)
	check(t, err, "weekly create", path)
	_, err = o.QueryTable("entry_ranking").Filter("entry", id).SetCond(weekly).Update(docount)
	check(t, err, "weekly", path)

	_, _, err = ranking.ReadOrCreateEntry(id, "monthly", month)
	check(t, err, "weekly create", path)
	_, err = o.QueryTable("entry_ranking").Filter("entry", id).SetCond(monthly).Update(docount)
	check(t, err, "monthly", path)

	_, _, err = ranking.ReadOrCreateEntry(id, "yearly", year)
	check(t, err, "yearly create", path)
	_, err = o.QueryTable("entry_ranking").Filter("entry", id).SetCond(yearly).Update(docount)
	check(t, err, "yearly", path)
}

func check(t *testing.T, err error, args ...interface{}) {
	if err != nil {
		t.Error(fmt.Sprintf("Update PageView by %v: ", args), err)
	}
}
