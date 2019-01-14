package database

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

var DBGcidList *sql.DB
var DBShub     *sql.DB
var DBGcidLixian *sql.DB

func init() {
    var err error
    DBGcidList, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/gcidlist?parseTime=true")
    if err != nil {
        log.Fatal(err.Error())
    }

    DBShub, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/special_data?parseTime=true")
    if err != nil {
        log.Fatal(err.Error())
    }

    DBGcidLixian, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/filter_gcid_lixian?parseTime=true")
    if err != nil {
        log.Fatal(err.Error())
    }
}
