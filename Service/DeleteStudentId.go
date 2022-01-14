package Service

import (
	"HealthReport/Model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func DeleteStudentId(c  * gin.Context)  {
	id,_:=c.GetPostForm("student_id")

	student_id,err:=strconv.Atoi(id);if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"msg": "请输入正确的学号"})
		return
	}

	var student Model.Student
	Model.DB.Where("student_id=? ", student_id).First(&student)
	if student.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "此学号已取消自动健康打卡"})
		return
	} else {
		err= Model.DB.Where("student_id=?",student_id).Delete(&student).Error;if err!=nil{
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "学号删除失败,请重试"})
			return
		}else {
			c.JSON(http.StatusOK, gin.H{"msg": "取消自动打卡成功"})
			return
		}

	}
}
