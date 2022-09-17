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
