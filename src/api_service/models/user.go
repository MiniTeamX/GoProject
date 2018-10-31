package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
	"time"
)

type User struct {
    UserId          int64  `orm:"column(user_id);pk"`	
	Username        string
	Password        string
	PayPassword     string
	NickName        string
	Gender          int64
	PhotoUrl        string
	Introduction    string
    PhraseNum       int64
    XwtBalance      int64
    XwtPowerValue   int64
	CreateTime      time.Time
}

func init() {
	orm.RegisterModel(new(User))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8")
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

func GetUserByName(username string) (res bool, u *User) {
	o := orm.NewOrm()
	var user User
	user.Username=username
	err := o.Read(&user, "Username")
	if err == orm.ErrNoRows {
		return false, nil
	}
	return true, &user
}

func Login(username, password string) (id int64, e error) {
	o := orm.NewOrm()
	if username == "" || password == "" {
		return -1, errors.New("username or password can not be empty.")
	}
	var user User
	user.Username=username
	err := o.Read(&user, "Username")

	if err == orm.ErrNoRows {
		return -1,errors.New("username not correct.")
	}
	if user.Password != password {
		return -1,errors.New("username or password not correct.")
	}
	return user.UserId, errors.New("login success.") 
}

func Register(u User) (res bool, e error) {
	r, _ := regexp.Compile("[0-9]{11}")
	if !r.MatchString(u.Username) {
		return false, errors.New("you should input phonenum")
	}
	if u.Password == "" {
		return false, errors.New("password can not be empty.")
	}
	res,_ = GetUserByName(u.Username)
	if res {
		return false, errors.New("user already exists.")
	}
	AddUser(u);
	return true, errors.New("register success.")
}

func AddUserLikeCount(uid string) {
	o := orm.NewOrm()
	o.Raw("UPDATE user SET phrase_num = phrase_num + 1 WHERE user_id = " + uid).Exec()
}