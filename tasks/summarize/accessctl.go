package summarize

import (
	"fmt"

	"bitbucket.org/ikeikeikeike/antenna/lib/accessctl"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func Showcounter() (err error) {
	o := orm.NewOrm()

	c := accessctl.NewShowCounter()
	c.Bootstrap()

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
			case "video":
				_, err = o.QueryTable("video").Filter("entry", r.ID).Update(docount)
			case "book":
				_, err = o.QueryTable("picture").Filter("entry", r.ID).Update(docount)
			}

			if err != nil {
				beego.Warn(fmt.Sprintf("Update PageView by path(%s): ", path), err)
			}
		}
	}

	return
}
