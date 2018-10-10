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

	beego.Router("/essay/?:eid", &controllers.EssayController{})
	beego.Router("/essay/recommand_tab", &controllers.EssayController{}, "get:GetRecommandEssays")
	beego.Router("/essay/hot_tab", &controllers.EssayController{}, "get:GetHotEssays")
	beego.Router("/essay/article_like_click", &controllers.EssayController{}, "get:EssayLike")
	beego.Router("/essay/article_collect_click", &controllers.EssayController{}, "post:EssayCollect")
}
