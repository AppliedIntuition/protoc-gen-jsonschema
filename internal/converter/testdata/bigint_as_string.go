package testdata

const BigIntAsString = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/BigIntAsString",
    "$fullRef": "#/definitions/samples.BigIntAsString",
    "definitions": {
        "BigIntAsString": {
            "properties": {
                "big_number": {
                    "oneOf": [
                        {
                            "type": "integer",
                            "format": "int64"
                        },
                        {
                            "type": "null"
                        }
                    ]
                }
            },
            "additionalProperties": true,
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "object"
                }
            ],
            "title": "Big Int As String"
        }
    }
}`

const BigIntAsStringFail = `{"big_number": "1827634182736443333"}`

const BigIntAsStringPass = `{"big_number": 1827634182736443333}`
