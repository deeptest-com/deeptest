
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
    status: [] | null,
    createUser: [] | null,
    title: string | null,
    categoryId: string | null,
    tagNames:[]|[]
}


interface NewEndpointFormState {
    title: string;
    categoryId: string | null | number;
    description: string | undefined;
    curl:string;
}

export interface QueryCaseTreeParams {
    currentProjectId:number
    serveId:number
  }
