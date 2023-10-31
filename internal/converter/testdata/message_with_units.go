package testdata

const MessageWithUnits = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/MessageWithUnits",
    "$fullRef": "#/definitions/samples.MessageWithUnits",
    "definitions": {
        "MessageWithUnits": {
            "properties": {
                "height": {
                    "type": "integer",
                    "options": {
                        "units": 2
                    }
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message With Units"
        }
    }
}`
