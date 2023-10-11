package testdata

const MessageUnusedMessage = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/MessageUnusedMessage",
    "$fullRef": "#/definitions/samples.MessageUnusedMessage",
    "definitions": {
        "MessageUnusedMessage": {
            "properties": {
                "var": {
                    "type": "boolean"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Unused Message"
        },
        "samples.MessageUnusedMessage.NestedMessage": {
            "properties": {
                "foo": {
                    "type": "boolean"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Nested Message"
        }
    }
}`
