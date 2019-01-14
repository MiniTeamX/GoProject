package controllers

import (
	"api_service/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)


type UserController struct {
	beego.Controller
}

type CollectArticle struct {
	UserAvatar      string
	UserNickName    string
	EssayId         int64
	EssayTitle      string
	EssayReadNum     int64
	EssayPhraseNum   int64
	EssayCommentNum  int64
	EssayWatchNum    int64
	EssayXwtReward  int64
	EssayPictureUrl string
}

// @Title Post
// @Description post new user
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 { "result":1 , "message":"post success" }
// @Failure 403 { "result":0, "message": "fail message"}
// @router /:uid [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	u.Data["json"] = map[string]int64{"uid": uid}
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
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

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
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

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid,_:= u.GetInt(":uid")
	models.DeleteUser(int64(uid))
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the systems
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 { "uid":uid , "message":"login success" }
// @Failure 403 { "uid":-1, "message": "fail message"}
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	uid,err := models.Login(username, password) 
    u.Data["json"] =  map[string]interface{}{"uid":uid, "message":err.Error()}
	u.ServeJSON()
}

// @Title Register
// @Description Register new user
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 { "result":1 , "message":"register success" }
// @Failure 403 { "result":0, "message": "fail message"}
// @router /register [post]
func (u *UserController) Register() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	res, err := models.Register(user)
	if res == true {
		u.Data["json"] = map[string]interface{} {"result":1, "message":err.Error()}
	} else {
		u.Data["json"] = map[string]interface{} {"result":0,"message":err.Error()}
	}
	u.ServeJSON()
}

// @Title Upload Avatar
// @Description upload avatar photo
// @Param	 avatar_name  form	 form  true	"multipart/form-data; filename:avatar_name"
// @Param	 uid	form 	form   true		"user id"
// @Success 200 { "result":1 , "message":"avatar url" }
// @Failure 403 { "result":0, "message": "upload failed"}
// @router /upload [post]
func (u *UserController) UploadAvatar() {
	f, h, err := u.GetFile("avatar_name")
	if err != nil {
		u.Data["json"] = map[string] interface{} {"result":0, "message":"upload failed"}
	}
	defer f.Close()
	uid := u.GetString("uid")
    suffix := strings.Split(h.Filename, ".")[1]
    filename := uid + "." + suffix;
	avatar_url := "http://192.168.26.193" + "/static/avatar/" + filename;
	u.SaveToFile("avatar_name", "static/avatar/" + filename)
	u.Data["json"] = map[string] interface{} {"result":1, "message":avatar_url}
	uid_int64,_ := strconv.ParseInt(uid, 10, 64)
	user, _ := models.GetUser(uid_int64)
	user.PhotoUrl = avatar_url
	models.UpdateUser(uid_int64, user)
	u.ServeJSON()
}

// @Title
// @Description 展示用户收藏文章集合
// @Param	UserId	query 	string   true	"user id"
// @Success 200 {object} models.Commodity
// @router /collect_articles [get]
func (u *UserController) CollectEssays() {
	uid,_ := u.GetInt("UserId")
	_,collects := models.GetCollectByUserId(int64(uid))
	cc := make([]CollectArticle, 0, 0)
	var c CollectArticle
	for _,v := range collects {
		essay,_ := models.GetEssay(v.EssayId)
		user,_ := models.GetUser(essay.UserId)
		c.UserNickName = user.NickName
		c.UserAvatar = user.PhotoUrl
		c.EssayId = essay.EssayId
		c.EssayPhraseNum = essay.PhraseNum
		c.EssayCommentNum = essay.CommentNum
		c.EssayWatchNum = essay.WatchNum
		c.EssayTitle = essay.Title
		c.EssayPictureUrl = essay.PictureUrl
		c.EssayXwtReward = essay.WxtReward
		cc = append(cc, c)
	}
	u.Data["json"] = cc
	u.ServeJSON()
}




