package main

import (
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
	r.POST("/student_id", Service.SaveStudentId)
	r.DELETE("/student_id", Service.DeleteStudentId)

	c := cron.New()
	_=c.AddFunc("* */2 6-23 * * *", controller.Report)
	_=c.AddFunc("* * 0-5 * * *",controller.ReportAll)

	c.Start()

	if err := r.Run(":9002"); err != nil {
		panic(err)
	}

	//select {}
}

