import request from '@/utils/request';

const apiDatapool = 'datapools'

export async function listDatapool(): Promise<any> {
    const params = {}

    return request({
        url: `/${apiDatapool}`,
        method: 'GET',
        params,
    });
}
export async function getDatapool(id: number, projectId: number): Promise<any> {
    const params = {projectId: projectId}
    return request({
        url: `/${apiDatapool}/${id}`,
        method: 'GET',
        params
    });
}

export async function saveDatapool(data): Promise<any> {
    return request({
        url: `/${apiDatapool}`,
        method: data.id ? 'PUT' : 'POST',
        data: data,
    });
}

export async function removeDatapool(id: number): Promise<any> {
    return request({
        url: `/${apiDatapool}/${id}`,
        method: 'DELETE',
    });
}
