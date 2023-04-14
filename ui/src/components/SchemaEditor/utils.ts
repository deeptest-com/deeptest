/**
 * 可参考：https://json-schema.apifox.cn/#%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B
 * */
export const JSONSchemaDataTypes = [
    {
        label: "string",
        value: "string",
        color: 'pink',
        active: false,
        props: {
            label: 'Properties',
            options: [
                {
                    label: 'enum',
                    name: 'enum',
                    component: 'selectTag',
                    type: 'array',
                    placeholder: '输入文本后按回车添加',
                    value: [],
                },
                {
                    label: 'format',
                    name: 'format',
                    type: 'string',
                    component: 'select',
                    placeholder: 'select a value',
                    value: null,
                    options: [
                        {
                            label: 'data-time',
                            value: 'data-time',
                        },
                        {
                            label: 'time',
                            value: 'time',
                        },
                        {
                            label: 'email',
                            value: 'email',
                        },
                        {
                            label: 'idn-email',
                            value: 'idn-email',
                        },
                        {
                            label: 'hostname',
                            value: 'hostname',
                        },
                        {
                            label: 'idn-hostname',
                            value: 'idn-hostname',
                        },
                        {
                            label: 'ipv4',
                            value: 'ipv4',
                        },
                        {
                            label: 'ipv6',
                            value: 'ipv6',
                        },
                        {
                            label: 'uri',
                            value: 'uri',
                        },
                        {
                            label: 'uri-reference',
                            value: 'uri-reference',
                        },
                        {
                            label: 'iri',
                            value: 'iri',
                        },
                        {
                            label: 'iri-reference',
                            value: 'iri-reference',
                        },
                        {
                            label: 'uri-template',
                            value: 'uri-template',
                        },
                        {
                            label: 'json-pointer',
                            value: 'json-pointer',
                        },
                        {
                            label: 'regex',
                            value: 'regex',
                        },
                        {
                            label: 'uuid',
                            value: 'uuid',
                        },
                        {
                            label: 'password',
                            value: 'password',
                        },
                        {
                            label: 'byte',
                            value: 'byte',
                        },
                    ],
                },
                {
                    label: 'behavior',
                    name: 'behavior',
                    type: 'boolean',
                    component: 'select',
                    placeholder: 'select a value',
                    value: null,
                    options: [
                        {
                            label: 'Read/Write',
                            value: 'Read/Write',
                        },
                        {
                            label: 'Read Only',
                            value: 'Read Only',
                        },
                        {
                            label: 'Write Only',
                            value: 'Write Only',
                        },
                    ]
                },
                {
                    label: 'default',
                    name: 'default',
                    component: 'input',
                    placeholder: 'default',
                    type: 'string',
                    value: '',
                },
                {
                    label: 'example',
                    name: 'example',
                    type: 'string',
                    component: 'input',
                    placeholder: 'example',
                    value: '',
                },
                {
                    label: 'pattern',
                    name: 'pattern',
                    type: 'string',
                    component: 'input',
                    placeholder: 'pattern',
                    value: '',
                },
                {
                    label: 'minLength',
                    name: 'minLength',
                    component: 'inputNumber',
                    placeholder: '>=0',
                    type: 'integer',
                    value: '',
                },
                {
                    label: 'maxLength',
                    name: 'maxLength',
                    type: 'integer',
                    component: 'inputNumber',
                    placeholder: '>=0',
                    value: '',
                },
                {
                    label: 'deprecated',
                    name: 'deprecated',
                    type: 'boolean',
                    component: 'switch',
                    value: false,
                },
            ]
        }
    },
    {
        label: "number",
        value: "number",
        color: 'cyan',
        active: false,
        props: {
            label: 'Properties',
            options: [
                {
                    label: 'enum',
                    name: 'enum',
                    component: 'selectTag',
                    type: 'array',
                    placeholder: '输入文本后按回车添加',
                    value: [],
                },
                {
                    label: 'format',
                    name: 'format',
                    type: 'string',
                    component: 'select',
                    placeholder: 'select a value',
                    value: null,
                    options: [
                        {
                            label: 'float',
                            value: 'float',
                        },
                        {
                            label: 'double',
                            value: 'double',
                        },
                    ]
                },
                {
                    label: 'behavior',
                    name: 'behavior',
                    type: 'boolean',
                    component: 'select',
                    placeholder: 'select a value',
                    value: null,
                    options: [
                        {
                            label: 'Read/Write',
                            value: 'Read/Write',
                        },
                        {
                            label: 'Read Only',
                            value: 'Read Only',
                        },
                        {
                            label: 'Write Only',
                            value: 'Write Only',
                        },
                    ]
                },
                {
                    label: 'default',
                    name: 'default',
                    component: 'input',
                    placeholder: 'default',
                    type: 'string',
                    value: '',
                },
                {
                    label: 'example',
                    name: 'example',
                    type: 'string',
                    component: 'input',
                    placeholder: 'example',
                    value: '',
                },
                {
                    label: 'minimum',
                    name: 'minimum',
                    type: 'number',
                    component: 'inputNumber',
                    placeholder: '>=0',
                    value: '',
                },
                {
                    label: 'maximum',
                    name: 'maximum',
                    type: 'number',
                    value: '',
                    component: 'inputNumber',
                    placeholder: '>=0',
                },
                {
                    label: 'maxLength',
                    name: 'maxLength',
                    type: 'integer',
                    component: 'inputNumber',
                    placeholder: '>=0',
                    value: '',
                },
                {
                    label: 'multipleOf',
                    name: 'multipleOf',
                    type: 'number',
                    component: 'inputNumber',
                    placeholder: '>=0',
                    value: '',
                },
                {
                    label: 'exclusiveMin',
                    name: 'exclusiveMin',
                    type: 'boolean',
                    component: 'switch',
                    value: false,
                },
                {
                    label: 'exclusiveMax',
                    name: 'exclusiveMax',
                    type: 'boolean',
                    component: 'switch',
                    value: false,
                },
                {
                    label: 'deprecated',
                    name: 'deprecated',
                    type: 'boolean',
                    component: 'switch',
                    value: false,
                },
            ]
        }
    },
    {
        label: "integer",
        value: "integer",
        color: 'green',
        active: false,
        props: {
            label: 'Properties',
            options: [
                {
                    label: 'enum',
                    name: 'enum',
                    component: 'selectTag',
                    type: 'array',
                    placeholder: '输入文本后按回车添加',
                    value: [],
                },
                {
                    label: 'format',
                    name: 'format',
                    type: 'string',
                    component: 'select',
                    placeholder: 'select a value',
                    value: null,
                    options: [
                        {
                            label: 'int32',
                            value: 'int32',
                        },
                        {
                            label: 'int64',
                            value: 'int64',
                        },
                    ]
                },
                {
                    label: 'behavior',
                    name: 'behavior',
                    type: 'boolean',
                    component: 'select',
                    placeholder: 'select a value',
                    value: null,
                    options: [
                        {
                            label: 'Read/Write',
                            value: 'Read/Write',
                        },
                        {
                            label: 'Read Only',
                            value: 'Read Only',
                        },
                        {
                            label: 'Write Only',
                            value: 'Write Only',
                        },
                    ]
                },
                {
                    label: 'default',
                    name: 'default',
                    component: 'input',
                    placeholder: 'default',
                    type: 'string',
                    value: '',
                },
                {
                    label: 'example',
                    name: 'example',
                    type: 'string',
                    component: 'input',
                    placeholder: 'example',
                    value: '',
                },
                {
                    label: 'minimum',
                    name: 'minimum',
                    type: 'number',
                    component: 'inputNumber',
                    placeholder: '>=0',
                    value: '',
                },
                {
                    label: 'maximum',
                    name: 'maximum',
                    type: 'number',
                    value: '',
                    component: 'inputNumber',
                    placeholder: '>=0',
                },
                {
                    label: 'maxLength',
                    name: 'maxLength',
                    type: 'integer',
                    component: 'inputNumber',
                    placeholder: '>=0',
                    value: '',
                },
                {
                    label: 'multipleOf',
                    name: 'multipleOf',
                    type: 'number',
                    component: 'inputNumber',
                    placeholder: '>=0',
                    value: '',
                },
                {
                    label: 'exclusiveMin',
                    name: 'exclusiveMin',
                    type: 'boolean',
                    component: 'switch',
                    value: false,
                },
                {
                    label: 'exclusiveMax',
                    name: 'exclusiveMax',
                    type: 'boolean',
                    component: 'switch',
                    value: false,
                },
                {
                    label: 'deprecated',
                    name: 'deprecated',
                    type: 'boolean',
                    component: 'switch',
                    value: false,
                },
            ]
        }
    },
    {
        label: "object",
        value: "object",
        color: 'blue',
        active: false,
        props: {
            label: 'Properties',
            options: [
                {
                    label: 'minProperties',
                    name: 'minProperties',
                    type: 'integer',
                    placeholder: '>=0',
                    component: 'inputNumber',
                    value: null,
                },
                {
                    label: 'maxProperties',
                    name: 'maxProperties',
                    type: 'integer',
                    component: 'inputNumber',
                    placeholder: '>=0',
                    value: null,
                },
                {
                    label: 'allow additional Properties',
                    name: 'additionalProperties',
                    type: 'boolean',
                    component: 'switch',
                    value: false,
                },
                {
                    label: 'deprecated',
                    name: 'deprecated',
                    type: 'boolean',
                    component: 'switch',
                    value: false,
                },
            ]
        },
    },
    {
        label: "array",
        value: "array",
        color: 'orange',
        active: false,
        props: {
            label: 'Properties',
            options: [
                {
                    label: 'minItems',
                    name: 'minItems',
                    type: 'integer',
                    placeholder: '>=0',
                    component: 'inputNumber',
                    value: null,
                },
                {
                    label: 'maxItems',
                    name: 'maxItems',
                    type: 'integer',
                    placeholder: '>=0',
                    component: 'inputNumber',
                    value: null,
                },
                {
                    label: 'uniqueItems',
                    name: 'additionalProperties',
                    component: 'switch',
                    type: 'boolean',
                    value: false,
                },
                {
                    label: 'deprecated',
                    name: 'deprecated',
                    type: 'boolean',
                    component: 'switch',
                    value: false,
                },
            ],
            subTypes: []
        },
    },
    {
        label: "boolean",
        value: "boolean",
        color: 'red',
        active: false,
        props: {
            label: 'Properties',
            options: [
                {
                    label: 'behavior',
                    name: 'behavior',
                    type: 'string',
                    component: 'select',
                    value: '',
                    options: [
                        {
                            label: 'Read/Write',
                            value: 'Read/Write',
                        },
                        {
                            label: 'Read Only',
                            value: 'Read Only',
                        },
                        {
                            label: 'Write Only',
                            value: 'Write Only',
                        },
                    ]
                },
                {
                    label: 'default',
                    name: 'default',
                    type: 'boolean',
                    component: 'select',
                    placeholder: 'select a value',
                    value: '',
                    options: [
                        {
                            label: 'true',
                            value: true,
                        },
                        {
                            label: 'false',
                            value: false,
                        },
                    ]
                },
                {
                    label: 'deprecated',
                    name: 'deprecated',
                    type: 'boolean',
                    component: 'switch',
                    value: false,
                },
            ],
        }
    },
];
/**
 * 设置schema模块数据类型
 * */
