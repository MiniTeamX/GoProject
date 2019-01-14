package controllers

import (
	"check_service/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryKeywordStatus(c *gin.Context) {
	keyword := c.Query("keyword")
	k, err:= models.GetKeyWordsByKeyWord(keyword)
	var status string
	if err == sql.ErrNoRows {
		status = "N"
	} else {
		status = k.Flag
	}
	c.JSON(http.StatusOK, gin.H{
			"keyword": keyword,
			"status": status,
			"ts":k.Ts,
		})
}

