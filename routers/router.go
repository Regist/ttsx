package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"ttsx/controllers"
)

func init() {
	// 设置路由
	beego.InsertFilter("/user/*", // 过滤规则
		beego.BeforeExec, // 执行时机,访问路由之后执行controller之前
		filterFunc)       // 过滤器函数
	// 展示主页
	beego.Router("/", // 请求路径
		&controllers.GoodsController{}, //控制器
		"get:ShowIndex")                // 请求方式和映射的方法
	// 注册路由
	beego.Router("/register", &controllers.UserController{}, "get:ShowRegister;post:HandleRegister")
	// 激活路由
	beego.Router("/activie", &controllers.UserController{}, "get:HandleActive")
	// 展示登陆页面
	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")
	// 退出登陆
	beego.Router("/user/logout", &controllers.UserController{}, "get:Logout")
}

// 检查登陆状态的过滤器
var filterFunc = func(ctx *context.Context) {
	// 获取session
	username := ctx.Input.Session("userName")
	if username == nil {
		ctx.Redirect(302, "/login")
	}

}
