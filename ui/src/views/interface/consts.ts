export const Methods = [
    "GET",
    "POST",
    "PUT",
    "PATCH",
    "DELETE",
    "HEAD",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "CUSTOM",
]


export enum ExtractorSrc {
    header = 'header',
    body = 'body',
}
export enum ExtractorType {
    fulltext = 'fulltext',
    regular = 'regular',
    xpath = 'xpath',
    jsonPath = 'jsonPath',
    cssSelector = 'cssSelector',
    boundary = 'boundary',
}
export enum CheckpointType {
    statusCode = 'statusCode',
    responseHeader = 'responseHeader',
    responseBody = 'responseBody',
    extractor = 'extractor',
}
export enum CheckpointOperator {
    contain = 'contain',
    equal = '=',
    notEqual = '!=',
    greaterThan = '>',
    lessThan = '<',
    greaterThanOrEqual = '>=',
    lessThanOrEqual = '<=',
}
