/**
 * 请求方法选择项
 * */
export const requestMethodOpts = [
    {
        label: "GET",
        value: "GET",
    },
    {
        label: "POST",
        value: "POST",
    },
    {
        label: "PUT",
        value: "PUT",
    },
    {
        label: "DELETE",
        value: "DELETE",
    },
    {
        label: "HEAD",
        value: "HEAD",
    },
    {
        label: "OPTION",
        value: "OPTION",
    },
    {
        label: "PATCH",
        value: "PATCH",
    },
]


/**
 * 响应码枚举
 * */
export const repCodeOpts = [
    {
        label: "200",
        value: "200",
    },
    {
        label: "404",
        value: "404",
    },
    {
        label: "500",
        value: "500",
    },
    {
        label: "501",
        value: "501",
    },
    {
        label: "502",
        value: "502",
    },
    {
        label: "503",
        value: "503",
    },
]

/**
 * path params 数据类型
 * */
export const pathParamsDataTypesOpts = [
    {
        label: "any",
        value: "any",
    },
    {
        label: "string",
        value: "string",
    },
    {
        label: "number",
        value: "number",
    },
    {
        label: "integer",
        value: "integer",
    },
    {
        label: "boolean",
        value: "boolean",
    },
    {
        label: "any",
        value: "any",
    },
    {
        label: "object",
        value: "object",
    },
]

/**
 * openApi 数据类型
 * */
export const openApiDataTypesOpts = [
    {
        label: "any",
        value: "any",
    },
    {
        label: "string",
        value: "string",
    },
    {
        label: "number",
        value: "number",
    },
    {
        label: "interger",
        value: "interger",
    },
    {
        label: "boolean",
        value: "boolean",
    },
    {
        label: "any",
        value: "any",
    },
    {
        label: "object",
        value: "object",
    },
]

export const paramsSchemaDataTypes: any = {
    "any": {
        label: "any",
        value: "any",
        props: {
            label: 'Properties',
            options: [
                {
                    label: 'enum',
                    name: 'enum',
                    component: 'selectTag',
                    type: 'array',
                    placeholder: '输入文本后按回车添加',
                },
                {
                    label: 'default',
                    name: 'default',
                    component: 'input',
                    placeholder: 'default',
                    type: 'string',
                },
                {
                    label: 'example',
                    name: 'example',
                    type: 'string',
                    component: 'input',
                    placeholder: 'example',
                },
                {
                    label: 'deprecated',
                    name: 'deprecated',
                    type: 'boolean',
                    component: 'switch',
                },
            ]
        }
    },
    "string": {
        label: "string",
        value: "string",
        props: {
            label: 'Properties',
            options: [
                {
                    label: 'enum',
                    name: 'enum',
                    component: 'selectTag',
                    type: 'array',
                    placeholder: '输入文本后按回车添加',
                },
                {
                    label: 'format',
                    name: 'format',
                    type: 'string',
                    component: 'select',
                    placeholder: 'select a value',
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
                    label: 'default',
                    name: 'default',
                    component: 'input',
                    placeholder: 'default',
                    type: 'string',
                },
                {
                    label: 'example',
                    name: 'example',
                    type: 'string',
                    component: 'input',
                    placeholder: 'example',
                },
                {
                    label: 'pattern',
                    name: 'pattern',
                    type: 'string',
                    component: 'input',
                    placeholder: 'pattern',
                },
                {
                    label: 'minLength',
                    name: 'minLength',
                    component: 'inputNumber',
                    placeholder: '>=0',
                    type: 'integer',
                },
                {
                    label: 'maxLength',
                    name: 'maxLength',
                    type: 'integer',
                    component: 'inputNumber',
                    placeholder: '>=0',
                },
                {
                    label: 'deprecated',
                    name: 'deprecated',
                    type: 'boolean',
                    component: 'switch',
                },
            ]
        }
    },
    "number": {
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
                },
                {
                    label: 'format',
                    name: 'format',
                    type: 'string',
                    component: 'select',
                    placeholder: 'select a value',
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
                    label: 'default',
                    name: 'default',
                    component: 'input',
                    placeholder: 'default',
                    type: 'string',
                },
                {
                    label: 'example',
                    name: 'example',
                    type: 'string',
                    component: 'input',
                    placeholder: 'example',
                },
                {
                    label: 'minimum',
                    name: 'minimum',
                    type: 'number',
                    component: 'inputNumber',
                    placeholder: '>=0',
                },
                {
                    label: 'maximum',
                    name: 'maximum',
                    type: 'number',
                    component: 'inputNumber',
                    placeholder: '>=0',
                },
                {
                    label: 'maxLength',
                    name: 'maxLength',
                    type: 'integer',
                    component: 'inputNumber',
                    placeholder: '>=0',
                },
                {
                    label: 'multipleOf',
                    name: 'multipleOf',
                    type: 'number',
                    component: 'inputNumber',
                    placeholder: '>=0',
                },
                {
                    label: 'exclusiveMin',
                    name: 'exclusiveMin',
                    type: 'boolean',
                    component: 'switch',
                },
                {
                    label: 'exclusiveMax',
                    name: 'exclusiveMax',
                    type: 'boolean',
                    component: 'switch',
                },
                {
                    label: 'deprecated',
                    name: 'deprecated',
                    type: 'boolean',
                    component: 'switch',
                },
            ]
        }
    },
    "boolean": {
        label: "boolean",
        value: "boolean",
        color: 'red',
        active: false,
        props: {
            label: 'Properties',
            options: [
                {
                    label: 'default',
                    name: 'default',
                    type: 'boolean',
                    component: 'select',
                    placeholder: 'select a value',
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
                    label: 'example',
                    name: 'example',
                    type: 'string',
                    component: 'input',
                    placeholder: 'example',
                },
                {
                    label: 'deprecated',
                    name: 'deprecated',
                    type: 'boolean',
                    component: 'switch',
                },
            ],
        }
    },
    "integer": {
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
                },
                {
                    label: 'format',
                    name: 'format',
                    type: 'string',
                    component: 'select',
                    placeholder: 'select a value',
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
                    label: 'default',
                    name: 'default',
                    component: 'input',
                    placeholder: 'default',
                    type: 'string',
                },
                {
                    label: 'example',
                    name: 'example',
                    type: 'string',
                    component: 'input',
                    placeholder: 'example',
                },
                {
                    label: 'minimum',
                    name: 'minimum',
                    type: 'number',
                    component: 'inputNumber',
                    placeholder: '>=0',
                },
                {
                    label: 'maximum',
                    name: 'maximum',
                    type: 'number',
                    component: 'inputNumber',
                    placeholder: '>=0',
                },
                {
                    label: 'maxLength',
                    name: 'maxLength',
                    type: 'integer',
                    component: 'inputNumber',
                    placeholder: '>=0',
                },
                {
                    label: 'multipleOf',
                    name: 'multipleOf',
                    type: 'number',
                    component: 'inputNumber',
                    placeholder: '>=0',
                },
                {
                    label: 'exclusiveMin',
                    name: 'exclusiveMin',
                    type: 'boolean',
                    component: 'switch',
                },
                {
                    label: 'exclusiveMax',
                    name: 'exclusiveMax',
                    type: 'boolean',
                    component: 'switch',
                },
                {
                    label: 'deprecated',
                    name: 'deprecated',
                    type: 'boolean',
                    component: 'switch',
                },
            ]
        }
    },
    "array": {
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
                },
                {
                    label: 'maxItems',
                    name: 'maxItems',
                    type: 'integer',
                    placeholder: '>=0',
                    component: 'inputNumber',
                },
                {
                    label: 'uniqueItems',
                    name: 'additionalProperties',
                    component: 'switch',
                    type: 'boolean',
                },
                {
                    label: 'deprecated',
                    name: 'deprecated',
                    type: 'boolean',
                    component: 'switch',
                },
            ],
        },
    },
}



