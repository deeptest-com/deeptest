interface BasicSchemaInfo {
    name: string;
    serveId: number | string;
    tags: string;
    id?: number | string;
    content?: string;
    examples?: string;
    type?: string;
    description?: string;
}

export interface FormState {
    name: string;
    description: string;
    serveId?: string,
}

export interface DataItem {
    key: string;
    name: string;
    age: number;
    address: string;
}

export interface VarsChangeState {
    field: string;
    index: number;
    e: any; 
    action?: string;
}

export interface ParamsChangeState {
    field: string;
    index: number;
    e: any; 
    type: string;
    action?: string;
}

export interface SaveSchemaReqParams {
    schemaInfo: BasicSchemaInfo,
    action: string;
    serveId: string | number;
    name?: string;
}

export interface SchemaListReqParams {
    serveId: number | string;
    page?: number;
    pageSize?: number;
    name?: string;
}

export interface BasicSchemaParams {
    name: string; // 搜索关键字
    serveId?: number | string;
    id?: string;
    data?: string;
}

// a schema of filterFormComponent
export interface Schema {
    type: string;
    stateName?: string;
    options?: any[];
    placeholder?: string;
    action?: any;
    text?: string;
    valueType?: string;
    title?: string;
}
export interface SearchInfo {
    keyword: string;
    placeholder: string;
    action: any;
}