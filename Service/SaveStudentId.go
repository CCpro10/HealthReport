package Service

import (
	"HealthReport/Model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func SaveStudentId(c * gin.Context)  {
	id,_:=c.GetPostForm("student_id")

	student_id,err:=strconv.Atoi(id);if err!=nil{
		log.Println(err,id,student_id)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "请输入正确的学号"})
		return
	}

	var student Model.Student
	Model.DB.Where("student_id=? ", student_id).First(&student)
	if student.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "此学号已设为自动健康打卡"})
		return
	} else {

		student.StudentId=student_id
		//这里可能会显示上一个句子的未找到的错误
		err= Model.DB.Create(&student).Error;if err==nil||err!=gorm.ErrRecordNotFound{
				c.JSON(http.StatusOK, gin.H{"msg": "添加成功"})
				return
		}else {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "学号添加失败,请重试"})
			return
		}

		}
	}
