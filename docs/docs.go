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
        "/example/helloworld": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "ping example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/hotels": {
            "post": {
                "description": "Add a hotel to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add a hotel",
                "parameters": [
                    {
                        "description": "Hotel",
                        "name": "hotel",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Hotel"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Hotel"
                        }
                    }
                }
            }
        },
        "/hotels/{name}": {
            "get": {
                "description": "Get a hotel from the database by its name",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a hotel by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hotel Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Hotel"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a hotel in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a hotel",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Hotel Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Hotel",
                        "name": "hotel",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Hotel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Hotel"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Hotel": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/api/v1",
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
