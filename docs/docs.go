// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/Huberyoung/blog-service",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/api/articles": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "文章名称",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.ArticleSwagger"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorcode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorcode.Error"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "文章名称",
                        "name": "title",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "maxLength": 65535,
                        "minLength": 100,
                        "description": "文章内容",
                        "name": "content",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "创建者",
                        "name": "created_by",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.ArticleSwagger"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorcode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorcode.Error"
                        }
                    }
                }
            }
        },
        "/api/api/articles/{id}": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "文章 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "文章名称",
                        "name": "title",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "maxLength": 65535,
                        "minLength": 100,
                        "description": "文章内容",
                        "name": "content",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "修改者",
                        "name": "modified_by",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.ArticleSwagger"
                            }
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorcode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorcode.Error"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "文章 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorcode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorcode.Error"
                        }
                    }
                }
            }
        },
        "/api/api/tags": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "maxLength": 100,
                        "type": "string",
                        "description": "标签名称",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "type": "integer",
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.TagSwagger"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorcode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorcode.Error"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "标签名称",
                        "name": "name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "创建者",
                        "name": "created_by",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/model.TagSwagger"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorcode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorcode.Error"
                        }
                    }
                }
            }
        },
        "/api/api/tags/{id}": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "标签 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "标签名称",
                        "name": "name",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "enum": [
                            0,
                            1
                        ],
                        "default": 1,
                        "description": "状态",
                        "name": "state",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "maxLength": 100,
                        "minLength": 3,
                        "description": "修改者",
                        "name": "modified_by",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.TagSwagger"
                            }
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorcode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorcode.Error"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "标签 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "请求错误",
                        "schema": {
                            "$ref": "#/definitions/errorcode.Error"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "$ref": "#/definitions/errorcode.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.Pager": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "total_rows": {
                    "type": "integer"
                }
            }
        },
        "errorcode.Error": {
            "type": "object"
        },
        "model.Article": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "cover_image_url": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "created_on": {
                    "type": "integer"
                },
                "delete_on": {
                    "type": "integer"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_del": {
                    "type": "integer"
                },
                "modified_by": {
                    "type": "string"
                },
                "modified_on": {
                    "type": "integer"
                },
                "state": {
                    "description": "state 状态 0 为禁用、1 为启用",
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.ArticleSwagger": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Article"
                    }
                },
                "paper": {
                    "$ref": "#/definitions/app.Pager"
                }
            }
        },
        "model.Tag": {
            "type": "object",
            "properties": {
                "created_by": {
                    "type": "string"
                },
                "created_on": {
                    "type": "integer"
                },
                "delete_on": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_del": {
                    "type": "integer"
                },
                "modified_by": {
                    "type": "string"
                },
                "modified_on": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "type": "integer"
                }
            }
        },
        "model.TagSwagger": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Tag"
                    }
                },
                "paper": {
                    "$ref": "#/definitions/app.Pager"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Hubery 的博客系统",
	Description:      "这只是一个测试.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
