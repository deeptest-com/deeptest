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

export enum OAuth2GrantTypes {
    'AuthorizationCode' = 'Authorization Code',
    'AuthorizationCodeWithPKCE' = 'Authorization Code (With PKCE)',
    'Implicit GrantType' = 'Implicit',
    'PasswordCredential' = 'Password Credential',
    'ClientCredential' = 'Client Credential',
}

export enum OAuth2ClientAuthenticationWay {
    'SendAsBasicAuthHeader' = 'Send as Basic Auth header',
    'SendClientCredentialsInBody' = 'Send client credentials in body',
}
