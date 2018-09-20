package main

import (
	_ "api_service/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/avatar", "static/avatar")
	beego.Run()
}
