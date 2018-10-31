package controllers

import (
	"api_service/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type XwtRequestController struct {
	beego.Controller
}

// @Title post request
// @Description 发送请求
// @Param	body	body 	models.Xwtrequest	true	"body for request contennt"
// @Success 200 { "result":1 , "message":"register success" }
// @router /post_request [post]
func (r *XwtRequestController) PostRequest() {
	var request models.Xwtrequest
	json.Unmarshal(r.Ctx.Input.RequestBody, &request)
	models.AddRequest(request)
	r.Data["json"] = map[string]interface{} {"result":1, "message":"post seccuss."}
	r.ServeJSON()
}

// @Title get unused request
// @Description  得到所有没有处理过的请求
// @Success 200 {object} models.XwtRequest
// @router /get_all_request [get]
func (r *XwtRequestController) GetAllRequest() {
	var rr []string
	requests := models.GetAllRequest()
	for _,v := range requests {
	//	v.IsUsed = 1
		v.IsUsed = 0
		models.UpdateRequest(v.RequestId, &v)
		rr = append(rr, v.Content)
	}
	r.Data["json"] = rr
	r.ServeJSON()
}

// @Title deal  request callback
// @Description  处理请求回调
// @Success 200 {"result" : 1, "message":"调用成功"}
// @router /deal_request [post]
func (r *XwtRequestController) DealRequestCallback() {
	f := r.GetString("from")
	fmt.Print("付款地址:" + f)
}





