package book

import (
	_ "bitbucket.org/ikeikeikeike/antenna/routers"

	"bitbucket.org/ikeikeikeike/antenna/controllers"
	"bitbucket.org/ikeikeikeike/antenna/controllers/book"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &book.EntriesController{}, "get:Home")

	beego.Router("/news.html", &book.EntriesController{}, "get:News")
	beego.Router("/hots.html", &book.EntriesController{}, "get:Hots")

	beego.Router(`/book/v:id([0-9]+)/`, &book.EntriesController{}, "get:Show")
	beego.Router(`/book/v:id([0-9]+)/:title(.*)?`, &book.EntriesController{}, "get:Show")
	beego.Router(`/book/viewer:id([0-9]+)/:title(.*)?`, &book.EntriesController{}, "get:Viewer")

	beego.Router("/ranking/dayly", &book.RankingsController{}, "get:Dayly")
	// beego.Router("/ranking/weekly", &book.RankingsController{}, "get:Weekly")
	// beego.Router("/ranking/monthly", &book.RankingsController{}, "get:Monthly")
	// beego.Router("/ranking/yearly", &book.RankingsController{}, "get:Yearly")

	beego.Router("/search.html", &book.EntriesController{}, "get:Search")

	beego.Router(`/blog/v:id([0-9]+)/`, &book.BlogsController{}, "get:Show")
	beego.Router(`/blog/v:id([0-9]+)/:title(.*)?`, &book.BlogsController{}, "get:Show")

	beego.Router("/tags.html", &book.TagsController{}, "get:Index")
	beego.Router(`/tag/:name`, &book.TagsController{}, "get:Show")

	beego.Router("/divas.html", &book.DivasController{}, "get:Index")
	beego.Router(`/diva/:name`, &book.DivasController{}, "get,post:Show")

	beego.Router("/animes.html", &book.AnimesController{}, "get:Index")
	beego.Router(`/anime/:name`, &book.AnimesController{}, "get,post:Show")

	beego.Router("/characters.html", &book.CharactersController{}, "get:Index")
	beego.Router(`/character/:name`, &book.CharactersController{}, "get,post:Show")

	beego.Router(`/feed/rdf.xml`, &book.FeedsController{}, "get:Rdf")
	beego.Router(`/feed/rss.xml`, &book.FeedsController{}, "get:Rss")
	beego.Router(`/feed/atom.xml`, &book.FeedsController{}, "get:Atom")

	// ext
	beego.Router(`/ok.ico`, &controllers.OkController{}, "get:Helthcheck")
}
