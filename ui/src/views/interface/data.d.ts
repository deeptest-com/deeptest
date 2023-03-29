
export interface Interfaces {
  "responseBodies": any[],
  "requestBody": any,
  params: any[],
  headers: any[],
  cookies: any[],
  method: string |undefined,
}

export interface Interface {
    id: number;
    name: string;
    desc: string;
    interfaces: Array<Interfaces>,
    pathParams: any[],
}


export interface QueryResult {
    list: Interface[];
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



interface InterfaceListReqParams {
    "prjectId"?: number,
    "page"?: number,
    "pageSize"?: number,
    "status"?: number,
    "userId"?: number,
    "title"?: string
}

interface SaveInterfaceReqParams {
    project_id?: number,
    serveId?: number,
    title?: string,
    path?: string
}


export interface Interface {
    id: number;
    name: string;
    desc: string;
}

export interface QueryResult {
    list: Interface[];
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


interface NewInterfaceFormState {
    title: string;
    categoryId: string | null;
    description: string | undefined;
}
