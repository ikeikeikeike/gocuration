package main

import (
	"testing"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"

	"bitbucket.org/ikeikeikeike/antenna/lib/accessctl"
	"bitbucket.org/ikeikeikeike/antenna/models"
	"bitbucket.org/ikeikeikeike/antenna/models/blog"
	"github.com/k0kubun/pp"
)

func TestInScoreLogic(t *testing.T) {
	var blogs []*models.Blog

	scoredin := accessctl.NewScoringInLog()
	scoredin.Cleaned = false
	scoredin.Bootstrap()

	blog.LivingBlogs().All(&blogs)
	for _, b := range blogs {

		// Calc score.
		weight, err := scoredin.Scoring(b.Url)
		pp.Println(b.Url, weight)

		if err != nil {
			continue
		} else if weight < 1 {
			continue
		}

	}
}
