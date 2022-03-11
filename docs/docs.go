// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/report": {
            "post": {
                "description": "!!!!!!!!!先点击右边的 Try it out ,然后输入网址 ,再点击下方的Execute(执行),即可开始每天自动打卡,点击后查看下面的信息看看有没有打卡成功",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "获取网址"
                ],
                "summary": "点这里开始自动健康打卡",
                "parameters": [
                    {
                        "type": "string",
                        "description": "这里下面填健康打卡界面的网址: 具体看图片操作",
                        "name": "Url",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "这里下面填打卡的详细地址, 可以不填, 默认为:江西省南昌大学",
                        "name": "AddressInfo",
                        "in": "formData"
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "先点击右上角的 try it out 输入网址 ,再点击下方的Execute(执行)即可帮您取消每天自动打卡,点击后查看下面的信息看看有没有打卡成功",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "获取网址"
                ],
                "summary": "如果你不再想要每天自动健康打卡了,点这里取消自动健康打卡",
                "parameters": [
                    {
                        "type": "string",
                        "description": "这里下面填健康打卡界面的网址, 进入每日健康打卡页面, 点击右上角, 再点击复制链接",
                        "name": "Url",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "tags": [
        {
            "description": "###  这里按步骤打开后,再按企业微信或微信上的提示复制\r\n\u003cdiv class=\"half\" style=\"text-align: center;\"\u003e\u003e\r\n    \u003cimg src=\"http://incu-campus-num.ncuos.com/health_report/b37b3cad8e3fb12fb6e0736bcf35355.jpg?x-oss-process=image/resize,m_lfit,h_200,w_200\" alt=\"提示\"\u003e\r\n    \u003cimg src=\"http://incu-campus-num.ncuos.com/health_report/7928c85e54dbde094066e4a5f1ae6a4.jpg?x-oss-process=image/resize,m_lfit,h_400,w_400\" alt=\"提示\"\u003e\r\n    \u003cimg src=\"http://incu-campus-num.ncuos.com/health_report/b7127ddb8a11df661dcba50bbb76cbf.jpg?x-oss-process=image/resize,m_lfit,h_400,w_400\" alt=\"提示\"\u003e\r\n\u003c/div\u003e",
            "name": "获取网址"
        }
    ]
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "测试版 v0.96",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "企业微信自动健康打卡脚本",
	Description: "这是一个获取你的网址, 然后就可以帮你每天健康打卡的脚本~ (以后辅导员再也不用催我打卡啦~)\n任何NCU的同学都可以使用",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
