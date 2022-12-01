/**
 * 自定义 request 网络请求工具,基于axios
 * @author LiQingSong
 */
import axios, { AxiosPromise, AxiosRequestConfig, AxiosResponse } from 'axios';
import router from '@/config/routes';
import i18n from "@/config/i18n";
import bus from "@/utils/eventBus";
import settings from '@/config/settings';
import { getToken } from '@/utils/localToken';
import { getCache } from '@/utils/localCache';

export interface ResponseData {
    code: number;
    data?: any;
    msg?: string;
    token?: string;
}
export interface ResultErr {
    httpCode: number;
    resultCode: number;
    resultMsg: string;
}

/**
 * 异常处理程序
 */
const errorHandler = (axiosResponse: AxiosResponse) => {
    console.log(axiosResponse)

    if (!axiosResponse) axiosResponse = {status: 500} as AxiosResponse

    if (axiosResponse.status === 200) {
        const result ={
            httpCode: axiosResponse.status,
            resultCode: axiosResponse.data.code,
            resultMsg: axiosResponse.data.msg
        } as ResultErr

        bus.emit(settings.eventNotify, result)

        const { config, data } = axiosResponse;
        const { url, baseURL } = config;
        const { code, msg } = data

        const reqUrl = (url + '').split("?")[0].replace(baseURL + '', '');
        const noNeedLogin = settings.ajaxResponseNoVerifyUrl.includes(reqUrl);
        if (code === 401 && !noNeedLogin) {
            router.replace('/user/login');
        }

    } else {
        bus.emit(settings.eventNotify, {
            httpCode: axiosResponse.status
        })
    }

    return Promise.reject({})
}

/**
 * 配置request请求时的默认参数
 */
const request = axios.create({
    baseURL: process.env.VUE_APP_API_SERVER,
    withCredentials: true, // 跨域请求时发送cookie
    timeout: 0
});

const requestAgent = axios.create({
    baseURL: process.env.VUE_APP_API_AGENT
});

// 全局设置 - post请求头
// request.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8';

/**
 * 请求拦截器
 */
const requestInterceptors = async (config: AxiosRequestConfig & { cType?: boolean, baseURL?: string }) => {
    // 如果设置了cType 说明是自定义 添加 Content-Type类型 为自定义post 做铺垫
    if (config['cType']) {
        config.headers['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8';
    }

    // 添加jwt token
    const jwtToken = await getToken();
    if (jwtToken) {
        config.headers[settings.ajaxHeadersTokenKey] = 'Bearer ' + jwtToken;
    }

    // 修改示例请求指向mock地址
    const url = config.url || '';
    if (url.indexOf('/home') > -1 || url.indexOf('/pages') > -1) {
        config.baseURL = '/api';
    }

    // 加随机数清除缓存
    config.params = { ...config.params, ts: Date.now() };
    if (!config.params.currProjectId) {
        const projectId = await getCache(settings.currProjectId);
        config.params = { ...config.params, currProjectId: projectId, lang: i18n.global.locale.value };
    }

    console.log('=== request ===', config)
    return config;
}
request.interceptors.request.use(
    requestInterceptors,
    /* error=> {} */ // 已在 export default catch
);
requestAgent.interceptors.request.use(
    requestInterceptors,
    /* error=> {} */ // 已在 export default catch
);

/**
 * 响应拦截器
 */
const responseInterceptors = async (axiosResponse: AxiosResponse) => {
    console.log('=== response ===', axiosResponse.config.url, axiosResponse)

    const res: ResponseData = axiosResponse.data;
    const { code, token } = res;

    // 自定义状态码验证
    if (code !== 0) {
        return Promise.reject(axiosResponse);
    }

    return axiosResponse;
}
request.interceptors.response.use(
    responseInterceptors,
    /* error => {} */ // 已在 export default catch
);
requestAgent.interceptors.response.use(
    responseInterceptors,
    /* error => {} */ // 已在 export default catch
);

/**
 * ajax 导出
 *
 * Method: get
 *     Request Headers
 *         无 - Content-Type
 *     Query String Parameters
 *         name: name
 *         age: age
 *
 * Method: post
 *     Request Headers
 *         Content-Type:application/json;charset=UTF-8
 *     Request Payload
 *         { name: name, age: age }
 *         Custom config parameters
 *             { cType: true }  Mandatory Settings Content-Type:application/json;charset=UTF-8
 * ......
 */
export default function(config: AxiosRequestConfig): AxiosPromise<any> {
    return request(config).then((response: AxiosResponse) => response.data).catch(error => errorHandler(error));
}

export function requestToAgent (config: AxiosRequestConfig): AxiosPromise<any> {
    return requestAgent(config).then((response: AxiosResponse) => response.data).catch(error => errorHandler(error));
}
