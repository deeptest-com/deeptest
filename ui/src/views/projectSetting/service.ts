import request from '@/utils/request';

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
        url: `/serves/save`,
        method: 'post',
        data: data
    });
}

/**
 * 服务列表
 * */
export async function getServeList(data: any): Promise<any> {
    return request({
        url: `/serves/index`,
        method: 'post',
        data: data,
    });
}

/**
 * 删除服务
 * */
export async function deleteServe(id: Number): Promise<any> {
    return request({
        url: `/serves/delete?id=${id}`,
        method: 'delete',
    });
}

/**
 * 禁用服务
 * */
export async function disableServe(id: any): Promise<any> {
    return request({
        url: `/serves/expire?id=${id}`,
        method: 'put',
    });
}

/**
 * 复制服务
 * */
export async function copyServe(id: any): Promise<any> {
    return request({
        url: `/serves/copy?id=${id}`,
        method: 'get',
    });
}


/**
 * 保存服务版本
 * */
export async function saveServeVersion(data: any): Promise<any> {
    return request({
        url: `/serves/version/save`,
        method: 'post',
        data: data,
    });
}


/**
 * 服务版本列表
 * */
export async function getServeVersionList(data): Promise<any> {
    return request({
        url: `/serves/version/list`,
        method: 'post',
        data: data
    });
}

/**
 * 服务版本禁用
 * */
export async function disableServeVersions(id: Number | String | any): Promise<any> {
    return request({
        url: `/serves/version/expire?id=${id}`,
        method: 'put',
    });
}

/**
 * 服务版本删除
 * */
export async function deleteServeVersion(id: Number | String | any): Promise<any> {
    return request({
        url: `/serves/version/delete?id=${id}`,
        method: 'delete',
    });
}


/**
 * 服务环境列表
 * */
export async function serverList(data: SaveInterfaceReqParams): Promise<any> {
    return request({
        url: `/serves/server/list`,
        method: 'post',
        data: data
    });
}

/**
 * 用户列表
 * */
export async function getUserList(name): Promise<any> {
    return request({
        url: `/users`,
        method: 'get',
    });
}

/**
 *  保存组件
 * */
export async function saveSchema(data): Promise<any> {
    return request({
        url: `/serves/schema/save`,
        method: 'post',
        data: data
    });
}

/**
 *  组件列表
 * */
export async function getSchemaList(data): Promise<any> {
    return request({
        url: `/serves/schema/list`,
        method: 'post',
        data: data
    });
}

/**
 * 删除服务
 * */
export async function deleteSchema(id: Number | String | any): Promise<any> {
    return request({
        url: `/serves/schema/delete?id=${id}`,
        method: 'delete',
    });
}

/**
 * 禁用schema
 * */
export async function disableSchema(id: any | Number | String): Promise<any> {
    return request({
        url: `/serves/schema/expire?id=${id}`,
        method: 'put',
    });
}

/**
 * 复制服务
 * */
export async function copySchema(id: any | Number | String): Promise<any> {
    return request({
        url: `/serves/schema/copy?id=${id}`,
        method: 'put',
    });
}


/**
 *  example转schema
 * */
export async function example2schema(data): Promise<any> {
    return request({
        url: `serves/schema/example2schema`,
        method: 'post',
        data: data
    });
}

/**
 *   schema转example
 * */
export async function schema2example(data): Promise<any> {
    return request({
        url: `/serves/schema/schema2example`,
        method: 'post',
        data: data
    });
}

/**
 *   schema转yaml
 * */
export async function schema2yaml(data): Promise<any> {
    return request({
        url: `/serves/schema/schema2yaml`,
        method: 'post',
        data: data
    });
}


/**
 *   保存环境
 * */
export async function saveEnv(data): Promise<any> {
    return request({
        url: `/environments/save`,
        method: 'post',
        data: data
    });
}

/**
 *   环境列表
 * */
export async function getEnvList(data): Promise<any> {
    return request({
        url: `/environments/list?projectId=${data.projectId}`,
        method: 'get',
    });
}

/**
 *   删除环境
 * */
export async function deleteEnv(data): Promise<any> {
    return request({
        url: `/environments/delete?id=${data.id}`,
        method: 'delete',
    });
}

/**
 *   复制环境信息
 * */
export async function copyEnv(data): Promise<any> {
    return request({
        url: `/environments/copy?id=${data.id}`,
        method: 'get',
    });
}


/**
 *   保存全局变量
 * */
export async function saveGlobalVars(data): Promise<any> {
    return request({
        url: `/environments/vars/global`,
        method: 'post',
        data: data
    });
}


/**
 *   全局变量列表
 * */
export async function getGlobalVarsList(data): Promise<any> {
    return request({
        url: `/environments/vars/global?projectId=${data.projectId}`,
        method: 'get',
    });
}


/**
 *   保存全局变量
 * */
export async function saveEnvironmentsParam(data): Promise<any> {
    return request({
        url: `/environments/param`,
        method: 'post',
        data: data
    });
}

/**
 *   获取全局变量
 * */
export async function getEnvironmentsParamList(data): Promise<any> {
    return request({
        url: `/environments/param?projectId=${data.projectId}`,
        method: 'get',
    });
}

/**
 *   保存授权信息
 * */
 export async function saveSecurityInfo(data): Promise<any> {
    return request({
        url: `/serves/security/save`,
        method: 'post',
        data: data
    });
}

/**
 *   授权列表
 * */
 export async function getSecurityList(data): Promise<any> {
    return request({
        url: `/serves/security/list`,
        method: 'post',
        data: data
    });
}

/**
 * 删除授权
 * */
 export async function deleteSecurity(id: Number | String | any): Promise<any> {
    return request({
        url: `/serves/security/delete?id=${id}`,
        method: 'delete',
    });
}