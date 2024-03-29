// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (

	"github.com/swaggo/swag"
)

func init() {

    // doc, _ := ioutil.ReadFile("swagger.json")

    const doc = `
    {
        "swagger": "2.0",
        "info": {
            "contact": {},
            "title": "Image Processing",
            "version": "1.0.0"
        },
        "basePath": "%7B%7BHost%7D%7D",
        "paths": {
            "/api/v1/compress": {
                "post": {
                    "consumes": [
                        "multipart/form-data"
                    ],
                    "parameters": [
                        {
                            "format": "binary",
                            "in": "formData",
                            "name": "image_file",
                            "type": "string"
                        }
                    ],
                    "responses": {
                        "200": {
                            "description": ""
                        }
                    },
                    "description": "compress",
                    "operationId": "compress",
                    "summary": "compress"
                }
            },
            "/api/v1/convert": {
                "post": {
                    "consumes": [
                        "multipart/form-data"
                    ],
                    "parameters": [
                        {
                            "format": "binary",
                            "in": "formData",
                            "name": "image_file",
                            "type": "string"
                        }
                    ],
                    "responses": {
                        "200": {
                            "description": ""
                        }
                    },
                    "description": "convert",
                    "operationId": "convert",
                    "summary": "convert"
                }
            },
            "/api/v1/resize": {
                "post": {
                    "consumes": [
                        "multipart/form-data"
                    ],
                    "parameters": [
                        {
                            "in": "formData",
                            "name": "height",
                            "type": "string"
                        },
                        {
                            "format": "binary",
                            "in": "formData",
                            "name": "image_file",
                            "type": "string"
                        },
                        {
                            "in": "formData",
                            "name": "width",
                            "type": "string"
                        }
                    ],
                    "responses": {
                        "200": {
                            "description": ""
                        }
                    },
                    "description": "resize",
                    "operationId": "resize",
                    "summary": "resize"
                }
            }
        },
        "tags": []
    }
    `

    // SwaggerInfo holds exported Swagger Info so clients can modify it
    var SwaggerInfo = &swag.Spec{
        Version:          "",
        Host:             "",
        BasePath:         "",
        Schemes:          []string{},
        Title:            "",
        Description:      "",
        InfoInstanceName: "swagger",
        SwaggerTemplate:  doc,
    }

	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
