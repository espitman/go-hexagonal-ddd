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
        "contact": {
            "name": "API Support",
            "email": "s.heidar@jabama.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/item": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Add a new item to list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "Add a new item to list",
                "parameters": [
                    {
                        "description": "Item object to Add",
                        "name": "list",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/appModel.NewItem"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/appModel.Item"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/commonModels.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/commonModels.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/item/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "remove Item from the list by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "remove Item from the list by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/commonModels.OkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/commonModels.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/commonModels.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/list": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get all lists",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "List"
                ],
                "summary": "Get all lists",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/appModel.List"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/commonModels.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/commonModels.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a new list with the specified name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "List"
                ],
                "summary": "Create a new list",
                "parameters": [
                    {
                        "description": "List object to create",
                        "name": "list",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/appModel.NewList"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/appModel.List"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/commonModels.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/commonModels.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/commonModels.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/list/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get a list using the provided ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "List"
                ],
                "summary": "Get list by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/appModel.ListWithItems"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/commonModels.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/commonModels.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "update a list using the provided ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "List"
                ],
                "summary": "update list by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "List object to update",
                        "name": "list",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/appModel.NewList"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/appModel.List"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/commonModels.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/commonModels.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/commonModels.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/commonModels.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "delete a list using the provided ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "List"
                ],
                "summary": "delete list by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/commonModels.OkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/commonModels.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/commonModels.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "appModel.Item": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "itemCode": {
                    "type": "integer"
                },
                "listId": {
                    "type": "string"
                }
            }
        },
        "appModel.List": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "appModel.ListWithItems": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/appModel.Item"
                    }
                },
                "name": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "appModel.NewItem": {
            "type": "object",
            "properties": {
                "itemCode": {
                    "type": "integer"
                },
                "listId": {
                    "type": "string"
                }
            }
        },
        "appModel.NewList": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "commonModels.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "commonModels.OkResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Jabama Wishes",
	Description:      "This is a sample hexagonal domain driven for golang.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
