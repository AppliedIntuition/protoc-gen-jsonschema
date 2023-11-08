package testdata

const NestedUnreferencedEnum = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/Outfit",
    "$fullRef": "#/definitions/samples.Outfit",
    "definitions": {
        "Outfit": {
            "additionalProperties": true,
            "type": "object",
            "title": "Outfit",
            "description": "A set of predefined gear ratios and torque maps."
        },
        "samples.Outfit.Pants": {
            "properties": {
                "pants_type": {
                    "$ref": "#/definitions/samples.Outfit.Pants.PantsType",
                    "type": "string",
                    "title": "Pants Type",
                    "description": "When pants_type is unset, no pants are selected. Not recommended."
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Pants"
        },
        "samples.Outfit.Pants.PantsType": {
            "enum": [
                "SHORTS",
                "LONG_PANTS"
            ],
            "enumValues": [
                0,
                1
            ],
            "type": "string",
            "title": "Pants Type"
        },
        "samples.Outfit.Shirt": {
            "properties": {
                "shirt_type": {
                    "$ref": "#/definitions/samples.Outfit.Shirt.ShirtType",
                    "type": "string",
                    "title": "Shirt Type",
                    "description": "When shirt_type is unset, no shirt is selected."
                }
            },
            "additionalProperties": true,
            "type": "object",
            "title": "Shirt"
        },
        "samples.Outfit.Shirt.ShirtType": {
            "enum": [
                "T_SHIRT",
                "LONG_SLEEVE",
                "NO_SHIRT"
            ],
            "enumValues": [
                0,
                1,
                2
            ],
            "type": "string",
            "title": "Shirt Type"
        }
    }
}`
