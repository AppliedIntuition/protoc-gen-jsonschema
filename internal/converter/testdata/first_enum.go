package testdata

const FirstEnum = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$fullRef": "#/definitions/samples.FirstEnum",
    "enum": [
        "VALUE_0",
        0,
        "VALUE_1",
        1,
        "VALUE_2",
        2,
        "VALUE_3",
        3
    ],
    "oneOf": [
        {
            "type": "string"
        },
        {
            "type": "integer"
        }
    ],
    "title": "First Enum"
}`

const FirstEnumFail = `5`

const FirstEnumPass = `3`
