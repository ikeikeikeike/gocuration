package summarize

import (
	"fmt"
	"net/url"
	"reflect"

	"bitbucket.org/ikeikeikeike/antenna/lib/models/blog"
	"bitbucket.org/ikeikeikeike/antenna/lib/models/entry"
	"bitbucket.org/ikeikeikeike/antenna/lib/models/picture"
	"bitbucket.org/ikeikeikeike/antenna/lib/models/video"
	"bitbucket.org/ikeikeikeike/antenna/models"

	rss "github.com/ikeikeikeike/go-pkg-rss"
	"github.com/ikeikeikeike/shuffler"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func RssFeed() (err error) {
	var (
		msg   string
		feed  *rss.Feed
		blogs []*models.Blog
		chs   []*rss.Channel
	)

	o := orm.NewOrm()
	c := orm.NewCondition()
	cond := c.And("rss__isnull", false).OrNot("rss", "")

	num, err := o.QueryTable("blog").SetCond(cond).RelatedSel().All(&blogs)
	if err != nil || num < 1 {
		return
	}

	// XXX: Bad performance random sort is no problem.
	//		Because less than 10 thousand blogs target.
	blogs = blog.WeightChoiceBlogs(blogs, 60)
	shuffle := shuffler.Shuffler(blogs).(reflect.Value)

	for i := 0; i < shuffle.Len(); i++ {
		b := shuffle.Index(i).Interface().(*models.Blog)

		if b.Rss == "" {
			beego.Warn("blank Rss field")
			continue
		} else if u, err := url.Parse(b.Rss); u.Scheme == "" {
			beego.Warn("scheme less Rss field: ", err)
			continue
		}

		feed = rss.New(20, true, chHandler, itemHandler)
		if err := feed.Fetch(b.Rss, nil); err != nil {
			beego.Warn("Rss fetch error: ", err)
			continue
		}

		chs = feed.Channels
		if len(chs) < 1 {
			continue
		}

		if err := blog.UpdateByChannel(b, chs[0], false); err != nil {
			beego.Error("update err Blog{id=", b.Id, "}: ", err)
		}

		// Q'ty: 2 chs
		entries, errs := entry.AddsByChannels(numberOf(chs, 2), b)
		if len(errs) > 0 {
			msg = fmt.Sprintf("Add err %s Entries by Blog{id=%s}: ",
				len(entries), b.Id)
			beego.Error(msg, errs)
		}

		if len(entries) > 0 {
			// Add video urls for video summary site
			video.AddsByEntries(entries)

			// Add picture urls for image summary site
			picture.AddsByEntries(entries)
		}
	}

	return
}

func chHandler(f *rss.Feed, chs []*rss.Channel)                   {} // fmt.Printf("%d new channel(s) in %s\n", len(chs), f.Url)
func itemHandler(f *rss.Feed, ch *rss.Channel, items []*rss.Item) {} // fmt.Printf("%d new item(s) in %s\n", len(items), f.Url)

func numberOf(channels []*rss.Channel, num int) (chs []*rss.Channel) {
	for i, ch := range channels {
		chs = append(chs, ch)
		if i >= num {
			break
		}
	}
	return chs
}
