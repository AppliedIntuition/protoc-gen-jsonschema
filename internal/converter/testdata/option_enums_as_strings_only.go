package testdata

const OptionEnumsAsStringsOnly = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$fullRef": "#/definitions/samples.Currency",
    "enum": [
        "NOT_SPECIFIED",
        "USD",
        "GBP",
        "EUR"
    ],
    "enumValues": [
        0,
        1,
        2,
        3
    ],
    "type": "string",
    "title": "Currency"
}`

const OptionEnumsAsStringsOnlyPass = `"NOT_SPECIFIED"`
const OptionEnumsAsStringsOnlyFail = `2`
