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
                    "type": "boolean",
                    "options": {
                        "DefaultValue": true
                    }
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Option Default Value"
        }
    }
}`

