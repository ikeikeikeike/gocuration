package image

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"time"

	"github.com/astaxie/beego/orm"

	"bitbucket.org/ikeikeikeike/antenna/lib/cache"
)

/*
	Cache expires: 3 days
*/
func CachedSources(size string) (sources []string) {
	key := fmt.Sprintf("models.image.cachedfallbackimage:%s", size)
	s := reflect.ValueOf(cache.Client.Get(key))

	if !cache.Client.IsExist(key) {
		qs := orm.NewOrm().QueryTable("image").OrderBy("-Id")

		if size == "small" {
			qs = qs.Filter("width__gt", 100).Filter("width__lt", 300)
		} else if size == "middle" {
			qs = qs.Filter("width__gt", 300).Filter("width__lt", 500)
		} else {
			qs = qs.Filter("width__gt", 500)
		}

		var list orm.ParamsList
		qs.Limit(5000).ValuesFlat(&list, "src")

		for _, src := range list {
			sources = append(sources, src.(string))
		}

		bytes, _ := json.Marshal(sources)
		cache.Client.Put(key, bytes, 60*60*24*3)
	} else {
		json.Unmarshal(s.Interface().([]uint8), &sources)
	}
	return
}

func CachedRandomSrc(size string) string {
	sources := CachedSources(size)
	return sources[randomNumber(1, len(sources)-1)]
}

func randomNumber(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}
