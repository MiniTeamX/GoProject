package models

import (
	"check_service/database"
	"database/sql"
	"log"
)

type GcidInfoLiXian struct {
    Gcid        string
    Cid         string
    FileName    string
    FileSize    int64
    Flag        int64
    InsertTime  int64
    UpdateTime  int64
    User        string
    Lx_flag     int64
}



var lixian_table_suffix = [16]string{"0","1","2","3","4","5","6","7","8","9","A","B","C","D","E","F"}

func GetUncheckList(lx_flag int64, limit int64)([]GcidInfoLiXian,error) {
	tableNamePrefix := "gcid_info_lixian_"
    gcidInfoLiXians := make([]GcidInfoLiXian,0)
    var err error
    for i:= 0; i < 16; i++ {
        for j := 0; j < 16; j++ {
            tableName := tableNamePrefix + lixian_table_suffix[i] + lixian_table_suffix[j]
            rows,err := database.DBGcidLixian.Query("SELECT gcid,filename,insert_time FROM " + tableName + " WHERE flag=9 and lx_flag=?",lx_flag)
            if err != nil {
                log.Fatalln(err.Error())
            }
            defer rows.Close()
            for rows.Next() {
                var gcidInfoLixian GcidInfoLiXian
                var fileNameNull sql.NullString
                err := rows.Scan(&gcidInfoLixian.Gcid, &fileNameNull, &gcidInfoLixian.InsertTime)
                //如果数据库没数据，怎么处理错误?
                if err != nil {
                    log.Println(err.Error())
                }
                if fileNameNull.Valid {
                    gcidInfoLixian.FileName = fileNameNull.String
                }
                gcidInfoLiXians = append(gcidInfoLiXians, gcidInfoLixian)
                if (len(gcidInfoLiXians) >= int(limit)) {
                    return gcidInfoLiXians,err
                }
            }
        }
    }
	return gcidInfoLiXians,err
}


func AddHumanCheck(gcidInfo GcidInfoLiXian) (error) {
    subgcid := string([]rune(gcidInfo.Gcid)[0:2])
    //flag=1 or flag=2 do not update
    _, err := database.DBGcidLixian.Exec("INSERT INTO  gcid_info_lixian_" + subgcid + "(gcid, cid, filename, filesize, flag, insert_time, user, lx_flag) " +
        "values (?, ?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE " +
        "flag=IF(flag!=1 and flag!=2,?,flag)," +
        "lx_flag=IF(flag!=1 and flag!=2,0,lx_flag)," +
        "update_time=IF(flag!=1 and flag!=2,?,update_time)",
        gcidInfo.Gcid, gcidInfo.Cid, gcidInfo.FileName, gcidInfo.FileSize, gcidInfo.Flag, gcidInfo.InsertTime, gcidInfo.User, 0, gcidInfo.Flag, gcidInfo.InsertTime)
    if err != nil {
        log.Println(err.Error())
    }
    return err
}
