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

	// 查询下部的商品
	var goodsSkus = make([]map[string]interface{}, len(goodTypes))

	// 查询商品类型
	for index, _ := range goodsSkus {
		temp := make(map[string]interface{})
		temp["types"] = goodTypes[index]
		goodsSkus[index] = temp
	}

	// 查询商品
	for _, goodsMap := range goodsSkus {
		var goodsImage []models.IndexTypeGoodsBanner
		var goodsText []models.IndexTypeGoodsBanner
		ormer.QueryTable("IndexTypeGoodsBanner").RelatedSel("GoodsType", "GoodsSKU").Filter("GoodsType", goodsMap["types"]).Filter("DisplayType", 1).All(&goodsImage)
		ormer.QueryTable("IndexTypeGoodsBanner").RelatedSel("GoodsType", "GoodsSKU").Filter("GoodsType", goodsMap["types"]).Filter("DisplayType", 0).All(&goodsText)

		goodsMap["goodsImage"] = goodsImage
		goodsMap["goodsText"] = goodsText
	}

	this.Data["goodsSkus"] = goodsSkus

	this.Data["goodTypes"] = goodTypes
	this.Data["indexGoodsBanners"] = indexGoodsBanners
	this.Data["indexPromotionBanners"] = indexPromotionBanners

	this.TplName = "index.html"
}
