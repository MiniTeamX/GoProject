package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Comment struct {
	CommentId       int64  `orm:"column(comment_id);pk"`
	EssayId         int64
	Content         string
	PhraseNum       int64
	FromId          int64
	ToId            int64
	CreateTime      time.Time
}

func init() {
	orm.RegisterModel(new(Comment))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:xcvvcx@/my_db?charset=utf8")
}

func AddComment(c Comment) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(&c)
	return id
}

func GetComment(cid int64) (c *Comment, err error) {
	o := orm.NewOrm()
	c = &Comment{CommentId: cid}
	err = o.Read(c)
	return c, err
}

func DeleteComment(cid int64) {
	o := orm.NewOrm()
	comment := Comment{CommentId: cid}
	o.Delete(&comment)
}

func GetCommentsByEssayId(eid string) (cc []Comment) {
	o := orm.NewOrm()
	comments := make([]Comment, 0, 0)
	o.Raw("SELECT * FROM comment WHERE essay_id = " + eid).QueryRows(&comments)
	return comments
}

