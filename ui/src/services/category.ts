import request from '@/utils/request';
const apiPath = 'categories';

export async function loadCategory(type): Promise<any> {
    const params = {type}
    return request({
        url: `/${apiPath}/load`,
        method: 'get',
        params,
    });
}
export async function getCategory(id: number): Promise<any> {
    return request({url: `/${apiPath}/${id}`});
}
export async function createCategory(data): Promise<any> {
    return request({
        url: `/${apiPath}`,
        method: 'POST',
        data: data,
    });
}
export async function updateCategory(id: number, params: any): Promise<any> {
    return request({
        url: `/${apiPath}/${id}`,
        method: 'PUT',
        data: params,
    });
}
export async function updateCategoryName(id: number, name: string): Promise<any> {
    const data = {id: id, name: name}

    return request({
        url: `/${apiPath}/${id}/updateName`,
        method: 'PUT',
        data: data,
    });
}
export async function removeCategory(id: number): Promise<any> {
    return request({
        url: `/${apiPath}/${id}`,
        method: 'delete',
    });
}
export async function moveCategory(data: any): Promise<any> {
    return request({
        url: `/${apiPath}/move`,
        method: 'post',
        data: data,
    });
}
