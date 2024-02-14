package testdata

const EnumWithComments = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$fullRef": "#/definitions/samples.EnumWithComments",
    "enum": [
        "VAL1",
        0,
        "VAL2",
        1,
        "VAL3",
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
    "enumDescription": {
        "VAL1": "This description discusses value number 1. Not much to say about value 1.",
        "VAL2": "Value 2 comes after value1. It's not as important.",
        "VAL3": "Value 3 is the most important. Here is a multiline description about the third value. Done talking about value 3 now."
    },
    "title": "Enum With Comments",
    "description": "This enum contains a lot of information."
}`
