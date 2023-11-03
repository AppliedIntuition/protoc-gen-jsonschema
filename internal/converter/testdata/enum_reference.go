package testdata

const EnumReference1 = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/MessageWithEnums",
    "$fullRef": "#/definitions/samples.MessageWithEnums",
    "definitions": {
        "MessageWithEnums": {
            "properties": {
                "enumFieldOne": {
                    "$ref": "#/definitions/samples.EnumOne",
                    "deprecated": true,
                    "title": "Enum One",
                    "description": "Comment about EnumOne."
                },
                "enumFieldTwo": {
                    "$ref": "#/definitions/samples.MessageWithEnums.NestedEnum",
                    "title": "Nested Enum"
                },
                "enum_with_manual": {
                    "$ref": "#/definitions/samples.EnumOne",
                    "options": {
                        "manualLink": "manual-link-1"
                    },
                    "title": "Enum One"
                },
                "enum_with_manual_array": {
                    "items": {
                        "$ref": "#/definitions/samples.EnumOne",
                        "title": "Enum One"
                    },
                    "type": "array",
                    "options": {
                        "manualLink": "manual-link-2"
                    }
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message With Enums"
        },
        "samples.EnumOne": {
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
            "title": "Enum One"
        },
        "samples.MessageWithEnums.DefinedUnusedEnum": {
            "enum": [
                "Var",
                0
            ],
            "oneOf": [
                {
                    "type": "string"
                },
                {
                    "type": "integer"
                }
            ],
            "title": "Defined Unused Enum"
        },
        "samples.MessageWithEnums.DefinedUnusedMessage": {
            "properties": {
                "var": {
                    "type": "boolean"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Defined Unused Message"
        },
        "samples.MessageWithEnums.DefinedUnusedMessage.NestedUnusedMessage": {
            "properties": {
                "foo": {
                    "type": "boolean"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Nested Unused Message"
        },
        "samples.MessageWithEnums.DefinedUsedMessage": {
            "properties": {
                "var": {
                    "type": "boolean"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Defined Used Message"
        },
        "samples.MessageWithEnums.NestedEnum": {
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

const EnumReference2 = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$fullRef": "#/definitions/samples.EnumOne",
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
    "title": "Enum One"
}
                `
