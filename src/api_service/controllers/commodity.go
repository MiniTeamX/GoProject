package controllers

import (
	"api_service/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type CommodityController struct {
	beego.Controller
}

func (e *CommodityController) Post() {
	var commodity models.Commodity
	json.Unmarshal(e.Ctx.Input.RequestBody, &commodity)
	cid := models.AddCommodity(commodity)
	e.Data["json"] = map[string]int64{"CommodityId": cid}
	e.ServeJSON()
}

func (e *CommodityController) Put() {
	eid,err := e.GetInt(":eid")
	if err == nil {
		var essay models.Essay
		json.Unmarshal(e.Ctx.Input.RequestBody, &essay)
		ee, err := models.UpdateEssay(int64(eid), &essay)
		if err != nil {
			e.Data["json"] = err
		} else {
			e.Data["json"] = ee
		}
	}
	e.ServeJSON()
}

// @Title Get
// @Description get commodity by commodity id
// @Param	cid		path 	string	true		"The key for commodity id"
// @Success 200 {object} models.Commodity
// @Failure 403 {string} error message
// @router /:cid [get]
func (e *CommodityController) Get() {
	cid,err := e.GetInt(":cid")
	if  err == nil {
		commodity, err := models.GetCommodity(int64(cid))
		if err != nil {
			e.Data["json"] = err.Error()
		} else {
			e.Data["json"] = commodity
		}
	}
	e.ServeJSON()
}

// @Title Delete
// @Description delete the commodity
// @Param	cid		path 	string	true	"The eid you want to delete"
// @Success 200 {string} delete success!
// @router /:cid [delete]
func (c *CommodityController) Delete() {
	cid,_:= c.GetInt(":cid")
	models.DeleteCommodity(int64(cid))
	c.Data["json"] = "delete success!"
	c.ServeJSON()
}


// @Title 信息验证
// @Description   验证用户余额和支付密码是否ok
// @Param	UserId	query 	string	true	"The user id"
// @Param	CommodityId	query 	string	true	"The commodity id"
// @Param	PayPassword	query 	string	true	"The pay password"
// @Success 200 { "result":1 , "message":"验证成功" }
// @Failure 403 { "result":0, "message": "验证失败原因"}
// @router /commodity_verify [get]
func (c *CommodityController) CommodityVerify() {
	uid, _ := c.GetInt("UserId")
	cid, _ := c.GetInt("CommodityId")
	paypassword := c.GetString("PayPassword")
	user,_ := models.GetUser(int64(uid))
	commodity,_ := models.GetCommodity(int64(cid))
	if user.XwtBalance < commodity.Price {
		c.Data["json"] = map[string] interface{} {"result":0, "message":"余额不足"}
	} else {
		if (paypassword != user.PayPassword) {
			c.Data["json"] = map[string] interface{} {"result:":0, "message": "支付密码不正确"}
		} else {
		//	user.XwtBalance -= commodity.Price
		//	commodity.LeftCount -= 1
		//	models.UpdateUser(int64(uid), user)
		//	models.UpdateCommodity(int64(cid), commodity)
			c.Data["json"] = map[string] interface{} {"result:":1, "message": "验证成功"}
		}
	}
	c.ServeJSON()
}










