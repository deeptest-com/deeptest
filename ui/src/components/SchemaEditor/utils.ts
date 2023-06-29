/**
 * 是否是对象类型
 * */
export function isObject(value: any): boolean {
    if (typeof value === 'object') {
        return value?.type === 'object'
    }
    return typeof value === 'string' && value === 'object';
}

/**
 * 是否是引用类型
 * */
export function isRef(obj: any): boolean {
    return !!obj?.ref || !!obj?.$ref;
}

/**
 * 是否是数组类型
 * */
export function isArray(value: any): boolean {
    if (typeof value === 'object') {
        return value?.type === 'array'
    }
    return typeof value === 'string' && value === 'array';
}

/**
 * 普通类型
 * */
export function isNormalType(type: string): boolean {
    return !['array', 'object'].includes(type);
}

function getExpandedValue(val: any, defaultVal: boolean) {
    return (typeof val?.extraViewInfo === 'object' && 'isExpand' in val?.extraViewInfo) ? val.extraViewInfo.isExpand : defaultVal
}

/**
 * 根据传入的 schema 结构信息，添加需要额外的渲染属性
 * */
export function addExtraViewInfo(val: Object | any | undefined | null): any {
    if (!val) {
        return null
    }
    val.extraViewInfo = {
        "isExpand": getExpandedValue(val, true),
        "depth": 1,
        "type": val.type ,
        "parent": null,
        "keyName": "root",
        "keyIndex": 0,
        "isFirst": true,
        "isLast": true,
        "isRef": isRef(val),
    };

    function traverse(obj: any, depth: number, parent: any, options: any = {}, isRefChildNode = false) {

        // base Case 普通类型，递归结束，
        if (isNormalType(obj.type) && !isRef(obj)) {
            obj.extraViewInfo = {
                ...obj.extraViewInfo || {},
                "isExpand": getExpandedValue(val, true),
                "depth": depth,
                "type": obj.type,
                "parent": parent,
                isRefChildNode,
                ...options
            }
            return;
        }
        // 处理对象类型
        if (isObject(obj.type) && !isRef(obj)) {
            obj.extraViewInfo = {
                ...obj.extraViewInfo || {},
                "isExpand": getExpandedValue(val, true),
                "depth": depth,
                "type": obj.type,
                "parent": parent,
                isRefChildNode,
                ...options
            }
            Object.entries(obj.properties || {}).forEach(([keyName, value]: any, keyIndex: number) => {
                traverse(value, depth + 1, obj, {
                        keyName,
                        keyIndex,
                        isFirst: keyIndex === 0,
                        "isRef": isRef(value),
                        ancestor: obj,
                        isLast: keyIndex === Object.keys(obj.properties).length - 1
                    },
                    isRefChildNode
                );
            })
        }
        // 处理数组类型
        if (isArray(obj.type) && !isRef(obj) && obj.items) {
            // 找到最后一个非数组类型的节点
            const {node, types} = findLastNotArrayNode(obj);
            if (node) {
                node.types = types;
                traverse(node, depth, obj, options, isRefChildNode);
            }
        }
        // 处理引用类型
        if (isRef(obj)) {
            obj.ref = obj.ref || obj.$ref;
            obj.type = obj.ref.type || val?.content?.type || 'object'
            obj.extraViewInfo = {
                ...obj.extraViewInfo || {},
                "isExpand": !!(obj?.content && obj.content?.type),
                "depth": depth,
                "type": obj.type,
                "parent": parent,
                isRef: true,
                isRefChildNode,
                ...options
            }
            if (obj?.content && obj.content?.type) {
                traverse(obj.content, depth + 1, obj, {
                    ...options,
                    isRefRootNode: true,
                }, true);
            }
        }
    }

    // array  object  ref 类型都需要递归
    if (!isNormalType(val.type) || isRef(val)) {
        traverse(val, 1, null, false);
    }
    return val;
}


/**
 * 根据传入的 schema 结构信息，删除额外的渲染属性
 * */
export function removeExtraViewInfo(val: Object | any, isRemoveRefContent = false): object | null {
    function traverse(obj: any) {
        if(obj?.extraViewInfo && isRemoveRefContent){
            delete obj?.extraViewInfo;
        }
        // base Case 普通类型，递归结束，
        if (isNormalType(obj.type) && !isRef(obj)) {
            delete obj?.extraViewInfo;
            // 切换成普通类型 之前可能是数组，所以可能有 items 属性
            delete obj?.items;
            return;
        }
        // 处理对象类型
        if (isObject(obj.type)) {
            delete obj?.extraViewInfo;
            // 切换类型之前可能是数组，所以可能有 items 属性
            delete obj?.items;
            Object.entries(obj.properties || {}).forEach(([keyName, value]: any, keyIndex: number) => {
                traverse(value);
            })
        }
        // 处理数组类型
        if (isArray(obj.type)) {
            (function fn(obj: any) {
                delete obj?.extraViewInfo;
                if (!isArray(obj.type)) {
                    traverse(obj);
                    return;
                }
                obj?.items?.type && fn(obj.items);
                if (isRemoveRefContent) {
                    // debugger;
                    // 兼容有可能是数组类型的 ref，但是且 type 属性
                    delete obj.items?.extraViewInfo;
                    // 直接删除 content 属性
                    delete obj?.content;
                } else if (obj?.content && obj.content?.type) {
                    obj?.content?.type && fn(obj.content);
                }
            })(obj);
        }
        if (isRef(obj)) {
            delete obj?.extraViewInfo;
            // 切换类型之前可能是数组，所以可能有 items 属性
            delete obj?.items;
            if (isRemoveRefContent) {
                // 直接删除 content 属性
                delete obj?.content;
            } else if (obj?.content && obj.content?.type) {
                traverse(obj.content);
            }
        }
    }

    try {
        if (!val) {
            return null
        }
        delete val?.extraViewInfo;
        traverse(val);
    } catch (e) {
        console.log(832,'removeExtraViewInfo error',e);
    }


    return val;
}

/**
 * 找到最后一个非数组类型的节点
 */
export function findLastNotArrayNode(tree: Object): any {
    const types: any = [];
    let node: any = null;

    function fn(tree: any, types: any[]) {
        if (!isArray(tree?.type) || isRef(tree)) {
            node = tree;
            return;
        }
        types.push('array');
        fn(tree.items, types);
    }

    fn(tree, types);
    return {
        node,
        types,
        parent: tree
    };
}

/**
 * 根据具体类型的数据组成的数组，生成对应的 schema 结构
 * */
export const generateSchemaByArray = (arr: any[]): any => {
    const res = {};
    arr.reduce((prev, next, index, array) => {
        if (index === 0) {
            prev = Object.assign(prev, next);
            return prev;
        } else {
            prev.items = Object.assign({}, next);
            return prev.items;
        }
    }, res);
    return res;
};

/**
 * 将 $ref 字段转成 ref
 * */
export const handleRef = (res) => {
    // 将$ref转换为ref
    function fn(obj) {
        if (!obj) return;
        if (typeof obj !== 'object') return;
        if (typeof obj === 'object') {
            Object.entries(obj).forEach(([key, value]) => {
                if (key === '$ref') {
                    obj.ref = value;
                    obj.type = obj.type || 'object';
                    delete obj.$ref;
                }
                if (typeof value === 'object') {
                    fn(value);
                }
            });
        }
    }

    fn(res);
    return res;
}




