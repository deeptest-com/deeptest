import request from '@/utils/request';

const apiPathJslib = 'jslib';
const apiPathAgent = 'agents';

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
        url: `/${apiPathJslib}/${id}`,
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


export async function listAgent(params): Promise<any> {
    return request({
        url: `/${apiPathAgent}`,
        method: 'get',
        params
    });
}

export async function getAgent(id): Promise<any> {
    const params = {id}
    return request({
        url: `/${apiPathAgent}/${id}`,
        method: 'get',
        params
    });
}

export async function saveAgent(data): Promise<any> {
    return request({
        url: `/${apiPathAgent}`,
        method: 'post',
        data
    });
}
export async function updateAgentName(data): Promise<any> {
    return request({
        url: `/${apiPathAgent}/updateName`,
        method: 'put',
        data
    });
}

export async function deleteAgent(id): Promise<any> {
    return request({
        url: `/${apiPathAgent}/${id}`,
        method: 'delete',
    });
}

export async function disableAgent(id): Promise<any> {
    return request({
        url: `/${apiPathAgent}/${id}/disable`,
        method: 'put',
    });
}