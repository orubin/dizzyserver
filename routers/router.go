package routers

import (
	"github.com/astaxie/beego"
	"github.com/orubin/dizzyserver/controllers"
	"github.com/orubin/dizzyserver/controllers/api"
)

func init() {
	api := beego.NewNamespace("/api", beego.NSInclude(
		&api.APIController{},
	))

	beego.AddNamespace(api)

	beego.Router("/", &controllers.MainController{})
}
