package testdata

const MessageUnusedEnum = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/MessageUnusedEnum",
    "$fullRef": "#/definitions/samples.MessageUnusedEnum",
    "definitions": {
        "MessageUnusedEnum": {
            "properties": {
                "var": {
                    "type": "boolean"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Unused Enum"
        },
        "samples.MessageUnusedEnum.NestedEnum": {
            "enum": [
                "Foo",
                0,
                "Bar",
                1,
                "Baz",
                2
            ],
            "oneOf": [
                {
                    "type": "string"
                },
                {
                    "type": "integer"
                }
            ],
            "title": "Nested Enum"
        }
    }
}`
