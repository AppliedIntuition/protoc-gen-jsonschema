package testdata

const NestedObject = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/NestedObject",
    "$fullRef": "#/definitions/samples.NestedObject",
    "definitions": {
        "NestedObject": {
            "properties": {
                "payload": {
                    "$ref": "#/definitions/samples.NestedObject.NestedPayload",
                    "additionalProperties": true
                },
                "description": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Nested Object"
        },
        "samples.NestedObject.NestedPayload": {
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
                    "$ref": "#/definitions/samples.NestedObject.NestedPayload.Topology",
                    "title": "Topology"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Nested Payload"
        },
        "samples.NestedObject.NestedPayload.Topology": {
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

const NestedObjectFail = `{"payload": false}`

const NestedObjectPass = `{
	"payload": {
	  "topology": "NESTED_OBJECT"
	}
}`
