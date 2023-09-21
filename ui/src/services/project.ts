import request from '@/utils/request';

const apiPath = 'projects';

export async function getByUser(currProjectId: number): Promise<any> {
    const params = {currProjectId: currProjectId}

    return request({
        url: `/${apiPath}/getByUser`,
        method: 'GET',
        params,
    });
}

export async function changeProject(projectId: number): Promise<any> {
    const data = {id: projectId}

    return request({
        url: `/${apiPath}/changeProject`,
        method: 'POST',
        data,
    });
}

export async function getPermissionMenuList(payload): Promise<any> {
    return request({
        url: `/${apiPath}/menus/userMenuList`,
        method: 'GET',
        params: payload,
    })
}

export async function queryMembers(params): Promise<any> {
    return request({
        url: `/${apiPath}/members`,
        method: 'get',
        params,
    });
}

export async function checkProjectAndUser(params) {
    return request({
        url: `/${apiPath}/checkProjectAndUser`,
        method: 'get',
        params,
    })
}
