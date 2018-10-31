package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Xwtrequest struct {
	RequestId      int64  `orm:"column(request_id);pk"`
	Content        string
	IsUsed         int64
}

func init() {
	orm.RegisterModel(new(Xwtrequest))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8")
}

func UpdateRequest(rid int64, r *Xwtrequest) (err error, rr *Xwtrequest) {
	o := orm.NewOrm()
	request := Xwtrequest{RequestId: rid}
	if o.Read(&request) == nil {
		_, err := o.Update(r)
		return  err,r
	}
	return errors.New("id not exits."),r
}

func AddRequest(r Xwtrequest){
	o := orm.NewOrm()
	o.Insert(&r)
}

func GetAllRequest() ( requests []Xwtrequest) {
	o := orm.NewOrm()
	o.Raw("SELECT * FROM xwtrequest WHERE is_used = 0").QueryRows(&requests)
	return
}
