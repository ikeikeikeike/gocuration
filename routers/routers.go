package routers

import (
	"bitbucket.org/ikeikeikeike/antenna/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// Errors
	beego.ErrorController(&controllers.ErrorController{})

	// var genErrHandler = func(status string) http.HandlerFunc {
	// return func(rw http.ResponseWriter, r *http.Request) {
	// t, _ := template.ParseFiles(beego.ViewsPath + "/errors/" + status + ".tpl")
	// data := make(map[string]interface{})
	// data["Meta"] = controllers.NewMeta()
	// t.Execute(rw, data)
	// }
	// }
	// beego.Errorhandler("400", genErrHandler("400"))
	// beego.Errorhandler("401", genErrHandler("401"))
	// beego.Errorhandler("403", genErrHandler("403"))
	// beego.Errorhandler("404", genErrHandler("404"))
	// beego.Errorhandler("500", genErrHandler("500"))
	// beego.Errorhandler("503", genErrHandler("503"))
}
