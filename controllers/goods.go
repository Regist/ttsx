package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"ttsx/models"
)

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

	ormer := orm.NewOrm()
	// 查询主页分类数据
	var goodTypes []models.GoodsType
	ormer.QueryTable("GoodsType").All(&goodTypes)

	// 查询主页轮播图
	var indexGoodsBanners []models.IndexGoodsBanner
	ormer.QueryTable("IndexGoodsBanner").OrderBy("Index").All(&indexGoodsBanners)

	// 查询主页推广
	var indexPromotionBanners []models.IndexPromotionBanner
	ormer.QueryTable("IndexPromotionBanner").OrderBy("Index").All(&indexPromotionBanners)

	this.Data["goodTypes"] = goodTypes
	this.Data["indexGoodsBanners"] = indexGoodsBanners
	this.Data["indexPromotionBanners"] = indexPromotionBanners

	this.TplName = "index.html"
}
