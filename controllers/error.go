package controllers

import (
	"fmt"

	"bitbucket.org/ikeikeikeike/antenna/lib/mailer"
	"github.com/astaxie/beego/context"
)

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error400() {
	mailer.SendStackTrace(traceMessage(c.Ctx))
	c.TplNames = "errors/400.tpl"
}

func (c *ErrorController) Error401() {
	mailer.SendStackTrace(traceMessage(c.Ctx))
	c.TplNames = "errors/401.tpl"
}

func (c *ErrorController) Error403() {
	mailer.SendStackTrace(traceMessage(c.Ctx))
	c.TplNames = "errors/403.tpl"
}

func (c *ErrorController) Error404() {
	c.TplNames = "errors/404.tpl"
}

func (c *ErrorController) Error500() {
	mailer.SendStackTrace(traceMessage(c.Ctx))
	c.TplNames = "errors/500.tpl"
}

func (c *ErrorController) Error503() {
	mailer.SendStackTrace(traceMessage(c.Ctx))
	c.TplNames = "errors/503.tpl"
}

func (c *ErrorController) ErrorDb() {
	mailer.SendStackTrace(traceMessage(c.Ctx))
	c.TplNames = "errors/dberror.tpl"
}

func traceMessage(ctx *context.Context) []string {
	return []string{
		fmt.Sprintf("Url: %s", ctx.Input.Url()),
		fmt.Sprintf("Uri: %s", ctx.Input.Uri()),
		fmt.Sprintf("Method: %s", ctx.Input.Method()),
		fmt.Sprintf("Params: %v", ctx.Input.Params),
		fmt.Sprintf("UserAgent: %s", ctx.Input.UserAgent()),
		"",
	}
}
