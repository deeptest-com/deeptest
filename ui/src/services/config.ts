import request from '@/utils/request';

const apiPath = 'configs';

export async function getServerConfig(): Promise<any> {
    return request({
        url: `/${apiPath}`,
        method: 'GET',
    })
}
