package Model

import "time"

func GetStudentIdByID(id uint) (int,error){
    var student Student
	err:=DB.Select("student_id").Where("ID=?",id).Find(&student).Error;if err!=nil{
		return student.StudentId, err
	}else {
		return 0 ,err
	}

}

//更新打卡时间
func NewReportedDay(student_id int)  {
	t:=time.Now().Day()
	DB.Model(&Student{}).Where("student_id=?",student_id).Update("reported_day",t)
}
