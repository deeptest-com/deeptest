import request from '@/utils/request';
import { Project, QueryParams } from './data.d';

const apiPath = 'projects';

export async function query(params?: QueryParams): Promise<any> {
    return request({
        url: `/${apiPath}`,
        method: 'get',
        params,
    });
}
export async function queryMembers(params): Promise<any> {
    return request({
        url: `/${apiPath}/members`,
        method: 'get',
        params,
    });
}

export async function save(params: Partial<Project>): Promise<any> {
    return request({
        url: `/${apiPath}`,
        method: params.id? 'PUT': 'POST',
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

export async function removeMember(userId: number, projectId: number): Promise<any> {
    return request({
        url: `/${apiPath}/removeMember`,
        method: 'post',
        data: {userId: userId, projectId: projectId}
    });
}


export async function getUserList(name: any): Promise<any> {
    return request({
        url: `/users`,
        method: 'get',
    });
}

export async function getRoles(): Promise<any> {
    return request({
        url: `/${apiPath}/perms/rolesList`,
        method: 'get',
    });
}


export async function changeRole(params:object): Promise<any> {
    return request({
        url: `/${apiPath}/changeUserRole`,
        method: 'post',
        data:params
    });
}
