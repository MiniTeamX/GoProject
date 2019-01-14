package controllers

import (
	"check_service/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//机器审核没结果需要人工审核
func AddHumanCheck(c *gin.Context) {
	gcid := c.Query("gcid")
	if gcid == "" || len(gcid) != 40 {
		c.JSON(http.StatusOK, gin.H{
			"result":-1,
			"message":"gcid is null or the len is not equal to 40",
		})
		return
	}
	cid := c.Query("cid")
	fileName := c.DefaultQuery("filename", "")
	fileSize := c.DefaultQuery("filesize", "0")
	user := c.DefaultQuery("user", "jiqishencha")
	flagString := c.DefaultQuery("flag","0")
	if flagString != "0" && flagString != "1" {
		c.JSON(http.StatusOK, gin.H{
			"result":-1,
			"message":"flag can only be 0 or 1",
		})
		return
	}

	var flag int64
	flagInt,_ := strconv.ParseInt(flagString, 10, 64)
	//在filter_gcid_lixian这张表中，flag=9代表人工审核。flag=4表示审核的结果不确定
	if flagInt == 0 {
		flag = 9
	} else {
		flag = 4
	}
	timeStamp := time.Now().Unix()
	var gcidInfo models.GcidInfoLiXian
	gcidInfo.Gcid = gcid
	gcidInfo.Cid = cid
	gcidInfo.FileName = fileName
	gcidInfo.FileSize,_= strconv.ParseInt(fileSize, 10, 64)
	gcidInfo.Flag = flag
	gcidInfo.User = user
	gcidInfo.InsertTime = timeStamp

	err := models.AddHumanCheck(gcidInfo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"result":-1,
			"message" : err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result" : 0,
			"message" : "",
		})
	}
}
