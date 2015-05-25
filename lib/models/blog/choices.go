package blog

import (
	"bitbucket.org/ikeikeikeike/antenna/models"
	"bitbucket.org/ikeikeikeike/antenna/models/score"
	"github.com/jmcvetta/randutil"
	"github.com/k0kubun/pp"
)

/*
	Return weighted items by blog's score.
*/
func WeightChoiceBlogs(in []*models.Blog, max int) []*models.Blog {
	var (
		choices []randutil.Choice
		blogs   []*models.Blog
	)

	for _, b := range in {
		s := score.RandomGetByBlog(b)
		pp.Println(s.Name)

		// If site had many blog, Remove code below.
		cnt := int(s.Count)
		if cnt < 1 {
			cnt = 1
		}

		// DMCA affect
		weight := cnt * b.VerifyScore()
		if b.IsPenalty {
			weight = 1
		}

		choices = append(choices, randutil.Choice{
			Weight: weight,
			Item:   b,
		})
	}

	limit := 0
	for len(blogs) < max && limit <= 1000000 {
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
