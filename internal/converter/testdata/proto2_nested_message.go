package testdata

const Proto2NestedMessage = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/Proto2NestedMessage",
    "$fullRef": "#/definitions/samples.Proto2NestedMessage",
    "definitions": {
        "Proto2NestedMessage": {
            "required": [
                "payload"
            ],
            "properties": {
                "payload": {
                    "$ref": "#/definitions/samples.Proto2PayloadMessage",
                    "additionalProperties": false
                },
                "description": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Proto 2 Nested Message"
        },
        "samples.Proto2PayloadMessage": {
            "required": [
                "name",
                "id"
            ],
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
                    "$ref": "#/definitions/samples.Proto2PayloadMessage.Topology",
                    "title": "Topology"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Proto 2 Payload Message"
        },
        "samples.Proto2PayloadMessage.Topology": {
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

const Proto2NestedMessageFail = `{
	"payload": {
		"topology": "FLAT"	
	}
}`

const Proto2NestedMessagePass = `{
	"payload": {
		"id": 1,
		"name": "something",
		"topology": "FLAT"
	}
}`
