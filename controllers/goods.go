package controllers

import "github.com/astaxie/beego"

type GoodsController struct {
	beego.Controller
}

// 显示主页
func (this *GoodsController) ShowIndex() {
	// 从session中获取用户,并传递给主页
	username := this.GetSession("userName")
	if username == nil {
		this.Data["username"] = ""
	} else {
		this.Data["username"] = username.(string)
	}

	this.TplName = "index.html"
}
