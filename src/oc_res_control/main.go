package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"oc_res_control/conf"
	"oc_res_control/controllers"
	"oc_res_control/database"
	"oc_res_control/util"
)

func main () {
	cfg := conf.NewConfig()
	if err := cfg.Load("conf/config.yaml"); err != nil {
		log.Fatalf("failed to load config file:%s", err)
	}
	//init log
	defer zap.L().Sync()
	if err := BuildLogger(&cfg.Logger); err != nil {
		log.Fatalf("failed to build log: %s", err)
	}

	//init db
	database.InitDb(cfg)
	zap.L().Info("init db success.")

	//init stat
	util.GlobalStat = &util.Statistics{}
	util.GlobalStat.Init(cfg)
	zap.L().Info("stat handle init success.")

	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	} else if cfg.Server.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
	}
	zap.L().Info(fmt.Sprintf("oc res_control service start! version: %s, Run mode:%s, port:%s",cfg.Server.Version, cfg.Server.Mode, cfg.Server.Port))

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/query_black_url", controllers.GetBlackUrlListHandle)
		v1.POST("/add_black_url", controllers.AddBlackUrlListHandle)
		v1.GET("/delete_black_url", controllers.DeleteBlackUrlListHandle)

		v1.GET("/query_black_infohash", controllers.GetBlackInfohashListHandle)
		v1.POST("/add_black_infohash", controllers.AddBlackInfohashListHandle)
		v1.GET("/delete_black_infohash", controllers.DeleteBlackInfohashListHandle)

		v1.GET("/query_black_emulehash", controllers.GetBlackEmulehashListHandle)
		v1.POST("/add_black_emulehash", controllers.AddBlackEmulehashListHandle)
		v1.GET("/delete_black_emulehash", controllers.DeleteBlackEmulehashListHandle)
	}

	router.Run(cfg.Server.Port)
}

