package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)


type Xwtlog struct {
	Id            int64  `orm:"column(id);pk"`
	Content       string
	Writed        int64
	Time          time.Time
}

func init() {
	orm.RegisterModel(new(Xwtlog))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8")
}


func Addlog(l Xwtlog){
	o := orm.NewOrm()
	o.Insert(&l)
}


func GetLogNotWrited() (logs []Xwtlog){
	o := orm.NewOrm()
	o.Raw("SELECT * FROM xwtlog WHERE writed = 0").QueryRows(logs)
	return
}