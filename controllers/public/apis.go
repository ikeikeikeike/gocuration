package public

import (
	"encoding/json"
	"fmt"
	"reflect"

	"bitbucket.org/ikeikeikeike/antenna/lib/cache"
	"bitbucket.org/ikeikeikeike/antenna/models"
	"github.com/ikeikeikeike/gopkg/convert"
)

// "github.com/k0kubun/pp"

type ApisController struct {
	BaseController
}

func (c *ApisController) NestPrepare() {
	c.EnableXSRF = false
}

func (c *ApisController) Parts() {
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	c.Ctx.Output.Header("Access-Control-Allow-Methods", "POST,OPTIONS")

	pers, err := convert.StrTo(c.GetString("per_item")).Int()
	if err != nil {
		c.Ctx.WriteString("parts does not exists")
		return
	}
	if pers > 50 {
		pers = 50
	}

	// TODO: リスティング広告も取り入れつつ、Accessトレードも考慮する形に変更
	//
	var sidx []*models.EntryIndex

	key := fmt.Sprintf("controller.public.apis pers:%d", pers)
	s := reflect.ValueOf(cache.Client.Get(key))

	if !cache.Client.IsExist(key) {
		qs := models.Summaries().RelatedSel()
		// qs = c.SetQ(qs, "")

		var summes []*models.Summary
		qs.Limit(pers).All(&summes)

		for _, s := range summes {
			s.LoadRelated()
			s.Entry.LoadRelated()

			sidx = append(sidx, s.Entry.SearchData())
		}

		bytes, _ := json.Marshal(sidx)
		cache.Client.Put(key, bytes, 60*30)
	} else {
		json.Unmarshal(s.Interface().([]uint8), &sidx)
	}

	c.Data["json"] = sidx
	c.ServeJson()
}
