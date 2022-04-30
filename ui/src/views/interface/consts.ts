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
    regular = 'regular',
    xpath = 'xpath',
    jsonPath = 'jsonPath',
    cssSelector = 'cssSelector',
    boundary = 'boundary',
}
export enum CheckpointOperator {
    equal = '=',
    notEqual = '!=',
    greaterThan = '>',
    lessThan = '>',
    greaterThanOrEqual = '>=',
    lessThanOrEqual = '<=',
}
