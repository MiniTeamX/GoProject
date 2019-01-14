package controllers

import (
	"check_service/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type  GcidInfo struct {
	Gcid        string      `json:"gcid"`
	Status      string      `json:"status"`
	User        string      `json:"user"`
	BeginTime   time.Time   `json:"begin_time"`
	EndTime     time.Time   `json:"end_time"`
}

func QueryGcidStatus(c *gin.Context) {
	var err error
	var gcidInfo GcidInfo
	gcid := c.Query("gcid")

	_,err = models.GetBlackGcidList(gcid)
	if err != sql.ErrNoRows {
		gcidInfo.Gcid = gcid
		gcidInfo.Status = "T"
		gcidInfo.User = ""
		gcidInfo.BeginTime = time.Time{}
		gcidInfo.EndTime = time.Time{}
	} else {
		var gcidIndex models.GcidIndex
		gcidIndex,err = models.GetGcidIndex(gcid)
		if err == sql.ErrNoRows {
			gcidInfo.Gcid = gcid
			gcidInfo.Status = "N"
			gcidInfo.User = ""
			gcidInfo.BeginTime = time.Time{}
			gcidInfo.EndTime = time.Time{}
		} else if err != nil {
			gcidInfo.Gcid = gcid
			if gcidIndex.Flag == "B" {
				gcidInfo.Status = "S"
			} else {
				gcidInfo.Status = gcidIndex.Flag
			}
			gcidInfo.User = gcidIndex.Checker
			gcidInfo.BeginTime = gcidIndex.Ts
		 	gcidInfo.EndTime = time.Time{}
		}
	}
	c.JSON(http.StatusOK, gcidInfo)
}
