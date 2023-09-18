import {requestToStatic} from '@/utils/request';

export async function getClientVersion(): Promise<any> {
    return requestToStatic({
        url: `/LeyanAPI/version.json`,
        method: 'GET',
    });
}

