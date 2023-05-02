import request from '@/utils/request';

const apiEnvironment = 'environments'
const apiEnvironmentVar = `${apiEnvironment}/vars`

export async function listEnvironment(): Promise<any> {
    const params = {}

    return request({
        url: `/${apiEnvironment}`,
        method: 'GET',
        params,
    });
}
export async function getEnvironment(id: number, projectId: number): Promise<any> {
    const params = {projectId: projectId}
    return request({
        url: `/${apiEnvironment}/${id}`,
        method: 'GET',
        params
    });
}

export async function changeEnvironment(id, projectId): Promise<any> {
    const params = {id, projectId}

    return request({
        url: `/${apiEnvironment}/changeEnvironment`,
        method: 'POST',
        params,
    });
}
export async function saveEnvironment(data): Promise<any> {
    return request({
        url: `/${apiEnvironment}`,
        method: data.id ? 'PUT' : 'POST',
        data: data,
    });
}
export async function copyEnvironment(id): Promise<any> {
    const params = {id: id}
    return request({
        url: `/${apiEnvironment}/copyEnvironment`,
        method: 'POST',
        params,
    });
}
export async function removeEnvironment(id: number): Promise<any> {
    return request({
        url: `/${apiEnvironment}/${id}`,
        method: 'DELETE',
    });
}

// environment var
export async function saveEnvironmentVar(data): Promise<any> {
    return request({
        url: `/${apiEnvironmentVar}`,
        method: data.id ? 'PUT' : 'POST',
        data: data,
    });
}
export async function removeEnvironmentVar(id: number): Promise<any> {
    return request({
        url: `/${apiEnvironmentVar}/${id}`,
        method: 'DELETE',
    });
}
export async function clearEnvironmentVar(environmentId: number): Promise<any> {
    const params = {environmentId: environmentId}
    return request({
        url: `/${apiEnvironmentVar}/clear`,
        method: 'POST',
        params,
    });
}
