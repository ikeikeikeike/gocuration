package main

import (
	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	_ "bitbucket.org/ikeikeikeike/antenna/routers/antenna"

	"github.com/astaxie/beego"
)

func main() {
	// 	if beego.RunMode == "dev" {
	// 		beego.DirectoryIndex = true
	// 		beego.StaticDir["/swagger"] = "swagger"
	// 	}

	beego.Run()
}
