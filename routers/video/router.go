package video

import (
	_ "bitbucket.org/ikeikeikeike/antenna/routers"

	"bitbucket.org/ikeikeikeike/antenna/controllers/video"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &video.EntriesController{}, "get:Home")

	beego.Router("/news.html", &video.EntriesController{}, "get:News")
	beego.Router("/hots.html", &video.EntriesController{}, "get:Hots")

	beego.Router(`/video/v:id([0-9]+)/`, &video.EntriesController{}, "get:Show")
	beego.Router(`/video/v:id([0-9]+)/:title(.*)?`, &video.EntriesController{}, "get:Show")

	beego.Router("/search.html", &video.EntriesController{}, "get:Search")
}