package main

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/ikeikeikeike/shuffler"
	"github.com/k0kubun/pp"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	libb "bitbucket.org/ikeikeikeike/antenna/lib/models/blog"
	"bitbucket.org/ikeikeikeike/antenna/models"
	"bitbucket.org/ikeikeikeike/antenna/models/blog"
	"bitbucket.org/ikeikeikeike/antenna/models/image"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestLivingBlogs(t *testing.T) {
	var blogs []*models.Blog
	blog.LivingBlogs().All(&blogs)

	blogs = libb.WhightChoiceBlogs(blogs, 100)

	for _, b := range blogs {
		pp.Println(b.Id, b.Rss, b.Url, b.Name, b.VerifyScore())
	}

	println("shuffle")

	shuffle := shuffler.Shuffler(blogs).(reflect.Value)
	for i := 0; i < shuffle.Len(); i++ {
		b := shuffle.Index(i).Interface().(*models.Blog)
		pp.Println(b.Id, b.Rss, b.Url, b.Name, b.VerifyScore())
	}

}

func TestParsingBlogs(t *testing.T) {
	var blogs []*models.Blog

	o := orm.NewOrm()
	c := orm.NewCondition()
	cond := c.And("rss__isnull", false).OrNot("rss", "")
	o.QueryTable("blog").SetCond(cond).RelatedSel().All(&blogs)

	for _, b := range libb.WhightChoiceBlogs(blogs, 100) {
		pp.Println(b.Id, b.Rss, b.Url, b.Name, b.VerifyScore())
	}

	pp.Println(image.CachedRandomSrc("small"))
	pp.Println(image.CachedRandomSrc("middle"))
	pp.Println(image.CachedRandomSrc("large"))

	pp.Println(image.CachedRandomSrc("small"))
	pp.Println(image.CachedRandomSrc("middle"))
	pp.Println(image.CachedRandomSrc("large"))

	pp.Println(image.CachedRandomSrc("small"))
	pp.Println(image.CachedRandomSrc("middle"))
	pp.Println(image.CachedRandomSrc("large"))

	pp.Println(image.CachedRandomSrc("small"))
	pp.Println(image.CachedRandomSrc("middle"))
	pp.Println(image.CachedRandomSrc("large"))

	pp.Println(image.CachedRandomSrc("small"))
	pp.Println(image.CachedRandomSrc("middle"))
	pp.Println(image.CachedRandomSrc("large"))

	pp.Println(image.CachedRandomSrc("small"))
	pp.Println(image.CachedRandomSrc("middle"))
	pp.Println(image.CachedRandomSrc("large"))

	pp.Println(image.CachedRandomSrc("small"))
	pp.Println(image.CachedRandomSrc("middle"))
	pp.Println(image.CachedRandomSrc("large"))

	pp.Println(image.CachedRandomSrc("small"))
	pp.Println(image.CachedRandomSrc("middle"))
	pp.Println(image.CachedRandomSrc("large"))

	pp.Println(image.CachedRandomSrc("small"))
	pp.Println(image.CachedRandomSrc("middle"))
	pp.Println(image.CachedRandomSrc("large"))

}

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}
