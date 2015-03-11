package main

import (
	"testing"

	"github.com/jmcvetta/randutil"

	"bitbucket.org/ikeikeikeike/antenna/models"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"
)

func TestBlogsChoices(t *testing.T) {
	var blogs []*models.Blog

	qs := models.Blogs()
	qs.All(&blogs)

	for _, b := range blogs {
		println(b.Id)
	}

	for _, b := range whightChoiceBlogs(blogs, 1000) {
		println(b.Id)
	}
}

/*
	whighted blog items by blog in score.

*/
func whightChoiceBlogs(in []*models.Blog, max int) []*models.Blog {
	var (
		choices []randutil.Choice
		blogs   []*models.Blog
	)

	for _, b := range in {
		s := &models.Score{Name: "in", Blog: b}
		_, _, _ = s.ReadOrCreate("Name", "Blog")

		cnt := int(s.Count)
		if cnt < 1 {
			cnt = 0
		}
		choices = append(choices, randutil.Choice{
			Weight: cnt,
			Item:   b,
		})
	}

	limit := 0
	for len(blogs) < max && limit <= 5000 {
		limit++

		wc, err := randutil.WeightedChoice(choices)
		if err == nil {
			blogs = appendIfMissing(blogs, wc.Item.(*models.Blog))
		}
	}

	return blogs
}

func appendIfMissing(blogs []*models.Blog, blog *models.Blog) []*models.Blog {
	for _, elm := range blogs {
		if elm == blog {
			return blogs
		}
	}
	return append(blogs, blog)
}
