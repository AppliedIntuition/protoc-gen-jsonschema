package testdata

const OptionEnumsTrimPrefix = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$fullRef": "#/definitions/samples.Scheme",
    "enum": [
        "UNSPECIFIED",
        "HTTP",
        "HTTPS"
    ],
    "enumValues": [
        0,
        1,
        2
    ],
    "type": "string",
    "title": "Scheme"
}`

const OptionEnumsTrimPrefixPass = `"HTTP"`

const OptionEnumsTrimPrefixFail = `4`
