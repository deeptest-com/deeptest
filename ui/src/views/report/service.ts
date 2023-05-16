import request from '@/utils/request';
import {QueryParams} from "./data";

const apiPath = 'scenarios/reports';
const planApiPath = 'plans/reports';

export async function query(params?: QueryParams): Promise<any> {
    return request({
        url: `/${planApiPath}`,
        method: 'get',
        params,
    });
}
export async function get(id: number): Promise<any> {
    return request({url: `/${planApiPath}/${id}`});
}

export async function remove(id: number): Promise<any> {
    return request({
        url: `/${planApiPath}/${id}`,
        method: 'delete',
    });
}

export async function members(id: number): Promise<any> {
    return request({
        url: `/projects/members?id=${id}`,
        method: 'get',
    })
}