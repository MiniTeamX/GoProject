package controllers

import (
"database/sql"
"github.com/gin-gonic/gin"
"net/http"
"oc_res_control/models"
)


type BlackInfohashListResponse struct {
	Infohash    string     `json:"infohash"`
	DwStrategy  int64      `json:"dw_strategy"`
	Cause       int64      `json:"cause"`
}


func GetBlackInfohashListHandle(c *gin.Context) {
	infohash := c.DefaultQuery("infohash", "")
	blackInfohashListResponse := BlackInfohashListResponse{ Infohash: infohash}

	blackInfohashList,err := models.GetBlackInfohash(infohash)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{
				"errcode":400,
				"msg":"请求的资源不存在",
				"result_info": blackInfohashListResponse,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"errcode": 500 ,
				"msg":"数据库查询出错",
				"result_info": blackInfohashListResponse,
			})
			return
		}
	} else {
		blackInfohashListResponse.DwStrategy = blackInfohashList.DwStrategy
		blackInfohashListResponse.Cause = blackInfohashList.Cause

		c.JSON(http.StatusOK, gin.H {
			"errcode":200,
			"msg":"SUCCESS",
			"result_info":blackInfohashListResponse,
		})
		return
	}
}

func AddBlackInfohashListHandle(c *gin.Context) {
	var blackInfohashListRequest models.BlackInfohashList
	if err := c.ShouldBindJSON(&blackInfohashListRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"errcode":400,
			"msg":"JSON请求参数解析出错",
			"result_info":"",
		})
		return
	}
	err := models.AddBlackInfohash(blackInfohashListRequest)
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


func DeleteBlackInfohashListHandle(c *gin.Context) {
	infohash := c.DefaultQuery("infohash", "")
	if err := models.DeleteBlackInfohash(infohash); err != nil {
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
