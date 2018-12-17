package controllers

import "github.com/astaxie/beego"

type GoodsController struct {
	beego.Controller
}

// 显示主页
func (this *GoodsController) ShowIndex() {
	this.TplName = "index.html"
}
