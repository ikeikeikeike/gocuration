package summary

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego/orm"
)

func RelatedSummaries(sid int64, in []string, summaries interface{}) {
	if len(in) <= 0 {
		in = append(in, "巨乳")
	}

	// models.Summaries().RelatedSel().
	// Filter("entry__tags__tag__name__in", in).
	// Limit(15).All(summaries)

	// 上記を `DISTINCT` 付きでやっている
	names := fmt.Sprintf("'%s'", strings.Join(in, "','"))
	q := fmt.Sprintf(`
	SELECT DISTINCT 
		s.* FROM summary as s 
	LEFT OUTER JOIN 
		entry e ON e.id = s.entry_id 
	LEFT OUTER JOIN 
		blog b ON b.id = e.blog_id 
	LEFT OUTER JOIN 
		entry_tag et ON et.entry_id = e.id 
	LEFT OUTER JOIN 
		tag tag ON tag.id = et.tag_id 
	WHERE 
		(tag.name IN (%s) OR e.q like '%%%s%%') 
	  AND 
		e.id != '%d'
	ORDER BY 
		s.sort DESC 
	LIMIT 3`, names, names[0], sid)
	orm.NewOrm().Raw(q).QueryRows(summaries)
}
