import request, {ResponseData} from '@/utils/request';




/**
 * 获取接口文档信息
 * */
export async function getDocs(data: any): Promise<any> {
    return request({
        url: `/document`,
        method: 'post',
        data: data
    });
}

/**
 * 获取接口文档版本列表，无分页
 * */
export async function getVersionList(data: any): Promise<any> {
    return request({
        url: `/document/version_list`,
        method: 'post',
        data: data
    });
}


/**
 * 发布接口文档
 * */
export async function publishDocument(data: any): Promise<any> {
    return request({
        url: `/document/publish`,
        method: 'post',
        data: data
    });
}


/**
 * 删除快照
 * */
export async function deleteDocumentVersion(data: any): Promise<any> {
    return request({
        url: `/document/delete`,
        method: 'delete',
        data: data
    });
}


/**
 * 更新文档版本名称
 * */
export async function updateDocumentVersion(data: any): Promise<any> {
    return request({
        url: `/document/update_version`,
        method: 'post',
        data: data
    });
}


/**
 * 获取分享文档的code
 * */
export async function shareDocs(data: any): Promise<any> {
    return request({
        url: `/document/share`,
        method: 'post',
        data: data
    });
}



/**
 * 通过获取分享的文档数据
 * */
export async function getShareContent(data: any): Promise<any> {
    return request({
        url: `/document/get_share_content?code=${data.code}`,
        method: 'get',
    });
}
