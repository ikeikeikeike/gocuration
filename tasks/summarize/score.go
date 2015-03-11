package summarize

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ikeikeikeike/gopkg/db/dialect"

	"bitbucket.org/ikeikeikeike/antenna/lib/models/entry"
	"bitbucket.org/ikeikeikeike/antenna/lib/models/score"
	"bitbucket.org/ikeikeikeike/antenna/lib/models/summary"
	"bitbucket.org/ikeikeikeike/antenna/models"
	"bitbucket.org/ikeikeikeike/antenna/models/blog"
)

func SocialScore() (err error) {
	var entries []*models.Entry
	t := time.Now().AddDate(0, 0, -7) // 7 days ago

	qs := models.Entries()
	qs = qs.Filter("published_at__gte", t)
	qs.All(&entries)

	for _, s := range entries {
		time.Sleep(500 * time.Millisecond)
		entry.UpdateSocials(s)
	}
	return
}

/*
  ブログを回して登録RSS or 登録URLを元にReferer名をfind
  UserAgent, Hostで UU(1セッション)とする

  ブログの累積Inアクセスを追加

  SummaryモデルにEntry(ID)を追加する
  Summary Tableは500件まで
*/
func InScore() (err error) {
	var blogs []*models.Blog

	bqs := blog.LivingBlogs()
	if !bqs.Exist() {
		err = errors.New("Blogs does not exists")
		return
	}
	bqs.All(&blogs)

	choices, err := score.ToScoredFrom(blogs)
	if err != nil {
		return
	}

	msg := "[tasks.summarize.score.InScore] "

	// TODO: Return update number and write log it.
	// Update total score
	score.UpdateBlogTotalScoreBy(choices)

	// TODO: Return Update number and write log it.
	//
	// The following is terminated
	//   - Add thirty record to summary
	//   - If stored 30 entities yet, no register.
	//   - 10,000 Loop.
	//   - 7 days ago.
	summary.WeightingPushEntryBy(choices)

	// If more than five hundred, Remove from old record.
	max := int64(500)
	sqs := models.Summaries()

	cnt, _ := sqs.Count()
	if max < cnt {
		var summs []*models.Summary
		sqs.OrderBy("created").Limit(cnt - max).All(&summs)
		for _, s := range summs {
			s.Delete()
		}
	}

	// Random summaries
	sql := fmt.Sprintf(`-- TODO: dialect random function
	UPDATE summary SET sort = round(%s * 10000);`, dialect.RandomBuiltinFunc())

	_, err = orm.NewOrm().Raw(sql).Exec()
	if err != nil {
		msg += "Missing randomize summary:"
		beego.Warn(msg, err)
	}

	return
}
