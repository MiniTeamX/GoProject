package main

import (
	_ "api_service/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/avatar", "static/avatar")
	beego.SetStaticPath("/commodity", "static/commodity")
	beego.SetStaticPath("/essay", "static/essay")
	beego.SetStaticPath("/swagger", "swagger")
	beego.Run()
}
