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
    return !['array', 'object', 'oneof', 'anyof', 'allOf'].includes(type);
}

/**
 * 复合类型
 * */
export function isCompositeType(type: string): boolean {
    return ['oneof', 'anyof', 'allOf'].includes(type);
}

function getExpandedValue(val: any, defaultVal: boolean) {
    return (typeof val?.extraViewInfo === 'object' && 'isExpand' in val?.extraViewInfo) ? val.extraViewInfo.isExpand : defaultVal
}

/**
 * 根据传入的 schema 结构信息，添加需要额外的渲染属性
 * */
export function addExtraViewInfo(val: Object | any | undefined | null): any {
    if (!val) {
        return null;
    }
    let type = val.type;
    if (!type) {
        // todo 先展示出来，但是还没实现这几个关键词
        if (val?.oneOf) {
            type = 'oneof';
        } else if (val?.anyOf) {
            type = 'anyof';
        } else if (val?.allOf) {
            type = 'allOf';
        } else {
            type = 'string';
        }
    }
    val.type = type;

    val.extraViewInfo = {
        "isExpand": getExpandedValue(val, true),
        "depth": 1,
        "type": type,
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
            // 需要兼容两种写法，三方导入的$ref
            obj.ref = obj.ref || obj.$ref;
            obj.name = obj.ref?.split('/')?.pop(),
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
        }
            // todo 先展示出来，但是还没实现这几个关键词
        // AND = all of XOR = one of OR = any of
        else if (tree.content?.anyof) {
            tree.content.type = 'anyof';
        } else if (tree.content?.oneof) {
            tree.content.type = 'oneof';
        } else if (tree.content?.allof) {
            tree.content.type = 'allof';
        } else {
            tree.content.type = 'string';
        }
    }
}

