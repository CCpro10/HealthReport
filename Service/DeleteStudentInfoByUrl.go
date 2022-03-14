package Service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"main/Model"
	"net/http"
	"net/url"
	"strings"
)

//通过url删除学生的自动打卡信息
func DeleteStudentInfoByUrl(URL string) (studentId string, e error) {
	Url, _ := url.Parse(URL)
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
		e = errors.New("你复制的网址失效啦(用浏览器打开网站再复制会失效哦哦!)    请重新按图片步骤,先退出,重新进入 a学生疫情常态化管理 复制网址")
		return
	}

	studentId = m.Data.UserId
	Model.DB.Where("student_id=?", studentId).Delete(&Model.Student{})

	return studentId, nil

}
