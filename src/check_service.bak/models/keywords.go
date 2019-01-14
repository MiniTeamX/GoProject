package models

import (
	db "check_service/database"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type KeyWord struct {
	Id       int64
	KeyWord  string
	Ts       time.Time
	Flag     string
}

func GetKeyWordsByKeyWord(word string) (keyword KeyWord, err error) {
	var ts mysql.NullTime
	err = db.DBGcidList.QueryRow("SELECT id,keyword,ts,flag FROM keywords WHERE keyword=?", word).Scan(
		&keyword.Id, &keyword.KeyWord, &ts, &keyword.Flag,
	)
	// maybe the ts is null time
	if ts.Valid {
		keyword.Ts = ts.Time
	}
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err.Error())
			return
		} else {
			log.Fatalln(err.Error())
		}
	}

    return
}
