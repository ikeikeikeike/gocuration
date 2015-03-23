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
	Cache expires: 1 days
*/
func CachedSources(size string) (sources []string) {
	key := fmt.Sprintf("models.image.cachedfallbackimage:%s", size)
	s := reflect.ValueOf(cache.Client.Get(key))

	if !cache.Client.IsExist(key) {
		var q string

		excludes := `
		src not like '%xvideos.%' AND 
		src not like '%erovideo.%' AND 
		src not like '%redtube.%' AND 
		src not like '%xhamster.%' AND 
		src not like '%videomega.%' AND 
		src not like '%fc2.png' AND 
		src not like '%fc2.jpg'
		`

		if size == "small" {
			q = fmt.Sprintf(`SELECT src FROM image WHERE width < 300 AND width > 100 AND %s ORDER BY id DESC LIMIT 3000`, excludes)
			// WHERE width < 300 AND width > 100
		} else if size == "middle" {
			q = fmt.Sprint(`SELECT src FROM image WHERE width < 500 AND width > 300 AND  %s ORDER BY id DESC LIMIT 3000`, excludes)
			// width < 500 AND width > 300
		} else {
			q = fmt.Sprint(`SELECT src FROM image WHERE width > 500 AND				 	 %s ORDER BY id DESC LIMIT 3000`, excludes)
			// width > 500
		}

		var list orm.ParamsList
		orm.NewOrm().Raw(q).ValuesFlat(&list, "src")

		for _, src := range list {
			sources = append(sources, src.(string))
		}

		bytes, _ := json.Marshal(sources)
		cache.Client.Put(key, bytes, 60*60*24*1)
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
