package character

import (
	"encoding/json"
	"reflect"

	"bitbucket.org/ikeikeikeike/antenna/lib/cache"
	"bitbucket.org/ikeikeikeike/antenna/models"
)

/*
	Cache expires: 1 days
*/
func CachedCharacters() (objs []*models.Character) {
	key := "models.character.cachedcharacters"
	s := reflect.ValueOf(cache.Client.Get(key))

	if !cache.Client.IsExist(key) {
		models.Characters().Limit(1000000).All(&objs)

		bytes, _ := json.Marshal(objs)
		cache.Client.Put(key, bytes, 60*60*24*1)
	} else {
		json.Unmarshal(s.Interface().([]uint8), &objs)
	}

	return
}
