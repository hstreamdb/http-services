// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/cluster/stats": {
            "get": {
                "summary": "Get cluster stats",
                "operationId": "statsGet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category",
                        "name": "category",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Metrics",
                        "name": "metrics",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "description": "Interval collection",
                        "name": "interval",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TableResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorno.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/cluster/status": {
            "get": {
                "summary": "Get server status of the cluster",
                "operationId": "statusGet",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TableResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorno.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/streams/": {
            "get": {
                "summary": "List all streams in the cluster",
                "operationId": "streamList",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Stream"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorno.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorno.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "summary": "Create a stream",
                "operationId": "streamCreate",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Stream"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorno.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorno.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/streams/{streamName}": {
            "get": {
                "summary": "Get specific stream by streamName",
                "operationId": "streamGet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Stream name",
                        "name": "streamName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Stream"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errorno.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorno.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "summary": "Append record to specific stream",
                "operationId": "streamAppend",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Record"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Stream name",
                        "name": "streamName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.RecordId"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorno.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorno.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "summary": "Delete specific stream by streamName",
                "operationId": "streamDelete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Stream Name",
                        "name": "streamName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorno.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/subscriptions/": {
            "get": {
                "summary": "List all subscriptions in the cluster",
                "operationId": "subscriptionList",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Subscription"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorno.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "summary": "Create a subscription",
                "operationId": "subscriptionCreate",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Subscription"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errorno.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorno.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/subscriptions/{subId}": {
            "get": {
                "summary": "Get specific subscription by subscription id",
                "operationId": "subscriptionGet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Subscription Id",
                        "name": "subId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Subscription"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errorno.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorno.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "summary": "Delete specific subscription by subscription id",
                "operationId": "subscriptionDelete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Subscription Id",
                        "name": "subId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errorno.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errorno.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "full_text": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.Record": {
            "type": "object",
            "required": [
                "data",
                "key",
                "type"
            ],
            "properties": {
                "data": {},
                "key": {
                    "type": "string"
                },
                "type": {
                    "description": "Record Type:\n* RAW - []byte payload\n* HRECORD - JSON payload",
                    "type": "string",
                    "enum": [
                        "RAW",
                        "HRECORD"
                    ]
                }
            }
        },
        "model.RecordId": {
            "type": "object",
            "properties": {
                "batch_id": {
                    "type": "integer"
                },
                "batch_index": {
                    "type": "integer"
                },
                "shard_id": {
                    "type": "integer"
                }
            }
        },
        "model.Stream": {
            "type": "object",
            "required": [
                "stream_name"
            ],
            "properties": {
                "backlog_duration": {
                    "type": "integer"
                },
                "replication_factor": {
                    "type": "integer"
                },
                "stream_name": {
                    "type": "string"
                }
            }
        },
        "model.Subscription": {
            "type": "object",
            "required": [
                "stream_name",
                "subscription_id"
            ],
            "properties": {
                "ack_timeout_seconds": {
                    "type": "integer"
                },
                "max_unacked_records": {
                    "type": "integer"
                },
                "stream_name": {
                    "type": "string"
                },
                "subscription_id": {
                    "type": "string"
                }
            }
        },
        "model.TableResult": {
            "type": "object",
            "properties": {
                "headers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "value": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "additionalProperties": {
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
	Version:          "0.1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "HStreamDB-Server API",
	Description:      "http server for HStreamDB",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
