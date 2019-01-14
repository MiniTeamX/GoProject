package models

import (
"database/sql"
"fmt"
"go.uber.org/zap"
db "oc_res_control/database"
)

type BlackEmulehashList struct {
	Emulehash      string       `json:"emulehash"`
	DwStrategy    int64        `json:"dw_strategy"`
	Cause         int64        `json:"cause"`
}

func GetBlackEmulehash(emulehash  string) (blackEmulehashList BlackEmulehashList, err error)  {
	var dwStrategy sql.NullInt64
	var cause sql.NullInt64
	err = db.DBShub.QueryRow("SELECT emulehash,dwStrategy,cause FROM black_emulehash_list WHERE hex(emulehash)=?", emulehash).Scan(
		&blackEmulehashList.Emulehash,&dwStrategy,&cause,
	)
	zap.L().Debug(fmt.Sprint("SELECT emulehash,dwStrategy,cause FROM black_emulehash_list WHERE hex(emulehash)=%s", emulehash))
	//deal with the null value in mysql
	if dwStrategy.Valid {
		blackEmulehashList.DwStrategy = dwStrategy.Int64
	}
	if cause.Valid {
		blackEmulehashList.Cause = cause.Int64
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

func AddBlackEmulehash(blackEmulehashList BlackEmulehashList) error{
	_,err := db.DBShub.Exec("INSERT INTO black_emulehash_list(emulehash,dwStrategy,cause) VALUES (UNHEX(?), ?, ?) ON DUPLICATE KEY UPDATE dwStrategy=?,cause=?",
		blackEmulehashList.Emulehash,blackEmulehashList.DwStrategy,blackEmulehashList.Cause, blackEmulehashList.DwStrategy, blackEmulehashList.Cause)

	zap.L().Debug(fmt.Sprint("INSERT INTO black_emulehash_list(emulehash,dwStrategy,cause) VALUES (UNHEX(%s), %d, %d) ON DUPLICATE KEY UPDATE dwStrategy=%d,cause=%d",blackEmulehashList.Emulehash,blackEmulehashList.DwStrategy,blackEmulehashList.Cause, blackEmulehashList.DwStrategy, blackEmulehashList.Cause))
	if err != nil {
		zap.L().Error(err.Error())
	}
	return err
}

func DeleteBlackEmulehash(emulehash string) error {
	_,err := db.DBShub.Exec("DELETE FROM black_emulehash_list WHERE HEX(emulehash)=?", emulehash)
	zap.L().Debug(fmt.Sprint("DELETE FROM black_url_list WHERE HEX(emulehash)=%s", emulehash))
	if err != nil {
		zap.L().Error(err.Error())
	}
	return err
}

func UpdateBlackEmulehash(blackEmulehashList BlackEmulehashList) error {
	return AddBlackEmulehash(blackEmulehashList)
}