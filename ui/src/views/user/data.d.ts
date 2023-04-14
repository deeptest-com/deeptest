export interface User {
    id: number;
    name: string;
    username: string;
    email: string;
    intro:string
    avatar:string
    password:string
}

export interface QueryResult {
    list: User[];
    pagination: PaginationConfig;
}

export interface QueryParams {
    username: string,
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
