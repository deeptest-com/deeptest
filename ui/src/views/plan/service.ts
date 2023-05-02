import request from '@/utils/request';
import {QueryParams} from "@/views/project/data";

const apiPath = 'plans';
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
export async function loadExecResult(planId): Promise<any> {
    const params = {planId}
    return request({
        url: `/${apiPathExec}/loadExecResult`,
        method: 'get',
        params,
    });
}

// for scenario selection
export async function addScenarios(planId, scenariosIds): Promise<any> {
    return request({
        url: `/${apiPath}/${planId}/addScenarios`,
        method: 'post',
        data: scenariosIds,
    });
}

export async function removeScenarioFromPlan(planId, scenarioId): Promise<any> {
    return request({
        url: `/${apiPath}/${planId}/removeScenario`,
        method: 'post',
        params: {scenarioId},
    });
}



export async function listScenario(): Promise<any> {
    return request({
        url: `/${apiPathScenario}/listByProject`,
        method: 'get',
    });
}