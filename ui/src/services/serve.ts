import request from '@/utils/request';

const apiPath = `serves`;

export async function listServe(): Promise<any> {
    return request({
        url: `/${apiPath}/listByProject`,
        method: 'get',
    });
}

export async function changeServe(id: number): Promise<any> {
    const data = {id}

    return request({
        url: `/${apiPath}/changeServe`,
        method: 'POST',
        data,
    });
}