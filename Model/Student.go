package Model

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	ID          uint   `gorm:"primary_key"`
	Token       string `json:"token"`
	StudentId   string `json:"student_id"`   //学号
	AddressInfo string `json:"address_info"` //打卡详细地址
	StartYear   int    `json:"start_year"`   //入学年份,如2020
	ReportedDay int    `json:"reported_day"` //上一次的打卡日,表示这个月的第几天
	CreatedAt   time.Time
}

func GetStudentInfoById(id uint) *Student {
	var s Student
	DB.First(&s, id)
	return &s
}

//通过id判断学生是否存在
func IsExistById(id uint) bool {
	var s Student
	e := DB.First(&s, id).Error
	if e == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

//通过id判断学生是否存在
func IsExistByStudentId(studentId string) bool {
	var s Student
	e := DB.Where("student_id=?", studentId).First(&s).Error
	if e == gorm.ErrRecordNotFound {
		return false
	}
	return true
}
