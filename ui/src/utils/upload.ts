
import axios, { AxiosPromise, AxiosRequestConfig, AxiosResponse } from 'axios';
import router from '@/config/routes';
import i18n from "@/config/i18n";
import bus from "@/utils/eventBus";
import settings from '@/config/settings';
import { getToken } from '@/utils/localToken';
import { getCache } from '@/utils/localCache';
import {getServerUrl} from "@/utils/request";
import {addSepIfNeeded} from "@/utils/url";

const baseUrl = addSepIfNeeded(getServerUrl())

export async function uploadRequest(file) {
    const data = new FormData()

    data.append('file', file, file.name)
    console.log(data.get('file'))

    const config = {
        headers: {
            'Content-Type': 'multipart/form-data',
        },
    }
    // 添加jwt token
    const jwtToken = await getToken();
    if (jwtToken) {
        config.headers[settings.ajaxHeadersTokenKey] = 'Bearer ' + jwtToken;
    }

    const url = baseUrl + 'upload'

    const resp = await axios.post(url, data, config)
    console.log(resp.data.code)

    if (resp.data.code === 0) {
        return resp.data.data.path
    }

    return 'UPLOAD ERROR'
}
