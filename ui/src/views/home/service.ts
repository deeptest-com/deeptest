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

