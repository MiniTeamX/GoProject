package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"strconv"
)

type Essay struct {
	EssayId        int64 `orm:"column(essay_id);pk"`
	UserId         int64
	Type           int64
	EssayUrl       string
	Title          string
	Content        string
	PhraseNum      int64
	PictureUrl     string
	EssayCount     int64
	CommentNum     int64
	WatchNum       int64
	RecommandNum   int64
	WxtReward      int64
	CreateTime     time.Time
}

type EssaySnapshot struct {
	EssayId      int64
	UserId       int64
	EssayCount   int64
	Type         int64
}

func init() {
	orm.RegisterModel(new(Essay))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8")
}

func AddEssay(e Essay) int64 {
	o := orm.NewOrm() 
	id, _ := o.Insert(&e)
	return id 
}

func GetEssay(eid int64) (e *Essay, err error) {
	o := orm.NewOrm()
	e = &Essay{EssayId: eid}
	err = o.Read(e)
	return e, err 
}

func UpdateEssay(eid int64, e *Essay) (err error, ee *Essay) {
	o := orm.NewOrm()
	essay := Essay{EssayId: eid}
	if o.Read(&essay) == nil {
		_, err := o.Update(e) 
		return  err,e 
	}
	return errors.New("id not exits."),e
}

func DeleteEssay(eid int64) {
	o := orm.NewOrm()
	essay := Essay{EssayId: eid}
	o.Delete(&essay)
}


func GetRecommandEssays(limit int64)(ee []Essay, n int64) {
	o := orm.NewOrm()
	var e []Essay
	num, _:= o.Raw("SELECT * FROM essay Limit 0, " + strconv.FormatInt(limit, 10)).QueryRows(&e)
	return e,num
}

func GetHotEssays()(ee []Essay, n int64) {
	return GetRecommandEssays(12)
}

func AddEssayLikeCount(eid string) {
	o := orm.NewOrm()
	o.Raw("UPDATE essay SET phrase_num = phrase_num + 1 WHERE essay_id = " + eid).Exec()
}

func GetEssaySnapshot() (e []EssaySnapshot){
	o := orm.NewOrm()
	o.Raw("SELECT essay_id,user_id,essay_count,type From essay").QueryRows(&e)
	return
}
