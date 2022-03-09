package Model

import (
	"time"
)

//更新打卡时间
func UpdateReportedDayById(id uint) {

	t := time.Now().Day()
	DB.Model(&Student{}).Where("id = ?", id).Update("reported_day", t)
}
