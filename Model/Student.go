package Model

type Student struct {
	ID  uint  `gorm:"primary_key"`
	StudentId  int
	ReportedDay int //上一次的打卡日,表示这个月的第几天
	isGraduate string
}

type Message struct {
	code int
	message string
}