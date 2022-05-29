import request from '@/utils/request';
import {Interface, OAuth20} from "@/views/interface/data";
import {isInArray} from "@/utils/array";
import {CheckpointOperator} from "@/views/interface/consts";
import {QueryParams} from "@/views/project/data";

const apiPath = 'scenarios';
const apiNodePath = `${apiPath}/nodes`;

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

export async function save(data: any): Promise<any> {
    return request({
        url: `/${apiPath}`,
        method: data.id? 'PUT': 'POST',
        data: data,
    });
}

export async function load(scenarioId): Promise<any> {
    const params = {scenarioId}
    return request({
        url: `/${apiPath}/load`,
        method: 'get',
        params,
    });
}

export async function getNode(id: number): Promise<any> {
    return request({url: `/${apiPath}/node/${id}`});
}

export async function createNode(data): Promise<any> {
    return request({
        url: `/${apiPath}/node`,
        method: 'POST',
        data: data,
    });
}

export async function updateNode(id: number, params: any): Promise<any> {
    return request({
        url: `/${apiPath}/node/${id}`,
        method: 'PUT',
        data: params,
    });
}

export async function updateNodeName(id: number, name: string): Promise<any> {
    const data = {id: id, name: name}

    return request({
        url: `/${apiNodePath}/${id}/updateName`,
        method: 'PUT',
        data: data,
    });
}

export async function removeNode(id: number): Promise<any> {
    return request({
        url: `/${apiPath}/node/${id}`,
        method: 'delete',
    });
}

export async function moveNode(data: any): Promise<any> {
    return request({
        url: `/${apiPath}/node/move`,
        method: 'post',
        data: data,
    });
}
