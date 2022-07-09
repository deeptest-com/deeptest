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

export enum ComparisonOperator {
    contain = 'contain',
    notContain = 'notContain',
    equal = 'equal',
    notEqual = 'notEqual',
    greaterThan = 'greaterThan',
    lessThan = 'lessThan',
    greaterThanOrEqual = 'greaterThanOrEqual',
    lessThanOrEqual = 'lessThanOrEqual',
}

export enum ExtractorSrc {
    header = 'header',
    body = 'body',
}
export enum ExtractorType {
    boundary = 'boundary',
    jsonquery = 'jsonquery',
    htmlquery = 'htmlquery',
    xmlquery = 'xmlquery',
    // regular = 'regular',
    // fulltext = 'fulltext',
}
export enum CheckpointType {
    responseStatus = 'responseStatus',
    responseHeader = 'responseHeader',
    responseBody = 'responseBody',
    extractor = 'extractor',
}

export enum AuthorizationTypes {
    '' = 'None',
    'basicAuth' = 'Basic Auth',
    'bearerToken' = 'Bearer Token',
    'oAuth2' = 'OAuth 2.0',
    'apiKey' = 'API Key',
}

export enum OAuth2GrantTypes {
    'authorizationCode' = 'Authorization Code',
    'authorizationCodeWithPKCE' = 'Authorization Code (With PKCE)',
    'implicit' = 'Implicit',
    'passwordCredential' = 'Password Credential',
    'clientCredential' = 'Client Credential',
}

export enum OAuth2ClientAuthenticationWay {
    'sendAsBasicAuthHeader' = 'Send as Basic Auth header',
    'sendClientCredentialsInBody' = 'Send client credentials in body',
}
