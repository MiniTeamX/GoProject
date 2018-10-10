package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Collect struct {
	UserId         int64 `orm:"column(user_id);pk"`
	EssayId        int64
}

func init() {
	orm.RegisterModel(new(Collect))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:xcvvcx@/my_db?charset=utf8")
}

func AddCollect(c Collect){
	o := orm.NewOrm() 
	o.Insert(&c)
}

func GetCollect(uid, eid int64) (c *Collect, err error) {
	o := orm.NewOrm()
	c = &Collect{UserId:uid, EssayId: eid}
	err = o.Read(c)
	return c, err 
}

func UpdateCollect(uid, eid int64, c *Collect) (err error, cc *Collect) {
	o := orm.NewOrm()
	collect := Collect{UserId:uid, EssayId: eid}
	if o.Read(&collect) == nil {
		_, err := o.Update(c) 
		return  err,c 
	}
	return errors.New("id not exits."),c
}

func DeleteCollect(uid, eid int64) {
	o := orm.NewOrm()
	collect := Collect{UserId: uid, EssayId: eid}
	o.Delete(&collect)
}
