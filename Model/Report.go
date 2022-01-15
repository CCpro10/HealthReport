package Model

import (
	"encoding/json"
	"errors"

	"io/ioutil"
	_ "log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func Report(student_id int) error {
	//发送打卡请求
	Url := "http://jc.ncu.edu.cn/gate/student/signIn"
	studentid := strconv.Itoa(student_id)

	payload := url.Values{}
	payload.Set("inChina", "是")
	payload.Set("addressProvince", "江西省")
	payload.Set("addressCity", "南昌市")
	payload.Set("temperatureStatus", "正常")
	payload.Set("isIll", "否")
	payload.Set("closeHb", "否")
	payload.Set("closeIll", "是")
	payload.Set("healthDetail", "无异常")
	payload.Set("isIsolation", "否")
	payload.Set("isolationPlace", "无")
	payload.Set("userId", studentid)
	payload.Set("addressInfo", "南昌大学")
	payload.Set("isGraduate", "否")
	payload.Set("healthStatus", "无异常")
	payload.Set("isIsolate", "否")
	payload.Set("isolatePlace", "无")

	//response, err := http.PostForm(targetUrl, payload)

	req, _ := http.NewRequest("POST", Url, strings.NewReader(payload.Encode()))
	req.Header.Set("token", Token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, _ := http.DefaultClient.Do(req)

	//检查返回体
	body, _ := ioutil.ReadAll(response.Body)
	var message Message
	_ = json.Unmarshal(body, &message)
	//log.Println(message)

	if message["code"] != "0" {
		return errors.New("打卡失败")
	}

	UpdateReportedDay(student_id)
	defer response.Body.Close()
	return nil

	//如果成功,把打卡时间修改为当天

}
