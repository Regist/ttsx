package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"regexp"
	"ttsx/models"
)

type UserController struct {
	beego.Controller
}

// 显示页面
func (this *UserController) ShowRegister() {
	this.TplName = "register.html"
}

// 处理注册请求
func (this *UserController) HandleRegister() {

	// 获取数据
	username := this.GetString("user_name")
	pwd := this.GetString("pwd")
	cpwd := this.GetString("cpwd")
	email := this.GetString("email")

	// 校验数据
	if username == "" || pwd == "" || cpwd == "" || email == "" {
		this.Data["error"] = "数据不能为空"
		this.TplName = "register.html"
		return
	}

	// 设置校验邮箱的正则表达式
	compile, _ := regexp.Compile("^[a-z0-9A-Z]+[- | a-z0-9A-Z . _]+@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)?\\.)+[a-z]{2,}$")
	// 校验邮箱
	// 校验成功,返回值就是获取到的邮箱
	// 校验失败,返回值为空
	compileEmailResult := compile.FindString(email)

	if compileEmailResult == "" {
		this.Data["error"] = "邮箱格式不正确,请重新输入"
		this.TplName = "register.html"
		return
	}

	// 处理数据
	ormer := orm.NewOrm()
	// 封装数据
	var user models.User
	user.Name = username
	user.PassWord = pwd
	user.Email = email
	// 插入数据
	_, err := ormer.Insert(&user)
	if err != nil {
		this.Data["error"] = "注册失败,请重试"
		this.TplName = "register.html"
		return
	}

	// 插入成功
	this.Ctx.WriteString("注册成功")

}
