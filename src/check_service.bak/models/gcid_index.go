package models

import (
	db "check_service/database"
	"database/sql"
	"log"
	"time"
)

type GcidIndex struct {
	Id          int64
	Gcid        string
	Flag        string
	Ts          time.Time
	Checker     string
}

func GetGcidIndex(gcid string) (gcidIndex GcidIndex, err error) {
	subgcid := string([]rune(gcid)[0:2])
	//querySql :="SELECT gcid,flag,ts,checker FROM gcidindex_" + subgcid + " WHERE gcid=" + "'" + gcid + "'"
	err = db.DBGcidList.QueryRow("SELECT gcid,flag,ts,checker FROM gcidindex_" + subgcid + " WHERE gcid=?",gcid).Scan(
		&gcidIndex.Gcid, &gcidIndex.Flag,&gcidIndex.Ts, &gcidIndex.Checker,
	)
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
