package controllers

import (
"database/sql"
"github.com/gin-gonic/gin"
"net/http"
"oc_res_control/models"
)

type BlackEmulehashListResponse struct {
	Emulehash    string     `json:"emulehash"`
	DwStrategy  int64      `json:"dw_strategy"`
	Cause       int64      `json:"cause"`
}


func GetBlackEmulehashListHandle(c *gin.Context) {
	emulehash := c.DefaultQuery("emulehash", "")
	blackEmulehashListResponse := BlackEmulehashListResponse{ Emulehash: emulehash}

	blackEmulehashList,err := models.GetBlackEmulehash(emulehash)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{
				"errcode":400,
				"msg":"请求的资源不存在",
				"result_info": blackEmulehashListResponse,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"errcode": 500 ,
				"msg":"数据库查询出错",
				"result_info": blackEmulehashListResponse,
			})
			return
		}
	} else {
		blackEmulehashListResponse.DwStrategy = blackEmulehashList.DwStrategy
		blackEmulehashListResponse.Cause = blackEmulehashList.Cause

		c.JSON(http.StatusOK, gin.H {
			"errcode":200,
			"msg":"SUCCESS",
			"result_info":blackEmulehashListResponse,
		})
		return
	}
}

func AddBlackEmulehashListHandle(c *gin.Context) {
	var blackEmulehashListRequest models.BlackEmulehashList
	if err := c.ShouldBindJSON(&blackEmulehashListRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"errcode":400,
			"msg":"JSON请求参数解析出错",
			"result_info":"",
		})
		return
	}
	err := models.AddBlackEmulehash(blackEmulehashListRequest)
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


func DeleteBlackEmulehashListHandle(c *gin.Context) {
	emulehash := c.DefaultQuery("emulehash", "")
	if err := models.DeleteBlackEmulehash(emulehash); err != nil {
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