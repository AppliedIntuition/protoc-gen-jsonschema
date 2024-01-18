package testdata

const OptionDefaultValue = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/OptionDefaultValue",
    "$fullRef": "#/definitions/samples.OptionDefaultValue",
    "definitions": {
        "OptionDefaultValue": {
            "properties": {
                "field_one": {
                    "type": "string",
                    "options": {
                        "DefaultValue": "sample_str"
                    }
                },
                "field_two": {
                    "items": {
                        "type": "boolean"
                    },
                    "type": "array",
                    "options": {
                        "DefaultValue": true
                    }
                },
                "field_three": {
                    "type": "integer",
                    "options": {
                        "DefaultValue": 7
                    }
                },
                "field_four": {
                    "type": "integer"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Option Default Value"
        }
    }
}`

