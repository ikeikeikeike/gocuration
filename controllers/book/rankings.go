package book

import "bitbucket.org/ikeikeikeike/antenna/ormapper"

type RankingsController struct {
	BaseController
}

func (c *RankingsController) NestFinish() {
	c.PushInAccessLog()
}

func (c *RankingsController) Dayly() {
	c.TplNames = "book/rankings/dayly.tpl"

	db := ormapper.PictureRankings().
		Where("picture_ranking.rank > ?", 0).
		Where("picture_ranking.begin_name = ?", "dayly").
		Where("picture_ranking.begin_time = ?", c.GetParamatedNow().BeginningOfDay()).
		Limit(100)

	var rankings []*ormapper.PictureRanking
	db.Order("picture_ranking.rank ASC").Find(&rankings)

	for _, s := range rankings {
		s.RankingsLoader()
	}

	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Weekly() {
	c.TplNames = "book/rankings/weekly.tpl"

	db := ormapper.PictureRankings().
		Where("picture_ranking.rank > ?", 0).
		Where("picture_ranking.begin_name = ?", "weekly").
		Where("picture_ranking.begin_time = ?", c.GetParamatedNow().BeginningOfWeek()).
		Limit(100)

	var rankings []*ormapper.PictureRanking
	db.Order("picture_ranking.rank ASC").Find(&rankings)

	for _, s := range rankings {
		s.RankingsLoader()
	}

	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Monthly() {
	c.TplNames = "book/rankings/monthly.tpl"

	db := ormapper.PictureRankings().
		Where("picture_ranking.rank > ?", 0).
		Where("picture_ranking.begin_name = ?", "monthly").
		Where("picture_ranking.begin_time = ?", c.GetParamatedNow().BeginningOfMonth()).
		Limit(100)

	var rankings []*ormapper.PictureRanking
	db.Order("picture_ranking.rank ASC").Find(&rankings)

	for _, s := range rankings {
		s.RankingsLoader()
	}

	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Yearly() {
	c.TplNames = "book/rankings/yearly.tpl"

	db := ormapper.PictureRankings().
		Where("picture_ranking.rank > ?", 0).
		Where("picture_ranking.begin_name = ?", "yearly").
		Where("picture_ranking.begin_time = ?", c.GetParamatedNow().BeginningOfYear()).
		Limit(100)

	var rankings []*ormapper.PictureRanking
	db.Order("picture_ranking.rank ASC").Find(&rankings)

	for _, s := range rankings {
		s.RankingsLoader()
	}

	c.Data["Rankings"] = rankings
}
