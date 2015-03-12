package diva

import (
	"encoding/json"
	"reflect"

	"bitbucket.org/ikeikeikeike/antenna/lib/cache"
	"bitbucket.org/ikeikeikeike/antenna/models"

	"github.com/robpike/filter"
)

/*
	Cache expires: 10 days
*/
func DivasName() (names []string) {
	key := "models.divas.name"
	s := reflect.ValueOf(cache.Client.Get(key))

	if !cache.Client.IsExist(key) || s.Len() <= 0 {
		var divas []*models.Diva
		models.Divas().Limit(1000000).All(&divas)

		names = filter.Apply(divas, func(d *models.Diva) string {
			return d.Name
		}).([]string)

		bytes, _ := json.Marshal(names)
		cache.Client.Put(key, bytes, 60*60*24*10)
	} else {
		json.Unmarshal(s.Interface().([]uint8), &names)
	}

	return names
}
