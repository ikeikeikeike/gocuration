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

	mtype := c.GetString("media_type")
	adtype := c.GetString("adsense_type")

	// TODO: リスティング広告も取り入れAccessトレードも考慮する形に変更
	//
	var sidx []*models.EntryIndex

	key := fmt.Sprintf(
		"controller.public.apis pers:%d:%s:%s", pers, mtype, adtype)
	s := reflect.ValueOf(cache.Client.Get(key))

	if !cache.Client.IsExist(key) {

		qs := models.Summaries().RelatedSel()
		if mtype != "" {
			qs = qs.Filter("entry__blog__mediatype", mtype)
		}
		if adtype != "" {
			qs = qs.Filter("entry__blog__adsensetype", adtype)
		}

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