export const endpointStatus = new Map([[0, '未知'], [1, '设计中'], [2, '开发中'], [3, '已发布'], [4, '已过时']])

export const serveStatus = new Map([[0, '未知'], [1, '新建'], [2, '设计中'], [3, '已发布'], [4, ' 已禁用']])

export const serveStatusTagColor = new Map([[0, 'default'], [1, 'default'], [2, 'processing'], [3, 'success'], [4, 'error']])

export const endpointStatusColor = new Map([[0, 'default'], [1, 'default'], [2, 'processing'], [3, 'success'], [4, 'error']])


export const endpointStatusOpts = [
    {
        label: "未知",
        value: 0,
    },
    {
        label: "设计中",
        value: 1,
    },
    {
        label: "开发中",
        value: 2,
    },
    {
        label: "已发布",
        value: 3,
    },
    {
        label: "已过时",
        value: 4,
    }
]


export const mediaTypesOpts = [
    {
        label: "application/json",
        value: "application/json",
    },
    {
        label: "application/EDI-X12",
        value: "application/EDI-X12",
    },
    {
        label: "application/zip",
        value: "application/zip",
    },
    {
        label: "application/octet-stream",
        value: "application/octet-stream",
    },
    {
        label: "multipart/form-data",
        value: "multipart/form-data",
    },
    {
        label: "application/x-www-form-urlencoded",
        value: "application/x-www-form-urlencoded",
    },

]


export const defaultPathParams = {
    name: '',
    type: 'string',
    description: '',
    required: false,
}

export const defaultQueryParams = {
    name: '',
    type: 'string',
    description: '',
    required: false,
}

export const defaultHeaderParams = {
    name: '',
    type: 'string',
    description: '',
    required: false,
}


export const defaultCookieParams = {
    name: '',
    type: 'string',
    description: '',
    required: false,
}

export const defaultCodeResponse = {
    "code": "",
    "endpointId": "",
    "mediaType": "",
    "description": "",
    "schemaRefId": null,
    "examples": "",
    "schemaItem": {
        "id": null,
        "name": "",
        "type": "object",
        "content": "",
        "ResponseBodyId": null
    },
    "headers": []
}


export const defaultEndpointDetail = {
    "name": "",
    "projectId": "",
    "serveId": "",
    "useId": "",
    "method": "",
    "description": "",
    "operationId": "",
    "security": "",
    "requestBody": {
        "id": null,
        "endpointId": null,
        "mediaType": "",
        "description": "",
        "schemaRefId": null,
        "examples": "",
        "schemaItem": {
            "id": null,
            "name": "",
            "type": "object",
            "content": "",
            "requestBodyId": null
        }
    },
    "responseBodies": [],
    "params": [],
    "headers": [],
    "cookies": []
}
