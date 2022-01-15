package Model

type Student struct {
	ID          uint `gorm:"primary_key"`
	StudentId   int
	ReportedDay int    //上一次的打卡日,表示这个月的第几天
	IsGraduate  string //
}

type Message map[string]interface {
	//code int
	//message string
	//data string
}
