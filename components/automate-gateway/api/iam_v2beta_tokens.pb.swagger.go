package api

func init() {
	Swagger.Add("iam_v2beta_tokens", `{
  "swagger": "2.0",
  "info": {
    "title": "components/automate-gateway/api/iam/v2beta/tokens.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/iam/v2beta/tokens": {
      "get": {
        "operationId": "ListTokens",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2beta.ListTokensResp"
            }
          }
        },
        "tags": [
          "Tokens"
        ]
      },
      "post": {
        "operationId": "CreateToken",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2beta.CreateTokenResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2beta.CreateTokenReq"
            }
          }
        ],
        "tags": [
          "Tokens"
        ]
      }
    },
    "/iam/v2beta/tokens/{id}": {
      "get": {
        "operationId": "GetToken",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2beta.GetTokenResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Tokens"
        ]
      },
      "delete": {
        "operationId": "DeleteToken",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2beta.DeleteTokenResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Tokens"
        ]
      },
      "put": {
        "operationId": "UpdateToken",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2beta.UpdateTokenResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "ID can't be changed; ID used to discover token",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.iam.v2beta.UpdateTokenReq"
            }
          }
        ],
        "tags": [
          "Tokens"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.iam.v2beta.CreateTokenReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "active": {
          "type": "boolean",
          "format": "boolean"
        },
        "value": {
          "type": "string"
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "chef.automate.api.iam.v2beta.CreateTokenResp": {
      "type": "object",
      "properties": {
        "token": {
          "$ref": "#/definitions/chef.automate.api.iam.v2beta.Token"
        }
      }
    },
    "chef.automate.api.iam.v2beta.DeleteTokenResp": {
      "type": "object"
    },
    "chef.automate.api.iam.v2beta.GetTokenResp": {
      "type": "object",
      "properties": {
        "token": {
          "$ref": "#/definitions/chef.automate.api.iam.v2beta.Token"
        }
      }
    },
    "chef.automate.api.iam.v2beta.ListTokensResp": {
      "type": "object",
      "properties": {
        "tokens": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.iam.v2beta.Token"
          }
        }
      }
    },
    "chef.automate.api.iam.v2beta.ResetAllTokenProjectsResp": {
      "type": "object"
    },
    "chef.automate.api.iam.v2beta.Token": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "value": {
          "type": "string"
        },
        "active": {
          "type": "boolean",
          "format": "boolean"
        },
        "created_at": {
          "type": "string"
        },
        "updated_at": {
          "type": "string"
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "chef.automate.api.iam.v2beta.UpdateTokenReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "title": "ID can't be changed; ID used to discover token"
        },
        "name": {
          "type": "string"
        },
        "active": {
          "type": "boolean",
          "format": "boolean"
        },
        "projects": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "chef.automate.api.iam.v2beta.UpdateTokenResp": {
      "type": "object",
      "properties": {
        "token": {
          "$ref": "#/definitions/chef.automate.api.iam.v2beta.Token"
        }
      }
    }
  }
}
`)
}
