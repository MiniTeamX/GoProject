package models

import (
"database/sql"
"fmt"
"go.uber.org/zap"
db "oc_res_control/database"
)

type BlackInfohashList struct {
	Infohash      string       `json:"infohash"`
	DwStrategy    int64        `json:"dw_strategy"`
	Cause         int64        `json:"cause"`
}

func GetBlackInfohash(infohash  string) (blackInfohashList BlackInfohashList, err error)  {
	var dwStrategy sql.NullInt64
	var cause sql.NullInt64
	err = db.DBShub.QueryRow("SELECT infohash,dwStrategy,cause FROM black_infohash_list WHERE hex(infohash)=?", infohash).Scan(
		&blackInfohashList.Infohash,&dwStrategy,&cause,
	)
	zap.L().Debug(fmt.Sprint("SELECT infohash,dwStrategy,cause FROM black_infohash_list WHERE hex(infohash)=%s", infohash))
	//deal with the null value in mysql
	if dwStrategy.Valid {
		blackInfohashList.DwStrategy = dwStrategy.Int64
	}
	if cause.Valid {
		blackInfohashList.Cause = cause.Int64
	}

	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Debug(err.Error())
			return
		} else {
			zap.L().Error(fmt.Sprintf("faild to query the database:%s", err.Error()))
		}
	}
	return
}

func AddBlackInfohash(blackInfohashList BlackInfohashList) error{
	_,err := db.DBShub.Exec("INSERT INTO black_infohash_list(infohash,dwStrategy,cause) VALUES(UNHEX(?), ?, ?) ON DUPLICATE KEY UPDATE dwStrategy=?,cause=?",
		blackInfohashList.Infohash,blackInfohashList.DwStrategy,blackInfohashList.Cause, blackInfohashList.DwStrategy, blackInfohashList.Cause)

	zap.L().Debug(fmt.Sprint("INSERT INTO black_infohash_list(infohash,dwStrategy,cause) VALUES (UNHEX(%s), %d, %d) ON DUPLICATE KEY UPDATE dwStrategy=%d,cause=%d",blackInfohashList.Infohash,blackInfohashList.DwStrategy,blackInfohashList.Cause, blackInfohashList.DwStrategy, blackInfohashList.Cause))
	if err != nil {
		zap.L().Error(err.Error())
	}
	return err
}

func DeleteBlackInfohash(infohash string) error {
	_,err := db.DBShub.Exec("DELETE FROM black_infohash_list WHERE HEX(infohash)=?", infohash)
	zap.L().Debug(fmt.Sprint("DELETE FROM black_url_list WHERE HEX(infohash)=%s", infohash))
	if err != nil {
		zap.L().Error(err.Error())
	}
	return err
}

func UpdateBlackInfohash(blackInfohashList BlackInfohashList) error {
	return AddBlackInfohash(blackInfohashList)
}