import request from '@/utils/request';
import { Script, QueryParams } from './data.d';
import {dataURItoBlob} from "@/utils/form";

const apiPath = 'scripts';

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

export async function create(params: Omit<Script, 'id'>): Promise<any> {
    return request({
        url: `/${apiPath}`,
        method: 'POST',
        data: params,
    });
}

export async function update(id: number, params: Omit<Script, 'id'>): Promise<any> {
    return request({
        url: `/${apiPath}/${id}`,
        method: 'PUT',
        data: params,
    });
}

export async function remove(id: number): Promise<any> {
    return request({
        url: `/${apiPath}/${id}`,
        method: 'delete',
    });
}

export async function createStep(params:any, dataURI: any): Promise<any> {
    const data = new FormData()

    data.append("file", dataURItoBlob(dataURI), 'step.jpg')

    Object.keys(params).forEach((k): void => {
        data.append(k, params[k])
    })

    console.log('===', data.get("file"))

    return request({
        url: `/steps`,
        method: 'POST',
        data: data,
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
        }
    });
}
