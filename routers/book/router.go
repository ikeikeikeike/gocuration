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

	beego.Router(`/book/b:id([0-9]+)/`, &book.EntriesController{}, "get:Show")
	beego.Router(`/book/b:id([0-9]+)/:title(.*)?`, &book.EntriesController{}, "get:Show")

	beego.Router("/search.html", &book.EntriesController{}, "get:Search")
}
