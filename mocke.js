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

declare type Response = {
    statusCode: string;
    statusContent: string;

    headers: Header[];
    cookies: ExecCookie[];

    data: any;
    contentType: string;

    contentLang: string;
    contentCharset: string;
    contentLength: number;
}

declare global {
    const dt: {
        response: Response,
    }

    const log : (obj: any) => {}
}

export default {};
