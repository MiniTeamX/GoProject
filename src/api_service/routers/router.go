package routers

import (
	"api_service/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/user/?:uid", &controllers.UserController{})
}
