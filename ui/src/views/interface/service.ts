import request from '@/utils/request';
import {QueryParams} from "@/views/project/data";

const apiPath = 'interfaces';
const apiPathExec = `${apiPath}/exec`;

const apiPathScenario = `scenarios`;
const apiPathServe = `serve`;

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

export async function listServe(): Promise<any> {
    return request({
        url: `/${apiPathServe}/listByProject`,
        method: 'get',
    });
}

export async function listScenario(serveId): Promise<any> {
    return request({
        url: `/${apiPathScenario}/listByServe`,
        method: 'get',
        params: {serveId},
    });
}