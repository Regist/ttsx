package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils"
	"regexp"
	"strconv"
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

	//发送邮件
	//设置发送人参数,具体设置请参考对应的邮件提供商
	emailConfig := `{"username":"alphatest1001@163.com","password":"alphatest1001","host":"smtp.163.com","port":25}`
	mail := utils.NewEMail(emailConfig)
	// 设置发件人
	mail.From = "alphatest1001@163.com"
	// 设置收件人
	mail.To = []string{email}
	// 设置邮件主题
	mail.Subject = "激活账号"
	// 设置邮件内容,普通文笔
	mail.Text = "感谢您注册天天生鲜,请复制链接到地址栏进行激活,http://localhost:8080/activie?id=" + strconv.Itoa(user.Id)
	// 设置邮件内容, html格式, 部分邮箱提供商可能会屏蔽超链接格式内容
	//	mail.HTML = "<a href=\"感谢您注册天天生鲜,请复制链接到地址栏进行激活,http://localhost:8080/activie?id=" + strconv.Itoa(user.Id) + ">点击激活</a>"

	// 发送邮件
	err = mail.Send()
	if err != nil {
		beego.Error(err)
	}

	// 插入成功
	this.Ctx.WriteString("注册成功,请查收激活邮件")

}

// 处理激活
func (this *UserController) HandleActive() {
	// 获取ID
	id, err := this.GetInt("id")
	if err != nil {
		this.Data["error"] = "激活失败,请重新注册"
		this.TplName = "register.html"
		return
	}

	ormer := orm.NewOrm()

	// 根据ID查询用户是否存在
	var user models.User
	user.Id = id
	err = ormer.Read(&user)
	if err != nil {
		this.Data["error"] = "激活失败,请重新注册"
		this.TplName = "register.html"
		return
	}

	// 激活用户
	user.Active = true
	_, err = ormer.Update(&user)
	if err != nil {
		this.Data["error"] = "激活失败,请重新注册"
		this.TplName = "register.html"
		return
	}

	// 激活成功,跳转到登陆界面
	this.Redirect("/login", 302)
}

// 展示登陆页面
func (this *UserController) ShowLogin() {
	// 从cookie中获取数据
	username := this.Ctx.GetCookie("username")
	if username == "" {
		// cookie中没有数据,清空数据
		this.Data["username"] = ""
		this.Data["checked"] = ""
	} else {
		// cookie中有数据,设置数据
		this.Data["username"] = username
		this.Data["checked"] = "checked"
	}

	this.TplName = "login.html"
}

// 处理登陆
func (this *UserController) HandleLogin() {
	// 获取用户输入的内容
	usr := this.GetString("username")
	psw := this.GetString("pwd")

	// 检查用户名和密码是否为空
	if usr == "" || psw == "" {
		this.Data["error"] = "用户名密码不能为空"
		this.TplName = "login.html"
		return
	}

	// 查询
	ormer := orm.NewOrm()
	var user models.User
	user.Name = usr
	user.PassWord = psw
	// 根据用户名查询用户是否存在
	err := ormer.Read(&user, "Name")
	if err != nil {
		this.Data["error"] = "用户不存在,请重新登陆"
		this.TplName = "login.html"
		return
	}

	// 判断用户是否激活
	if !user.Active {
		this.Data["error"] = "用户未激活,请激活"
		this.TplName = "login.html"
		return
	}

	// 比对密码
	if user.PassWord != psw {
		this.Data["error"] = "密码错误,请重试"
		this.TplName = "login.html"
		return
	}

	// 判断用户是否勾选了记住用户名
	check := this.GetString("check")
	beego.Info(check)
	// 勾选了,记住一小时
	if check == "on" {
		this.Ctx.SetCookie("username", usr, 3600)
	} else {
		// 没有勾选,清除数据
		this.Ctx.SetCookie("username", usr, -1)
	}

	// 跳转到主页
	this.Redirect("/", 302)
}
