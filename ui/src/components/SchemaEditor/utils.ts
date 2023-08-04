/*
* todo 改文件的很多公共方法需要提出去
* */

import {message} from "ant-design-vue";

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
    return !['array', 'object', 'oneOf', 'anyOf', 'allOf'].includes(type);
}

/**
 * 复合类型
 * */
export function isCompositeType(type: string): boolean {
    return ['oneOf', 'anyOf', 'allOf'].includes(type);
}

/**
 * 获取该节点的展开状态
 * */
function getExpandedValue(val: any, defaultVal: boolean) {
    return (typeof val?.extraViewInfo === 'object' && 'isExpand' in val?.extraViewInfo) ? val.extraViewInfo.isExpand : defaultVal;
}

/**
 * 处理组合类型的节点的类型
 * Notice: 该函数会修改传入的参数
 * */
function handleCompositeChildNode(val) {
    let type = val.type;
    if (!type) {
        if (val?.oneOf) {
            type = 'oneOf';
        } else if (val?.anyOf) {
            type = 'anyOf';
        } else if (val?.allOf) {
            type = 'allOf';
        } else {
            type = 'string';
        }
    }
    val.type = type;
}

/**
 * 根据传入的 schema 结构信息，添加需要额外的渲染属性
 * */
export function addExtraViewInfo(val: Object | any | undefined | null): any {
    console.log(8321, 'addExtraViewInfo', val);
    if (!val) {
        return null;
    }
    handleCompositeChildNode(val);
    val.extraViewInfo = {
        "isExpand": getExpandedValue(val, true),
        "depth": 1,
        "type": val.type,
        "parent": null,
        "keyName": "root",
        "keyIndex": 0,
        "isFirst": true,
        "isLast": true,
        "isRef": isRef(val),
    };

    /**
     *  @description 递归遍历 schema 结构，添加额外的渲染属性
     *  @param obj 当前节点
     *  @param depth 深度
     *  @param parent 父节点
     *  @param options 额外的参数
     *  @param isRefChildNode 是否是引用类型的子节点
     *  @param isCompositeChildNode 是否是复合类型的子节点
     * */
    function traverse(obj: any, depth: number, parent: any, options: any = {}, isRefChildNode = false, isCompositeChildNode = false) {
        // 普通类型，递归结束
        if (isNormalType(obj.type) && !isRef(obj)) {
            obj.extraViewInfo = {
                ...obj.extraViewInfo || {},
                "depth": depth,
                "type": obj.type,
                "parent": parent,
                isRefChildNode,
                isCompositeChildNode,
                ...options,
                "isExpand": getExpandedValue(val, true),
            }
            return;
        }
        // 处理对象类型
        if (isObject(obj.type) && !isRef(obj)) {
            obj.extraViewInfo = {
                ...obj.extraViewInfo || {},
                "depth": depth,
                "type": obj.type,
                "parent": parent,
                isRefChildNode,
                isCompositeChildNode,
                ...options,
                "isExpand": getExpandedValue(val, true),
            }
            Object.entries(obj.properties || {}).forEach(([keyName, value]: any, keyIndex: number) => {
                // 处理引用类型，添加 type 属性
                handleCompositeChildNode(value);
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
                // 处理组合类型，添加 type 属性
                handleCompositeChildNode(node);
                traverse(node, depth, obj, options, isRefChildNode, isCompositeChildNode);
            }
        }
        // 处理数组类型
        if (isCompositeType(obj.type) && !isRef(obj)) {
            obj.extraViewInfo = {
                ...obj.extraViewInfo || {},
                "depth": depth,
                "type": obj.type,
                "parent": parent,
                isRefChildNode,
                isCompositeChildNode,
                ...options,
                "isExpand": getExpandedValue(val, true),
            }
            const combines = {
                oneOf: obj.oneOf,
                anyOf: obj.anyOf,
                allOf: obj.allOf,
            }
            if (combines[obj.type]?.length) {
                combines[obj.type].forEach((item: any, index: number) => {
                    traverse(item, depth + 1, obj, {
                        keyName: index,
                        keyIndex: index,
                        isFirst: index === 0,
                        isLast: index === combines[obj.type].length - 1,
                        ancestor: obj,
                        isRef: isRef(item),
                    }, isRefChildNode, true);
                })
            }
        }
        // 处理引用类型
        if (isRef(obj)) {
            // 需要兼容两种写法，三方导入的$ref
            obj.ref = obj.ref || obj.$ref;
            obj.name = obj.ref?.split('/')?.pop();
            obj.extraViewInfo = {
                ...obj.extraViewInfo || {},
                "depth": depth,
                "type": obj.type,
                "parent": parent,
                isRef: true,
                isRefChildNode,
                isCompositeChildNode,
                ...options,
                "isExpand": !!(obj?.content && obj.content?.type),
            }
            if (obj?.content && obj.content?.type) {
                traverse(obj.content, depth + 1, obj, {
                    ...options,
                    isRefRootNode: true,
                }, true, isCompositeChildNode);
            }
        }
    }

    // array  object  ref 类型都需要递归
    if (!isNormalType(val.type) || isRef(val)) {
        traverse(val, 1, null, null, isRef(val), isCompositeType(val.type));
    }

    console.log(8322, 'addExtraViewInfo', val)

    return val;
}


/**
 * 根据传入的 schema 结构信息，删除额外的渲染属性
 * */
export function removeExtraViewInfo(val: Object | any, isRemoveRefContent = false): object | null {
    function traverse(obj: any) {
        if (obj?.extraViewInfo && isRemoveRefContent) {
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
        // 处理 composite 类型
        if (isCompositeType(obj.type)) {
            delete obj?.extraViewInfo;
            delete obj?.items;
            const combines = {
                oneOf: obj.oneOf,
                anyOf: obj.anyOf,
                allOf: obj.allOf,
            }
            if (combines[obj.type]?.length) {
                combines[obj.type].forEach((item: any, index: number) => {
                    traverse(item);
                })
            }
        }
        // 处理数组类型
        if (isArray(obj.type)) {
            (function fn(obj: any) {
                delete obj?.extraViewInfo;
                if (!isArray(obj.type)) {
                    traverse(obj);
                    return;
                }
                obj?.items && fn(obj.items);
                if (isRemoveRefContent) {
                    // debugger;
                    // 兼容有可能是数组类型的 ref，但是且 type 属性
                    delete obj.items?.extraViewInfo;
                    // 直接删除 content 属性
                    delete obj?.content;
                } else if (obj?.content && obj.content?.type) {
                    obj?.content && fn(obj.content);
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
        console.log(832, 'removeExtraViewInfo error', e);
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
 * @description: 根据传入的 schema 结构信息，生成对应的 ref 信息
 * @param {Object} tree 传入的 schema 结构信息
 * @param {Object} result 获取到的的 ref 信息
 * notice : 有副作用，会修改 tree
 * */
export const handleRefInfo = (tree: any, result: any) => {

    // 兼容，返回的值为空字符串的情况，则直接不展开
    if (!result?.content) {
        tree.extraViewInfo.isExpand = false;
        message.warning(`引用的字段的详情数据为空`);
        return;
    }

    tree.content = JSON.parse(result.content);

    // 兼容获取引用详情时，没有 type 字段的情况
    // 如果外层 result 有 type 字段，则直接使用
    // 否则，根据 content 的结构，判断 type
    tree.content.type = tree.content.type || result.type;
    if (!tree.content?.type) {
        if (result?.properties) {
            tree.content.type = 'object';
        } else if (result?.items) {
            tree.content.type = 'array';
        } else if (result?.allOf) {
            tree.content.type = 'array';
        } else if (tree.content?.anyOf) {
            tree.content.type = 'anyOf';
        } else if (tree.content?.oneOf) {
            tree.content.type = 'oneOf';
        } else if (tree.content?.allOf) {
            tree.content.type = 'allOf';
        } else {
            tree.content.type = 'string';
        }
    }
}

