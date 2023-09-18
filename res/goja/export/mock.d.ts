declare type Param = {
    name: string;
    value: string;
    type: string;
}
declare type Header = {
    name: string;
    value: string;
    type: string;
}
declare type Cookie = {
    name: string;
    value: any;
    path: string;

    domain: string;
    expireTime: Date;
}
declare type FormItem = {
    name: string;
    value: string;
    type: string;
}

declare type Request = {
    method: string;
    url: string;
    queryParams: Param[];
    pathParams: Param[];
    headers: Header[];
    cookies: Cookie[];

    body: string;
    formData:       FormItem[];
    bodyType: string;
};
declare type Response = {
    statusCode: number;

    data: any;
    contentType: string;
}

declare global {
    const dt: {
        request: Request,
        response: Response,

        getParam: (name) => string,
        getHeader: (name) => string,
        getCookie: (name) => Cookie,
    }

    const log : (obj: any) => {}
}

export default {};