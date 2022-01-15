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
