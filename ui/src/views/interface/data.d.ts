export interface Request {
    url: string;
    params: Param[];
    headers: Header[];
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

export interface Response {
    header: any;
    body: string;
}