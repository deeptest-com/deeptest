import request from '@/utils/request';

const apiPath = 'account';

export async function forgotPassword(usernameOrPassword): Promise<any> {
    return request({
        url: `/${apiPath}/forgotPassword`,
        method: 'POST',
        params: {usernameOrPassword: usernameOrPassword},
    });
}
