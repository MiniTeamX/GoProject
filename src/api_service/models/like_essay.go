package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)


type LikeEssay struct {
	UserId         int64 `orm:"column(user_id);pk"`
	EssayId        int64
	Time    time.Time
}


func init() {
	orm.RegisterModel(new(LikeEssay))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8")
}

func AddEssayLike(l LikeEssay){
	o := orm.NewOrm()
	o.Insert(&l)
}

func GetAllLikeEssayByData(now time.Time) (e []LikeEssay){
	 o := orm.NewOrm()
	 o.Raw("SELECT user_id,essay_id FROM like_essay").QueryRows(&e)
	 return
}



