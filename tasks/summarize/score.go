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
  We execute the following is sequence.

	  1. Provide score to blogs.
	  2. Update blog's total score to the storage.
	  3. Add entry(id) to the summary model: could register up to 1000 entry.

  Note:
	Regard 1 Session(e.g. UU) to combine UserAgent and Host.

*/
func InScore() (err error) {
	var blogs []*models.Blog

	bqs := blog.LivingBlogs()
	if !bqs.Exist() {
		err = errors.New("Blogs does not exists")
		return
	}
	bqs.All(&blogs)

	// Find referer by Stored URL(site host)
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
	// We terminate following is conditions.
	//   - Added 30 record to summary.
	//   - If 30 entities stored already, no register.
	//   - 10,000 Loop.
	//   - 7 days ago.
	//
	summary.WeightingPushEntryBy(choices)

	// If more than five hundred, We remove from the old record.
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
