
import axios, { AxiosRequestConfig } from 'axios';
import settings from '@/config/settings';
import { getToken } from '@/utils/localToken';
import {getUrls} from "@/utils/request";
import {addSepIfNeeded} from "@/utils/url";

const urls = getUrls()
const baseUrl = addSepIfNeeded(urls.serverUrl)

export async function uploadRequest(file: any, params?: any) {
    const data = new FormData()

    data.append('file', file, file.name)
    console.log(data.get('file'))

    const config:AxiosRequestConfig = {
        headers: {
            'Content-Type': 'multipart/form-data',
        },
        params: params,
    }
    // 添加jwt token
    const jwtToken = await getToken();
    if (jwtToken) {
        config.headers[settings.ajaxHeadersTokenKey] = 'Bearer ' + jwtToken;
    }

    const url = baseUrl + 'upload'

    const resp = await axios.post(url, data, config)
    console.log(resp.data.code)

    const ret = {} as any

    if (resp.data.code === 0) {
        ret.path = resp.data.data.path
        ret.data = resp.data.data.data
    }

    return ret
}
