// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/event": {
            "get": {
                "description": "获取活动列表",
                "tags": [
                    "活动"
                ],
                "summary": "活动",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "请输入当前页，默认第一页",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页大小",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "更新活动信息",
                "tags": [
                    "活动"
                ],
                "summary": "更新活动信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "title",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "primogems",
                        "name": "primogems",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "award",
                        "name": "award",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "preTask",
                        "name": "preTask",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "coPlay",
                        "name": "coPlay",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "todoNum",
                        "name": "todoNum",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "startTime",
                        "name": "startTime",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "endTime",
                        "name": "endTime",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "新增活动信息",
                "tags": [
                    "活动"
                ],
                "summary": "新增活动信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "title",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "primogems",
                        "name": "primogems",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "award",
                        "name": "award",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "preTask",
                        "name": "preTask",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "coPlay",
                        "name": "coPlay",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "todoNum",
                        "name": "todoNum",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "startTime",
                        "name": "startTime",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "endTime",
                        "name": "endTime",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/event-delete": {
            "post": {
                "description": "删除活动",
                "tags": [
                    "活动"
                ],
                "summary": "删除活动",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/event-with-tag": {
            "get": {
                "description": "获取活动与Tag",
                "tags": [
                    "活动"
                ],
                "summary": "活动和活动与Tag",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "请输入当前页，默认第一页",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页大小",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "创建活动与Tag",
                "tags": [
                    "活动"
                ],
                "summary": "创建活动与Tag",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "tagId",
                        "name": "tagId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "EventId",
                        "name": "eventId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/event-with-tag-delete": {
            "post": {
                "description": "删除活动与Tag",
                "tags": [
                    "活动"
                ],
                "summary": "删除活动和活动与Tag",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "tagId",
                        "name": "tagId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "eventId",
                        "name": "eventId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "登录",
                "tags": [
                    "用户"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "password",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "ping",
                "tags": [
                    "测试"
                ],
                "summary": "ping",
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"pong\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "注册",
                "tags": [
                    "用户"
                ],
                "summary": "注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "获取用户列表",
                "tags": [
                    "用户"
                ],
                "summary": "获取用户列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "请输入当前页，默认第一页",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页大小",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "修改用户信息",
                "tags": [
                    "用户"
                ],
                "summary": "修改用户信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "role",
                        "name": "role",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "新增用户信息",
                "tags": [
                    "用户"
                ],
                "summary": "新增用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "role",
                        "name": "role",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user-delete": {
            "post": {
                "description": "删除用户",
                "tags": [
                    "用户"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/version": {
            "get": {
                "description": "获取版本列表",
                "tags": [
                    "版本"
                ],
                "summary": "版本",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "请输入当前页，默认第一页",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页大小",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "更新版本信息",
                "tags": [
                    "版本"
                ],
                "summary": "更新版本信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "num",
                        "name": "num",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "startTime",
                        "name": "startTime",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "endTime",
                        "name": "endTime",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "title",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Img",
                        "name": "img",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "新增版本信息",
                "tags": [
                    "版本"
                ],
                "summary": "新增版本信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "num",
                        "name": "num",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "startTime",
                        "name": "startTime",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "endTime",
                        "name": "endTime",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "title",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Img",
                        "name": "img",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/version-delete": {
            "post": {
                "description": "删除版本",
                "tags": [
                    "版本"
                ],
                "summary": "删除版本",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/version-event": {
            "get": {
                "description": "获取版本活动",
                "tags": [
                    "版本活动"
                ],
                "summary": "根据版本获取版本活动",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "version_num",
                        "name": "version_num",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "修改版本活动",
                "tags": [
                    "版本活动"
                ],
                "summary": "修改版本活动",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "versionNum",
                        "name": "versionNum",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "title",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "content",
                        "name": "content",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "level",
                        "name": "level",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "startTime",
                        "name": "startTime",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "endTime",
                        "name": "endTime",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "todoNum",
                        "name": "todoNum",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "新增活动",
                "tags": [
                    "版本活动"
                ],
                "summary": "新增活动",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "versionNum",
                        "name": "versionNum",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "title",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "content",
                        "name": "content",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "level",
                        "name": "level",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "startTime",
                        "name": "startTime",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "endTime",
                        "name": "endTime",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "todoNum",
                        "name": "todoNum",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/version-event-delete": {
            "post": {
                "description": "删除版本活动",
                "tags": [
                    "版本活动"
                ],
                "summary": "删除版本活动",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/version-event-with-tag": {
            "get": {
                "description": "获取版本活动和tag",
                "tags": [
                    "版本活动"
                ],
                "summary": "根据版本获取版本活动和tag",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "version_num",
                        "name": "version_num",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "创建版本活动与Tag",
                "tags": [
                    "版本活动"
                ],
                "summary": "创建版本活动与Tag",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "tagId",
                        "name": "tagId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "versionEventId",
                        "name": "versionEventId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/version-event-with-tag-delete": {
            "post": {
                "description": "删除版本活动与Tag",
                "tags": [
                    "版本活动"
                ],
                "summary": "删除版本活动与Tag",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "tagId",
                        "name": "tagId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "versionEventId",
                        "name": "versionEventId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/version-with-tag": {
            "get": {
                "description": "获取版本与Tag",
                "tags": [
                    "版本"
                ],
                "summary": "版本和版本与Tag",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "请输入当前页，默认第一页",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页大小",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "创建版本与Tag",
                "tags": [
                    "版本"
                ],
                "summary": "创建版本与Tag",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "tagId",
                        "name": "tagId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "versionId",
                        "name": "versionId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/version-with-tag-delete": {
            "post": {
                "description": "删除版本与Tag",
                "tags": [
                    "版本"
                ],
                "summary": "删除版本和版本与Tag",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "tagId",
                        "name": "tagId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "versionId",
                        "name": "versionId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"msg\":\"\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
