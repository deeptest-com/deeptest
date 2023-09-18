import request from '@/utils/request';
import {
    BasicSchemaInfo,
    BasicSchemaParams,
    EnvDataItem,
    EnvReqParams,
    GlobalParamsReqData,
    SaveVersionParams,
    SchemaDetailReqParams,
    SchemaListReqParams,
    ServeListParams,
    ServeReqParams,
    VarsReqParams,
    VersionListReqParams,
    SwaggerSync
} from './data';

/**
 * 保存服务
 * */
export async function saveServe(data: ServeReqParams): Promise<any> {
    return request({
        url: `/serves/save`,
        method: 'post',
        data: data
    });
}

/**
 * 服务列表
 * */
export async function getServeList(data: ServeListParams): Promise<any> {
    return request({
        url: `/serves/index`,
        method: 'post',
        data: data,
    });
}

/**
 * 删除服务
 * */
export async function deleteServe(id: Number | String | undefined): Promise<any> {
    return request({
        url: `/serves/delete?id=${id}`,
        method: 'delete',
    });
}

/**
 * 禁用服务
 * */
export async function disableServe(id: Number | String | undefined): Promise<any> {
    return request({
        url: `/serves/expire?id=${id}`,
        method: 'put',
    });
}

/**
 * 复制服务
 * */
export async function copyServe(id: Number | String | undefined): Promise<any> {
    return request({
        url: `/serves/copy?id=${id}`,
        method: 'get',
    });
}


/**
 * 保存服务版本
 * */
export async function saveServeVersion(data: SaveVersionParams): Promise<any> {
    return request({
        url: `/serves/version/save`,
        method: 'post',
        data: data,
    });
}


/**
 * 服务版本列表
 * */
