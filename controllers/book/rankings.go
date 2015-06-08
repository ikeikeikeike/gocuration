package book

import (
	"bitbucket.org/ikeikeikeike/antenna/ormapper"
	"bitbucket.org/ikeikeikeike/antenna/ormapper/ranking"
)

type RankingsController struct {
	BaseController
}

func (c *RankingsController) NestFinish() {
	c.PushInAccessLog()
}

func (c *RankingsController) Dayly() {
	c.TplNames = "book/rankings/dayly.tpl"

	db := ormapper.EntryRankings().
		Scopes(ranking.RankMoreThanZero).
		Scopes(ranking.HasDayly).
		Scopes(ranking.FilterBeginTime(c.GetParamatedNow().BeginningOfDay())).
		Limit(100)

	var rankings []*ormapper.EntryRanking
	db.Order("entry_ranking.rank ASC").Find(&rankings)

	// for _, s := range rankings { s.NewsLoader() }

	c.Data["Rankings"] = rankings
}
