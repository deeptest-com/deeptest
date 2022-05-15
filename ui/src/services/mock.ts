import request from '@/utils/request';
import {LoginParamsType} from "@/views/user/login/data";

export async function getOAuth2AccessToken(params): Promise<any> {
    return request({
        url: '/auth/getOAuth2AccessToken',
        method: 'post',
        params
    });
}

export async function useOAuth2AccessToken(token, tokenType): Promise<any> {
    const params = {token: token, tokenType: tokenType}

    return request({
        url: '/auth/useOAuth2AccessToken',
        method: 'post',
        params
    });
}
