// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/items": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "商品一覧取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controller.ItemListResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "list": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/controller.ItemResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "商品新規作成",
                "parameters": [
                    {
                        "description": "アイテム作成リクエスト",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/controller.ItemRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ItemResponse"
                        }
                    }
                }
            },
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "商品更新",
                "parameters": [
                    {
                        "description": "アイテム作成リクエスト",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/controller.ItemRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ItemResponse"
                        }
                    }
                }
            }
        },
        "/items/{item_id}/stocks": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item_stock"
                ],
                "summary": "商品詳細取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ItemStockResponse"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item_stock"
                ],
                "summary": "商品詳細登録",
                "parameters": [
                    {
                        "description": "商品詳細作成リクエスト",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/controller.ItemStockRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ItemStockResponse"
                        }
                    }
                }
            },
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item_stock"
                ],
                "summary": "商品詳細更新",
                "parameters": [
                    {
                        "description": "商品詳細更新リクエスト",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/controller.ItemStockRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ItemStockResponse"
                        }
                    }
                }
            }
        },
        "/itemsStocks": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item_stock"
                ],
                "summary": "商品詳細一覧取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controller.ItemStockListResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "list": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/controller.ItemStockResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ItemListResponse": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.ItemResponse"
                    }
                }
            }
        },
        "controller.ItemRequest": {
            "type": "object",
            "properties": {
                "janCode": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "controller.ItemResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "janCode": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "controller.ItemStockListResponse": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.ItemStockResponse"
                    }
                }
            }
        },
        "controller.ItemStockRequest": {
            "type": "object",
            "properties": {
                "place": {
                    "type": "string"
                },
                "stock": {
                    "type": "integer"
                },
                "stock_min": {
                    "type": "integer"
                }
            }
        },
        "controller.ItemStockResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "item_id": {
                    "type": "string"
                },
                "place": {
                    "type": "string"
                },
                "stock": {
                    "type": "integer"
                },
                "stock_min": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8810",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "StockerAPI",
	Description:      "stocker application server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}