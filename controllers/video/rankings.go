package video

import "bitbucket.org/ikeikeikeike/antenna/ormapper"

type RankingsController struct {
	BaseController
}

func (c *RankingsController) NestFinish() {
	c.PushInAccessLog()
}

func (c *RankingsController) Dayly() {
	c.TplNames = "video/rankings/dayly.tpl"

	db := ormapper.VideoRankings().
		Where("video_ranking.rank > ?", 0).
		Where("video_ranking.begin_name = ?", "dayly").
		Where("video_ranking.begin_time = ?", c.GetParamatedNow().BeginningOfDay()).
		Limit(100)

	var rankings []*ormapper.VideoRanking
	db.Order("video_ranking.rank ASC").Find(&rankings)

	for _, s := range rankings {
		s.RankingsLoader()
	}

	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Weekly() {
	c.TplNames = "video/rankings/weekly.tpl"

	db := ormapper.VideoRankings().
		Where("video_ranking.rank > ?", 0).
		Where("video_ranking.begin_name = ?", "weekly").
		Where("video_ranking.begin_time = ?", c.GetParamatedNow().BeginningOfWeek()).
		Limit(100)

	var rankings []*ormapper.VideoRanking
	db.Order("video_ranking.rank ASC").Find(&rankings)

	for _, s := range rankings {
		s.RankingsLoader()
	}

	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Monthly() {
	c.TplNames = "video/rankings/monthly.tpl"

	db := ormapper.VideoRankings().
		Where("video_ranking.rank > ?", 0).
		Where("video_ranking.begin_name = ?", "monthly").
		Where("video_ranking.begin_time = ?", c.GetParamatedNow().BeginningOfMonth()).
		Limit(100)

	var rankings []*ormapper.VideoRanking
	db.Order("video_ranking.rank ASC").Find(&rankings)

	for _, s := range rankings {
		s.RankingsLoader()
	}

	c.Data["Rankings"] = rankings
}

func (c *RankingsController) Yearly() {
	c.TplNames = "video/rankings/yearly.tpl"

	db := ormapper.VideoRankings().
		Where("video_ranking.rank > ?", 0).
		Where("video_ranking.begin_name = ?", "yearly").
		Where("video_ranking.begin_time = ?", c.GetParamatedNow().BeginningOfYear()).
		Limit(100)

	var rankings []*ormapper.VideoRanking
	db.Order("video_ranking.rank ASC").Find(&rankings)

	for _, s := range rankings {
		s.RankingsLoader()
	}

	c.Data["Rankings"] = rankings
}
