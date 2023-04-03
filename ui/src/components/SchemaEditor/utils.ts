

export function isObject(type: string) {
    return type === 'object';
}

export function isArray(type: string) {
    return type === 'array';
}

/**
 * 根据传入的 schema 结构信息，添加需要额外的渲染属性
 * */
// 适配数据结构
export function addExtraInfo(val: any) {
    if (!val) {
        return null
    }
    val.extraViewInfo = {
        "isExpand": true,
        "isRoot": true,
        "name": "root",
        "depth": 1,
    };
    function fn(obj: any, depth: number) {
        if (obj.properties && obj.type === 'object') {
            Object.entries(obj.properties).forEach(([key, value]: any) => {
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
        if (obj.type === 'array' && obj.items) {
            Object.entries(obj.items.properties).forEach(([key, value]: any) => {
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
    console.log(222,val);
    return val;
}


/**
 * 根据传入的 schema 结构信息，删除额外的渲染属性
 * */
export function removeExtraInfo(val) {
    if (!val) {
        return null
    }
    if (val.extraViewInfo) {
        delete val.extraViewInfo;
    }
    function fn(obj: any) {
        if (obj.properties && obj.type === 'object') {
            Object.entries(obj.properties).forEach(([key, value]: any) => {
                if (value.extraViewInfo) {
                    delete value.extraViewInfo;
                }
                if (value.type === 'object') {
                    fn(value);
                }
            })
        }
    }
    fn(val);
    return val;
}

