package Service

import (
	"main/Model"
	"main/utils"
	"time"
)

//到晚上过了12点就运行这个,给所有人健康打卡
func ReportInOrder() {

	student := Model.Student{}
	Model.DB.Last(&student)
	count := student.ID //最后一个学生的id

	var id uint
	for id = 1; id <= count; id++ {

		if !Model.IsExistById(id) { //不存在则跳过,存在就打卡
			continue
		}

		if Model.AlreadyReport(id) { //已经打卡了就不用打卡了
			continue
		}
		_ = utils.LoginAndReportById(id)
		time.Sleep(time.Second / 10) //降低并发
	}
}

//对所有未打卡的同学打卡
func ReportAll() {
	//找到今天还未打卡的同学
	t := time.Now().Day()
	var students []Model.Student
	Model.DB.Where("reported_day != ?", t).Find(&students)

	for _, student := range students {
		_ = utils.LoginAndReportById(student.ID)
	}

}
