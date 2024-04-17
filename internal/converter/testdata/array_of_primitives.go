package testdata

const ArrayOfPrimitives = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/ArrayOfPrimitives",
    "$fullRef": "#/definitions/samples.ArrayOfPrimitives",
    "definitions": {
        "ArrayOfPrimitives": {
            "properties": {
                "description": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "string"
                        }
                    ]
                },
                "luckyNumbers": {
                    "items": {
                        "oneOf": [
                            {
                                "type": "null"
                            },
                            {
                                "type": "integer"
                            }
                        ]
                    },
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "array"
                        }
                    ]
                },
                "luckyBigNumbers": {
                    "items": {
                        "oneOf": [
                            {
                                "type": "string",
                                "format": "int64"
                            },
                            {
                                "type": "null"
                            }
                        ]
                    },
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "array"
                        }
                    ]
                },
                "keyWords": {
                    "items": {
                        "oneOf": [
                            {
                                "type": "null"
                            },
                            {
                                "type": "string"
                            }
                        ]
                    },
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "array"
                        }
                    ]
                },
                "big_number": {
                    "oneOf": [
                        {
                            "type": "string",
                            "format": "int64"
                        },
                        {
                            "type": "null"
                        }
                    ]
                }
            },
            "additionalProperties": true,
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "object"
                }
            ],
            "title": "Array Of Primitives"
        }
    }
}`

const ArrayOfPrimitivesFail = `{"luckyNumbers": ["false"]}`

const ArrayOfPrimitivesPass = `{"luckyNumbers": [1,2,3]}`

const ArrayOfPrimitivesDouble = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/ArrayOfPrimitives",
    "$fullRef": "#/definitions/samples.ArrayOfPrimitives",
    "definitions": {
        "ArrayOfPrimitives": {
            "properties": {
                "description": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "string"
                        }
                    ]
                },
                "luckyNumbers": {
                    "items": {
                        "oneOf": [
                            {
                                "type": "null"
                            },
                            {
                                "type": "integer"
                            }
                        ]
                    },
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "array"
                        }
                    ]
                },
                "luckyBigNumbers": {
                    "items": {
                        "oneOf": [
                            {
                                "type": "string",
                                "format": "int64"
                            },
                            {
                                "type": "null"
                            }
                        ]
                    },
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "array"
                        }
                    ]
                },
                "keyWords": {
                    "items": {
                        "oneOf": [
                            {
                                "type": "null"
                            },
                            {
                                "type": "string"
                            }
                        ]
                    },
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "array"
                        }
                    ]
                },
                "big_number": {
                    "oneOf": [
                        {
                            "type": "string",
                            "format": "int64"
                        },
                        {
                            "type": "null"
                        }
                    ]
                },
                "bigNumber": {
                    "oneOf": [
                        {
                            "type": "string",
                            "format": "int64"
                        },
                        {
                            "type": "null"
                        }
                    ]
                }
            },
            "additionalProperties": true,
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "object"
                }
            ],
            "title": "Array Of Primitives"
        }
    }
}`

const ArrayOfPrimitivesDoubleFail = `{"bigNumber": false}`

const ArrayOfPrimitivesDoublePass = `{"bigNumber": "2"}`
