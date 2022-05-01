import {CheckpointOperator, ExtractorSrc, ExtractorType} from "@/views/interface/consts";

export interface Interface {
    id: number;
    url: string;
    method: string;
    headers: Header[];
    params: Param[];

    body: string;
    bodyType: string;

    authorizationType: string
    basicAuth: BasicAuth;
    bearerToken: BearerToken;
    oAuth20: OAuth20;
    apiKey: ApiKey;

    extractors: Extractor[];
    checkpoints: Checkpoint[];
}

export interface Response {
    code: number
    headers: Header[];

    content: string;
    contentType: string;

    contentLang: string;
    contentCharset: string;
}

export interface Param {
    name: string;
    value: any;
    disabled: boolean;
}

export interface Header {
    name: string;
    value: any;
    disabled: boolean;
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
    authURL: string; // Authentication URL
    accessTokenURL: string; // Access Token URL
    clientID: string; // Client ID
    scope: string; // Scope
}
export interface ApiKey {
    username: string;
    value: string;
    transferMode: string;
}

export interface Extractor {
    id: number;
    src: string;
    type: string;
    expression: string;
    variable: string;
    interfaceId: number;
}
export interface Checkpoint {
    id: number;
    src: ExtractorSrc;
    type: ExtractorType;
    expression: string;

    operator: CheckpointOperator;
    value: any;
    interfaceId: number;
}