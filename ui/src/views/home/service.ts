import request from '@/utils/request';
import {  QueryParams } from './data.d';

const apiPath = 'summary';

export async function query(params: QueryParams): Promise<any> {
    return request({
        url: `/${apiPath}/details`,
        method: 'get',
        params,
        
    });
}
export async function queryCardData(params): Promise<any> {
    return request({
        url: `/${apiPath}/card/${params.projectId}`,
        method: 'get',
        params,
        
    });
}
export async function queryPieData(params): Promise<any> {
    return request({
        url: `/${apiPath}/bugs/${params.projectId}`,
        method: 'get',
        params,
        
    });
}