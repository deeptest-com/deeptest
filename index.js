const data = {
    "type": "object",
    // 额外的视图属性，仅有type为 "object" 有该属性
    "extraViewInfo": {
        "isExpand": true,
        "isRoot": true,
        "name": "root",
        "depth": 1,
    },
    "required": [
        "name"
    ],
    "properties": {
        "name": {
            "type": "string",
            "extraViewInfo": {
                "depth": 2,
            },
        },
        "age": {
            "type": "integer",
            "format": "int32",
            "minimum": 0,
            "extraViewInfo": {
                "depth": 2,
            },
        },
        'obj1': {
            "type": "object",
            "extraViewInfo": {
                "isExpand": true,
                "isRoot": false,
                "name": "obj1",
                "depth": 2,
            },
            "required": [
                "name"
            ],
            "properties": {
                "name1": {
                    "type": "string",
                    "extraViewInfo": {
                        "depth": 3,
                    },
                },
                "age1": {
                    "type": "integer",
                    "format": "int32",
                    "minimum": 0,
                    "extraViewInfo": {
                        "depth": 3,
                    },
                },
                'obj2': {
                    "type": "object",
                    "extraViewInfo": {
                        "isExpand": true,
                        "isRoot": false,
                        "name": "obj2",
                        "depth": 3,
                    },
                    "required": [
                        "name11"
                    ],
                    "properties": {
                        "name3232": {
                            "type": "string",
                            "extraViewInfo": {
                                "depth": 4,
                            },
                        },
                        // "address": {
                        //     "$ref": "#/components/schemas/Address"
                        // },
                        "age3333": {
                            "type": "integer",
                            "format": "int32",
                            "minimum": 0,
                            "extraViewInfo": {
                                "depth": 4,
                            },
                        },
                        'obj4341': {
                            "type": "object",
                            "extraViewInfo": {
                                "isExpand": true,
                                "isRoot": false,
                                "name": 'obj4341',
                                "depth": 4,
                            },
                            "required": [
                                "name"
                            ],
                            "properties": {
                                "name1332323": {
                                    "type": "string",
                                    "extraViewInfo": {
                                        "depth": 5,
                                    },
                                },
                                "age22221": {
                                    "type": "integer",
                                    "format": "int32",
                                    "minimum": 0,
                                    "extraViewInfo": {
                                        "depth": 5,
                                    },
                                },
                            }
                        }
                    }
                }
            }
        }
    }
};

const obj = {
    "type": "object",
    "required": [
        "name"
    ],
    "properties": {
        "name": {
            "type": "string",
        },
        "age": {
            "type": "integer",
            "format": "int32",
            "minimum": 0,

        },
        'obj1': {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name1": {
                    "type": "string",
                },
                "age1": {
                    "type": "integer",
                    "format": "int32",
                    "minimum": 0,
                },
                'obj2': {
                    "type": "object",
                    "required": [
                        "name11"
                    ],
                    "properties": {
                        "name3232": {
                            "type": "string",

                        },
                        "age3333": {
                            "type": "integer",
                            "format": "int32",
                            "minimum": 0,

                        },
                        'obj4341': {
                            "type": "object",

                            "required": [
                                "name"
                            ],
                            "properties": {
                                "name1332323": {
                                    "type": "string",

                                },
                                "age22221": {
                                    "type": "integer",
                                    "format": "int32",
                                    "minimum": 0,

                                },
                            }
                        }
                    }
                }
            }
        }
    }
};

function gen(val) {
    val.extraViewInfo = {
        "isExpand": true,
        "isRoot": true,
        "name": "root",
        "depth": 1,
    };
    function fn(obj, depth) {
        if (obj.properties && obj.type === 'object') {
            Object.entries(obj.properties).forEach(([key, value]) => {
                value.extraViewInfo = {
                    "isExpand": true,
                    "isRoot": false,
                    "name": key,
                    "depth": depth,
                }
                if (value.type === 'object') {
                    fn(value, depth + 1);
                }
            })
        }
    }
    fn(val, 2);
    return val;
}
console.log(gen(obj));
