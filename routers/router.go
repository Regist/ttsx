package routers

import (
	"github.com/astaxie/beego"
	"ttsx/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// 注册路由
	beego.Router("/register", &controllers.UserController{}, "get:ShowRegister;post:HandleRegister")
	// 激活路由
	beego.Router("/activie", &controllers.UserController{}, "get:HandleActive")
}
