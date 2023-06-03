import { Ref } from 'vue';

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

export interface VarDataItem {
    name: string;
    localValue: string | number;
    remoteValue: string | number;
    rightValue?: string;
}

export interface EnvReqParams {
    id?: string | number;
    activeEnvId?: string | number;
    projectId?: string | number;
}

export interface EnvDataItem {
    id: string | number;
    projectId: string | number;
    name: string;
    serveServers: any[];
    vars: any[];
}

export interface VarsReqParams {
    projectId: string | number;
}

export interface GlobalParamsReqData {
    header: any[];
    body: any[];
    query: any[];
    cookie: any[];
    projectId?: string | number;
}

export interface ServeListParams {
    "projectId"?: number | string,
    "page"?: number,
    "pageSize"?: number,
    "status"?: number,
    "userId"?: number,
    "title"?: string,
    "name"?: string;
}

export interface StoreDatapoolParams {
    projectId: string | number;
    formState?: any;
    id?: string | number | undefined;
    action?: string;
}
export interface DatapoolListParams {
    "projectId"?: number | string,
    "page"?: number,
    "pageSize"?: number,
    "status"?: number,
    "userId"?: number,
    "title"?: string,
    "name"?: string;
}
export interface DatapoolReqParams {
    projectId: string | number;
    name: string;
    description?: string;
    id?: number;
}
export interface DatapoolDetail {
    id: number;
    name: string;
    path: string,
    data: any,
    description: string;
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


export interface SchemaDetailReqParams {
    serveId: number | string;
    ref: string;
}

export interface BasicSchemaParams {
    name?: string; // 搜索关键字
    serveId?: number | string;
    id?: string | number | undefined;
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
    mode?: string;
    focusType?: string;
    required?: boolean;
    message?: string;
}
export interface SearchInfo {
    keyword: string;
    placeholder: string;
    action: any;
}
export interface SecurityListReqParams {
    serveId: number | string;
    page?: number;
    pageSize?: number;
    id?: string;

}

export interface StoreServeParams {
    projectId: string | number;
    formState?: any;
    id?: string | number | undefined;
    action?: string;
}

export interface ServeReqParams {
    projectId: string | number;
    name: string;
    description?: string;
    id?: number;
}

export interface ServeDetail {
    name: string;
    id: string;
    description: string;
}

export interface SaveVersionParams {
    serveId: string | number;
    value: string;
    createUser?: string;
    description?: string;
}

export interface EnvHookParams {
    isShowGlobalParams: Ref<boolean>;
    isShowGlobalVars: Ref<boolean>;
}

export interface EnvReturnData {
    isShowAddEnv: Ref<boolean>;
    isShowEnvDetail: Ref<boolean>;
    activeEnvDetail: any;
    getEnvsList: () => Promise<void>;
    showEnvDetail: (item: any, isAdd?: boolean) => void;
    addVar: () => void;
    addEnvData: () => void;
    deleteEnvData: (...arg: any) => void;
    copyEnvData: (...arg: any) => void;
    handleEnvChange: (type: string, field: string, index: number, e: any, action?: string) => void;
    handleEnvNameChange: (e: any) => void;
}

export interface GlobalVarsProps {
    isShowAddEnv: Ref<boolean>,
    isShowEnvDetail: Ref<boolean>,
    activeEnvDetail: Ref<any>,
    isShowGlobalParams: Ref<boolean>,
    isShowGlobalVars: Ref<boolean>,
    globalParamsActiveKey: Ref<string>
}

export interface VarsReturnData {
    showGlobalParams: () => void;
    showGlobalVars: () => void;
    addGlobalVar: () => void;
    addGlobalParams: (data: any) => void;
    handleSaveGlobalParams: () => void;
    handleSaveGlobalVars: () => void;
    handleGlobalVarsChange: (field: string, index: number, e: any, action?: string) => void;
    handleGlobalParamsChange: (type: string, field: string, index: number, e: any, action?: string) => void;
}

export interface VersionListReqParams {
    serveId: string | number;
    createUser: string;
    version?: string;
    page?: number;
    pageSize?: number;
}
