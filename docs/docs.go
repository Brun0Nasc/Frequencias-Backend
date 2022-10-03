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
            "name": "Equipe Komanda",
            "url": "http://www.swagger.io/support",
            "email": "caiosousafernandesferreira@hotmail.com"
        },
        "license": {
            "name": "Mozilla Public License 2.0",
            "url": "https://www.mozilla.org/en-US/MPL/2.0/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/frequencia": {
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "POST um novo usuario. O unico parametro que deve ser passado é o id do usuario em questão",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Frequencias"
                ],
                "summary": "POST uma nova frequencia",
                "parameters": [
                    {
                        "description": "IDUsuario",
                        "name": "IDUsuario",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/frequencia/lista": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "POST um novo usuario. O unico parametro que deve ser passado é o id do usuario em questão",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Frequencias"
                ],
                "summary": "POST uma nova frequencia",
                "parameters": [
                    {
                        "type": "string",
                        "description": "data",
                        "name": "data",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "nome",
                            "entrada",
                            "saida"
                        ],
                        "type": "string",
                        "description": "OrderBy",
                        "name": "orderBy",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "description": "order",
                        "name": "order",
                        "in": "query"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/frequencia_usuario.ListaUsuarioFrequencia"
                            }
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Autentica um usuário com as credenciais enviadas e se OK gera um token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuarios"
                ],
                "summary": "Login para usuário",
                "operationId": "Authentication",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "Login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/usuarios": {
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "POST um novo usuario. Os parametros que devem ser passados são, \"nome\", \"email\", \"senha\", são necessários.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuarios"
                ],
                "summary": "POST um novo usuario",
                "parameters": [
                    {
                        "description": "NovoUsuario",
                        "name": "NovoUsuario",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/usuarios/inativar/{id}": {
            "delete": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Inativa o usuario a partir do id dele",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuarios"
                ],
                "summary": "Inativar Usuario",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
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
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/usuarios/list_user": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "GET todos os usuarios com orderBy \u0026 || order (desc, cresc)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuarios"
                ],
                "summary": "GET todos os usuarios",
                "parameters": [
                    {
                        "enum": [
                            "id",
                            "nome",
                            "created_at"
                        ],
                        "type": "string",
                        "description": "column",
                        "name": "orderBy",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "order desc",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "true",
                            "false"
                        ],
                        "type": "string",
                        "description": "removido",
                        "name": "removido",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/usuarios.ListaUsuarios"
                            }
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "frequencia_usuario.ListaUsuarioFrequencia": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/frequencia_usuario.ResFrequenciaUsuario"
                    }
                }
            }
        },
        "frequencia_usuario.ResFrequenciaUsuario": {
            "type": "object",
            "properties": {
                "data_frequencia": {
                    "type": "string"
                },
                "entrada": {
                    "type": "string"
                },
                "frequencia_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                },
                "saida": {
                    "type": "string"
                },
                "usuario_id": {
                    "type": "integer"
                }
            }
        },
        "usuarios.ListaUsuarios": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/usuarios.Usuario"
                    }
                }
            }
        },
        "usuarios.Usuario": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                },
                "removed_at": {
                    "type": "string"
                },
                "senha": {
                    "type": "string"
                },
                "tipo": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
