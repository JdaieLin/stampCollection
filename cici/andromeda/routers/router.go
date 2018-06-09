package routers

import (
	"cici/andromeda/controllers"

	"github.com/astaxie/beego"
)

func InitRouters() {
	beego.Router("/", &controllers.MainController{}, "get:Get")
	beego.AutoPrefix("/api", &controllers.UserController{})
	beego.AutoPrefix("/api", &controllers.StampController{})
	beego.AutoPrefix("/api", &controllers.ChestController{})
	beego.AutoPrefix("/api", &controllers.SweepController{})
	beego.AutoPrefix("/api", &controllers.DealController{})
	beego.AutoPrefix("/api", &controllers.DebugController{})
}
