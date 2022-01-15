package Model

import (
	"errors"
	"time"
)

var Token = "eyJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJlY2hpc2FuIiwic3ViIjoiNTgwMTEyMDAxNiIsImlhdCI6MTY0MjIzNTY5Mn0.CSd4IbnL5zJadwwZosOfL6_M9O1-TRq1xUIS0m2LmXsCxcUGALcIHHuoeiB31exP7YZwU4l3N9guoxUhV6th4w"

func GetStudentIdByID(id uint) (int, error) {
	var student Student
	DB.Select("student_id").Where("ID = ?", id).Find(&student)
	if student.StudentId != 0 {
		return student.StudentId, nil
	} else {
		return 0, errors.New("查询失败")
	}

}

//更新打卡时间
func UpdateReportedDay(student_id int) {

	t := time.Now().Day()
	DB.Model(&Student{}).Where("student_id = ?", student_id).Update("reported_day", t)
}
