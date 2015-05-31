package verify

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/ikeikeikeike/gopkg/extract/image"
)

func AvailableImageChecker() *image.Info {
	info := image.NewInfo()
	info.Header("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	info.Header("Accept-Encoding", "gzip, deflate, sdch")
	info.Header("Accept-Language", "ja,en-US;q=0.8,en;q=0.6,zh;q=0.4,zh-TW;q=0.2,ko;q=0.2,es;q=0.2,ru;q=0.2")
	info.Header("Cache-Control", "no-cache")
	info.Header("Connection", "keep-alive")
	info.Header("Pragma", "no-cache")
	info.Header("Referer", fmt.Sprintf("http://%s", beego.AppConfig.String("domain"))) // Own host(domain) referer.
	info.Header("User-Agent", beego.AppConfig.String("UserAgent"))
	return info
}
