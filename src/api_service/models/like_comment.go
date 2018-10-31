package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)


type LikeCommnet struct {
	UserId         int64 `orm:"column(user_id);pk"`
	CommentId        int64
	Time    time.Time
}

func init() {
	orm.RegisterModel(new(LikeCommnet))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8")
}

func AddCommentLike(l LikeCommnet){
	o := orm.NewOrm()
	o.Insert(&l)
}
