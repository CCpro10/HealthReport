package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	_ "log"
	"main/Model"
	"net/http"
	"net/url"
	"strings"
)

func IsGraduate(startYear int) string {
	if startYear <= 2018 {
		return "是"
	} else {
		return "否"
	}
}

//通过id打卡,在打卡前会先登录
func LoginAndReportById(id uint) error {
	//登录地址
	LoginUrl := "http://jc.ncu.edu.cn/system/auth/loginByToken"

	//打卡地址
	Url := "http://jc.ncu.edu.cn/gate/student/signIn"

	student := Model.GetStudentInfoById(id)

	//打卡前先登录
	LoginReq, _ := http.NewRequest("GET", LoginUrl, strings.NewReader(""))
	LoginReq.Header.Set("token", student.Token)
	_, _ = http.DefaultClient.Do(LoginReq)

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
	payload.Set("userId", student.StudentId)
	payload.Set("addressInfo", student.AddressInfo)
	payload.Set("isGraduate", IsGraduate(student.StartYear))
	payload.Set("healthStatus", "无异常")
	payload.Set("isIsolate", "否")
	payload.Set("isolatePlace", "无")

	req, _ := http.NewRequest("POST", Url, strings.NewReader(payload.Encode()))
	req.Header.Set("token", student.Token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var message Model.Message
	_ = json.Unmarshal(body, &message)

	if message["code"] != "0" {
		return errors.New("打卡失败")
	}

	Model.UpdateReportedDayById(id)

	return nil

	//如果成功,把打卡时间修改为当天

}
