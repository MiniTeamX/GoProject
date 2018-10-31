package routers

import (
	"api_service/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace( "/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/essay",
			beego.NSInclude(
				&controllers.EssayController{},
			),
		),
		beego.NSNamespace("/commodity",
			beego.NSInclude(
				&controllers.CommodityController{},
			),
		),
		beego.NSNamespace("/request",
			beego.NSInclude(
				&controllers.XwtRequestController{},
			),
		),
	)
	beego.AddNamespace(ns)
}


