package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//用户表
type User struct {
	Id       int
	Name     string     `orm:"size(20);unique"` //用户名
	PassWord string     `orm:"size(20)"`        //登陆密码
	Email    string     `orm:"size(50)"`        //邮箱
	Active   bool       `orm:"default(false)"`  //是否激活
	Power    int        `orm:"default(0)"`      //权限设置  0 表示普通用户  1表示管理员用户
	Address  []*Address `orm:"reverse(many)"`   //收货地址
}

//地址表
type Address struct {
	Id        int
	Receiver  string `orm:"size(20)"`       //收件人
	Addr      string `orm:"size(50)"`       //收件地址
	Zipcode   string `orm:"size(20)"`       //邮编
	Phone     string `orm:"size(20)"`       //联系方式
	Isdefault bool   `orm:"default(false)"` //是否默认 false 为非默认  true为默认
	User      *User  `orm:"rel(fk)"`        //用户ID `
}

//商品SPU表
type Goods struct {
	Id       int
	Name     string      `orm:"size(20)"`  //商品名称
	Detail   string      `orm:"size(200)"` //详细描述
	GoodsSKU []*GoodsSKU `orm:"reverse(many)"`
}

//商品类型表
type GoodsType struct {
	Id                   int
	Name                 string                  //种类名称
	Logo                 string                  //logo
	Image                string                  //图片
	GoodsSKU             []*GoodsSKU             `orm:"reverse(many)"`
	IndexTypeGoodsBanner []*IndexTypeGoodsBanner `orm:"reverse(many)"`
}

//商品SKU表
type GoodsSKU struct {
	Id                   int
	Goods                *Goods                  `orm:"rel(fk)"` //商品SPU
	GoodsType            *GoodsType              `orm:"rel(fk)"` //商品所属种类
	Name                 string                  //商品名称
	Desc                 string                  //商品简介
	Price                int                     //商品价格
	Unite                string                  //商品单位
	Image                string                  //商品图片
	Stock                int                     `orm:"default(1)"`   //商品库存
	Sales                int                     `orm:"default(0)"`   //商品销量
	Status               int                     `orm:"default(1)"`   //商品状态
	Time                 time.Time               `orm:"auto_now_add"` //添加时间
	GoodsImage           []*GoodsImage           `orm:"reverse(many)"`
	IndexGoodsBanner     []*IndexGoodsBanner     `orm:"reverse(many)"`
	IndexTypeGoodsBanner []*IndexTypeGoodsBanner `orm:"reverse(many)"`
}

//商品图片表
type GoodsImage struct {
	Id       int
	Image    string    //商品图片
	GoodsSKU *GoodsSKU `orm:"rel(fk)"` //商品SKU
}

//首页轮播商品展示表
type IndexGoodsBanner struct {
	Id       int
	GoodsSKU *GoodsSKU `orm:"rel(fk)"` //商品sku
	Image    string    //商品图片
	Index    int       `orm:"default(0)"` //展示顺序
}

//首页分类商品展示表
type IndexTypeGoodsBanner struct {
	Id          int
	GoodsType   *GoodsType `orm:"rel(fk)"`    //商品类型
	GoodsSKU    *GoodsSKU  `orm:"rel(fk)"`    //商品sku
	DisplayType int        `orm:"default(1)"` //展示类型 0代表文字，1代表图片
	Index       int        `orm:"default(0)"` //展示顺序
}

//首页促销商品展示表
type IndexPromotionBanner struct {
	Id    int
	Name  string `orm:"size(20)"` //活动名称
	Url   string `orm:"size(50)"` //活动链接
	Image string //活动图片
	Index int    `orm:"default(0)"` //展示顺序
}

func init() {
	// 注册数据库
	orm.RegisterDataBase("default", //别名
		"mysql", //驱动
		"root:root@tcp(127.0.0.1:3306)/ttsx?charset=utf8") // 连接字符串/数据源
	// 注册模型
	orm.RegisterModel(new(User), new(Address), new(Goods), //
		new(GoodsSKU), new(GoodsType), new(GoodsImage), //
		new(IndexTypeGoodsBanner), new(IndexPromotionBanner), new(IndexGoodsBanner))
	// 运行
	orm.RunSyncdb("default", // 别名
		false, //是否强制重新建表
		true)  // 是否现实运行日志

}
