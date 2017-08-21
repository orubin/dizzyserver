package routers

import (
	"github.com/astaxie/beego"
	"github.com/orubin/dizzyserver/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
