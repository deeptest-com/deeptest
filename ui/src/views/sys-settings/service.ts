import request from '@/utils/request';

const apiPathJslib = 'jslib';

export async function listJslib(params): Promise<any> {
    return request({
        url: `/${apiPathJslib}`,
        method: 'get',
        params
    });
}

export async function getJslib(id): Promise<any> {
    const params = {id}
    return request({
        url: `/${apiPathJslib}`,
        method: 'get',
        params
    });
}

export async function saveJslib(data): Promise<any> {
    return request({
        url: `/${apiPathJslib}`,
        method: 'post',
        data
    });
}
export async function updateJsLibName(data): Promise<any> {
    return request({
        url: `/${apiPathJslib}/updateName`,
        method: 'put',
        data
    });
}

export async function deleteJslib(id): Promise<any> {
    return request({
        url: `/${apiPathJslib}/${id}`,
        method: 'delete',
    });
}

export async function disableJslib(id): Promise<any> {
    return request({
        url: `/${apiPathJslib}/${id}/disable`,
        method: 'put',
    });
}