package diva

import (
	"reflect"
	"strings"

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

		cache.Client.Put(key, names, 60*60*24*10)
	} else {
		divas := strings.Split(string(s.Interface().([]uint8)), " ")
		for _, name := range divas {
			names = append(names, name)
		}
	}

	return names
}
