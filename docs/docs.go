// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1beta1/book": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book"
                ],
                "summary": "获取书籍列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Response"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book"
                ],
                "summary": "新增书籍",
                "parameters": [
                    {
                        "description": "测试请求json参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Response"
                        }
                    }
                }
            }
        },
        "/api/v1beta1/example": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "Get请求",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Response"
                        }
                    }
                }
            }
        },
        "/api/v1beta1/example/err/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "Err请求",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id值",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Response"
                        }
                    }
                }
            }
        },
        "/api/v1beta1/example/form": {
            "post": {
                "description": "处理application/x-www-form-urlencoded类型的POST请求",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "form表单请求",
                "parameters": [
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Response"
                        }
                    }
                }
            }
        },
        "/api/v1beta1/example/json": {
            "post": {
                "description": "邮箱、用户名校验",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "JSON参数请求",
                "parameters": [
                    {
                        "description": "测试请求json参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.JSONBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Response"
                        }
                    }
                }
            }
        },
        "/api/v1beta1/example/query": {
            "get": {
                "description": "查询参数，匹配 query?id=xxx",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "query参数查询",
                "parameters": [
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Response"
                        }
                    }
                }
            }
        },
        "/api/v1beta1/example/uri/{id}": {
            "get": {
                "description": "路径参数，匹配 /uri/{id}",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "uri参数请求",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id值",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "req.Book": {
            "type": "object",
            "required": [
                "name",
                "url"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "req.JSONBody": {
            "type": "object",
            "required": [
                "email",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "admin@qq.com"
                },
                "username": {
                    "type": "string",
                    "example": "admin"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	Title:       "接口文档",
	Description: "🐐 基于 gin + gorm 的轻量级工程项目",
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
