package controller

import (
	"HealthReport/Model"
	"time"
)

//到晚上过了12点就运行这个,给所有人健康打卡
func ReportAll() {

	student := Model.Student{}
	Model.DB.Last(&student)
	count := student.ID

	var i uint
	for i = 1; i <= count; i++ {
		StudentId, err := Model.GetStudentIdByID(i)
		if err == nil { //没有错误就打卡,有错误则跳过
			_ = Model.Report(StudentId)
		}

	}
}

//对未打卡的同学打卡
func Report() {
	//找到今天还未打卡的同学
	t := time.Now().Day()
	var students []Model.Student
	Model.DB.Where("reported_day != ?", t).Find(&students)
	//log.Println(students)
	for _, student := range students {
		studentid := student.StudentId
		//log.Println(studentid)
		_ = Model.Report(studentid)
	}

}
