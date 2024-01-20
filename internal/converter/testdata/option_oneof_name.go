package testdata

const OptionOneofName = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/OptionOneofName",
    "$fullRef": "#/definitions/samples.OptionOneofName",
    "definitions": {
        "OptionOneofName": {
            "properties": {
                "foo": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array",
                    "options": {
                        "oneofName": "oneof_field"
                    }
                },
                "bar": {
                    "items": {
                        "type": "string"
                    },
                    "type": "array",
                    "options": {
                        "oneofName": "oneof_field"
                    }
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Option Oneof Name"
        }
    }
}`
