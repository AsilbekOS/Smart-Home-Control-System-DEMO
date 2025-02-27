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
        "/control": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Bu endpoint qurulmani boshqarish uchun ishlatiladi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Create control",
                "parameters": [
                    {
                        "description": "ControlRequest",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/proto.ControlRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Commad sccessfully created",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {
                            "$ref": "#/definitions/models.StandartError"
                        }
                    },
                    "403": {
                        "description": "Forbidden error",
                        "schema": {
                            "$ref": "#/definitions/models.ForbiddenError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.StandartError"
                        }
                    }
                }
            }
        },
        "/devices": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Bu endpoint yangi qurulma qo'shish uchun ishlatiladi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Add Devices",
                "parameters": [
                    {
                        "description": "AddDeviceRequestSwag",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/proto.AddDeviceRequestSwag"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Device successfully added",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {
                            "$ref": "#/definitions/models.StandartError"
                        }
                    },
                    "403": {
                        "description": "Forbidden error",
                        "schema": {
                            "$ref": "#/definitions/models.ForbiddenError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.StandartError"
                        }
                    }
                }
            }
        },
        "/devices/delete": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Bu endpoint qurulmani o'chirish uchun ishlatiladi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Add Devices",
                "parameters": [
                    {
                        "description": "DeleteDeviceRequest",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/proto.DeleteDeviceRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Device successfully deleted",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {
                            "$ref": "#/definitions/models.StandartError"
                        }
                    },
                    "403": {
                        "description": "Forbidden error",
                        "schema": {
                            "$ref": "#/definitions/models.ForbiddenError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.StandartError"
                        }
                    }
                }
            }
        },
        "/devices/update": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Bu endpoint qurulmani yangilash uchun ishlatiladi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Add Devices",
                "parameters": [
                    {
                        "description": "UpdateDeviceRequestSwag",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/proto.UpdateDeviceRequestSwag"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Device successfully updated",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {
                            "$ref": "#/definitions/models.StandartError"
                        }
                    },
                    "403": {
                        "description": "Forbidden error",
                        "schema": {
                            "$ref": "#/definitions/models.ForbiddenError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.StandartError"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Bu endpoint foydalanuvchi profilini olish uchun ishlatiladi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "LOGIN User",
                "parameters": [
                    {
                        "description": "LoginRequest",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/proto.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successful",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {
                            "$ref": "#/definitions/models.StandartError"
                        }
                    },
                    "403": {
                        "description": "Forbidden error",
                        "schema": {
                            "$ref": "#/definitions/models.ForbiddenError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.StandartError"
                        }
                    }
                }
            }
        },
        "/user/profile": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Bu endpoint foydalanuvchi profilini olish uchun ishlatiladi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "PROFILE User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token: ",
                        "name": "token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User ID: ",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {
                            "$ref": "#/definitions/models.StandartError"
                        }
                    },
                    "403": {
                        "description": "Forbidden error",
                        "schema": {
                            "$ref": "#/definitions/models.ForbiddenError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.StandartError"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Bu endpoint yangi foydalanuvchini yaratish uchun ishlatiladi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "CREATE User",
                "parameters": [
                    {
                        "description": "UserSwag",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/proto.UserSwag"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User successfully created",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request error",
                        "schema": {
                            "$ref": "#/definitions/models.StandartError"
                        }
                    },
                    "403": {
                        "description": "Forbidden error",
                        "schema": {
                            "$ref": "#/definitions/models.ForbiddenError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.StandartError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ForbiddenError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.StandartError": {
            "type": "object",
            "properties": {
                "error": {}
            }
        },
        "proto.AddDeviceRequestSwag": {
            "type": "object",
            "properties": {
                "device": {
                    "$ref": "#/definitions/proto.SwagDevice"
                }
            }
        },
        "proto.BrightnessCommandPayload": {
            "type": "object",
            "properties": {
                "brightness": {
                    "type": "integer"
                },
                "color": {
                    "type": "string"
                }
            }
        },
        "proto.Configuration": {
            "type": "object",
            "properties": {
                "brightness": {
                    "type": "integer"
                },
                "color": {
                    "type": "string"
                }
            }
        },
        "proto.ControlCommand": {
            "type": "object",
            "properties": {
                "command_id": {
                    "type": "string"
                },
                "command_payload": {
                    "$ref": "#/definitions/proto.BrightnessCommandPayload"
                },
                "command_type": {
                    "type": "string"
                },
                "device_id": {
                    "type": "string"
                },
                "device_status": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "proto.ControlRequest": {
            "type": "object",
            "properties": {
                "command": {
                    "$ref": "#/definitions/proto.ControlCommand"
                }
            }
        },
        "proto.DeleteDeviceRequest": {
            "type": "object",
            "properties": {
                "device_id": {
                    "type": "string"
                }
            }
        },
        "proto.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "proto.Profile": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "proto.SwagDevice": {
            "type": "object",
            "properties": {
                "configuration": {
                    "$ref": "#/definitions/proto.Configuration"
                },
                "connectivity_status": {
                    "type": "string"
                },
                "device_name": {
                    "type": "string"
                },
                "device_status": {
                    "type": "string"
                },
                "device_type": {
                    "type": "string"
                },
                "firmware_version": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "proto.UpdateDeviceRequestSwag": {
            "type": "object",
            "properties": {
                "device": {
                    "$ref": "#/definitions/proto.UpdateSwagDevice"
                }
            }
        },
        "proto.UpdateSwagDevice": {
            "type": "object",
            "properties": {
                "configuration": {
                    "$ref": "#/definitions/proto.Configuration"
                },
                "connectivity_status": {
                    "type": "string"
                },
                "device_id": {
                    "type": "string"
                },
                "device_name": {
                    "type": "string"
                },
                "device_status": {
                    "type": "string"
                },
                "device_type": {
                    "type": "string"
                },
                "firmware_version": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                }
            }
        },
        "proto.UserSwag": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password_hash": {
                    "type": "string"
                },
                "profile": {
                    "$ref": "#/definitions/proto.Profile"
                },
                "username": {
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
	Host:             "localhost:9000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Project: SMART HOME CONTROL SYSTEM",
	Description:      "This swagger UI was created for Exam",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
