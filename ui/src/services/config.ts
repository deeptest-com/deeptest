import request from '@/utils/request';

const apiPath = 'configs';

export async function getServerConfig(): Promise<any> {
    return request({
        url: `/${apiPath}`,
        method: 'GET',
    })
}

/**
 * 根据 key 获取配置
 * */
export async function getConfigByKey(key): Promise<any> {
    return request({
        url: `/${apiPath}/getValue?key=${key}`,
        method: 'GET',
    })
}

