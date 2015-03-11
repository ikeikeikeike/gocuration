package controllers

import (
	"fmt"

	"github.com/ikeikeikeike/gopkg/convert"
	// "github.com/k0kubun/pp"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller

	Meta *Meta
}

type NestPreparer interface {
	NestPrepare()
}

type NestFinisher interface {
	NestFinish()
}

func (c *BaseController) Prepare() {
	c.SetParams()

	c.Data["HeadMetas"] = []map[string]string{}
	c.Data["HeadStyles"] = []string{}
	c.Data["HeadScripts"] = []string{}
	c.Data["HeadBullets"] = []string{}

	c.Data["Ctx"] = c.Ctx

	c.Meta = NewMeta()
	c.Meta.Init(c)
	c.Data["Meta"] = c.Meta

	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

func (c *BaseController) Finish() {
	if app, ok := c.AppController.(NestFinisher); ok {
		app.NestFinish()
	}
}

func (c *BaseController) SetParams() {
	c.Data["Params"] = make(map[string]string)
	for k, v := range c.Input() {
		c.Data["Params"].(map[string]string)[k] = v[0]
	}
}

func (c *BaseController) BuildRequestUrl(uri string) string {
	if uri == "" {
		uri = c.Ctx.Input.Uri()
	}
	return fmt.Sprintf("%s:%s%s",
		c.Ctx.Input.Site(), convert.ToStr(c.Ctx.Input.Port()), uri)
}
