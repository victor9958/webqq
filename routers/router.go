package routers

import (
	"webqq/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/admin", &controllers.UserController{})
    beego.Router("/login", &controllers.LoginController{})
    beego.Router("/index", &controllers.IndexController{})
    beego.Router("/reptile", &controllers.ReptileController{})
	beego.Router("/workers", &controllers.WorkersController{})
	beego.Router("/weblist", &controllers.WeblistController{})
}
