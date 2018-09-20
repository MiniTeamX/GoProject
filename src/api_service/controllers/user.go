package controllers

import (
	"api_service/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	u.Data["json"] = map[string]int64{"uid": uid}
	u.ServeJSON()
}

func (u *UserController) Get() {
	uid,err := u.GetInt(":uid")
	if  err == nil {
		user, err := models.GetUser(int64(uid))
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

func (u *UserController) Put() {
	uid,err := u.GetInt(":uid")
	if err == nil {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(int64(uid), &user)
		if err != nil {
			u.Data["json"] = err
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

func (u *UserController) Delete() {
	uid,_:= u.GetInt(":uid")
	models.DeleteUser(int64(uid))
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	uid,err := models.Login(username, password) 
    u.Data["json"] =  map[string]string {"uid":strconv.FormatInt(uid, 10), "message":err.Error()}
	u.ServeJSON()
}

func (u *UserController) Register() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	res, err := models.Register(user)
	if res == true {
		u.Data["json"] = map[string]string {"result":"1", "message":err.Error()}
	} else {
		u.Data["json"] = map[string]string {"result":"0","message":err.Error()}
	}
	u.ServeJSON()
}

func (u *UserController) Upload() {
	jsoninfo := u.GetString("testjson")
	u.Ctx.WriteString(jsoninfo)
}

func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

