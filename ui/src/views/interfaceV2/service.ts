import request from '@/utils/request';
import {requestToAgent} from '@/utils/request';
import {
    ApiKey,
    BasicAuth, BearerToken,
    BodyFormDataItem,
    BodyFormUrlEncodedItem, Checkpoint, Extractor,
    Header,
    Interface,
    OAuth20,
    Param
} from "@/views/interface/data";
import {isInArray} from "@/utils/array";
import {UsedBy} from "@/utils/enum";

const apiPath = 'interfaces';
const apiImport = 'import';
const apiSpec = 'spec';
const apiInvocation = 'invocations';
const apiAuth = 'auth';
const apiEnvironment = 'environments'
const apiEnvironmentVar = `${apiEnvironment}/vars`
const apiShareVar = `${apiEnvironment}/shareVars`
const apiSnippets = 'snippets'

const apiExtractor = 'extractors'
const apiCheckpoint = 'checkpoints'

const apiParser = 'parser'


interface InterfaceListReqParams {
    "prjectId"?: number,
    "page"?: number,
    "pageSize"?: number,
    "status"?: number,
    "userId"?: number,
    "title"?: string
}

// todo liguwe 待整理
interface SaveInterfaceReqParams {
    project_id?: number,
    serve_id?: number,
}


/**
 * 接口列表
 * */
export async function getInterfaceList(data: InterfaceListReqParams): Promise<any> {
    return request({
        url: `/endpoint/index`,
        method: 'post',
        data: data
    });
}

/**
 * 接口详情
 * */
export async function getInterfaceDetail(id: Number | String | any): Promise<any> {
    return request({
        url: `/endpoint/detail?id=${id}`,
        method: 'get',
    });
}

/**
 * 删除接口
 * */
export async function deleteInterface(id: Number): Promise<any> {
    return request({
        url: `/endpoint/delete?id=${id}`,
        method: 'delete',
    });
}


/**
 * 接口过时
 * */
export async function expireInterface(id: Number): Promise<any> {
    return request({
        url: `/endpoint/expire?id=${id}`,
        method: 'put',
    });
}


/**
 * 保存接口
 * */
export async function saveInterface(data: SaveInterfaceReqParams): Promise<any> {
    return request({
        url: `/endpoint/save`,
        method: 'post',
        data: data
    });
}
