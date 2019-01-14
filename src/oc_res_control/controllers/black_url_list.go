package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"oc_res_control/models"
)


type BlackUrlListResponse struct {
	FileUrl     string     `json:"file_url"`
	DwStrategy  int64      `json:"dw_strategy"`
	Cause       int64      `json:"cause"`
}


func GetBlackUrlListHandle(c *gin.Context) {
	fileUrl := c.DefaultQuery("file_url", "")
	blackUrlListResponse := BlackUrlListResponse{ FileUrl:fileUrl}
	blackUrlList,err := models.GetBlackUrl(fileUrl)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{
				"errcode":400,
				"msg":"请求的资源不存在",
				"result_info": blackUrlListResponse,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"errcode": 500 ,
				"msg":"数据库查询出错",
				"result_info": blackUrlListResponse,
			})
			return
		}
	} else {
		blackUrlListResponse.DwStrategy = blackUrlList.DwStrategy
		blackUrlListResponse.Cause = blackUrlList.Cause

		c.JSON(http.StatusOK, gin.H {
			"errcode":200,
			"msg":"SUCCESS",
			"result_info":blackUrlListResponse,
		})
		return
	}
}

func AddBlackUrlListHandle(c *gin.Context) {
	 var blackUrlListRequest models.BlackUrlList
	 if err := c.ShouldBindJSON(&blackUrlListRequest); err != nil {
	 	c.JSON(http.StatusOK, gin.H{
	 		"errcode":400,
	 		"msg":"JSON请求参数解析出错",
	 		"result_info":"",
		})
	 	return
	 }
	 err := models.AddBlackUrl(blackUrlListRequest)
	 if err != nil {
	 	c.JSON(http.StatusOK, gin.H{
	 		"errcode":500,
	 		"msg":"数据库查询出错",
	 		"result_info":"",
		})
	 	return
	 }

	 c.JSON(http.StatusOK, gin.H{
		 "errcode":200,
		 "msg":"SUCCESS",
		 "result_info":"",
	 })
}


func DeleteBlackUrlListHandle(c *gin.Context) {
	fileUrl := c.DefaultQuery("file_url", "")
	if err := models.DeleteBlackUrl(fileUrl); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"errcode":500,
			"msg":"数据库查询出错",
			"result_info":"",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errcode":200,
		"msg":"SUCCESS",
		"result_info":"",
	})
}

