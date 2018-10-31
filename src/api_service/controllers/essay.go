package controllers

import (
	"api_service/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type RecommandContent struct 
{
	UserId     int64
	ArticleId  int64
	AuthorPic  string
	AuthorNickname string
	ArticleTitle string
	ArticleDescription string
}

type HotContent struct
{
	ArticleId        int64
	ArticleTitle     string
	PictureUrl       string
	ArticleDescription    string
	ArticleWxtReward  int64
}

type EssayInformation struct
{
	models.Essay
	Comments []models.Comment
}

type EssayController struct {
	beego.Controller
}

// @Title upload essay
// @Description post essay
// @Param	body	body 	models.Essay	true	"essay"
// @Success 200 { "eid":eid  }
// @router /upload_article [post]
func (e *EssayController) Post() {
	var essay models.Essay
	json.Unmarshal(e.Ctx.Input.RequestBody, &essay)
	time := time.Now()
	essay.CreateTime = time

	eid := models.AddEssay(essay)
	e.Data["json"] = map[string]int64{"eid": eid}
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

// @Title Get
// @Description get essay by essay id
// @Param	eid		path 	string	true		"The key for essayid"
// @Success 200 {object} models.Essay
// @Failure 403 {string} error message
// @router /:eid [get]
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

// @Title Delete
// @Description delete the essay
// @Param	uid		path 	string	true		"The eid you want to delete"
// @Success 200 {string} delete success!
// @router /:eid [delete]
func (e *EssayController) Delete() {
	eid,_:= e.GetInt(":eid")
	models.DeleteEssay(int64(eid))
	e.Data["json"] = "delete success!"
	e.ServeJSON()
}

// @Title GetRecommandEssays
// @Description GetRecommand Essays
// @Success 200 {object} controllers.RecommandContent
// @Failure 403 {string} "num == 0"
// @router /recommand_tab [Get]
func (e *EssayController) GetRecommandEssays() {
	var recommand []RecommandContent
	essays, num := models.GetRecommandEssays(12)
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
		r.ArticleId = v.EssayId
		r.ArticleTitle = v.Title
		r.ArticleDescription = v.Content
		recommand = append(recommand, r)
	}
	e.Data["json"] = recommand 
	e.ServeJSON()
}

// @Title GetHotEssays
// @Description GetHot Essays
// @Success 200 {object} controllers.HotContent
// @Failure 403 {string} "num == 0"
// @router /hot_tab [Get]
func (e *EssayController) GetHotEssays() {
	hot := make([]HotContent, 0, 0)
	essays, num := models.GetHotEssays()
	if num == 0 {
		e.Data["json"] = "num == 0"
		return
	} 
	var h HotContent
	for _, v := range essays {
		h.ArticleId = v.EssayId
		h.ArticleTitle = v.Title
		h.ArticleDescription = v.Content
		h.PictureUrl = v.PictureUrl
		h.ArticleWxtReward = v.WxtReward
		hot = append(hot, h)
	}
	e.Data["json"] = hot  
	e.ServeJSON()
}

// @Title EssayLike
// @Description user like the essay
// @Param	UserId	  query 	int	true		"The user id"
// @Param	EssayId	  query 	int	true		"The essay id"
// @router /article_like_click [get]
func (e *EssayController) EssayLike() {
	userid, _:= e.GetInt("UserId")
	essayid, _:= e.GetInt("EssayId")
	t := time.Now()
	essay_like := models.LikeEssay{UserId:int64(userid), EssayId:int64(essayid), Time: t}
	models.AddEssayLike(essay_like)

	models.AddEssayLikeCount(strconv.Itoa(essayid))
	essay,_:= models.GetEssay(int64(essayid))
	models.AddUserLikeCount(strconv.FormatInt(essay.UserId, 10))
}


// @Title CommentLike
// @Description user like the  comment
// @Param	UserId	  query 	int	true		"The user id"
// @Param	CommentId	  query 	int	true		"The essay id"
// @router /comment_like_click [get]
func (e *EssayController) CommentLike() {
	userid, _:= e.GetInt("UserId")
	commentid, _:= e.GetInt("CommentId")
	t := time.Now()
	comment_like := models.LikeCommnet{UserId:int64(userid), CommentId:int64(commentid), Time: t}
	models.AddCommentLike(comment_like)

	c,_ := models.GetComment(int64(commentid))
	c.PhraseNum += 1
	models.UpdateComment(c.CommentId, c)
	from_id := c.FromId
	models.AddUserLikeCount(strconv.FormatInt(from_id, 10))
}



// @Title EssayCollect
// @Description user collect the essay
// @Param	UserId	 body  int  	true		"The user id"
// @Param	EssayId	 body  int	   true		"The essay id"
// @router /article_collect_click [post]
func (e *EssayController) EssayCollect() {
	var c models.Collect
	json.Unmarshal(e.Ctx.Input.RequestBody, &c)
	models.AddCollect(c)
	e.Data["json"] = "Add success"
	e.ServeJSON()
}

// @Title Essay comment
// @Description 评论功能
// @Param	body	 body   models.Comment   true		"The essay id"
// @router /article_comment_click [post]
func (e *EssayController) EssayComment() {
	var c models.Comment
	json.Unmarshal(e.Ctx.Input.RequestBody, &c)
	fmt.Print(c)
	models.AddComment(c)
	e.Data["json"] = "comment Add success"
	e.ServeJSON()
}


// @Title Essay information
// @Description essay information
// @Param	EssayId	 query  string   true		"The essay id"
// @Success 200 {object} Controllers.EssayInformation
// @router /article_information [Get]
func (e *EssayController) EssayInformation() {
	var essay_information  EssayInformation
	comments := make([]models.Comment, 0 , 0)
	eid := e.GetString("EssayId")
	eid_int,_ := strconv.ParseInt(eid, 10, 64)
	essay,_ := models.GetEssay(eid_int)

	essay_information.Essay = *essay
	comments = models.GetCommentsByEssayId(eid)
	essay_information.Comments = comments
	e.Data["json"] = essay_information
	e.ServeJSON()
}

// @Title Upload picture
// @Description upload essay photo
// @Param	 picture_name  form	 form  true	"multipart/form-data; filename:picture_name"
// @Success 200 { "result":1 , "message":"avatar url" }
// @Failure 403 { "result":0, "message": "upload failed"}
// @router /upload_picture [post]
func (u *EssayController) UploadPicture() {
	f, h, err := u.GetFile("picture_name")
	if err != nil {
		u.Data["json"] = map[string] interface{} {"result":0, "message":"upload failed"}
	}
	defer f.Close()
//	suffix := strings.Split(h.Filename, ".")[1]
//	filename := uid + "." + suffix;
	picture_url := "http://192.168.26.193" + "/static/avatar/" + h.Filename;
	u.SaveToFile("picture_name", "static/avatar/" + h.Filename)
	u.Data["json"] = map[string] interface{} {"result":1, "message":picture_url}
	u.ServeJSON()
}


