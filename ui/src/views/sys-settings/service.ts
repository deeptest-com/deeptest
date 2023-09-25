import request from '@/utils/request';

const apiPathJsLibs = 'jsLibs';

export async function listJsLib(): Promise<any> {
    return request({
        url: `/${apiPathJsLibs}`,
        method: 'get',
    });
}
