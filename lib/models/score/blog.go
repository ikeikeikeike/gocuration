package score

import (
	"errors"

	"bitbucket.org/ikeikeikeike/antenna/lib/accessctl"
	"bitbucket.org/ikeikeikeike/antenna/models"
	"github.com/jinzhu/now"
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
	Update each score.
*/
func UpdateBlogEachScoreBy(choices []randutil.Choice) {
	for _, wc := range choices {
		// Update total score.
		total := &models.Score{Name: "in", Blog: wc.Item.(*models.Blog)}
		_, _, _ = total.ReadOrCreate("Name", "Blog")

		month := &models.Score{Name: "month_in", Blog: wc.Item.(*models.Blog)}
		_, _, _ = month.ReadOrCreate("Name", "Blog")

		week := &models.Score{Name: "week_in", Blog: wc.Item.(*models.Blog)}
		_, _, _ = week.ReadOrCreate("Name", "Blog")

		day := &models.Score{Name: "day_in", Blog: wc.Item.(*models.Blog)}
		_, _, _ = day.ReadOrCreate("Name", "Blog")

		if !(now.BeginningOfMonth().Unix() < month.Updated.Unix() && month.Updated.Unix() < now.EndOfMonth().Unix()) {
			month.Count = 0
			month.Update()
		}
		if !(now.BeginningOfWeek().Unix() < week.Updated.Unix() && week.Updated.Unix() < now.EndOfWeek().Unix()) {
			week.Count = 0
			week.Update()
		}
		if !(now.BeginningOfDay().Unix() < day.Updated.Unix() && day.Updated.Unix() < now.EndOfDay().Unix()) {
			day.Count = 0
			day.Update()
		}

		if wc.Weight > 0 {
			total.Count += int64(wc.Weight)
			total.Update()

			month.Count += int64(wc.Weight)
			month.Update()

			week.Count += int64(wc.Weight)
			week.Update()

			day.Count += int64(wc.Weight)
			day.Update()
		}
	}
}
