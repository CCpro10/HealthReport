package main

import (
	"main/Model"
	"net/http"
	"strconv"
)

var C int64

func main() {

	Model.InitMySQL()

	http.HandleFunc("/", CountHandler)
	_ = http.ListenAndServe(":8889", nil)
}

// handler函数
func CountHandler(w http.ResponseWriter, r *http.Request) {
	Model.DB.Model(&Model.Student{}).Count(&C)
	w.Write([]byte("总共有" + strconv.FormatInt(C, 10) + "人使用了自动健康打卡"))
}
