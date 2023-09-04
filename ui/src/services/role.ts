import request from '@/utils/request';

export async function getAllSysRoles(): Promise<any> {
    return request({
        url: '/roles/all',
        method: 'get'
    });
}
