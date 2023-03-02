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


export const interfaceStatus = new Map([[0, '未知'], [1, '设计中'], [2, '开发中'], [3, '已发布'], [4, '已过时']])

export const serveStatus = new Map([[0, '未知'],[1, '新建'], [2, '设计中'], [3, '已发布'], [4, ' 已禁用']])


export const interfaceStatusOpts = [
    {
        label: "未知",
        value: "0",
    },
    {
        label: "设计中",
        value: "1",
    },
    {
        label: "开发中",
        value: "2",
    },
    {
        label: "已发布",
        value: "3",
    },
    {
        label: "已过时",
        value: "4",
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
