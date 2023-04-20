import request from '@/utils/request';
import {requestToAgent} from '@/utils/request';
import {DebugInfo, Interface, OAuth20} from "./data";
import {isInArray} from "@/utils/array";
import {UsedBy} from "@/utils/enum";
import {getToken} from "@/utils/localToken";

const apiPath = 'debugs';
const apiPathInterface = `${apiPath}/interface`;
const apiPathInvoke = `${apiPath}/invoke`;

const apiAgentExec = 'exec';

const apiSpec = 'spec';
const apiAuth = 'auth';
const apiShareVar = `shareVars`
const apiSnippets = 'snippets'

const apiExtractor = 'extractors'
const apiCheckpoint = 'checkpoints'

const apiParser = 'parser'

// debug interface
export async function loadData(data): Promise<any> {
    return request({
        url: `/${apiPathInterface}/load`,
        method: 'post',
        data,
    });
}
export async function save(interf: Interface): Promise<any> {
    return request({
        url: `/${apiPathInterface}/save`,
        method: 'post',
        data: interf,
    });
}

// agent debug invoke
export async function call(data): Promise<any> {
    // call agent api
    return requestToAgent({
        url: `/${apiAgentExec}/call`,
        method: 'POST',
        data,
    });
}

// debug invoke
export async function listInvocation(params: DebugInfo): Promise<any> {
    return request({
        url: `/${apiPathInvoke}`,
        method: 'GET',
        params,
    });
}
export async function getInvocationAsInterface(id: number): Promise<any> {
    return request({url: `/${apiPathInvoke}/${id}`});
}
export async function getLastInvocationResp(params: any): Promise<any> {
    return request({
        url: `/${apiPathInvoke}/getLastResp`,
        params
    });
}
export async function removeInvocation(id: number): Promise<any> {
    return request({
        url: `/${apiPathInvoke}/${id}`,
        method: 'DELETE',
    });
}

export async function parseSpecInLocalAgent(data, targetId): Promise<any> {
    data.targetId = targetId
    data.serverUrl = process.env.VUE_APP_API_SERVER // used by agent to request server
    data.token = await getToken()

    return requestToAgent({
        url: `/${apiSpec}/parseSpec`,
        method: 'POST',
        params: {targetId: targetId},
        data: data,
    });
}

// auth
export async function genOAuth2AccessToken(oauth: OAuth20): Promise<any> {
    return request({
        url: `/${apiAuth}/oauth2Authorization`,
        method: 'post',
        data: oauth,
    });
}
export async function listOAuth2Token(projectId): Promise<any> {
    const params = {projectId}

    return request({
        url: `/${apiAuth}/listOAuth2Token`,
        method: 'get',
        params,
    });
}
export async function removeOAuth2Token(id): Promise<any> {
    const params = {id}

    return request({
        url: `/${apiAuth}/removeToken`,
        method: 'get',
        params,
    });
}

// share var
export async function listShareVar(data: any, usedBy: UsedBy): Promise<any> {
    return request({
        url: `/${apiShareVar}/list`,
        method: 'POST',
        data,
    });
}
export async function removeShareVar(id: number): Promise<any> {
    return request({
        url: `/${apiShareVar}/${id}`,
        method: 'DELETE',
    });
}
export async function clearShareVar(endpointOrProcessorId: number, usedBy?: string): Promise<any> {
    const params = {endpointOrProcessorId, usedBy}
    return request({
        url: `/${apiShareVar}/clear`,
        method: 'POST',
        params,
    });
}

// helper
export function prepareDataForRequest(data: any) {
    if (data.headers) {
        data.headers = data.headers.filter((item) => {
            return !!item.name
        })
    }

    if (data.params) {
        data.params = data.params.filter((item) => {
            return !!item.name
        })
    }

    if (data.bodyFormData) {
        data.bodyFormData = data.bodyFormData.filter((item) => {
            return !!item.name
        })
    }
    if (data.bodyFormUrlencoded) {
        data.bodyFormUrlencoded = data.bodyFormUrlencoded.filter((item) => {
            return !!item.name
        })
    }

    data.body = data.body.replaceAll('\n', '').replaceAll(' ', '')

    return data
}

export function getCodeLangByType(type) {
    type = type.split('/')[1]

    if (isInArray(type, ['json', 'xml', 'html'])) {
        return type
    } else {
        return 'plaintext'
    }
}

export const getEnumSelectItems = (enumDef) => {
    const arr : any[] = []

    for (const item in enumDef) {
        arr.push({label: enumDef[item], value: item})
    }

    return arr
}

// extractor
export async function listExtractor(interfaceId: number, usedBy: UsedBy): Promise<any> {
    const params = {interfaceId, usedBy}

    return request({
        url: `/${apiExtractor}`,
        method: 'GET',
        params,
    });
}
export async function getExtractor(id: number): Promise<any> {
    return request({
        url: `/${apiExtractor}/${id}`,
        method: 'GET',
    });
}
export async function saveExtractor(data): Promise<any> {
    return request({
        url: `/${apiExtractor}`,
        method: data.id ? 'PUT' : 'POST',
        data: data,
    });
}
export async function createExtractorOrUpdateResult(data): Promise<any> {
    return request({
        url: `/${apiExtractor}/createOrUpdateResult`,
        method: 'POST',
        data: data,
    });
}

export async function removeExtractor(id: number): Promise<any> {
    return request({
        url: `/${apiExtractor}/${id}`,
        method: 'DELETE',
    });
}
export async function listExtractorVariable(interfaceId: number): Promise<any> {
    const params = {interfaceId}
    return request({
        url: `/${apiExtractor}/listExtractorVariableForCheckpoint`,
        method: 'GET',
        params,
    });
}

// checkpoint
export async function listCheckpoint(interfaceId: number, usedBy: UsedBy): Promise<any> {
    const params = {interfaceId, usedBy}

    return request({
        url: `/${apiCheckpoint}`,
        method: 'GET',
        params,
    });
}
export async function getCheckpoint(id: number): Promise<any> {
    return request({
        url: `/${apiCheckpoint}/${id}`,
        method: 'GET',
    });
}
export async function saveCheckpoint(data): Promise<any> {
    return request({
        url: `/${apiCheckpoint}`,
        method: data.id ? 'PUT' : 'POST',
        data: data,
    });
}
export async function removeCheckpoint(id: number): Promise<any> {
    return request({
        url: `/${apiCheckpoint}/${id}`,
        method: 'DELETE',
    });
}

export async function parseHtml(data): Promise<any> {
    return request({
        url: `/${apiParser}/parseHtml`,
        method: 'POST',
        data
    });
}
export async function parseXml(data): Promise<any> {
    return request({
        url: `/${apiParser}/parseXml`,
        method: 'POST',
        data
    });
}
export async function parseJson(data): Promise<any> {
    return request({
        url: `/${apiParser}/parseJson`,
        method: 'POST',
        data
    });
}
export async function parseText(data): Promise<any> {
    return request({
        url: `/${apiParser}/parseText`,
        method: 'POST',
        data
    });
}
export async function testExpr(data): Promise<any> {
    return request({
        url: `/${apiParser}/testExpr`,
        method: 'POST',
        data
    });
}

export async function getSnippet(name): Promise<any> {
    const params = {name: name}

    return request({
        url: `/${apiSnippets}`,
        method: 'GET',
        params
    });
}

export function getContextMenuStyle(e) {
    console.log('getContextMenuStyle', e.clientY)

    const style = {
        left: e.clientX + 'px',
        top: (e.clientY - 12 > 6 ? e.clientY - 12 : 6)  + 'px',
    }

    return style
}