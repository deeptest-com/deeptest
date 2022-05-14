import {CheckpointOperator, CheckpointType, ExtractorSrc, ExtractorType} from "@/views/interface/consts";

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
    oauth20: OAuth20;
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
    key: string;
    value: string;
    transferMode: string;
}

export interface Extractor {
    id?: number;
    src: ExtractorSrc;
    type: ExtractorType;
    expression: string;

    boundaryStart: string;
    boundaryEnd: string;
    boundaryIndex: number;
    boundaryIncluded: boolean;

    variable: string;
    interfaceId?: number;
}

export interface Checkpoint {
    id?: number;
    type: CheckpointType;
    expression: string;

    operator: CheckpointOperator;
    value: any;
    interfaceId?: number;
}