import request from '@/utils/request';
import {QueryParams} from "@/views/project/data";

const apiPath = 'interfaces';
const apiPathExec = `${apiPath}/exec`;

const apiPathScenario = `scenarios`;

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
        method: data.id? 'PUT': 'POST',
        data: data,
    });
}
export async function remove(id: number): Promise<any> {
    return request({
        url: `/${apiPath}/${id}`,
        method: 'delete',
    });
}
export async function loadExecResult(interfaceId): Promise<any> {
    const params = {interfaceId}
    return request({
        url: `/${apiPathExec}/loadExecResult`,
        method: 'get',
        params,
    });
}

// for scenario selection
export async function addScenarios(interfaceId, scenariosIds): Promise<any> {
    return request({
        url: `/${apiPath}/${interfaceId}/addScenarios`,
        method: 'post',
        data: scenariosIds,
    });
}

export async function removeScenarioFromInterface(interfaceId, scenarioId): Promise<any> {
    return request({
        url: `/${apiPath}/${interfaceId}/removeScenario`,
        method: 'post',
        params: {scenarioId},
    });
}

export async function listScenario(serveId): Promise<any> {
    return request({
        url: `/${apiPathScenario}/listByServe`,
        method: 'get',
        params: {serveId},
    });
}


/*************************************************
 * :::: add
 ************************************************/


/**
 * 接口列表
 * */
export async function getInterfaceList(data: any): Promise<any> {
    return request({
        url: `/endpoint/index`,
        method: 'post',
        data: data
    });
}

/**
 * 接口详情
 * */
export async function getInterfaceDetail(id: Number | String | any): Promise<any> {
    return request({
        url: `/endpoint/detail?id=${id}`,
        method: 'get',
    });
}

/**
 * 删除接口
 * */
export async function deleteInterface(id: Number): Promise<any> {
    return request({
        url: `/endpoint/delete?id=${id}`,
        method: 'delete',
    });
}


/**
 * 复制接口
 * */
export async function copyInterface(id: Number): Promise<any> {
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
 * 接口过时
 * */
export async function expireInterface(id: Number): Promise<any> {
    return request({
        url: `/endpoint/expire?id=${id}`,
        method: 'put',
    });
}

/**
 * 保存接口
 * */
export async function saveInterface(data: any): Promise<any> {
    return request({
        url: `/endpoint/save`,
        method: 'post',
        data: data
    });
}

