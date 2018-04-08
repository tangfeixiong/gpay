package pb 

const (
swagger = `{
  "swagger": "2.0",
  "info": {
    "title": "pb/service.proto",
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
    "/api/v1/crd": {
      "get": {
        "operationId": "ReapCrd",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbCrdReqResp"
            }
          }
        },
        "tags": [
          "SimpleGRpcService"
        ]
      },
      "post": {
        "operationId": "CreateCrd",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbCrdReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbCrdReqResp"
            }
          }
        ],
        "tags": [
          "SimpleGRpcService"
        ]
      }
    }
  },
  "definitions": {
    "CrdRecipientResourceScope": {
      "type": "string",
      "enum": [
        "Cluster",
        "Namespaced"
      ],
      "default": "Cluster"
    },
    "pbCrdRecipient": {
      "type": "object",
      "properties": {
        "group": {
          "type": "string",
          "format": "string"
        },
        "kind": {
          "type": "string",
          "format": "string"
        },
        "plural": {
          "type": "string",
          "format": "string"
        },
        "resource_scope": {
          "$ref": "#/definitions/CrdRecipientResourceScope"
        },
        "scope": {
          "type": "string",
          "format": "string"
        },
        "singular": {
          "type": "string",
          "format": "string"
        },
        "version": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbCrdReqResp": {
      "type": "object",
      "properties": {
        "recipe": {
          "$ref": "#/definitions/pbCrdRecipient"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string",
          "format": "string"
        }
      }
    }
  }
}
`
)
