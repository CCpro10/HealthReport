package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"main/Service"
	"net/http"
)

// @Summary 点这里取消自动健康打卡
// @Description 先点击右上角的 try it out 输入网址 ,再点击下方的Execute(执行)即可帮您取消每天自动打卡
// @Tags    健康打卡
// @Produce json
// @Param Url formData string true "这里填健康打卡界面的网站地址"
// @Success 200 {string} string "{"信息": "取消成功"}"
// @Router /report [delete]
func EndReport(c *gin.Context) {
	url, _ := c.GetPostForm("Url")

	//参数检验
	validate := validator.New()
	e := validate.Var(url, "required,url")
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"信息": "请输入正确的地址"})
		return
	}
	studentId, e := Service.DeleteStudentInfoByUrl(url)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"信息": e.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"信息": "学号为:" + studentId + "的同学,已帮您结束自动打卡"})
	return

}
