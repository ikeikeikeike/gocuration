package summarize

import (
	"fmt"
	"reflect"
	"time"

	"bitbucket.org/ikeikeikeike/antenna/lib/accessctl"
	"bitbucket.org/ikeikeikeike/antenna/models"
	"bitbucket.org/ikeikeikeike/antenna/models/ranking"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/jinzhu/now"
)

func Showcounter() (err error) {
	o := orm.NewOrm()
	cond := orm.NewCondition()

	c := accessctl.NewShowCounter()
	c.Bootstrap()

	jst := time.Duration(9) * time.Hour
	day := now.BeginningOfDay().Add(jst)
	week := now.BeginningOfWeek().Add(jst)
	month := now.BeginningOfMonth().Add(jst)
	year := now.BeginningOfYear().Add(jst)

	dayly := cond.And("begin_time", day).And("begin_name", "dayly")
	weekly := cond.And("begin_time", week).And("begin_name", "weekly")
	monthly := cond.And("begin_time", month).And("begin_name", "monthly")
	yearly := cond.And("begin_time", year).And("begin_name", "yearly")

	for _, path := range []string{"elog", "video", "book"} {
		results, err := c.Counting(path)
		if err != nil {
			continue
		}

		for _, r := range results {
			var err error

			docount := orm.Params{"page_view": orm.ColValue(orm.Col_Add, r.Count)}

			switch path {
			case "elog":
				_, err = o.QueryTable("entry").Filter("id", r.ID).Update(docount)
				check(err, path)

				_, _, err = ranking.ReadOrCreateEntry(r.ID, "dayly", day)
				check(err, "dayly create", path)
				_, err = o.QueryTable("entry_ranking").Filter("entry", r.ID).SetCond(dayly).Update(docount)
				check(err, "dayly ranking", path)

				_, _, err = ranking.ReadOrCreateEntry(r.ID, "weekly", week)
				check(err, "weekly create", path)
				_, err = o.QueryTable("entry_ranking").Filter("entry", r.ID).SetCond(weekly).Update(docount)
				check(err, "weekly ranking", path)

				_, _, err = ranking.ReadOrCreateEntry(r.ID, "monthly", month)
				check(err, "weekly create", path)
				_, err = o.QueryTable("entry_ranking").Filter("entry", r.ID).SetCond(monthly).Update(docount)
				check(err, "monthly ranking", path)

				_, _, err = ranking.ReadOrCreateEntry(r.ID, "yearly", year)
				check(err, "yearly create", path)
				_, err = o.QueryTable("entry_ranking").Filter("entry", r.ID).SetCond(yearly).Update(docount)
				check(err, "yearly ranking", path)
			case "video":
				_, err = o.QueryTable("video").Filter("entry", r.ID).Update(docount)
				check(err, path)

				_, _, err = ranking.ReadOrCreateVideo(r.ID, "dayly", day)
				check(err, "dayly create", path)
				_, err = o.QueryTable("video_ranking").Filter("video__entry", r.ID).SetCond(dayly).Update(docount)
				check(err, "dayly ranking", path)

				_, _, err = ranking.ReadOrCreateVideo(r.ID, "weekly", week)
				check(err, "weekly create", path)
				_, err = o.QueryTable("video_ranking").Filter("video__entry", r.ID).SetCond(weekly).Update(docount)
				check(err, "weekly ranking", path)

				_, _, err = ranking.ReadOrCreateVideo(r.ID, "monthly", month)
				check(err, "monthly create", path)
				_, err = o.QueryTable("video_ranking").Filter("video__entry", r.ID).SetCond(monthly).Update(docount)
				check(err, "monthly ranking", path)

				_, _, err = ranking.ReadOrCreateVideo(r.ID, "yearly", year)
				check(err, "yearly create", path)
				_, err = o.QueryTable("video_ranking").Filter("video__entry", r.ID).SetCond(yearly).Update(docount)
				check(err, "yearly ranking", path)
			case "book":
				_, err = o.QueryTable("picture").Filter("entry", r.ID).Update(docount)
				check(err, path)

				_, _, err = ranking.ReadOrCreatePicture(r.ID, "dayly", day)
				check(err, "dayly create", path)
				_, err = o.QueryTable("picture_ranking").Filter("picture__entry", r.ID).SetCond(dayly).Update(docount)
				check(err, "dayly ranking", path)

				_, _, err = ranking.ReadOrCreatePicture(r.ID, "weekly", week)
				check(err, "weekly create", path)
				_, err = o.QueryTable("picture_ranking").Filter("picture__entry", r.ID).SetCond(weekly).Update(docount)
				check(err, "weekly ranking", path)

				_, _, err = ranking.ReadOrCreatePicture(r.ID, "monthly", month)
				check(err, "monthly create", path)
				_, err = o.QueryTable("picture_ranking").Filter("picture__entry", r.ID).SetCond(monthly).Update(docount)
				check(err, "monthly ranking", path)

				_, _, err = ranking.ReadOrCreatePicture(r.ID, "yearly", year)
				check(err, "yearly create", path)
				_, err = o.QueryTable("picture_ranking").Filter("picture__entry", r.ID).SetCond(yearly).Update(docount)
				check(err, "yearly ranking", path)
			}
		}
	}

	type rankUpdater interface {
		UpdateRank(rank int64) error
	}
	var updateRank = func(ranks interface{}) {
		values := reflect.ValueOf(ranks)

		for i := 0; i < values.Len(); i++ {
			r := values.Index(i).Interface().(rankUpdater)
			r.UpdateRank(int64(i + 1))
		}
	}

	var eranks []*models.EntryRanking
	o.QueryTable("entry_ranking").SetCond(dayly).OrderBy("-page_view").All(&eranks)
	updateRank(eranks)

	eranks = make([]*models.EntryRanking, 0)
	o.QueryTable("entry_ranking").SetCond(weekly).OrderBy("-page_view").All(&eranks)
	updateRank(eranks)

	eranks = make([]*models.EntryRanking, 0)
	o.QueryTable("entry_ranking").SetCond(monthly).OrderBy("-page_view").All(&eranks)
	updateRank(eranks)

	eranks = make([]*models.EntryRanking, 0)
	o.QueryTable("entry_ranking").SetCond(yearly).OrderBy("-page_view").All(&eranks)
	updateRank(eranks)

	var vranks []*models.VideoRanking
	o.QueryTable("video_ranking").SetCond(dayly).OrderBy("-page_view").All(&vranks)
	updateRank(vranks)

	vranks = make([]*models.VideoRanking, 0)
	o.QueryTable("video_ranking").SetCond(weekly).OrderBy("-page_view").All(&vranks)
	updateRank(vranks)

	vranks = make([]*models.VideoRanking, 0)
	o.QueryTable("video_ranking").SetCond(monthly).OrderBy("-page_view").All(&vranks)
	updateRank(vranks)

	vranks = make([]*models.VideoRanking, 0)
	o.QueryTable("video_ranking").SetCond(yearly).OrderBy("-page_view").All(&vranks)
	updateRank(vranks)

	var pranks []*models.PictureRanking
	o.QueryTable("picture_ranking").SetCond(dayly).OrderBy("-page_view").All(&pranks)
	updateRank(pranks)

	pranks = make([]*models.PictureRanking, 0)
	o.QueryTable("picture_ranking").SetCond(weekly).OrderBy("-page_view").All(&pranks)
	updateRank(pranks)

	pranks = make([]*models.PictureRanking, 0)
	o.QueryTable("picture_ranking").SetCond(monthly).OrderBy("-page_view").All(&pranks)
	updateRank(pranks)

	pranks = make([]*models.PictureRanking, 0)
	o.QueryTable("picture_ranking").SetCond(yearly).OrderBy("-page_view").All(&pranks)
	updateRank(pranks)

	return
}

func check(err error, args ...interface{}) {
	if err != nil {
		beego.Warn(fmt.Sprintf("Update PageView by %v: ", args), err)
	}
}
