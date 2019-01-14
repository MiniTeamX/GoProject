package main

import (
	_ "api_service/routers"
	"api_service/xwtreward"
	"github.com/astaxie/beego"
)

func main() {
	xwtreward.CreatEssayReward()
	xwtreward.EssayLikeReward()
	xwtreward.WriteLogToFile()
	beego.SetStaticPath("/avatar", "static/avatar")
	beego.SetStaticPath("/commodity", "static/commodity")
	beego.SetStaticPath("/essay", "static/essay")
	beego.SetStaticPath("/swagger", "swagger")
	beego.Run()
}
