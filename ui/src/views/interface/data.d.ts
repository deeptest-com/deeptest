export interface Request {
    url: string;
    params: Param[];
}

export interface Param {
    name: string;
    value: any;
    disabled: boolean
}

export interface Response {
    header: any;
    body: string;
}