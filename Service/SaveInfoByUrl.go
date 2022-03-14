package Service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"main/Model"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//获取token,如果此用户为新用户则会保存新用户信息
func SaveInfo(URL string, addr string) (id uint, studentId string, e error) {
	Url, _ := url.ParseRequestURI(URL)
	//获取query中的code
	q := Url.Query()
	code := q.Get("code")

	if code == "" {
		e = errors.New("输入的地址不正确哦")
		return
	}

	//从这个网站获取token
	Url2 := "http://jc.ncu.edu.cn/system/auth/getWebChat"

	//设置载荷为获取的code,生成一个http请求
	payload := url.Values{}
	payload.Set("code", code)
	req, _ := http.NewRequest("POST", Url2, strings.NewReader(payload.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var m Model.LoginMessage
	_ = json.Unmarshal(body, &m)

	if m.Code != "308" {
		e = errors.New("你复制的网址失效啦(不能用浏览器打开网站再复制哦!)\n请先退出,重新按图片步骤, 重新进入 a学生疫情常态化管理 复制网址")
		return
	}

	studentId = m.Data.UserId

	//已经存了此用户的信息
	if Model.IsExistByStudentId(studentId) {
		Model.DB.Model(&Model.Student{}).Where("student_id=?", studentId).Update("address_info", addr)
		e = errors.New("已经帮你健康打卡啦")
		return
	}

	startYear, _ := strconv.ParseInt(m.Data.Grade, 10, 64)
	s := Model.Student{
		Token:       response.Header.Get("Token"),
		StudentId:   studentId,
		AddressInfo: addr,
		StartYear:   int(startYear),
	}

	Model.DB.Create(&s)

	return s.ID, studentId, nil

}
