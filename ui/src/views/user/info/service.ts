import request from '@/utils/request';

const apiPath = 'users';

export async function inviteUser(data, projectId): Promise<any> {
    data.projectId = projectId

    return request({
        url: `/${apiPath}/invite`,
        method: 'POST',
        data,
    });
}
