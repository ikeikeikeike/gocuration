package main

import (
	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	_ "bitbucket.org/ikeikeikeike/antenna/routers/book"

	"github.com/astaxie/beego"
)

func main() {
	beego.AppName = "book"
	beego.HttpPort = 8888
	beego.EnableAdmin = false
	beego.Run()
}
