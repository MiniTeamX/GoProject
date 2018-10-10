package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Commodity struct {
	CommodityId     int64 `orm:"column(Commodity_id);pk"`
	name            string
	description     string
	picture_url     string
	pricce          int64
	left_count      int64
}

func init() {
	orm.RegisterModel(new(Commodity))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:xcvvcx@/my_db?charset=utf8")
}

func AddCommodity(c Commodity) int64 {
	o := orm.NewOrm() 
	id, _ := o.Insert(&c)
	return id 
}

func GetCommodity(cid int64) (c *Commodity, err error) {
	o := orm.NewOrm()
	c = &Commodity{CommodityId: cid}
	err = o.Read(c)
	return c, err 
}

func UpdateCommodity(cid int64, c *Commodity) (err error, cc *Commodity) {
	o := orm.NewOrm()
	commodity := Commodity{CommodityId: cid}
	if o.Read(&commodity) == nil {
		_, err := o.Update(c) 
		return  err,c 
	}
	return errors.New("id not exits."),c
}

func DeleteCommodity(cid int64) {
	o := orm.NewOrm()
	commodity := Commodity{CommodityId: cid}
	o.Delete(&commodity)
}
