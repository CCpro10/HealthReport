package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	}
	MYSQL struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Addr     string `yaml:"addr"`
		Database string `yaml:"database"`
	}
	Server struct {
		Port string `yaml:"port"`
	}
}

//获取配置
func GetConf() *Conf {
	var c = Conf{}
	yamlFile, err := ioutil.ReadFile("./config/conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return &c
}

//获取配置文件
var Config = GetConf()
