import request from '@/utils/request';
import {LoginParamsType} from "@/views/user/login/data";

export async function getOAuth2AccessToken(params): Promise<any> {
    return request({
        url: '/auth/getOAuth2AccessToken',
        method: 'post',
        params
    });
}

export async function useOAuth2AccessToken(name, token, tokenType, interfaceId, projectId): Promise<any> {
    const params = {name, token, tokenType, interfaceId, projectId}

    return request({
        url: '/auth/useOAuth2AccessToken',
        method: 'post',
        params
    });
}
