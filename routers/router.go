package routers

import (
	"beego_forxiner/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/checktoken", &controllers.AutoReturnController{}, "post:AutoReturnHandler")
}
