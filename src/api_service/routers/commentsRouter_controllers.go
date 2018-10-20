package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["api_service/controllers:CommodityController"] = append(beego.GlobalControllerRouter["api_service/controllers:CommodityController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:cid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_service/controllers:CommodityController"] = append(beego.GlobalControllerRouter["api_service/controllers:CommodityController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:cid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_service/controllers:CommodityController"] = append(beego.GlobalControllerRouter["api_service/controllers:CommodityController"],
		beego.ControllerComments{
			Method: "CommodityVerify",
			Router: `/commodity_verify`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_service/controllers:EssayController"] = append(beego.GlobalControllerRouter["api_service/controllers:EssayController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:eid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_service/controllers:EssayController"] = append(beego.GlobalControllerRouter["api_service/controllers:EssayController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:eid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_service/controllers:EssayController"] = append(beego.GlobalControllerRouter["api_service/controllers:EssayController"],
		beego.ControllerComments{
			Method: "EssayCollect",
			Router: `/article_collect_click`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_service/controllers:EssayController"] = append(beego.GlobalControllerRouter["api_service/controllers:EssayController"],
		beego.ControllerComments{
			Method: "EssayInformation",
			Router: `/article_information`,
			AllowHTTPMethods: []string{"Get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_service/controllers:EssayController"] = append(beego.GlobalControllerRouter["api_service/controllers:EssayController"],
		beego.ControllerComments{
			Method: "EssayLike",
			Router: `/article_like_click`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_service/controllers:EssayController"] = append(beego.GlobalControllerRouter["api_service/controllers:EssayController"],
		beego.ControllerComments{
			Method: "GetHotEssays",
			Router: `/hot_tab`,
			AllowHTTPMethods: []string{"Get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_service/controllers:EssayController"] = append(beego.GlobalControllerRouter["api_service/controllers:EssayController"],
		beego.ControllerComments{
			Method: "GetRecommandEssays",
			Router: `/recommand_tab`,
			AllowHTTPMethods: []string{"Get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_service/controllers:UserController"] = append(beego.GlobalControllerRouter["api_service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_service/controllers:UserController"] = append(beego.GlobalControllerRouter["api_service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_service/controllers:UserController"] = append(beego.GlobalControllerRouter["api_service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_service/controllers:UserController"] = append(beego.GlobalControllerRouter["api_service/controllers:UserController"],
		beego.ControllerComments{
			Method: "CollectEssays",
			Router: `/collect_articles`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_service/controllers:UserController"] = append(beego.GlobalControllerRouter["api_service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_service/controllers:UserController"] = append(beego.GlobalControllerRouter["api_service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_service/controllers:UserController"] = append(beego.GlobalControllerRouter["api_service/controllers:UserController"],
		beego.ControllerComments{
			Method: "UploadPicture",
			Router: `/upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_service/controllers:UserController"] = append(beego.GlobalControllerRouter["api_service/controllers:UserController"],
		beego.ControllerComments{
			Method: "UploadAvatar",
			Router: `/upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
