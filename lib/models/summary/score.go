package summary

import (
	"time"

	"bitbucket.org/ikeikeikeike/antenna/models"
	"github.com/astaxie/beego/orm"
	"github.com/jmcvetta/randutil"
)

/*
	Find entry 14 days ago and not pushed entry yet.

	The following is terminated ::

		- Add thirty record to summary
		- If stored 30 entities already, no register.
		- 10,000 Loop.
		- 7 days ago.

*/
func WeightingPushEntryBy(choices []randutil.Choice) {
	var entry *models.Entry
	var blog *models.Blog

	i := 0
	m := 0
	o := orm.NewOrm()
	t := time.Now().AddDate(0, 0, -5) // 7 days ago
	sqs := models.Summaries()

	// entry and summary one to one relation.
	for i <= 30 && m <= 5000 {
		m++
		time.Sleep(100 * time.Millisecond)

		wc, _ := randutil.WeightedChoice(choices)
		blog = wc.Item.(*models.Blog)

		cnt, _ := sqs.Filter("entry__blog__id", blog).Count()
		if cnt > 15 {
			continue
		}

		sql := `-- XXX: Need summary left join
		SELECT entry.* FROM entry 
		INNER JOIN 
		  entry_image ON entry_image.entry_id = entry.id 
		INNER JOIN 
		  image ON image.id = entry_image.image_id 
		INNER JOIN 
		  blog ON blog.id = entry.blog_id 
		LEFT OUTER JOIN 
		  summary ON summary.entry_id = entry.id 
		WHERE 
		  entry.blog_id = ? 
		    AND 
		  entry.published_at >= ? 
			AND 
		  summary.id IS NULL 
			AND 
		  image.width > 280 
		    AND 
		  image.height > 200
		ORDER BY entry.id DESC LIMIT 1
		`
		err := o.Raw(sql, blog, t).QueryRow(&entry)
		if err != nil {
			continue
		}

		s := &models.Summary{Entry: entry}
		created, _, _ := s.ReadOrCreate("Entry")
		if created {
			i++
		}
	}
}
