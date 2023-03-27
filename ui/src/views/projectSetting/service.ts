import request from '@/utils/request';
import {requestToAgent} from '@/utils/request';
import {
    ApiKey,
    BasicAuth, BearerToken,
    BodyFormDataItem,
    BodyFormUrlEncodedItem, Checkpoint, Extractor,
    Header,
    Interface,
    OAuth20,
    Param
} from "@/views/interface1/data";
import {isInArray} from "@/utils/array";
import {UsedBy} from "@/utils/enum";


const apiPath = 'interfaces';
const apiImport = 'import';
const apiSpec = 'spec';
const apiInvocation = 'invocations';
const apiAuth = 'auth';
const apiEnvironment = 'environments'
const apiEnvironmentVar = `${apiEnvironment}/vars`
const apiShareVar = `${apiEnvironment}/shareVars`
const apiSnippets = 'snippets'

const apiExtractor = 'extractors'
const apiCheckpoint = 'checkpoints'

const apiParser = 'parser'


interface InterfaceListReqParams {
    "prjectId"?: number,
    "page"?: number,
    "pageSize"?: number,
    "status"?: number,
    "userId"?: number,
    "title"?: string
}

// todo liguwe 待整理
interface SaveInterfaceReqParams {
    // project_id?: number,
    serveId?: number,
    title?: string,
    path?: string
}


/**
 * 保存服务
 * */
export async function saveServe(data: any): Promise<any> {
    return request({
        url: `/serve/save`,
        method: 'post',
        data: data
    });
}

/**
 * 服务列表
 * */
export async function getServeList(data: any): Promise<any> {
    return request({
        url: `/serve/index`,
        method: 'post',
        data: data,
    });
}

/**
 * 删除服务
 * */
export async function deleteServe(id: Number): Promise<any> {
    return request({
        url: `/serve/delete?id=${id}`,
        method: 'delete',
    });
}

/**
 * 禁用服务
 * */
export async function disableServe(id: any): Promise<any> {
    return request({
        url: `/serve/expire?id=${id}`,
        method: 'put',
    });
}

/**
 * 复制服务
 * */
export async function copyServe(id: any): Promise<any> {
    return request({
        url: `/serve/copy?id=${id}`,
        method: 'get',
    });
}


/**
 * 保存服务版本
 * */
export async function saveServeVersion(data: any): Promise<any> {
    return request({
        url: `/serve/version/save`,
        method: 'post',
        data: data,
    });
}


/**
 * 服务版本列表
 * */
export async function getServeVersionList(data: any): Promise<any> {
    return request({
        url: `/serve/version/list`,
        method: 'post',
        data: data
    });
}

/**
 * 服务版本禁用
 * */
export async function disableServeVersions(id: any): Promise<any> {
    return request({
        url: `/serve/version/expire?id=${id}`,
        method: 'put',
    });
}

/**
 * 服务版本删除
 * */
export async function deleteServeVersion(id: any): Promise<any> {
    return request({
        url: `/serve/version/delete?id=${id}`,
        method: 'delete',
    });
}


/**
 * 服务环境列表
 * */
export async function serverList(data: any): Promise<any> {
    return request({
        url: `/serves/server/list?serveId=${data.serveId}`,
        method: 'get',
    });
}

/**
 * 用户列表
 * */
export async function getUserList(name: any): Promise<any> {
    return request({
        url: `/users`,
        method: 'get',
    });
}

/**
 *  保存组件
 * */
export async function saveSchema(data: any): Promise<any> {
    return request({
        url: `/serve/schema/save`,
        method: 'post',
        data: data
    });
}

/**
 *  组件列表
 * */
export async function getSchemaList(data: any): Promise<any> {
    return request({
        url: `/serve/schema/list`,
        method: 'post',
        data: data
    });
}

/**
 * 删除服务
 * */
export async function deleteSchema(id: Number): Promise<any> {
    return request({
        url: `/serve/schema/delete?id=${id}`,
        method: 'delete',
    });
}

/**
 * 禁用schema
 * */
export async function disableSchema(id: any): Promise<any> {
    return request({
        url: `/serve/schema/expire?id=${id}`,
        method: 'put',
    });
}

/**
 * 复制服务
 * */
export async function copySchema(id: any): Promise<any> {
    return request({
        url: `/serve/schema/copy?id=${id}`,
        method: 'put',
    });
}


/**
 *  example转schema
 * */
export async function example2schema(data: any): Promise<any> {
    return request({
        url: `serve/schema/example2schema`,
        method: 'post',
        data: data
    });
}

/**
 *   schema转example
 * */
export async function schema2example(data: any): Promise<any> {
    return request({
        url: `/serve/schema/schema2example`,
        method: 'post',
        data: data
    });
}

/**
 *   schema转yaml
 * */
export async function schema2yaml(data: any): Promise<any> {
    return request({
        url: `/serve/schema/schema2yaml`,
        method: 'post',
        data: data
    });
}


/**
 *   保存环境
 * */
export async function saveEnv(data: any): Promise<any> {
    return request({
        url: `/environments/save`,
        method: 'post',
        data: data
    });
}

/**
 *   环境列表
 * */
export async function getEnvList(data: any): Promise<any> {
    return request({
        url: `/environments/list?projectId=${data.projectId}`,
        method: 'get',
    });
}

/**
 *   删除环境
 * */
export async function deleteEnv(data: any): Promise<any> {
    return request({
        url: `/environments/delete?id=${data.id}`,
        method: 'delete',
    });
}

/**
 *   复制环境信息
 * */
export async function copyEnv(data: any): Promise<any> {
    return request({
        url: `/environments/copy?id=${data.id}`,
        method: 'get',
    });
}


/**
 *   保存全局变量
 * */
export async function saveGlobalVars(data: any): Promise<any> {
    return request({
        url: `/environments/vars/global`,
        method: 'post',
        data: data
    });
}


/**
 *   全局变量列表
 * */
export async function getGlobalVarsList(data: any): Promise<any> {
    return request({
        url: `/environments/vars/global?projectId=${data.projectId}`,
        method: 'get',
    });
}


/**
 *   保存全局变量
 * */
export async function saveEnvironmentsParam(data: any): Promise<any> {
    return request({
        url: `/environments/param`,
        method: 'post',
        data: data
    });
}

/**
 *   获取全局变量
 * */
export async function getEnvironmentsParamList(data: any): Promise<any> {
    return request({
        url: `/environments/param?projectId=${data.projectId}`,
        method: 'get',
    });
}
