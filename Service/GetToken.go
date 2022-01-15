package Service

import (
	"HealthReport/Model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func GetToken(c *gin.Context) {
	Url1, _ := c.GetPostForm("url")
	Url, err := url.ParseRequestURI(Url1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "地址解析错误"})
		return
	}
	q := Url.Query()
	code := q.Get("code")

	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "地址不正确哦"})
		return
	}

	//从这个网站获取token
	Url2 := "http://jc.ncu.edu.cn/system/auth/getWebChat"
	payload := url.Values{}
	payload.Set("code", code)
	req, _ := http.NewRequest("POST", Url2, strings.NewReader(payload.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(response.Body)

	var message Model.Message
	_ = json.Unmarshal(body, &message)
	log.Printf("%+v", message)
	if message["code"] != "308" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "获取信息失败,很可能是网站已过期"})
		return
	}

	Model.Token = response.Header.Get("Token")

	c.JSON(http.StatusOK, gin.H{"msg": "成功啦"})
	defer response.Body.Close()

}
