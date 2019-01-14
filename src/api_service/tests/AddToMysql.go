package test

import (
	"api_service/models"
	"strconv"
	"time"
)

func init() {

}

func AddLikeEssay() {
	var u models.LikeEssay
	for uid:= 2;uid <= 9; uid++ {
		for eid := 1;eid <= 35;eid++ {
			u.UserId = int64(uid)
			u.EssayId = int64(eid)
			u.Time = time.Now()
			models.AddEssayLike(u)
		}
	}
}

func AddComment() {
	var c models.Comment
	for uid:= 1;uid <= 25; uid++ {
		for eid := 1;eid <= 35;eid++ {
			c.EssayId = int64(eid)
			c.Content = "好文章啊"
			c.PhraseNum = 100 + int64(uid)
			c.FromId = int64(uid)
			c.ToId = 25 - int64(uid)
			c.CreateTime = time.Now()
			models.AddComment(c)
		}
	}

}


func AddCollect() {
	var u models.Collect
	for uid:= 1;uid <= 20; uid++ {
		for eid := 1;eid <= 30;eid++ {
			u.UserId = int64(uid)
			u.EssayId = int64(eid)
			u.CollectTime = time.Now()
			models.AddCollect(u)
		}
	}
}

func AddUser () {
	var u models.User
	for id := 10;id < 17;id++ {
		u.UserId = int64(id)
		u.Username = "14522233344"
		u.Password = "123456"
		u.PayPassword = "123455"
		u.NickName = "周杰伦" + strconv.FormatInt(int64(id), 10)
		u.Gender = 1
		u.PhotoUrl = "http://119.23.255.171/static/avatar/2.jpg"
		u.Introduction = "我是杰伦"
		u.PhraseNum = 100
		u.CreateTime = time.Now()
		models.AddUser(u)
	}
}
