package summarize

import "bitbucket.org/ikeikeikeike/antenna/models"

// TODO: To unite
func DivaVideoscount() (err error) {
	var divas []*models.Diva

	qs := models.Divas().RelatedSel()
	qs.Limit(10000000).All(&divas)

	for _, d := range divas {
		nums, err := models.Videos().Filter("divas__diva", d).Count()
		if err == nil {
			if d.VideosCount < int(nums) {
				d.VideosCount = int(nums)
				d.Update("VideosCount", "Updated")
			}

		}
	}
	return
}
