package testdata

const OptionDisallowAdditionalProperties = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/OptionDisallowAdditionalProperties",
    "$fullRef": "#/definitions/samples.OptionDisallowAdditionalProperties",
    "definitions": {
        "OptionDisallowAdditionalProperties": {
            "properties": {
                "name2": {
                    "type": "string"
                },
                "timestamp2": {
                    "type": "string"
                },
                "id2": {
                    "type": "integer"
                },
                "rating2": {
                    "type": "number"
                },
                "complete2": {
                    "type": "boolean"
                }
            },
            "additionalProperties": false,
            "type": "object",
            "title": "Option Disallow Additional Properties"
        }
    }
}`

const OptionDisallowAdditionalPropertiesFail = `{"something": 12345}`

const OptionDisallowAdditionalPropertiesPass = `{"rating2": 12345}`
