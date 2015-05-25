package score

import (
	"bitbucket.org/ikeikeikeike/antenna/models"
	"github.com/ikeikeikeike/shuffler"
)

func RandomGetByBlog(blog *models.Blog) *models.Score {
	names := []string{"in", "year_in", "month_in", "week_in"}
	shuffler.Shuffle(names)

	s := &models.Score{Name: names[0], Blog: blog}
	_, _, _ = s.ReadOrCreate("Name", "Blog")
	return s
}
