package controllers

import (
	"api_service/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
)

type RecommandContent struct 
{
	UserId     int64
	AuthorPic  string
	AuthorNickname string
	ArticleTitle string
	ArticleDescription string
}

type HotContent struct
{
	ArticleTitle     string
	ArticleDescription    string
	ArticleWxtReward  int64
}



type EssayController struct {
	beego.Controller
}

func (e *EssayController) Post() {
	var essay models.Essay
	json.Unmarshal(e.Ctx.Input.RequestBody, &essay)
	eid := models.AddEssay(essay)
	e.Data["json"] = map[string]int64{"eid": eid}
	e.ServeJSON()
}

func (e *EssayController) Get() {
	eid,err := e.GetInt(":eid")
	if  err == nil {
		essay, err := models.GetEssay(int64(eid))
		if err != nil {
			e.Data["json"] = err.Error()
		} else {
			e.Data["json"] = essay 
		}
	}
	e.ServeJSON()
}

func (e *EssayController) Put() {
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

func (e *EssayController) Delete() {
	eid,_:= e.GetInt(":eid")
	models.DeleteEssay(int64(eid))
	e.Data["json"] = "delete success!"
	e.ServeJSON()
}

func (e *EssayController) GetRecommandEssays() {
	var recommand []RecommandContent
	essays, num := models.GetRecommandEssays(3)
	if num == 0 {
		e.Data["json"] = "num == 0"
		return
	} 
	var r RecommandContent
	for _, v := range essays {
		user, _ := models.GetUser(v.UserId)
		r.UserId = user.UserId
		r.AuthorPic = user.PhotoUrl
		r.AuthorNickname = user.NickName
		r.ArticleTitle = v.Title
		r.ArticleDescription = v.Content
		recommand = append(recommand, r)
	}
	e.Data["json"] = recommand 
	e.ServeJSON()
}

func (e *EssayController) GetHotEssays() {
	var hot []HotContent
	essays, num := models.GetHotEssays()
	if num == 0 {
		e.Data["json"] = "num == 0"
		return
	} 
	var h HotContent
	for _, v := range essays {
		h.ArticleTitle = v.Title
		h.ArticleDescription = v.Content
		h.ArticleWxtReward = v.WxtReward
		hot = append(hot, h)
	}
	e.Data["json"] = hot  
	e.ServeJSON()
}

func (e *EssayController) EssayLike() {
//	userid := e.GetString("user_id")
	essayid:= e.GetString("article_id")
	models.AddEssayLikeCount(essayid)
	id,_ := strconv.ParseInt(essayid, 10, 64)
	essay,_:= models.GetEssay(id)
	models.AddUserLikeCount(strconv.FormatInt(essay.UserId, 10))
}

func (e *EssayController) EssayCollect() {
	var c models.Collect
	json.Unmarshal(e.Ctx.Input.RequestBody, &c)
	cid := models.AddCollect(c)
	e.Data["json"] = map[string]int64{"cid": cid}
	e.ServeJSON()
}