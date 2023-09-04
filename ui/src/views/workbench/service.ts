import request from '@/utils/request';
import {  QueryParams } from './data.d';

const apiPath = 'summary';

export async function query(params: QueryParams): Promise<any> {
    return request({
        url: `/${apiPath}/projectUserRanking/${params.cycle}/${params.projectId}`,
        method: 'get',
        params,
        
    });
}

