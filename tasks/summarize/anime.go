package summarize

import "bitbucket.org/ikeikeikeike/antenna/models"

func AnimePicturescount() (err error) {
	var animes []*models.Anime

	qs := models.Animes().RelatedSel()
	qs.Limit(10000000).All(&animes)

	for _, m := range animes {
		nums, err := models.Pictures().Filter("anime", m).Count()
		if err == nil {
			if m.PicturesCount < int(nums) {
				m.PicturesCount = int(nums)
				m.Update("PicturesCount")
			}
		}
	}
	return
}
