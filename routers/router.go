package routers

import (
	"github.com/astaxie/beego"
	"ttsx/controllers"
)

func init() {
	// 展示主页
	beego.Router("/", &controllers.GoodsController{}, "get:ShowIndex")
	// 注册路由
	beego.Router("/register", &controllers.UserController{}, "get:ShowRegister;post:HandleRegister")
	// 激活路由
	beego.Router("/activie", &controllers.UserController{}, "get:HandleActive")
	// 展示登陆页面
	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")
}
