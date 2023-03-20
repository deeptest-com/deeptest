
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
    keywords:  string,
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
