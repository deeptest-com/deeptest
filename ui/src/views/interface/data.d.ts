export interface Request {
    url: string;
    headers: Header[];
    params: Param[];

    authorizationType: string
    basicAuth: BasicAuth
    bearerToken: BearerToken
    oAuth20: OAuth20
    apiKey: ApiKey
}

export interface Param {
    name: string;
    value: any;
    disabled: boolean
}

export interface Header {
    name: string;
    value: any;
    disabled: boolean
}

export interface BasicAuth {
    username: string;
    password: string;
}
export interface BearerToken {
    username: string;
}
export interface OAuth20 {
    key: string; // key
    oidcDiscoveryURL: string; // OpenID Connect Discovery URL
    authURL: string // Authentication URL
    accessTokenURL: string // Access Token URL
    clientID: string // Client ID
    scope: string // Scope
}
export interface ApiKey {
    username: string;
    value: string;
    transferMode: string
}

export interface Response {
    header: any;
    body: string;
}