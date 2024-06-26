package testdata

const MessageKind12 = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/MessageKind12",
    "$fullRef": "#/definitions/samples.MessageKind12",
    "definitions": {
        "MessageKind12": {
            "oneofNames": [
                "iface"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "f": {
                    "$ref": "#/definitions/samples.MessageKind11",
                    "additionalProperties": true
                },
                "kind5": {
                    "$ref": "#/definitions/samples.MessageKind5",
                    "additionalProperties": true
                },
                "kind6": {
                    "$ref": "#/definitions/samples.MessageKind6",
                    "additionalProperties": true
                },
                "kind7": {
                    "$ref": "#/definitions/samples.MessageKind7",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 12"
        },
        "samples.MessageKind1": {
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
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 1"
        },
        "samples.MessageKind11": {
            "oneofNames": [
                "iface"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "ones": {
                    "items": {
                        "$ref": "#/definitions/samples.MessageKind1"
                    },
                    "type": "array"
                },
                "kind2": {
                    "$ref": "#/definitions/samples.MessageKind2",
                    "additionalProperties": true
                },
                "kind3": {
                    "$ref": "#/definitions/samples.MessageKind3",
                    "additionalProperties": true
                },
                "kind4": {
                    "$ref": "#/definitions/samples.MessageKind4",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 11"
        },
        "samples.MessageKind2": {
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
                "isa": {
                    "type": "boolean"
                },
                "hasa": {
                    "type": "boolean"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 2"
        },
        "samples.MessageKind3": {
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
                "someProp": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 3"
        },
        "samples.MessageKind4": {
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
                "special": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 4"
        },
        "samples.MessageKind5": {
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
                "foo": {
                    "type": "number"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 5"
        },
        "samples.MessageKind6": {
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
                "bar": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 6"
        },
        "samples.MessageKind7": {
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
                "baz": {
                    "type": "string",
                    "format": "int64"
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Message Kind 7"
        }
    }
}`
