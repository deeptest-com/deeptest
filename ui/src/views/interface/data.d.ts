export interface Request {
    url: string;
    params: Param[];
}

export interface Param {
    name: string;
    value: any;
    disable: boolean
}

export interface Response {
    header: any;
    body: string;
}