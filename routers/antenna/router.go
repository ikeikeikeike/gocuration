package antenna

import (
	"bitbucket.org/ikeikeikeike/antenna/controllers/admin"
	"bitbucket.org/ikeikeikeike/antenna/controllers/public"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"
	"github.com/astaxie/beego"
)

func init() {
	// Public

	beego.Router("/", &public.EntriesController{}, "get:Home")

	beego.Router("/news.html", &public.EntriesController{}, "get:News")
	beego.Router("/hots.html", &public.EntriesController{}, "get:Hots")

	beego.Router("/search.html", &public.EntriesController{}, "get:Search")

	beego.Router(`/elog/v:id([0-9]+)/`, &public.EntriesController{}, "get:Show")
	beego.Router(`/elog/v:id([0-9]+)/:title(.*)?`, &public.EntriesController{}, "get:Show")

	beego.Router("/blogs.html", &public.BlogsController{}, "get:Index")
	beego.Router(`/blog/v:id([0-9]+)/`, &public.BlogsController{}, "get:Show")
	beego.Router(`/blog/v:id([0-9]+)/:title(.*)?`, &public.BlogsController{}, "get:Show")

	beego.Router("/tags.html", &public.TagsController{}, "get:Index")
	beego.Router(`/tag/:name`, &public.TagsController{}, "get:Show")

	beego.Router("/divas.html", &public.DivasController{}, "get:Index")
	beego.Router(`/diva/:name`, &public.DivasController{}, "get,post:Show")

	beego.Router("/animes.html", &public.AnimesController{}, "get:Index")
	beego.Router(`/anime/:name`, &public.AnimesController{}, "get,post:Show")

	beego.Router("/characters.html", &public.CharactersController{}, "get:Index")
	beego.Router(`/character/:name`, &public.CharactersController{}, "get,post:Show")

	beego.Router(`/register.html`, &public.RegistersController{}, "get:Index")
	beego.Router(`/register/parts.html`, &public.RegistersController{}, "get:Parts")
	beego.Router(`/register/links.html`, &public.RegistersController{}, "get:Links")
	beego.Router(`/register/rss.html`, &public.RegistersController{}, "get:Rss")

	beego.Router(`/feed/rdf.xml`, &public.FeedsController{}, "get:Rdf")
	beego.Router(`/feed/rss.xml`, &public.FeedsController{}, "get:Rss")
	beego.Router(`/feed/atom.xml`, &public.FeedsController{}, "get:Atom")

	beego.Router(`/v1/parts.js`, &public.ApisController{}, "post:Parts")

	// Admin

	beego.Router("/admin/signup", &admin.LoginController{}, "get,post:Signup")
	beego.Router("/admin/signup_confirm", &admin.LoginController{}, "get:SignupConfirm")
	beego.Router("/admin/login", &admin.LoginController{}, "get,post:Login")
	beego.Router("/admin/logout", &admin.LoginController{}, "get:Logout")
	beego.Router("/admin/password_reset", &admin.LoginController{}, "get,post:PasswordReset")
	beego.Router("/admin/reset_confirm", &admin.LoginController{}, "get,post:ResetConfirm")

	beego.Router(`/admin`, &admin.UsersController{}, "get:Index")
	beego.Router(`/admin/blog`, &admin.UsersController{}, "get,post:BlogRegister")
	beego.Router(`/admin/blog/:id([0-9]+)`, &admin.UsersController{}, "get,post:Blog")
	beego.Router(`/admin/blogs`, &admin.UsersController{}, "get:Blogs")
	beego.Router(`/admin/change_email`, &admin.UsersController{}, "get,post:ChangeMail")
	beego.Router("/admin/mail_confirm", &admin.UsersController{}, "get:MailConfirm")

	beego.Router("/admin/blog_remove", &admin.UsersController{}, "delete:BlogRemove")
	beego.Router("/admin/account_remove", &admin.UsersController{}, "delete:AccountRemove")
	// beego.Router(`/admin/change_password`, &admin.UsersController{}, "get,post:ChangePassword")

	beego.Router(`/admin/terms.html`, &admin.LicensesController{}, "get:Terms")
	beego.Router(`/admin/privacy.html`, &admin.LicensesController{}, "get:Privacy")
	beego.Router(`/admin/guideline.html`, &admin.LicensesController{}, "get:Guideline")
}
