package anime

import (
	"encoding/json"
	"reflect"

	"bitbucket.org/ikeikeikeike/antenna/lib/cache"
	"bitbucket.org/ikeikeikeike/antenna/models"
)

/*
	Cache expires: 1 days
*/
func CachedAnimes() (objs []*models.Anime) {
	key := "models.anime.cachedanimes"
	s := reflect.ValueOf(cache.Client.Get(key))

	if !cache.Client.IsExist(key) {
		models.Animes().Limit(1000000).All(&objs)

		bytes, _ := json.Marshal(objs)
		cache.Client.Put(key, bytes, 60*60*24*1)
	} else {
		json.Unmarshal(s.Interface().([]uint8), &objs)
	}
	return
}
