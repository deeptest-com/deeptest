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
    responseStatus = 'responseStatus',
    responseHeader = 'responseHeader',
    responseBody = 'responseBody',
    extractor = 'extractor',
}
export enum CheckpointOperator {
    contain = 'contain',
    equal = 'equal',
    notEqual = 'notEqual',
    greaterThan = 'greaterThan',
    lessThan = 'lessThan',
    greaterThanOrEqual = 'greaterThanOrEqual',
    lessThanOrEqual = 'lessThanOrEqual',
}

export enum AuthorizationTypes {
    '' = 'None',
    'basicAuth' = 'Basic Auth',
    'bearerToken' = 'Bearer Token',
    'oAuth2' = 'OAuth 2.0',
    'apiKey' = 'API Key',
}
