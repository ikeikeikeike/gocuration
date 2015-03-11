package entry

import (
	"strconv"

	"github.com/ikeikeikeike/gopkg/convert"

	"bitbucket.org/ikeikeikeike/antenna/lib/cache"
	"bitbucket.org/ikeikeikeike/antenna/models"
)

/*
	Cache expires: 1 hour
*/
func CachedMaxLength() (length int64) {
	key := "models.entry.cachedmaxlength"

	if !cache.Client.IsExist(key) {
		length, _ = models.Entries().Count()
		cache.Client.Put(key, length, 60*60)
	} else {
		i, _ := strconv.Atoi(convert.ToStr(cache.Client.Get(key)))
		length = int64(i)
	}
	return
}
