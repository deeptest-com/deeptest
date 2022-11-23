import request from '@/utils/request';

const apiPath = 'users';

export async function inviteUser(data): Promise<any> {
    return request({
        url: `/${apiPath}/inviteUser`,
        method: 'POST',
        data,
    });
}

export async function changeInfo(field, value): Promise<any> {
    return request({
        url: `/${apiPath}/changeInfo`,
        method: 'POST',
        data: {field, value},
    });
}