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

	t := time.Now().UTC()
	day := now.New(t).BeginningOfDay()
	week := now.New(t).BeginningOfWeek()
	month := now.New(t).BeginningOfMonth()
	year := now.New(t).BeginningOfYear()

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
				check(err, "dayly create", path, r.ID)
				_, err = o.QueryTable("entry_ranking").SetCond(dayly).Filter("entry", r.ID).Update(docount)
				check(err, "dayly ranking", path, r.ID)

				_, _, err = ranking.ReadOrCreateEntry(r.ID, "weekly", week)
				check(err, "weekly create", path, r.ID)
				_, err = o.QueryTable("entry_ranking").SetCond(weekly).Filter("entry", r.ID).Update(docount)
				check(err, "weekly ranking", path, r.ID)

				_, _, err = ranking.ReadOrCreateEntry(r.ID, "monthly", month)
				check(err, "weekly create", path, r.ID)
				_, err = o.QueryTable("entry_ranking").SetCond(monthly).Filter("entry", r.ID).Update(docount)
				check(err, "monthly ranking", path, r.ID)

				_, _, err = ranking.ReadOrCreateEntry(r.ID, "yearly", year)
				check(err, "yearly create", path, r.ID)
				_, err = o.QueryTable("entry_ranking").SetCond(yearly).Filter("entry", r.ID).Update(docount)
				check(err, "yearly ranking", path, r.ID)
			case "video":
				_, err = o.QueryTable("video").Filter("entry", r.ID).Update(docount)
				check(err, path)

				_, _, err = ranking.ReadOrCreateVideo(r.ID, "dayly", day)
				check(err, "dayly create", path, r.ID)
				_, err = o.QueryTable("video_ranking").SetCond(dayly).Filter("video__entry__id", r.ID).Update(docount)
				check(err, "dayly ranking", path, r.ID)

				_, _, err = ranking.ReadOrCreateVideo(r.ID, "weekly", week)
				check(err, "weekly create", path, r.ID)
				_, err = o.QueryTable("video_ranking").SetCond(weekly).Filter("video__entry__id", r.ID).Update(docount)
				check(err, "weekly ranking", path, r.ID)

				_, _, err = ranking.ReadOrCreateVideo(r.ID, "monthly", month)
				check(err, "monthly create", path, r.ID)
				_, err = o.QueryTable("video_ranking").SetCond(monthly).Filter("video__entry__id", r.ID).Update(docount)
				check(err, "monthly ranking", path, r.ID)

				_, _, err = ranking.ReadOrCreateVideo(r.ID, "yearly", year)
				check(err, "yearly create", path, r.ID)
				_, err = o.QueryTable("video_ranking").SetCond(yearly).Filter("video__entry__id", r.ID).Update(docount)
				check(err, "yearly ranking", path, r.ID)
			case "book":
				_, err = o.QueryTable("picture").Filter("entry", r.ID).Update(docount)
				check(err, path)

				_, _, err = ranking.ReadOrCreatePicture(r.ID, "dayly", day)
				check(err, "dayly create", path, r.ID)
				_, err = o.QueryTable("picture_ranking").SetCond(dayly).Filter("picture__entry__id", r.ID).Update(docount)
				check(err, "dayly ranking", path, r.ID)

				_, _, err = ranking.ReadOrCreatePicture(r.ID, "weekly", week)
				check(err, "weekly create", path, r.ID)
				_, err = o.QueryTable("picture_ranking").SetCond(weekly).Filter("picture__entry__id", r.ID).Update(docount)
				check(err, "weekly ranking", path, r.ID)

				_, _, err = ranking.ReadOrCreatePicture(r.ID, "monthly", month)
				check(err, "monthly create", path, r.ID)
				_, err = o.QueryTable("picture_ranking").SetCond(monthly).Filter("picture__entry__id", r.ID).Update(docount)
				check(err, "monthly ranking", path, r.ID)

				_, _, err = ranking.ReadOrCreatePicture(r.ID, "yearly", year)
				check(err, "yearly create", path, r.ID)
				_, err = o.QueryTable("picture_ranking").SetCond(yearly).Filter("picture__entry__id", r.ID).Update(docount)
				check(err, "yearly ranking", path, r.ID)
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
		beego.Notice(fmt.Sprintf("Update PageView by %v: ", args), err)
	}
}
