package book

import (
	_ "bitbucket.org/ikeikeikeike/antenna/routers"

	"bitbucket.org/ikeikeikeike/antenna/controllers/book"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &book.EntriesController{}, "get:Home")

	beego.Router("/news.html", &book.EntriesController{}, "get:News")
	beego.Router("/hots.html", &book.EntriesController{}, "get:Hots")

	beego.Router("/search.html", &book.EntriesController{}, "get:Search")
}
