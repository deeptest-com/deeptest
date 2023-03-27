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