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

// @title           企业微信自动健康打卡脚本
// @version         测试版 v0.95
// @description     这是一个获取你的网址, 然后就可以帮你每天健康打卡的脚本~ (以后辅导员再也不用催我打卡啦~)
// @description     任何需要打卡的同学都可以使用(ncu)
// @Tag.name  获取网址
// @tag.description.markdown
func main() {
	log.SetFlags(log.Lshortfile)

	Model.InitMySQL()

	r := gin.Default()
	r.Use(middleware.Cors())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.POST("/report", api.BeginReport)

	r.DELETE("/report", api.EndReport)

	c := cron.New()
	_ = c.AddFunc("* */10 4-23 * * *", Service.ReportAll)
	_ = c.AddFunc("* * 0-4/4 * * *", Service.ReportInOrder)

	c.Start()

	if err := r.Run(config.Config.Server.Port); err != nil {
		panic(err)
	}

	//select {}
}

//## <center> 获取打卡界面的网址
//
//### <center> 这里打开后再复制 ![第一步](http://incu-campus-num.ncuos.com/health_report/b37b3cad8e3fb12fb6e0736bcf35355.jpg?x-oss-process=image/resize,m_lfit,h_80,w_80 "图片")
//### <center>  这里复制 ![第二步](http://incu-campus-num.ncuos.com/health_report/7928c85e54dbde094066e4a5f1ae6a4.jpg?x-oss-process=image/resize,m_lfit,h_140,w_140 "")
//###  <center> 或者是这里复制 ![第二步](http://incu-campus-num.ncuos.com/health_report/b7127ddb8a11df661dcba50bbb76cbf.jpg?x-oss-process=image/resize,m_lfit,h_150,w_150 "")
