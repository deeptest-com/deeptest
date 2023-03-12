import request from '@/utils/request';
import {QueryParams} from "@/views/project/data";

const apiPath = 'plans';
const apiPathCategoryNodes = `${apiPath}/categories`;
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
export async function loadExecResult(planId): Promise<any> {
    const params = {planId}
    return request({
        url: `/${apiPathExec}/loadExecResult`,
        method: 'get',
        params,
    });
}

// category tree
export async function loadCategory(): Promise<any> {
    const params = {}
    return request({
        url: `/${apiPathCategoryNodes}/load`,
        method: 'get',
        params,
    });
}
export async function getCategory(id: number): Promise<any> {
    return request({url: `/${apiPathCategoryNodes}/${id}`});
}
export async function createCategory(data): Promise<any> {
    return request({
        url: `/${apiPathCategoryNodes}`,
        method: 'POST',
        data: data,
    });
}
export async function updateCategory(id: number, data: any): Promise<any> {
    return request({
        url: `/${apiPathCategoryNodes}/${id}`,
        method: 'PUT',
        data: data,
    });
}
export async function updateCategoryName(id: number, name: string): Promise<any> {
    const data = {id: id, name: name}

    return request({
        url: `/${apiPathCategoryNodes}/${id}/updateName`,
        method: 'PUT',
        data: data,
    });
}
export async function removeCategory(id: number): Promise<any> {
    return request({
        url: `/${apiPathCategoryNodes}/${id}`,
        method: 'delete',
    });
}
export async function moveCategory(data: any): Promise<any> {
    return request({
        url: `/${apiPathCategoryNodes}/move`,
        method: 'post',
        data: data,
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