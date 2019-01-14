package models

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	db "oc_res_control/database"
)

type BlackUrlList struct {
	FileUrl       string       `json:"file_url"`
	DwStrategy    int64        `json:"dw_strategy"`
	Cause         int64        `json:"cause"`
}

func GetBlackUrl(fileUrl  string) (blackUrlList BlackUrlList, err error)  {
	var dwStrategy sql.NullInt64
	var cause sql.NullInt64
	err = db.DBShub.QueryRow("SELECT file_url,dwStrategy,cause FROM black_url_list WHERE file_url=?", fileUrl).Scan(
		&blackUrlList.FileUrl,&dwStrategy,&cause,
	)
	zap.L().Debug(fmt.Sprint("SELECT file_url,dwStrategy,cause FROM black_url_list WHERE file_url='%s' ", fileUrl))
	//deal with the null value in mysql
	if dwStrategy.Valid {
		blackUrlList.DwStrategy = dwStrategy.Int64
	}
	if cause.Valid {
		blackUrlList.Cause = cause.Int64
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

func AddBlackUrl(blackUrlList BlackUrlList) error{
	_,err := db.DBShub.Exec("INSERT INTO black_url_list(file_url,dwStrategy,cause) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE dwStrategy=?,cause=?",
		blackUrlList.FileUrl, blackUrlList.DwStrategy,blackUrlList.Cause, blackUrlList.DwStrategy, blackUrlList.Cause)

	zap.L().Debug(fmt.Sprint("INSERT INTO black_url_list(file_url,dwStrategy,cause) VALUES (%s, %d, %d) ON DUPLICATE KEY UPDATE dwStrategy=%d,cause=%d", blackUrlList.FileUrl, blackUrlList.DwStrategy,blackUrlList.Cause, blackUrlList.DwStrategy, blackUrlList.Cause))
	if err != nil {
		zap.L().Error(err.Error())
	}
	return err
}

func DeleteBlackUrl(fileUrl string) error {
	_,err := db.DBShub.Exec("DELETE FROM black_url_list WHERE file_url=?",fileUrl)
	zap.L().Debug(fmt.Sprint("DELETE FROM black_url_list WHERE file_url=%s", fileUrl))
	if err != nil {
		zap.L().Error(err.Error())
	}
	return err
}

func UpdateBlackUrl(blackUrlList BlackUrlList) error {
	return AddBlackUrl(blackUrlList)
}
