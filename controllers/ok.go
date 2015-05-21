package controllers

import "github.com/astaxie/beego"

type OkController struct {
	beego.Controller
}

func (c *OkController) Helthcheck() {
	c.Ctx.Output.Header("Content-Type", "image/gif")
	c.Ctx.Output.Body([]byte(""))
}
