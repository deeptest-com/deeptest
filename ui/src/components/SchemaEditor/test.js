const obj1 = {
    "type": "array",
    "items": {
        "type": "array",
        "items": {
            "type": "array",
            items: {
                "type": "object",
                "properties": {
                    "name1": {
                        "type": "string",
                    },
                    "arr1": {
                        "type": "array",
                        "items": {
                            "type": "string",
                        }
                    },
                }
            }
        }
    }
};

function addExtraInfo(val) {
    if (!val) {
        return null
    }
    val.extraViewInfo = {
        "isExpand": true,
        "name": "root",
        "depth": 1,
        "type": val.type,
        "parent": null,
    };

    function fn(obj, depth) {
        if (obj.properties && obj.type === 'object') {
            Object.entries(obj.properties).forEach(([key, value]) => {
                if (value.type === 'array') {
                    fn(value.items, depth);
                } else {
                    value.extraViewInfo = {
                        "isExpand": true,
                        "name": key,
                        "depth": depth,
                        "type": obj.type,
                        "parent": obj,
                    }
                    if (value.type === 'object') {
                        fn(value, depth + 1);
                    }
                }
            })
        }
        if (obj.type === 'array' && obj.items) {
            if (obj.items.type === 'object') {
                Object.entries(obj.items.properties).forEach(([key, value]) => {
                    value.extraViewInfo = {
                        "isExpand": true,
                        "name": key,
                        "depth": depth,
                        "type": value.type,
                        "items": obj.items,
                        "parent": obj,
                    }
                    if (value.type === 'object') {
                        fn(value, depth + 1);
                    } else if (value.type === 'array') {
                        fn(value.items, depth);
                    }
                })
            } else if (obj.items.type === 'array') {
                fn(obj.items, depth)
            }
        }
    }

    fn(val, 2);
    return val;
}

console.log(addExtraInfo(obj1));


