package controllers

import (
	"check_service/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
    "time"
)

type GcidInfoList struct {
    Gcid        string    `json:"gcid"`
    FileName    string    `json:"filename"`
    InsertTime  time.Time `json:"insert_time"`
}

func QueryUnCheck(c *gin.Context) {
	var err error
    gcidInfoLists := make([]GcidInfoList,0)
	lxFlagString := c.Query("lx_flag")
	limitString := c.DefaultQuery("limit", "250")
	lxFlag,err := strconv.ParseInt(lxFlagString,10,64)
	limit,err := strconv.ParseInt(limitString, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"error" : "lx_flag and limit must be int",
		})
		return
	}
	gcidInfoLiXians,_ := models.GetUncheckList(lxFlag, limit)
    var gcidInfoList GcidInfoList
    for _,v := range gcidInfoLiXians {
        gcidInfoList.Gcid = v.Gcid
        gcidInfoList.FileName = v.FileName
        gcidInfoList.InsertTime = time.Unix(v.InsertTime, 0)
        gcidInfoLists = append(gcidInfoLists, gcidInfoList)
    }
	c.JSON(http.StatusOK, gin.H{
		"result" : gcidInfoLists,
	})
}
