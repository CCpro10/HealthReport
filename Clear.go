package main

import (
	"main/Model"
	"time"
)

func Clear() {

	Model.InitMySQL()

	Model.DB.Model(&Model.Student{}).Where("reported_day=?", time.Now().Day()).Update("reported_day", 0)
}

func main() {
	Clear()
}
