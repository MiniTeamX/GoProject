package models

import (
	db "check_service/database"
	"database/sql"
	"log"
)

type BlackGcidList struct {
	Gcid        string
	DwStrategy  int64
	Cause       int64
}

func GetBlackGcidList(gcid string) (blackGcidList BlackGcidList, err error) {
	var dwStrategy sql.NullInt64
	var cause sql.NullInt64
	err = db.DBShub.QueryRow("SELECT hex(gcid),dwStrategy,cause FROM black_gcid_list WHERE hex(gcid)=?", gcid).Scan(
		&blackGcidList.Gcid,&dwStrategy,&cause,
	)
	//deal with the null value in mysql
	if dwStrategy.Valid {
		blackGcidList.DwStrategy = dwStrategy.Int64
	}
	if cause.Valid {
		blackGcidList.Cause = cause.Int64
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


