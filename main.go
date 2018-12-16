package main

import (
	"github.com/astaxie/beego"
	_ "ttsx/models"
	_ "ttsx/routers"
)

func main() {
	beego.Run()
}
