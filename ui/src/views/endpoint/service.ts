import request, {ResponseData} from '@/utils/request';
import {QueryParams} from "@/types/data";
import {Interface} from "@/views/component/debug/data";

const apiPath = 'endpoints';
const apiPathInterface = 'endpoints/interfaces';
const apiPathCase = 'endpoints/cases';

export async function query(params?: QueryParams): Promise<any> {
    return request({
        url: `/${apiPath}`,
        method: 'get',
        params,
    });
}

export async function get(id: number): Promise<any> {
    return request({url: `/${apiPath}/${id}`});
}

export async function getDetail(id: number): Promise<any> {
    const params = {
        detail: true,
    }
    return request({url: `/${apiPath}/${id}`, params});
}

export async function save(data: any): Promise<any> {
    return request({
        url: `/${apiPath}`,
        method: data.id ? 'PUT' : 'POST',
        data: data,
    });
}

export async function remove(id: number): Promise<any> {
    return request({
        url: `/${apiPath}/${id}`,
        method: 'delete',
    });
}

/**
 * 接口列表
 * */
export async function getEndpointList(data: any): Promise<any> {
    return request({
        url: `/endpoint/index`,
        method: 'post',
        data: data
    });
}

export async function listEndpointInterface(data: any, pagination: any) {
    const resp: ResponseData = (await request({
        url: `/${apiPathInterface}/listForSelection`,
        method: 'post',
        data: {
            ...pagination,
            ...data,
        }
    })) as any;

    if (resp.code != 0) return;

    const ret = {
        list: resp.data.result || [],
        total: resp.data.total || 0,
    }

    return ret
}

/**
 * 接口详情
 * */
export async function getEndpointDetail(id: Number | String | any): Promise<any> {
    return request({
        url: `/endpoint/detail?id=${id}`,
        method: 'get',
    });
}

/**
 * 删除接口
 * */
export async function deleteEndpoint(id: Number): Promise<any> {
    return request({
        url: `/endpoint/delete?id=${id}`,
        method: 'delete',
    });
}


/**
 * 复制接口
 * */
export async function copyEndpoint(id: Number): Promise<any> {
    return request({
        url: `/endpoint/copy?id=${id}`,
        method: 'get',
    });
}


/**
 * 获取yaml展示
 * */
export async function getYaml(data: any): Promise<any> {
    return request({
        url: `/endpoint/yaml`,
        method: 'post',
        data: data
    });
}


/**
 * 接口过期
 * */
export async function expireEndpoint(id: Number): Promise<any> {
    return request({
        url: `/endpoint/expire?id=${id}`,
        method: 'put',
    });
}

/**
 * 保存接口
 * */
export async function saveEndpoint(data: any): Promise<any> {
    return request({
        url: `/endpoint/save`,
        method: 'post',
        data: data
    });
}


/**
 * 更新接口状态
 * */
export async function updateStatus(data: any): Promise<any> {
    return request({
        url: `/endpoint/updateStatus?id=${data.id}&status=${data.status}`,
        method: 'put',
    });
}


/**
 * 获取接口文档信息
 * */
export async function getDocs(data: any): Promise<any> {
    return request({
        url: `/document`,
        method: 'post',
        data: data
    });
}


/**
 * 导入接口 - 上传文件
 * */
export async function upload(data: any): Promise<any> {
    return request({
        url: `/upload`,
        method: 'post',
        headers: {
            'Content-Type': 'multipart/form-data'
        },
        data: data.file
    });
}

/**
 * 导入接口 - 导入接口数据
 * */
export async function importEndpointData(data: any): Promise<any> {
    return request({
        url: `/endpoints/interfaces/importEndpointData`,
        method: 'post',
        data: {
            ...data,
        }
    });
}

export async function listEndpointCase(endpointId: number): Promise<any> {
    const params = {endpointId}
    return request({
        url: `/${apiPathCase}/list`,
        method: 'GET',
        params
    })
}
export async function getEndpointCase(id: Number | String | any): Promise<any> {
    return request({
        url: `/${apiPathCase}/${id}`,
        method: 'GET',
    });
}
export async function saveEndpointCase(data: any): Promise<any> {
    return request({
        url: `/${apiPathCase}/${data.id?data.id:0}`,
        method: 'post',
        data
    });
}
export async function saveEndpointCaseDebugData(data: any): Promise<any> {
    return request({
        url: `/${apiPathCase}/saveDebugData`,
        method: 'post',
        data,
    });
}
export async function updateEndpointCaseName(data): Promise<any> {
    return request({
        url: `/${apiPathCase}/${data.id}`,
        method: 'put',
        data
    });
}
export async function removeEndpointCase(data): Promise<any> {
    return request({
        url: `/${apiPathCase}/${data.id}`,
        method: 'delete',
        data
    });
}

/**
 * 批量修改接口字段的值
 * @param data
 */
export async function batchUpdateField(data: any): Promise<any> {
    return request({
        url: `/endpoint/batchUpdateField`,
        method: 'post',
        data: data
    });
}

/**
 * 获取tags
 * @param data
 */
export async function tagList(): Promise<any> {
    return request({
        url: `/endpoint/tags`,
        method: 'get'
    });
}

/**
 * 更新标签
 * @param data
 */
export async function updateTag(data: any): Promise<any> {
    return request({
        url: `/endpoint/updateTag`,
        method: 'put',
        data: data
    });
}