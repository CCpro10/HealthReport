package main

import (
	middleware "HealthReport/Middleware"
	"HealthReport/Model"
	"HealthReport/Service"
	"HealthReport/controller"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	Model.InitMySQL()

	r := gin.Default()
	r.Use(middleware.Cors())

	r.POST("/student_id", Service.SaveStudentId)
	r.POST("/delete_student_id", Service.DeleteStudentId)
	r.POST("/url", Service.GetToken)

	c := cron.New()
	_ = c.AddFunc("* */2 6-23 * * *", controller.Report)
	_ = c.AddFunc("* * 0-5/5 * * *", controller.ReportAll)

	c.Start()

	if err := r.Run(":9002"); err != nil {
		panic(err)
	}

	//select {}
}
