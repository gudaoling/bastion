// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-06-23 20:26:18.397307 +0800 CST m=+0.049221151

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/nanzm",
        "contact": {
            "name": "nan",
            "url": "https://nancode.cn",
            "email": "msg@nancode.cn"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/stat/admin/create": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "创建管理员"
            }
        },
        "/api/stat/admin/info": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "管理员信息"
            }
        },
        "/api/stat/admin/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "所有管理员"
            }
        },
        "/api/stat/admin/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "管理员登录"
            }
        },
        "/api/stat/admin/update": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "管理员信息更新"
            }
        },
        "/api/stat/behavior": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "上报打点"
            }
        },
        "/api/stat/behavior/:id": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "根据id查找打点"
            }
        },
        "/api/stat/behaviors": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "查找所有打点"
            }
        },
        "/api/stat/device/:uid": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "根据id查找错误"
            }
        },
        "/api/stat/error/:id": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "根据id查找错误"
            }
        },
        "/api/stat/errors": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "查找所有错误"
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "errlock上报"
            }
        },
        "/api/stat/project": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "创建项目"
            }
        },
        "/api/stat/project/:id": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "查找项目"
            }
        },
        "/api/stat/projects": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "查找所有项目"
            }
        },
        "/go/sys/info": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "系统信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "response.Response": {
            "type": "object",
            "properties": {
                "_trace_id": {
                    "type": "object"
                },
                "_ts": {
                    "type": "integer"
                },
                "code": {
                    "type": "integer"
                },
                "error": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                },
                "result": {
                    "type": "object"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    }
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "bastion server",
	Description: "personal projects",
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
	swag.Register(swag.Name, &s{})
}
