package summarize

import "bitbucket.org/ikeikeikeike/antenna/models"

// TODO: To unite
func CharacterPicturescount() (err error) {
	var characters []*models.Character

	qs := models.Characters().RelatedSel()
	qs.Limit(10000000).All(&characters)

	for _, c := range characters {
		nums, err := models.Pictures().Filter("characters__character", c).Count()
		if err == nil {
			if c.PicturesCount < int(nums) {
				c.PicturesCount = int(nums)
				c.Update("PicturesCount", "Updated")
			}

		}
	}
	return
}