export const schemaSettingInfo = [
    {
        label: 'Type',
        subLabel: 'SubType',
        type: 'type',
        value: 'string',
        active: true,
        props: JSONSchemaDataTypes
    },
    {
        label: 'Components',
        type: '$ref',
        value: '',
        active: false,
        subLabel: 'Components',
    },
]
export const typeOpts = ['string', 'number', 'boolean', 'array', 'object', 'integer'];
// 树形结构的层级递进宽度
export const treeLevelWidth = 24;

/**
 * 是否是对象类型
 * */
export function isObject(type: string): boolean {
    return type === 'object';
}

/**
 * 是否是数组类型
 * */
export function isArray(type: string): boolean {
    return type === 'array';
}

/**
 * 普通类型
 * */
export function isNormalType(type: string): boolean {
    return !['array', 'object'].includes(type);
}

/**
 * 根据传入的 schema 结构信息，添加需要额外的渲染属性
 * */
export function addExtraViewInfo(val: Object | any | undefined | null): any {
    if (!val) {
        return null
    }
    val.extraViewInfo = {
        "isExpand": true,
        "depth": 1,
        "type": val.type,
        "parent": null,
        "keyName": "root",
        "keyIndex": 0,
        "isFirst": true,
        "isLast": true,
    };

    function traverse(obj: any, depth: number, parent: any, options: any = {}) {
        // base Case 普通类型，递归结束，
        if (isNormalType(obj.type)) {
            obj.extraViewInfo = {
                "isExpand": true,
                "depth": depth,
                "type": obj.type,
                "parent": parent,
                ...options
            }
            return;
        }
        // 处理对象类型
        if (isObject(obj.type)) {
            obj.extraViewInfo = {
                "isExpand": true,
                "depth": depth,
                "type": obj.type,
                "parent": parent,
                ...options
            }
            Object.entries(obj.properties || {}).forEach(([keyName, value]: any, keyIndex: number) => {
                traverse(value, depth + 1, obj, {
                        keyName,
                        keyIndex,
                        isFirst: keyIndex === 0,
                        ancestor: obj,
                        isLast: keyIndex === Object.keys(obj.properties).length - 1
                    }
                );
            })
        }
        // 处理数组类型
        if (isArray(obj.type) && obj.items) {
            // 找到最后一个非数组类型的节点
            const {node, types} = findLastNotArrayNode(obj);
            if (node) {
                node.types = types;
                traverse(node, depth, obj, options);
            }
        }
    }

    if (!isNormalType(val.type)) {
        traverse(val, 1, null);
    }
    return val;
}


