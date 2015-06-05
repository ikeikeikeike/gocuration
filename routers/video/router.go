package video

import (
	_ "bitbucket.org/ikeikeikeike/antenna/routers"

	"bitbucket.org/ikeikeikeike/antenna/controllers"
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

	beego.Router(`/blog/v:id([0-9]+)/`, &video.BlogsController{}, "get:Show")
	beego.Router(`/blog/v:id([0-9]+)/:title(.*)?`, &video.BlogsController{}, "get:Show")

	beego.Router("/ranking.html", &video.RankingsController{}, "get:Index")

	beego.Router("/tags.html", &video.TagsController{}, "get:Index")
	beego.Router(`/tag/:name`, &video.TagsController{}, "get:Show")

	beego.Router("/divas.html", &video.DivasController{}, "get:Index")
	beego.Router(`/diva/:name`, &video.DivasController{}, "get,post:Show")

	beego.Router("/animes.html", &video.AnimesController{}, "get:Index")
	beego.Router(`/anime/:name`, &video.AnimesController{}, "get,post:Show")

	beego.Router("/characters.html", &video.CharactersController{}, "get:Index")
	beego.Router(`/character/:name`, &video.CharactersController{}, "get,post:Show")

	beego.Router(`/feed/rdf.xml`, &video.FeedsController{}, "get:Rdf")
	beego.Router(`/feed/rss.xml`, &video.FeedsController{}, "get:Rss")
	beego.Router(`/feed/atom.xml`, &video.FeedsController{}, "get:Atom")

	// ext
	beego.Router(`/ok.ico`, &controllers.OkController{}, "get:Helthcheck")
}
