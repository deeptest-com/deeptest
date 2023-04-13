import {User, QueryParams} from "@/views/user/data";
import request from "@/utils/request";

const apiPath = 'users';

export async function query(params?: QueryParams): Promise<any> {
    return request({
        url: `/${apiPath}`,
        method: 'get',
        params,
    });
}

export async function save(params: Partial<User>): Promise<any> {
    return request({
        url: params.id? `/${apiPath}/${params.id}`:`/${apiPath}`,
        method: 'POST',
        data: params,
    });
}

export async function remove(id: number): Promise<any> {
    return request({
        url: `/${apiPath}/${id}`,
        method: 'delete',
    });
}

export async function detail(id: number): Promise<any> {
    return request({url: `/${apiPath}/${id}`});
}
