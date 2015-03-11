package entry

import (
	"strings"

	"bitbucket.org/ikeikeikeike/antenna/models"
	"github.com/ikeikeikeike/gopkg/sharecount"
)

func UpdateSocials(entry *models.Entry) {
	shares := sharecount.Fetch(entry.Url)

	for _, c := range shares {
		cnt := int64(c.Count)
		if 0 < cnt {
			s := &models.Score{
				Entry: entry,
				Name:  strings.ToLower(c.Service),
			}
			_, _, _ = s.ReadOrCreate("Entry", "Name")

			if s.Count < cnt {
				s.Count = cnt
				s.Update()
			}
		}
	}
}
