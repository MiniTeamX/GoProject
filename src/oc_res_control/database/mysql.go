package database

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "go.uber.org/zap"
    "oc_res_control/conf"
)


var DBShub     *sql.DB

func InitDb(cfg *conf.Config) {
	var err error
	fmt.Sprint("%s,%s,%s,%d,%s", cfg.Mysql.Db.SpecialData.User,cfg.Mysql.Db.SpecialData.Passwd, cfg.Mysql.Db.SpecialData.Host,cfg.Mysql.Db.SpecialData.Port,cfg.Mysql.Db.SpecialData.Name)
    DBShub,err = sql.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
                                                            cfg.Mysql.Db.SpecialData.User,
                                                            cfg.Mysql.Db.SpecialData.Passwd,
                                                            cfg.Mysql.Db.SpecialData.Host,
                                                            cfg.Mysql.Db.SpecialData.Port,
                                                            cfg.Mysql.Db.SpecialData.Name))
    DBShub.SetMaxOpenConns(500)
    DBShub.SetMaxIdleConns(1000)
    if err != nil {
        zap.L().Error(fmt.Sprintf("failed to connect to mysql, the database is: %s", cfg.Mysql.Db.SpecialData.Name))
    }
}


