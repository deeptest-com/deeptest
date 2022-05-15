import request from '@/utils/request';
import {LoginParamsType} from "@/views/user/login/data";

export async function getOAuth2AccessToken(params): Promise<any> {
    return request({
        url: '/auth/getOAuth2AccessToken',
        method: 'post',
        params
    });
}
