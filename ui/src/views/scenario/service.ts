import request from '@/utils/request';
import {QueryParams} from "./data";

import {
    ProcessorCookie, ProcessorData,
    ProcessorExtractor,
    ProcessorLogic,
    ProcessorLoop, ProcessorGroup, ProcessorTimer, ProcessorPrint,
    ProcessorCategory,
    ProcessorVariable, ProcessorAssertion, RequestBodyType, UsedBy, ProcessorAction
} from "@/utils/enum";
import {Interface} from "@/views/component/debug/data";

const apiPath = 'scenarios';
const apiPathNodes = `${apiPath}/nodes`;

const apiPathProcessors = `${apiPath}/processors`;
const apiPathExec = `${apiPath}/exec`;

const apiInvocation = `processors/invocations`;
const apiPathInterface = `processors/interfaces`

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
export async function getLastInvocationResp(id: number): Promise<any> {
    const params = {id : id}
    return request({
        url: `/${apiInvocation}/getLastResp`,
        params
    });
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

export async function loadExecResult(scenarioId): Promise<any> {
    const params = {scenarioId}
    return request({
        url: `/${apiPathExec}/loadExecResult`,
        method: 'get',
        params,
    });
}

// scenario tree
export async function loadScenario(scenarioId): Promise<any> {
    const params = {scenarioId}
    return request({
        url: `/${apiPath}/load`,
        method: 'get',
        params,
    });
}
export async function getNode(id: number): Promise<any> {
    return request({url: `/${apiPathProcessors}/${id}`});
}
export async function createNode(data): Promise<any> {
    return request({
        url: `/${apiPathNodes}`,
        method: 'POST',
        data: data,
    });
}
export async function updateNode(id: number, params: any): Promise<any> {
    return request({
        url: `/${apiPathNodes}/${id}`,
        method: 'PUT',
        data: params,
    });
}
export async function updateNodeName(id: number, name: string): Promise<any> {
    const data = {id: id, name: name}

    return request({
        url: `/${apiPathNodes}/${id}/updateName`,
        method: 'PUT',
        data: data,
    });
}
export async function removeNode(id: number): Promise<any> {
    return request({
        url: `/${apiPathNodes}/${id}`,
        method: 'delete',
    });
}
export async function moveNode(data: any): Promise<any> {
    return request({
        url: `/${apiPathNodes}/move`,
        method: 'post',
        data: data,
    });
}
export async function addInterfacesFromDefine(data): Promise<any> {
    return request({
        url: `/${apiPathNodes}/addInterfacesFromDefine`,
        method: 'POST',
        data: data,
    });
}
export async function addInterfacesFromTest(data): Promise<any> {
    return request({
        url: `/${apiPathNodes}/addInterfacesFromTest`,
        method: 'POST',
        data: data,
    });
}
export async function addProcessor(data): Promise<any> {
    return request({
        url: `/${apiPathNodes}/addProcessor`,
        method: 'POST',
        data: data,
    });
}
export async function saveInterface(interf: any): Promise<any> {
    return request({
        url: `/${apiPathInterface}/saveInterface`,
        method: 'post',
        data: interf,
    });
}
export async function saveProcessor(data: any): Promise<any> {
    return request({
        url: `/${apiPathProcessors}/${data.processorCategory}/save`,
        method: 'PUT',
        data: data,
    });
}
export async function saveProcessorName(data: any): Promise<any> {
    return request({
        url: `/${apiPathProcessors}/updateName`,
        method: 'PUT',
        data: data,
    });
}
export async function saveProcessorInfo(data: any): Promise<any> {
    return request({
        url: `/${apiPathProcessors}/saveProcessorInfo`,
        method: 'PUT',
        data: data,
    });
}

/**
 * 执行场景历史
 * */
export async function getScenariosReports(data: any): Promise<any> {
    return request({
        url: `/scenarios/reports`,
        method: 'POST',
        data: data,
    });
}


/**
 * 执行场景执行的详情
 * */
export async function getScenariosReportsDetail(data: any): Promise<any> {
    return request({
        url: `/scenarios/reports/${data.id}`,
        method: 'GET',
    });
}

/**
 * 生成报告
 * */
export async function genReport(data: any): Promise<any> {
    return request({
        url: `/scenarios/reports/${data.id}`,
        method: 'PUT',
    });
}

/**
 * 添加 关联测试计划
 * */
export async function addPlans(payload: any): Promise<any> {
    return request({
        url: `/scenarios/${payload.id}/addPlans`,
        method: 'POST',
        data: payload.data,
    });
}

/**
 * 移除 关联测试计划
 * */
export async function removePlans(payload: any): Promise<any> {
    return request({
        url: `/scenarios/${payload.id}/removePlans`,
        method: 'POST',
        data: payload.data,
    });
}

/**
 * 计划列表
 * */
export async function getPlans(payload: any): Promise<any> {
    return request({
        url: `/scenarios/${payload.id}/plans?currProjectId=${payload.currProjectId}`,
        method: 'POST',
        data: payload.data
    });
}

/**
 * 更新优先级
 * */
export async function updatePriority(payload: any): Promise<any> {
    return request({
        url: `/scenarios/${payload.id}/updatePriority?priority=${payload.priority}`,
        method: 'PUT',
        // data: payload.data
    });
}

/**
 * 更新场景状态
 * */
export async function updateStatus(payload: any): Promise<any> {
    return request({
        url: `/scenarios/${payload.id}/updateStatus?status=${payload.status}`,
        method: 'PUT',
        // data: payload.data
    });
}

// interface
export function getRequestBodyTypes() {
    return getEnumSelectItems(RequestBodyType)
}

export function getProcessorCategories() {
    return getEnumSelectItems(ProcessorCategory)
}

export function getProcessorTypeNames() {
    return {
        // ...getEnumMap(ProcessorThread),
        ...getEnumMap(ProcessorGroup),
        ...getEnumMap(ProcessorTimer),
        ...getEnumMap(ProcessorLogic),

        ...getEnumMap(ProcessorLoop),
        ...getEnumMap(ProcessorExtractor),
        ...getEnumMap(ProcessorVariable),
        ...getEnumMap(ProcessorCookie),
        ...getEnumMap(ProcessorData),
    }
}

export const getEnumMap = (enumDef) => {
    const ret = {}

    for (const item in enumDef) {
        ret[item] = enumDef[item]
    }

    return ret
}

export function getProcessorTypeMap() {
    return {
        // processor_thread: getEnumSelectItems(ProcessorThread),
        processor_group: getEnumSelectItems(ProcessorGroup),
        processor_timer: getEnumSelectItems(ProcessorTimer),
        processor_print: getEnumSelectItems(ProcessorPrint),
        processor_logic: getEnumSelectItems(ProcessorLogic),

        processor_loop: getEnumSelectItems(ProcessorLoop),
        processor_extractor: getEnumSelectItems(ProcessorExtractor),
        processor_variable: getEnumSelectItems(ProcessorVariable),
        processor_assertion: getEnumSelectItems(ProcessorAssertion),
        processor_cookie: getEnumSelectItems(ProcessorCookie),
        processor_data: getEnumSelectItems(ProcessorData),
    }
}


export const isRoot = (type) => {
    return type === 'processor_root'
}
export const isProcessor = (type) => {
    return type !==  'processor_interface' && type !== 'processor_root'
}
export const isInterface = (type) => {
    return type ===  'processor_interface'
}

export const getEnumSelectItems = (enumDef) => {
    const arr : any[] = []

    for (const item in enumDef) {
        arr.push({label: enumDef[item], value: item})
    }

    return arr
}

export async function saveDebugData(interf: Interface): Promise<any> {
    return request({
        url: `/scenarios/interface/saveDebugData`,
        method: 'post',
        data: interf,
    });
}

export async function syncDebugData(scenarioProcessorId: number): Promise<any> {
    const params = {scenarioProcessorId}

    return request({
        url: `/scenarios/interface/resetDebugData`,
        method: 'post',
        params,
    });
}

export function getMenu(entityCategory: ProcessorCategory): ProcessorAction[] {
    const ret:ProcessorAction[] = []

    if (isInterface(entityCategory)) {
        ret.push(
            ProcessorAction.ActionEdit,
            ProcessorAction.ActionRemove,
            ProcessorAction.ActionInInterface,)

    } else if (isRoot(entityCategory)) {
        ret.push(
            ProcessorAction.ActionAddProcessor,
            ProcessorAction.ActionImportInterface,)

    } else { // no-root dir
        ret.push(
            ProcessorAction.ActionEdit,
            ProcessorAction.ActionRemove,
            ProcessorAction.ActionAddProcessor,
            ProcessorAction.ActionImportInterface,)
    }

    return ret
}
