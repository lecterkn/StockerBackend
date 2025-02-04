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
        "/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "ユーザーログイン",
                "parameters": [
                    {
                        "description": "ユーザーログインリクエスト",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/controller.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.UserLoginResponse"
                        }
                    }
                }
            }
        },
        "/products/{janCode}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jancode"
                ],
                "summary": "製品情報取得",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JANコード",
                        "name": "janCode",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.JancodeResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "ユーザー作成",
                "parameters": [
                    {
                        "description": "ユーザー作成リクエスト",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/controller.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.UserResponse"
                        }
                    }
                }
            }
        },
        "/stores": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "store"
                ],
                "summary": "店舗一覧取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controller.StoreListResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "list": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/controller.StoreResponse"
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
                    "store"
                ],
                "summary": "店舗新規作成",
                "parameters": [
                    {
                        "description": "店舗新規作成リクエスト",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/controller.StoreCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.StoreResponse"
                        }
                    }
                }
            }
        },
        "/stores/{storeId}": {
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "store"
                ],
                "summary": "店舗更新",
                "parameters": [
                    {
                        "type": "string",
                        "description": "店舗ID",
                        "name": "storeId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "店舗更新リクエスト",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/controller.StoreUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.StoreResponse"
                        }
                    }
                }
            }
        },
        "/stores/{storeId}/itemStocks": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item_stock"
                ],
                "summary": "商品詳細一覧取得",
                "parameters": [
                    {
                        "type": "string",
                        "description": "店舗ID",
                        "name": "storeId",
                        "in": "path",
                        "required": true
                    }
                ],
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
                        "type": "string",
                        "description": "店舗ID",
                        "name": "storeId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "商品詳細作成リクエスト",
                        "name": "request",
                        "in": "body",
                        "required": true,
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
        "/stores/{storeId}/itemStocks/jancodes/{jancode}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item_stock"
                ],
                "summary": "Jancodeから商品詳細取得",
                "parameters": [
                    {
                        "type": "string",
                        "description": "店舗ID",
                        "name": "storeId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "JANコード",
                        "name": "jancode",
                        "in": "path",
                        "required": true
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
        "/stores/{storeId}/itemStocks/{itemId}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item_stock"
                ],
                "summary": "商品詳細取得",
                "parameters": [
                    {
                        "type": "string",
                        "description": "店舗ID",
                        "name": "storeId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "商品ID",
                        "name": "itemId",
                        "in": "path",
                        "required": true
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
                        "type": "string",
                        "description": "店舗ID",
                        "name": "storeId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "商品ID",
                        "name": "itemId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "商品詳細更新リクエスト",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ItemStockUpdateRequest"
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
        "/stores/{storeId}/items": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "商品一覧取得",
                "parameters": [
                    {
                        "type": "string",
                        "description": "店舗ID",
                        "name": "storeId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "商品名",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Janコード",
                        "name": "janCode",
                        "in": "query"
                    }
                ],
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
                        "description": "商品作成リクエスト",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ItemRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "店舗ID",
                        "name": "storeId",
                        "in": "path",
                        "required": true
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
        "/stores/{storeId}/items/{itemId}": {
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
                        "type": "string",
                        "description": "店舗ID",
                        "name": "storeId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "商品ID",
                        "name": "itemId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "商品更新リクエスト",
                        "name": "request",
                        "in": "body",
                        "required": true,
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
        "/stores/{storeId}/stockIns": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stock_ins"
                ],
                "summary": "入荷履歴一覧取得",
                "parameters": [
                    {
                        "type": "string",
                        "description": "店舗ID",
                        "name": "storeId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.StockInListResponse"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stock_ins"
                ],
                "summary": "入荷履歴作成取得",
                "parameters": [
                    {
                        "type": "string",
                        "description": "店舗ID",
                        "name": "storeId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "入荷履歴作成リクエスト",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.StockInCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.StockInResponse"
                        }
                    }
                }
            }
        },
        "/stores/{storeId}/stockOuts": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stock_outs"
                ],
                "summary": "販売履歴一覧取得取得",
                "parameters": [
                    {
                        "type": "string",
                        "description": "店舗ID",
                        "name": "storeId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controller.StockOutListResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "list": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/controller.StockOutResponse"
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
                    "stock_outs"
                ],
                "summary": "販売履歴作成取得",
                "parameters": [
                    {
                        "type": "string",
                        "description": "店舗ID",
                        "name": "storeId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "販売履歴作成リクエスト",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.StockOutCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.StockOutResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ItemListResponse": {
            "type": "object",
            "required": [
                "list"
            ],
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
            "required": [
                "janCode",
                "name"
            ],
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
            "required": [
                "createdAt",
                "id",
                "janCode",
                "name",
                "storeId",
                "updatedAt"
            ],
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
                "storeId": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "controller.ItemStockListResponse": {
            "type": "object",
            "required": [
                "list"
            ],
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
            "required": [
                "janCode",
                "name",
                "stock"
            ],
            "properties": {
                "janCode": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                },
                "stockMin": {
                    "type": "integer"
                }
            }
        },
        "controller.ItemStockResponse": {
            "type": "object",
            "required": [
                "createdAt",
                "itemId",
                "janCode",
                "name",
                "price",
                "stock",
                "stockMin",
                "storeId",
                "updatedAt"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "itemId": {
                    "type": "string"
                },
                "janCode": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                },
                "stockMin": {
                    "type": "integer"
                },
                "storeId": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "controller.ItemStockUpdateRequest": {
            "type": "object",
            "required": [
                "stock"
            ],
            "properties": {
                "price": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                },
                "stockMin": {
                    "type": "integer"
                }
            }
        },
        "controller.JancodeResponse": {
            "type": "object",
            "required": [
                "brandName",
                "makerName",
                "name"
            ],
            "properties": {
                "brandName": {
                    "type": "string"
                },
                "makerName": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "controller.StockInCreateRequest": {
            "type": "object",
            "required": [
                "itemId",
                "place",
                "price",
                "stocks"
            ],
            "properties": {
                "itemId": {
                    "type": "string"
                },
                "place": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "stocks": {
                    "type": "integer"
                }
            }
        },
        "controller.StockInListResponse": {
            "type": "object",
            "required": [
                "list"
            ],
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.StockInResponse"
                    }
                }
            }
        },
        "controller.StockInResponse": {
            "type": "object",
            "required": [
                "createdAt",
                "id",
                "itemId",
                "name",
                "place",
                "price",
                "stocks",
                "storeId",
                "updatedAt"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "itemId": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "place": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "stocks": {
                    "type": "integer"
                },
                "storeId": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "controller.StockOutCreateRequest": {
            "type": "object",
            "required": [
                "itemId",
                "price",
                "stocks"
            ],
            "properties": {
                "itemId": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "stocks": {
                    "type": "integer"
                }
            }
        },
        "controller.StockOutListResponse": {
            "type": "object",
            "required": [
                "list"
            ],
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.StockOutResponse"
                    }
                }
            }
        },
        "controller.StockOutResponse": {
            "type": "object",
            "required": [
                "createdAt",
                "id",
                "itemId",
                "name",
                "price",
                "stocks",
                "storeId",
                "updatedAt"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "itemId": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "stocks": {
                    "type": "integer"
                },
                "storeId": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "controller.StoreCreateRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "controller.StoreListResponse": {
            "type": "object",
            "required": [
                "list"
            ],
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.StoreResponse"
                    }
                }
            }
        },
        "controller.StoreResponse": {
            "type": "object",
            "required": [
                "createdAt",
                "id",
                "name",
                "updatedAt",
                "userId"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "controller.StoreUpdateRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "controller.UserLoginResponse": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "controller.UserRequest": {
            "type": "object",
            "required": [
                "name",
                "password"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "controller.UserResponse": {
            "type": "object",
            "required": [
                "createdAt",
                "id",
                "name",
                "updatedAt"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
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
