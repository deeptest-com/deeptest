import request from '@/utils/request';





/**
 * 接口列表
 * */
export async function getInterfaceList(data: any): Promise<any> {
    return request({
        url: `/endpoint/index`,
        method: 'post',
        data: data
    });
}

/**
 * 接口详情
 * */
export async function getInterfaceDetail(id: Number | String | any): Promise<any> {
    return request({
        url: `/endpoint/detail?id=${id}`,
        method: 'get',
    });
}

/**
 * 删除接口
 * */
export async function deleteInterface(id: Number): Promise<any> {
    return request({
        url: `/endpoint/delete?id=${id}`,
        method: 'delete',
    });
}


/**
 * 复制接口
 * */
export async function copyInterface(id: Number): Promise<any> {
    return request({
        url: `/endpoint/copy?id=${id}`,
        method: 'get',
    });
}


/**
 * 获取yaml展示
 * */
export async function getYaml(data: any): Promise<any> {
    return request({
        url: `/endpoint/yaml`,
        method: 'post',
        data: data
    });
}


/**
 * 接口过时
 * */
export async function expireInterface(id: Number): Promise<any> {
    return request({
        url: `/endpoint/expire?id=${id}`,
        method: 'put',
    });
}

/**
 * 保存接口
 * */
export async function saveInterface(data: any): Promise<any> {
    return request({
        url: `/endpoint/save`,
        method: 'post',
        data: data
    });
}

