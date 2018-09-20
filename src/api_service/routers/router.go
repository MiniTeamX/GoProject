package routers

import (
	"api_service/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/user/?:uid", &controllers.UserController{})
	beego.Router("/login", &controllers.UserController{}, "get:Login")
	beego.Router("/register", &controllers.UserController{}, "post:Register")
	beego.Router("/upload", &controllers.UserController{}, "post:Upload")
}
