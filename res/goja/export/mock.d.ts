declare type Header = {
    name: string;
    value: string;
    disabled:    boolean;
    type: string;
}
declare type ExecCookie = {
    name: string;
    value: any;
    path: string;

    domain: string;
    expireTime: Date;
}
declare type MockResponseHeader = {
    name: string;
    value: any;
}

declare type Response = {
    statusCode: number;
    contentType: string;

    data: any;
    // used by adv mock
    headers: MockResponseHeader[];
}

declare global {
    const dt: {
        response: Response,
    }

    const log : (obj: any) => {}
}

export default {};