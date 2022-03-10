package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"main/Service"
	"main/utils"
	"net/http"
)

// @Summary 点这里开始自动健康打卡
// @Description !!!!!!!!!先点击右边的 Try it out ,然后输入网址 ,再点击下方的Execute(执行),即可开始每天自动打卡,点击后查看下面的信息看看有没有打卡成功
// @Tags    获取网址
// @Produce json
// @Param Url formData string true "这里下面填健康打卡界面的网址: 具体看图片操作"
// @Param AddressInfo formData string false "这里下面填打卡的详细地址, 可以不填, 默认为:江西省南昌大学"
// @Router /report [post]
func BeginReport(c *gin.Context) {
	url, _ := c.GetPostForm("Url")
	addressInfo, _ := c.GetPostForm("AddressInfo")

	//参数检验
	validate := validator.New()
	e := validate.Var(url, "required,url")
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"信息": "请输入正确的地址"})
		return
	}

	e = validate.Var(addressInfo, "max=150")
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"信息": "详细地址请不要填太长"})
		return
	}

	if addressInfo == "" {
		addressInfo = "江西省南昌大学"
	}

	//检查Url,通过URL获取token,生成并保存Student数据
	id, studentId, e := Service.SaveInfo(url, addressInfo)
	if e != nil {
		if e.Error() == "已经帮你健康打卡啦" {
			c.JSON(http.StatusOK, gin.H{"信息": "已经帮你健康打卡啦,打卡地址修改为:" + addressInfo})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"信息": e.Error()})
		return
	}

	//登录和打卡
	e = utils.LoginAndReportById(id)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"信息": e.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"信息": "学号为:" + studentId + "的同学,已开始帮您每天自动打卡"})
	return

}
