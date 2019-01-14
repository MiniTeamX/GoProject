package main

import (
	"check_service/controllers"
	"github.com/gin-gonic/gin"
)

func main () {
	r := gin.Default()
	r.GET("/query_keyword_status", controllers.QueryKeywordStatus)
	r.GET("/query_gcid_status",controllers.QueryGcidStatus)
	r.GET("/query_uncheck", controllers.QueryUnCheck)
	r.GET("/add_human_check", controllers.AddHumanCheck)
	r.Run(":8088")
}

