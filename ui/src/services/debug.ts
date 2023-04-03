import request from '@/utils/request';
import {QueryParams} from "@/views/project/data";

const apiPath = 'debugs';

export async function loadData(data): Promise<any> {
    return request({
        url: `/${apiPath}/loadData`,
        method: 'POST',
        data,
    });
}
