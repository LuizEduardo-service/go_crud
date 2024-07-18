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
        "/createUser": {
            "post": {
                "description": "Create a new user with the provided user information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User information for registration",
                        "name": "userRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest_err.RestErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest_err.RestErr"
                        }
                    }
                }
            }
        },
        "/deleteUser/{userId}": {
            "delete": {
                "description": "Deletes a user based on the ID provided as a parameter.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the user to be deleted",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest_err.RestErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest_err.RestErr"
                        }
                    }
                }
            }
        },
        "/getUserById/{userId}": {
            "get": {
                "description": "Retrieves user details based on the user ID provided as a parameter.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Find User by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the user to be retrieved",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User information retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/response.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Error: Invalid user ID",
                        "schema": {
                            "$ref": "#/definitions/rest_err.RestErr"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/rest_err.RestErr"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Allows a user to log in and receive an authentication token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "description": "User login credentials",
                        "name": "userLogin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successful, authentication token provided",
                        "schema": {
                            "$ref": "#/definitions/response.UserResponse"
                        },
                        "headers": {
                            "Authorization": {
                                "type": "string",
                                "description": "Authentication token"
                            }
                        }
                    },
                    "403": {
                        "description": "Error: Invalid login credentials",
                        "schema": {
                            "$ref": "#/definitions/rest_err.RestErr"
                        }
                    }
                }
            }
        },
        "/updateUser/{userId}": {
            "put": {
                "description": "Updates user details based on the ID provided as a parameter.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the user to be updated",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User information for update",
                        "name": "userRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserUpdateRequest"
                        }
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest_err.RestErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest_err.RestErr"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.UserLogin": {
            "description": "Structure containing the necessary fields for user login.",
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "description": "User's email (required and must be a valid email address).",
                    "type": "string",
                    "example": "test@test.com"
                },
                "password": {
                    "description": "User's password (required, minimum of 6 characters, and must contain at least one of the characters: !@#$%*).",
                    "type": "string",
                    "minLength": 6,
                    "example": "password#@#@!2121"
                }
            }
        },
        "request.UserRequest": {
            "description": "Structure containing the required fields for creating a new user.",
            "type": "object",
            "required": [
                "age",
                "email",
                "name",
                "password"
            ],
            "properties": {
                "age": {
                    "description": "User's age (required, must be between 1 and 140).\n@json\n@jsonTag age\n@jsonExample 30",
                    "type": "integer",
                    "maximum": 140,
                    "minimum": 1,
                    "example": 30
                },
                "email": {
                    "description": "User's email (required and must be a valid email address).\nExample: user@example.com\n@json\n@jsonTag email\n@jsonExample user@example.com\n@binding required,email",
                    "type": "string",
                    "example": "test@test.com"
                },
                "name": {
                    "description": "User's name (required, minimum of 4 characters, maximum of 100 characters).\nExample: John Doe\n@json\n@jsonTag name\n@jsonExample John Doe\n@binding required,min=4,max=100",
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 4,
                    "example": "John Doe"
                },
                "password": {
                    "description": "User's password (required, minimum of 6 characters, and must contain at least one of the characters: !@#$%*).\n@json\n@jsonTag password\n@jsonExample P@ssw0rd!\n@binding required,min=6,containsany=!@#$%*",
                    "type": "string",
                    "minLength": 6,
                    "example": "password#@#@!2121"
                }
            }
        },
        "request.UserUpdateRequest": {
            "type": "object",
            "required": [
                "age",
                "name"
            ],
            "properties": {
                "age": {
                    "description": "User's age (required, must be between 1 and 140).\n@json\n@jsonTag age\n@jsonExample 30\n@binding required,min=1,max=140",
                    "type": "integer",
                    "maximum": 140,
                    "minimum": 1,
                    "example": 30
                },
                "name": {
                    "description": "User's name (required, minimum of 4 characters, maximum of 100 characters).\nExample: John Doe\n@json\n@jsonTag name\n@jsonExample John Doe\n@binding required,min=4,max=100",
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 4,
                    "example": "John Doe"
                }
            }
        },
        "response.UserResponse": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 30
                },
                "email": {
                    "type": "string",
                    "example": "test@test.com"
                },
                "id": {
                    "type": "string",
                    "example": "82bdd399-321b-41d8-8b40-9a0116db9e92"
                },
                "name": {
                    "type": "string",
                    "example": "John Doe"
                }
            }
        },
        "rest_err.Causes": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "rest_err.RestErr": {
            "type": "object",
            "properties": {
                "causes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/rest_err.Causes"
                    }
                },
                "code": {
                    "type": "integer"
                },
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
