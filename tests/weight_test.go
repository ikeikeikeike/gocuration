package main

import (
	"testing"

	"github.com/k0kubun/pp"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	libb "bitbucket.org/ikeikeikeike/antenna/lib/models/blog"
	"bitbucket.org/ikeikeikeike/antenna/models"
	"bitbucket.org/ikeikeikeike/antenna/models/blog"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestWeight(t *testing.T) {
	var blogs []*models.Blog
	blog.LivingBlogs().All(&blogs)

	blogs = libb.WeightChoiceBlogs(blogs, 10)

	for _, b := range blogs {
		pp.Println(b.Id, b.Rss, b.Url, b.Name, b.VerifyScore())
	}
}