export async function getServeVersionList(data: VersionListReqParams): Promise<any> {
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
export async function serverList(data: any): Promise<any> {
    return request({
        url: `/serves/server/list`,
        method: 'post',
        data,
    });
}

/**
 * 切换服务环境
 * @param data
 * @returns
 */
export async function changeServe(data: any): Promise<any> {
    return request({
        url: '/serves/server/changeServer',
        method: 'post',
        data,
    });
}

/**
 * 获取授权列表
 * */
export async function getSecurityList(data: any): Promise<any> {
    return request({
        url: `/serves/security/list`,
        method: 'post',
        data,
    });
}

/**
 * 用户列表
 * */
export async function getUserList(): Promise<any> {
    return request({
        url: `/users`,
        method: 'get',
    });
}

/**
 *  保存组件
 * */
export async function saveSchema(data: BasicSchemaInfo): Promise<any> {
    return request({
        url: `/serves/schema/save`,
        method: 'post',
        data: data
    });
}

/**
 *  组件列表
 * */
export async function getSchemaList(data: SchemaListReqParams): Promise<any> {
    return request({
        url: `/serves/schema/list`,
        method: 'post',
        data: data
    });
}

/**
 *  组件详情
 * */
export async function getSchemaDetail(data: SchemaDetailReqParams): Promise<any> {
    return request({
        url: `/serves/schema/detail`,
        method: 'post',
        data: data
    });
}
/**
 * 删除服务
 * */
export async function deleteSchema(id: Number | String | undefined): Promise<any> {
    return request({
        url: `/serves/schema/delete?id=${id}`,
        method: 'delete',
    });
}

/**
 * 禁用schema
 * */
export async function disableSchema(id: Number | String | undefined): Promise<any> {
    return request({
        url: `/serves/schema/expire?id=${id}`,
        method: 'put',
    });
}

/**
 * 复制schema
 * */
export async function copySchema(id: Number | String | undefined): Promise<any> {
    return request({
        url: `/serves/schema/copy?id=${id}`,
        method: 'put',
    });
}


/**
 *  example转schema
 * */
export async function example2schema(data: BasicSchemaParams): Promise<any> {
    return request({
        url: `serves/schema/example2schema`,
        method: 'post',
        data: data
    });
}

/**
 *   schema转example
 * */
export async function schema2example(data: BasicSchemaParams): Promise<any> {
    return request({
        url: `/serves/schema/schema2example`,
        method: 'post',
        data: data
    });
}

/**
 *   schema转yaml
 * */
export async function schema2yaml(data: { data: string }): Promise<any> {
    return request({
        url: `/serves/schema/schema2yaml`,
        method: 'post',
        data: data
    });
}


/**
 *   保存环境
 * */
export async function saveEnv(data: EnvDataItem): Promise<any> {
    return request({
        url: `/environments/save`,
        method: 'post',
        data: data
    });
}

/**
 *   环境列表
 * */
export async function getEnvList(data: EnvReqParams): Promise<any> {
    return request({
        url: `/environments/list?projectId=${data.projectId}`,
        method: 'get',
    });
}

/**
 *   删除环境
 * */
export async function deleteEnv(data: EnvReqParams): Promise<any> {
    return request({
        url: `/environments/delete?id=${data.id}`,
        method: 'delete',
    });
}

/**
 *   复制环境信息
 * */
export async function copyEnv(data: EnvReqParams): Promise<any> {
    return request({
        url: `/environments/copy?id=${data.id}`,
        method: 'get',
    });
}

/**
 * 排序环境列表
 * @param data 环境的id列表
 * @returns promise
 */
export async function sortEnv(data: number[]): Promise<any> {
    return request({
        url: `/environments/order`,
        method: 'post',
        data
    })
}

/**
 *   保存全局变量
 * */
export async function saveGlobalVars(data: any[]): Promise<any> {
    return request({
        url: `/environments/vars/global`,
        method: 'post',
        data: data
    });
}


/**
 *   全局变量列表
 * */
export async function getGlobalVarsList(): Promise<any> {
    return request({
        url: `/environments/vars/global`,
        method: 'get',
    });
}


/**
 *   保存全局变量
 * */
export async function saveEnvironmentsParam(data: GlobalParamsReqData): Promise<any> {
    return request({
        url: `/environments/param`,
        method: 'post',
        data: data
    });
}

/**
 *   获取全局变量
 * */
export async function getEnvironmentsParamList(data: VarsReqParams): Promise<any> {
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
 * 删除授权
 * */
 export async function deleteSecurity(id: Number | String | any): Promise<any> {
    return request({
        url: `/serves/security/delete?id=${id}`,
        method: 'delete',
    });
}


/**
 * 数据池
 * */
export async function listDatapool(data: ServeListParams): Promise<any> {
    return request({
        url: `/datapools/index`,
        method: 'post',
        data: data,
    });
}
export async function getDatapool(id: number): Promise<any> {
    return request({
        url: `/datapools/${id}`,
        method: 'get',
    });
}
export async function saveDatapool(data: ServeReqParams): Promise<any> {
    return request({
        url: `/datapools/save`,
        method: 'post',
        data: data
    });
}
export async function deleteDatapool(id: Number | String | undefined): Promise<any> {
    return request({
        url: `/datapools/${id}`,
        method: 'delete',
    });
}
export async function disableDatapool(id: Number | String | undefined): Promise<any> {
    return request({
        url: `/datapools/${id}/disable`,
        method: 'put',
    });
}

export async function saveSwaggerSync(data: SwaggerSync): Promise<any> {
    return request({
        url: `/projectSettings/saveSwaggerSync`,
        method: 'post',
        data: data
    });
}
export async function getSwaggerSync(): Promise<any> {
    return request({
        url: `/projectSettings/swaggerSyncDetail`,
        method: 'get',
    });
}

export async function saveMock(data): Promise<any> {
    return request({
        url: `/projectSettings/saveMock`,
        method: 'post',
        data: data
    });
}
export async function getMock(): Promise<any> {
    return request({
        url: `/projectSettings/getMock`,
        method: 'get',
    });
}

export async function getVarsByEnv(envId: number): Promise<any> {
    return request({
        url: `/environments/envVars/byEnv?envId=${envId}`,
        method: 'get'
    })
}