// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-02-15 00:11:08.4685964 +0100 CET m=+0.128500101

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/{platform}/darvo/": {
            "get": {
                "description": "get platform  Darvo Deal by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DarvoDeals"
                ],
                "summary": "Show active Darvo Deal",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Platform",
                        "name": "platform",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/parser.DarvoDeals"
                        }
                    }
                }
            }
        },
        "/{platform}/news/": {
            "get": {
                "description": "get platform  News",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Newsdata"
                ],
                "summary": "Show curent News",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Platform",
                        "name": "platform",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "language",
                        "name": "Accept-Language",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/parser.News"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "parser.DarvoDeals": {
            "type": "object",
            "properties": {
                "dealPrice": {
                    "type": "integer"
                },
                "discountPercent": {
                    "type": "integer"
                },
                "ends": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "item": {
                    "type": "string"
                },
                "itemtest": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "sold": {
                    "type": "integer"
                },
                "start": {
                    "type": "string"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "parser.News": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "priority": {
                    "type": "boolean"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:9090",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Mybitti Warframe API",
	Description: "This is the  simple REST Version of Mybitti's Warframe API.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}