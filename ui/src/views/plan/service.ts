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

export async function clonePlan(planId: number): Promise<any> {
    return request({
        url: `/${apiPath}/${planId}/clone`,
        method: 'post'
    })
}

export async function loadExecResult(planId: number): Promise<any> {
    const params = { planId };
    return request({
        url: `/${apiPathExec}/loadExecResult`,
        method: 'get',
        params,
    });
}

export async function listScenario(): Promise<any> {
    return request({
        url: `/${apiPathScenario}/listByProject`,
        method: 'get',
    });
}

export async function queryMembers(params): Promise<any> {
    return request({
        url: `/projects/members`,
        method: 'get',
        params,
    });
}

// 与计划关联的场景列表
interface PlanScenariosParam {
    planId: number;
    createUserId?: number | undefined | null;
    priority?: string | undefined | null;
    keywords?: string;
}
export async function getPlanScenarioList(params: PlanScenariosParam): Promise<any> {
    return request({
        url: `/${apiPath}/planScenariosList`,
        method: 'get',
        params
    })
}

// 关联测试场景
interface ScenariosIdsData {
    scenarioIds: number[]
}
export async function addScenarios(planId: number, payload: ScenariosIdsData): Promise<any> {
    return request({
        url: `/${apiPath}/${planId}/addScenarios`,
        method: 'post',
        data: payload,
    });
}

// 移除与计划关联的场景
export async function removeScenarios(planId: number, payload: ScenariosIdsData): Promise<any> {
    return request({
        url: `/${apiPath}/${planId}/removeScenarios`,
        method: 'post',
        data: payload,
    });
}
