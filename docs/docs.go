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
        "termsOfService": "http://swagger.io/terms/w",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/login": {
            "post": {
                "description": "用户登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "邮箱 密码必填",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.DataResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.DataResponse"
                        }
                    }
                }
            }
        },
        "/user": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "注册的表单",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.DataResponse"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.DataResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.DataResponse"
                        }
                    }
                }
            }
        },
        "/user/rule/add": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rule"
                ],
                "summary": "增加规则",
                "parameters": [
                    {
                        "description": "必须给定server_id 、content，uri_id可选 type为get\\post\\ua\\header\\cookie之一",
                        "name": "rulePageForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.RulePage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.DataResponse"
                        }
                    }
                }
            }
        },
        "/user/rule/delete": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rule"
                ],
                "summary": "删除规则",
                "parameters": [
                    {
                        "description": "必须给定server_id ，uri_id可选 type为get\\post\\ua\\header\\cookie之一",
                        "name": "rulePageForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.RulePage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.DataResponse"
                        }
                    }
                }
            }
        },
        "/user/rule/get": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rule"
                ],
                "summary": "获取规则",
                "parameters": [
                    {
                        "description": "page为页号，limit为一页的最大数量，类型为get\\post\\ua\\header\\cookie之一",
                        "name": "rulePageForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.RulePage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.DataResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/common.Rule"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/server/delete": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "server"
                ],
                "summary": "删除服务",
                "parameters": [
                    {
                        "description": "server_id为必要",
                        "name": "servers",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.GetServerForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.DataResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.DataResponse"
                        }
                    }
                }
            }
        },
        "/user/server/get": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "server"
                ],
                "summary": "添加Server",
                "parameters": [
                    {
                        "description": "服务器列表",
                        "name": "servers",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/common.Server"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.DataResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/common.Server"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.DataResponse"
                        }
                    }
                }
            }
        },
        "/user/switch/change": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "switch"
                ],
                "summary": "修改Server Switch或者URI Switch",
                "parameters": [
                    {
                        "description": "必须填写config_name；不给uri_id则修改server",
                        "name": "switchForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.SwitchOperation"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.DataResponse"
                        }
                    }
                }
            }
        },
        "/user/switch/waf": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "switch"
                ],
                "summary": "查询waf开关",
                "parameters": [
                    {
                        "type": "string",
                        "description": "服务器id",
                        "name": "server_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.DataResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "switch"
                ],
                "summary": "控制waf开关",
                "parameters": [
                    {
                        "description": "可以不填写config_name",
                        "name": "switchForm",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.SwitchOperation"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.DataResponse"
                        }
                    }
                }
            }
        },
        "/user/uri/add": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "uri"
                ],
                "summary": "增加URI",
                "parameters": [
                    {
                        "description": "server_id、path必填",
                        "name": "uri",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.URI"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.DataResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/common.Server"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.DataResponse"
                        }
                    }
                }
            }
        },
        "/user/uri/delete": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "uri"
                ],
                "summary": "删除URI",
                "parameters": [
                    {
                        "description": "server_id、id必填",
                        "name": "uri",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.URI"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.DataResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.DataResponse"
                        }
                    }
                }
            }
        },
        "/user/uri/get": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "uri"
                ],
                "summary": "获取URI",
                "parameters": [
                    {
                        "description": "server_id必填",
                        "name": "uri",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.GetURIForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.DataResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/common.URI"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.DataResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.DataResponse": {
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
        "common.GetServerForm": {
            "type": "object",
            "properties": {
                "server_id": {
                    "type": "integer"
                }
            }
        },
        "common.GetURIForm": {
            "type": "object",
            "properties": {
                "server_id": {
                    "type": "integer"
                }
            }
        },
        "common.Option": {
            "type": "object",
            "properties": {
                "ccrate": {
                    "type": "string"
                },
                "proxyPass": {
                    "type": "string"
                }
            }
        },
        "common.Rule": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "flag": {
                    "type": "string"
                },
                "hit": {
                    "type": "integer"
                },
                "host": {
                    "type": "string"
                },
                "server_id": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "uri": {
                    "type": "string"
                },
                "uri_id": {
                    "type": "integer"
                }
            }
        },
        "common.RulePage": {
            "type": "object",
            "properties": {
                "flag": {
                    "type": "string"
                },
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "server_id": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "uri_id": {
                    "type": "integer"
                }
            }
        },
        "common.Server": {
            "type": "object",
            "properties": {
                "domain": {
                    "type": "string"
                },
                "ip": {
                    "type": "string"
                },
                "option": {
                    "type": "object",
                    "$ref": "#/definitions/common.Option"
                },
                "uri": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/common.URI"
                    }
                },
                "user_id": {
                    "type": "integer"
                },
                "waf_status": {
                    "type": "boolean"
                }
            }
        },
        "common.SwitchOperation": {
            "type": "object",
            "properties": {
                "config_name": {
                    "type": "string"
                },
                "config_value": {
                    "type": "boolean"
                },
                "server_id": {
                    "type": "integer"
                },
                "uri_id": {
                    "type": "integer"
                }
            }
        },
        "common.URI": {
            "type": "object",
            "properties": {
                "host": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "server_id": {
                    "type": "integer"
                }
            }
        },
        "common.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
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
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Scarlet Backend",
	Description: "This is a sample server celler server.",
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
