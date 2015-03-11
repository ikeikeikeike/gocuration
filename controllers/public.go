package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	actl "bitbucket.org/ikeikeikeike/antenna/lib/accessctl"
	"bitbucket.org/ikeikeikeike/antenna/models/entry"

	"github.com/astaxie/beego/orm"
	"github.com/ikeikeikeike/gopkg/accessctl"
	"github.com/ikeikeikeike/gopkg/convert"
)

type PublicController struct {
	BaseController

	DefaultPers int
}

func (c *PublicController) Prepare() {
	c.BaseController.Prepare()

	c.Data["QURL"] = c.UrlFor("EntriesController.Home")
	c.Data["Q"] = strings.TrimSpace(c.GetString("q"))

	c.Data["EntryMaxLength"] = entry.CachedMaxLength()

	c.DefaultPers = 20
	if c.Meta.UA.Category == "pc" {
		c.DefaultPers = 25
	}
}

func (c *PublicController) PushInAccessLog() (err error) {
	if !strings.Contains(c.Ctx.Request.Header.Get("Accept"), "html") {
		return errors.New("Does not html request, Maybe request to css,js")
	}

	inlog := &actl.InLog{
		RemoteHost: c.Ctx.Input.IP(),
		Time:       time.Now().Add(time.Duration(9) * time.Hour).Format(actl.Timeformat),
		Method:     c.Ctx.Input.Method(),
		Url:        c.BuildRequestUrl(""),
		Status:     strconv.Itoa(c.Ctx.Output.Status),
		Referer:    c.Ctx.Input.Refer(),
		UserAgent:  c.Ctx.Input.UserAgent(),
	}

	if !actl.IsInAccessLog(inlog) {
		return errors.New("could not push access log")
	}

	counter, err := accessctl.NewCounter()
	if err != nil {
		return
	}

	str, _ := json.Marshal(*inlog)
	err = counter.Push(string(str))
	return
}

func (c *PublicController) SetQ(qs orm.QuerySeter, q string) orm.QuerySeter {
	v := c.GetString("q")
	if v != "" {
		for _, word := range convert.StrTo(v).MultiWord() {
			qs = qs.Filter(q+"q__icontains", word)
		}
	}
	return qs
}

func (c *PublicController) SetMt(qs orm.QuerySeter, q string) orm.QuerySeter {
	v := c.GetString("mt")
	if v != "" {
		qs = qs.Filter(q+"blog__mediatype", v)
	}
	return qs
}

func (c *PublicController) SetAt(qs orm.QuerySeter, q string) orm.QuerySeter {
	v := c.GetString("at")
	if v != "" {
		qs = qs.Filter(q+"blog__adsensetype", v)
	}
	return qs
}

var PrefixLines = map[string][]string{
	"あ": []string{"a", "i", "u", "e", "o"},
	"か": []string{"ka", "ki", "ku", "ke", "ko"},
	"さ": []string{"sa", "si", "su", "se", "so"},
	"た": []string{"ta", "ti", "tu", "te", "to"},
	"な": []string{"na", "ni", "nu", "ne", "no"},
	"は": []string{"ha", "hi", "hu", "he", "ho"},
	"ま": []string{"ma", "mi", "mu", "me", "mo"},
	"や": []string{"ya", "yu", "yo"},
	"ら": []string{"ra", "ri", "ru", "re", "ro"},
	"わ": []string{"wa"},
}

func (c *PublicController) SetPrefixLines(qs orm.QuerySeter, q string) orm.QuerySeter {
	v := c.GetString("line")
	if v != "" {
		qs = qs.Filter(q+"gyou__in", PrefixLines[v])
	}
	return qs
}

var BracupLines = map[string][]string{
	"C": []string{"AAA", "AA", "A", "B", "C"},
	"D": []string{"D"},
	"E": []string{"E"},
	"F": []string{"F"},
	"G": []string{"G"},
	"H": []string{"H"},
	"I": []string{"I", "J", "K", "L", "M", "N", "O", "P"},
}

func (c *PublicController) SetBracupLines(qs orm.QuerySeter) orm.QuerySeter {
	v := c.GetString("cup")
	if v != "" {
		qs = qs.Filter("bracup__in", BracupLines[v])
	}
	return qs
}

func (c *PublicController) SetBracup(qs orm.QuerySeter, q string) orm.QuerySeter {
	v := c.GetStrings("cup")
	if len(v) > 0 {
		qs = qs.Filter(q+"bracup__in", v)
	}
	return qs
}

func (c *PublicController) SetBlood(qs orm.QuerySeter, q string) orm.QuerySeter {
	v := c.GetString("blood")
	if v != "" {
		qs = qs.Filter(q+"blood", v)
	}
	return qs
}

func (c *PublicController) SetAdvancedSearch(qs orm.QuerySeter, prefix string) orm.QuerySeter {
	qs = c.SetBlood(qs, prefix+"video__divas__diva__")
	qs = c.SetBracup(qs, prefix+"video__divas__diva__")
	qs = c.SetPrefixLines(qs, prefix+"video__divas__diva__")
	qs = c.SetAt(c.SetMt(c.SetQ(qs, prefix), prefix), prefix)
	return qs
}
