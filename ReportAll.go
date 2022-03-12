package main

import (
	"main/Model"
	"main/Service"
)

func main() {
	Model.InitMySQL()
	Service.ReportAll()
}
