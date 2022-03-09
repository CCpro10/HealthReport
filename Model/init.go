package Model

import (
	"main/config"

	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//mysql,gorm配置
var DB *gorm.DB

func InitMySQL() {
	dsn := config.Config.MYSQL.Username + ":" +
		config.Config.MYSQL.Password + "@tcp(" +
		config.Config.MYSQL.Addr + ")/" +
		config.Config.MYSQL.Database + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) //这里用短变量声明会有歧义
	//DB, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//先创建表
	if err = DB.AutoMigrate(

		Student{},
	); err != nil {
		log.Panicln(err)
	}

}
