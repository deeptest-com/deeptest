import request from '@/utils/request';
import {QueryParams} from "@/types/data";

const apiPath = 'testInterfaces';

export async function query(data: any): Promise<any> {
    return request({
        url: `/${apiPath}/load`,
        method: 'post',
        data,
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

export async function clone(id: number): Promise<any> {
    return request({
        url: `/${apiPath}/${id}/clone`,
        method: 'post'
    })
}
