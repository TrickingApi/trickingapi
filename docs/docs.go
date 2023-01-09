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
        "termsOfService": "https://github.com/TrickingApi/trickingapi",
        "contact": {
            "name": "Mikael Mantis",
            "url": "https://github.com/TrickingApi/trickingapi",
            "email": "mikael@trickingapi.dev"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/TrickingApi/trickingapi/blob/main/LICENSE.md"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/categories": {
            "get": {
                "description": "Reads and returns all existing categories of tricks from the tricks.json file at https://github.com/TrickingApi/trickingapi",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/categories/:name": {
            "get": {
                "description": "Reads and returns a list of tricks for a specific category from the tricks.json file at https://github.com/TrickingApi/trickingapi",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories",
                    "tricks"
                ],
                "summary": "Get All Tricks Grouped Under A Category from TrickingApi/data/tricks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Trick"
                            }
                        }
                    }
                }
            }
        },
        "/categories/tricks": {
            "get": {
                "description": "Reads and returns a mapping of categories to list of tricks from the tricks.json file at https://github.com/TrickingApi/trickingapi",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories",
                    "tricks"
                ],
                "summary": "Get All Tricks Grouped by Categories from TrickingApi/data/tricks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/models.Trick"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/tricks": {
            "get": {
                "description": "Reads and returns list of tricks from the static tricks.json file at https://github.com/TrickingApi/trickingapi",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tricks"
                ],
                "summary": "Get All Tricks from TrickingApi/data/tricks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Trick"
                            }
                        }
                    }
                }
            }
        },
        "/tricks/names": {
            "get": {
                "description": "Reads and returns the names of all tricks from the static tricks.json file at https://github.com/TrickingApi/trickingapi",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tricks"
                ],
                "summary": "Get All Trick Names from TrickingApi/data/tricks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/tricks:name": {
            "get": {
                "description": "reads list of known Trick objects and returns trick matching the name param in the request",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tricks"
                ],
                "summary": "Get Trick by Specific Name from TrickingApi/data/tricks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Trick"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.TrickError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Trick": {
            "type": "object",
            "properties": {
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.TrickCategory"
                    }
                },
                "description": {
                    "type": "string"
                },
                "difficultyRank": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nextTricks": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "prereqs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models.TrickCategory": {
            "type": "string",
            "enum": [
                "Flip",
                "Vert Kick",
                "Twist",
                "Pseudo Double Flip",
                "Single",
                "Double",
                "Triple",
                "Quad"
            ],
            "x-enum-varnames": [
                "FLIP",
                "VERT_KICK",
                "TWIST",
                "PSEUDO_DUB",
                "SING",
                "DUB",
                "TRIP",
                "QUAD"
            ]
        },
        "models.TrickError": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "trickingapi.dev",
	BasePath:         "/api",
	Schemes:          []string{"https"},
	Title:            "Tricking Api",
	Description:      "Consumption-only API for barebones Tricking Terminology Data",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
