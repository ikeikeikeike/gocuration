package score

import (
	"errors"

	"bitbucket.org/ikeikeikeike/antenna/lib/accessctl"
	"bitbucket.org/ikeikeikeike/antenna/models"
	"github.com/jmcvetta/randutil"
)

/*
	To scored(Include randutil.Choice)
*/
func ToScoredFrom(blogs []*models.Blog) (choices []randutil.Choice, err error) {

	scoredin := accessctl.NewScoringInLog()
	scoredin.Bootstrap()

	for _, b := range blogs {

		// Calc score.
		weight, err := scoredin.Scoring(b.Url)
		if err != nil {
			continue
		} else if weight < 1 {
			continue
		}

		// Push WeightedChoice element
		choices = append(choices, randutil.Choice{
			Weight: weight,
			Item:   b,
		})
	}

	if len(choices) < 1 {
		err = errors.New("No scoring accesslog: length=0")
		return
	}

	return
}

/*
	Update total score.
*/
func UpdateBlogTotalScoreBy(choices []randutil.Choice) {
	for _, wc := range choices {
		s := &models.Score{Name: "in", Blog: wc.Item.(*models.Blog)}
		_, _, _ = s.ReadOrCreate("Name", "Blog")

		// Update total score.
		if wc.Weight > 0 {
			s.Count += int64(wc.Weight)
			s.Update()
		}
	}
}
