package testdata

const NestedMessage = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/NestedMessage",
    "$fullRef": "#/definitions/samples.NestedMessage",
    "definitions": {
        "NestedMessage": {
            "properties": {
                "payload": {
                    "$ref": "#/definitions/samples.PayloadMessage",
                    "additionalProperties": true
                },
                "description": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Nested Message"
        },
        "samples.PayloadMessage": {
            "properties": {
                "name": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                },
                "complete": {
                    "type": "boolean"
                },
                "topology": {
                    "$ref": "#/definitions/samples.PayloadMessage.Topology",
                    "title": "Topology"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Payload Message"
        },
        "samples.PayloadMessage.Topology": {
            "enum": [
                "FLAT",
                0,
                "NESTED_OBJECT",
                1,
                "NESTED_MESSAGE",
                2,
                "ARRAY_OF_TYPE",
                3,
                "ARRAY_OF_OBJECT",
                4,
                "ARRAY_OF_MESSAGE",
                5
            ],
            "oneOf": [
                {
                    "type": "string"
                },
                {
                    "type": "integer"
                }
            ],
            "title": "Topology"
        }
    }
}`

const NestedMessageFail = `{
	"payload": {
		"topology": "ROUND"
	}
}`

const NestedMessagePass = `{
	"payload": {
		"topology": "FLAT"
	}
}`
