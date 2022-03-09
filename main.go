package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"main/Model"
	"main/Service"
	"main/api"
	"main/config"
	_ "main/docs" //必需
	"main/middleware"
)

// @title           自动健康打卡脚本
// @version         测试版 v0.9
// @description     可以帮NCU students每天健康打卡的脚本
func main() {
	log.SetFlags(log.Lshortfile)

	Model.InitMySQL()

	r := gin.Default()
	r.Use(middleware.Cors())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.POST("/report", api.BeginReport)

	r.DELETE("/report", api.EndReport)

	c := cron.New()
	_ = c.AddFunc("* */5 4-23 * * *", Service.ReportAll)
	_ = c.AddFunc("* * 0-4/5 * * *", Service.ReportInOrder)

	c.Start()

	if err := r.Run(config.Config.Server.Port); err != nil {
		panic(err)
	}

	//select {}
}
