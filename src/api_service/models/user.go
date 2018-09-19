package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
    UserId          int64  `orm:"column(user_id);pk"`	
	Username        string
	Password        string
	NickName        string
	Gender          int64
	PhotoUrl        string
	Introduction    string
}

func init() {
	orm.RegisterModel(new(User))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:xcvvcx@/my_db?charset=utf8")
}

func AddUser(u User) int64 {
	o := orm.NewOrm() 
	id, _ := o.Insert(&u)
	return id 
}

func GetUser(uid int64) (u *User, err error) {
	o := orm.NewOrm()
	u = &User{UserId: uid}
	err = o.Read(u)
	return u, err 
}

func UpdateUser(uid int64, u *User) (err error, uu *User) {
	o := orm.NewOrm()
	user := User{UserId: uid}
	if o.Read(&user) == nil {
		_, err := o.Update(u) 
		return  err,u 
	}
	return errors.New("id not exits."),u
}

func DeleteUser(uid int64) {
	o := orm.NewOrm()
	user := User{UserId: uid}
	o.Delete(&user)
}

func Login(username, password string) bool {
	return false
}

