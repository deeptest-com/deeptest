
export interface Endpoints {
  "responseBodies": any[],
  "requestBody": any,
  params: any[],
  headers: any[],
  cookies: any[],
  method: string |undefined,
}

export interface Endpoint {
    id: number;
    name: string;
    desc: string;
    status: number;
    endpoints: Array<Endpoints>,
    pathParams: any[],
}

export interface QueryResult {
    list: Endpoint[];
    pagination: PaginationConfig;
}

export interface QueryParams {
    keywords: string,
    enabled: string,
    page: number,
    pageSize: number,
}

export interface PaginationConfig {
    total: number;
    current: number;
    pageSize: number;
    showSizeChanger: boolean;
    showQuickJumper: boolean;
}



interface EndpointListReqParams {
    "prjectId"?: number,
    "page"?: number,
    "pageSize"?: number,
    "status"?: number,
    "userId"?: number,
    "title"?: string
}

interface SaveEndpointReqParams {
    project_id?: number,
    serveId?: number,
    title?: string,
    path?: string
}


export interface Endpoint {
    id: number;
    name: string;
    desc: string;
}

export interface QueryResult {
    list: Endpoint[];
    pagination: PaginationConfig;
}

export interface QueryParams {
    keywords: string,
    enabled: string,
    page: number,
    pageSize: number,
}

export interface PaginationConfig {
    total: number;
    current: number;
    pageSize: number;
    showSizeChanger: boolean;
    showQuickJumper: boolean;
}


export interface filterFormState {
    status: string | null,
    createUser: string | null,
    title: string | null,
}


interface NewEndpointFormState {
    title: string;
    categoryId: string | null;
    description: string | undefined;
}
