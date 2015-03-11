package blog

import (
	"bitbucket.org/ikeikeikeike/antenna/models"
	"github.com/astaxie/beego/orm"
)

// RSS and URL was living
func LivingBlogs() orm.QuerySeter {
	c := orm.NewCondition()
	c = c.And("rss__isnull", false).AndNot("rss", "")
	c = c.And("url__isnull", false).AndNot("url", "")

	return models.Blogs().SetCond(c)
}