/**
 * 根据传入的 schema 结构信息，删除额外的渲染属性
 * */
export function removeExtraViewInfo(val: Object | any): object | null {
    if (!val) {
        return null
    }
    delete val?.extraViewInfo;

    function traverse(obj: any) {
        // base Case 普通类型，递归结束，
        if (isNormalType(obj.type)) {
            delete obj?.extraViewInfo;
            // 普通类型，需要删除 items 属性
            delete obj?.items;
            return;
        }
        // 处理对象类型
        if (isObject(obj.type)) {
            delete obj?.extraViewInfo;
            Object.entries(obj.properties || {}).forEach(([keyName, value]: any, keyIndex: number) => {
                traverse(value);
            })
        }
        // 处理数组类型
        if (isArray(obj.type)) {
            (function fn(obj: any) {
                if (!isArray(obj.type)) {
                    traverse(obj);
                    return;
                }
                delete obj?.extraViewInfo;
                fn(obj.items);
            })(obj);
        }
    }

    traverse(val);
    return val;
}

/**
 * 找到最后一个非数组类型的节点
 */
export function findLastNotArrayNode(tree: Object): any {
    const types: any = [];
    let node: any = null;

    function fn(tree: any, types: any[]) {
        if (!isArray(tree.type)) {
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
 * 移动光标在最后一个位置
 *
 * */
export const moveCursorToEnd = (dom: Object): void => {
    let range;
    if (window.getSelection) {//ie11 10 9 ff safari
        range = window.getSelection();//创建range
        range.selectAllChildren(dom);//range 选择obj下所有子内容
        range.collapseToEnd();//光标移至最后
    }
};
