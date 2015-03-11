package cache

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var Client cache.Cache

func init() {
	var err error

	if beego.RunMode == "test" {
		Client, err = cache.NewCache("memory", `{"interval":60}`)
	} else {
		conn := beego.AppConfig.String("RedisConn")
		Client, err = cache.NewCache("redis", fmt.Sprintf(`{"conn": "%s", "interval": 60}`, conn))
	}

	if err != nil {
		panic(err)
	}
}
