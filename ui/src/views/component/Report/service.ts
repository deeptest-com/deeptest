import request from '@/utils/request';
import {QueryParams} from "./data";

const apiPath = 'scenarios/reports';

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

export async function remove(id: number): Promise<any> {
    return request({
        url: `/${apiPath}/${id}`,
        method: 'delete',
    });
}

export async function members(id: number): Promise<any> {
    return request({
        url: `/projects/members?id=${id}`,
        method: 'get',
    })
}