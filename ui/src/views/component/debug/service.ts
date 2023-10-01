import request from '@/utils/request';
import {requestToAgent} from '@/utils/request';
import {DebugInfo, Interface, OAuth20} from "./data";
import {isInArray} from "@/utils/array";
import {ConditionCategory, UsedBy} from "@/utils/enum";
import {getToken} from "@/utils/localToken";
import {getAgentUrl} from '@/utils/agentEnv';
const apiPath = 'debugs';
const apiPathInterface = `${apiPath}/interface`;
const apiPathInvoke = `${apiPath}/invoke`;

const apiAgentExec = 'exec';

const apiSpec = 'spec';
const apiAuth = 'auth';
const apiShareVar = `shareVars`
const apiSnippets = 'snippets'

const apiPreConditions = 'preConditions'
const apiPostConditions = 'postConditions'
const apiExtractor = 'extractors'
const apiCookie = 'cookies'
const apiCheckpoint = 'checkpoints'
const apiScript = 'scripts'

const apiParser = 'parser'

const apiResponseDefine = 'responseDefine'

// debug interface
export async function loadData(data): Promise<any> {
    return request({
        url: `/${apiPathInterface}/load`,
        method: 'post',
        data,
    });
}
export async function save(data: Interface): Promise<any> {
    return request({
        url: `/${apiPathInterface}/save`,
        method: 'post',
        data,
    });
}
export async function saveAsCase(data: Interface): Promise<any> {
    return request({
        url: `/${apiPathInterface}/saveAsCase`,
        method: 'post',
        data,
    });
}
export async function generateCases(data: Interface): Promise<any> {
    return request({
        url: `/${apiPathInterface}/generateCases`,
        method: 'post',
        data,
    });
}

// agent debug invoke
export async function call(data): Promise<any> {
   const agentUrl = getAgentUrl() || null
    console.log('===call===', agentUrl, data)

    // call agent api
    return requestToAgent({
        url: `/${apiAgentExec}/call`,
        agentUrl: agentUrl,
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
export async function getLastInvocationResp(params: DebugInfo): Promise<any> {
    return request({
        url: `/${apiPathInvoke}/getLastResp`,
        params
    });
}
export async function getInvocationResult(invokeId: number): Promise<any> {
    const params = {invokeId}
    return request({
        url: `/${apiPathInvoke}/getResult`,
        params
    });
}
export async function getInvocationLog(invokeId: number): Promise<any> {
    const params = {invokeId}
    return request({
        url: `/${apiPathInvoke}/getLog`,
        params
    });
}
export async function getInvocationAsInterface(id: number): Promise<any> {
    return request({url: `/${apiPathInvoke}/${id}`});
}

export async function removeInvocation(id: number): Promise<any> {
    return request({
        url: `/${apiPathInvoke}/${id}`,
        method: 'DELETE',
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
export async function clearShareVar(data: any): Promise<any> {
    return request({
        url: `/${apiShareVar}/clear`,
        method: 'POST',
        data,
    });
}

// helper
export function prepareDataForRequest(data: any) {
    if (data.headers) {
        data.headers = data.headers.filter((item) => {
            return !!item.name
        })
    }

    if (data.queryParams) {
        data.queryParams = data.queryParams.filter((item) => {
            return !!item.name
        })
    }
    if (data.pathParams) {
        data.pathParams = data.pathParams.filter((item) => {
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

// conditions
export async function getPreConditionScript(debugInterfaceId, endpointInterfaceId: number): Promise<any> {
    const params = {debugInterfaceId, endpointInterfaceId}

    return request({
        url: `/${apiPreConditions}/getScript`,
        method: 'GET',
        params,
    });
}
export async function createPreConditions(data): Promise<any> {
    return request({
        url: `/${apiPreConditions}`,
        method: data.id ? 'PUT' : 'POST',
        data: data,
    });
}
export async function disablePreConditions(id): Promise<any> {
    return request({
        url: `/${apiPreConditions}/${id}/disable`,
        method: 'POST',
    });
}
export async function removePreConditions(id): Promise<any> {
    return request({
        url: `/${apiPreConditions}/${id}`,
        method: 'DELETE',
    });
}
export async function movePreConditions(data): Promise<any> {
    return request({
        url: `/${apiPreConditions}/move`,
        method: 'POST',
        data: data,
    });
}

export async function listPostConditions(debugInterfaceId, endpointInterfaceId: number, category: ConditionCategory): Promise<any> {
    const params = {debugInterfaceId, endpointInterfaceId, category}

    return request({
        url: `/${apiPostConditions}`,
        method: 'GET',
        params,
    });
}
export async function createPostConditions(data): Promise<any> {
    return request({
        url: `/${apiPostConditions}`,
        method: data.id ? 'PUT' : 'POST',
        data: data,
    });
}
export async function disablePostConditions(id): Promise<any> {
    return request({
        url: `/${apiPostConditions}/${id}/disable`,
        method: 'POST',
    });
}
export async function removePostConditions(id): Promise<any> {
    return request({
        url: `/${apiPostConditions}/${id}`,
        method: 'DELETE',
    });
}
export async function movePostConditions(data): Promise<any> {
    return request({
        url: `/${apiPostConditions}/move`,
        method: 'POST',
        data: data,
    });
}

// extractor
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
export async function quickCreateExtractor(data): Promise<any> {
    return request({
        url: `/${apiExtractor}/quickCreate`,
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
export async function listExtractorVariable(data: any): Promise<any> {
    return request({
        url: `/${apiExtractor}/listExtractorVariableForCheckpoint`,
        method: 'POST',
        data,
    });
}

// cookie
export async function getCookie(id: number): Promise<any> {
    return request({
        url: `/${apiCookie}/${id}`,
        method: 'GET',
    });
}
export async function saveCookie(data): Promise<any> {
    return request({
        url: `/${apiCookie}`,
        method: data.id ? 'PUT' : 'POST',
        data: data,
    });
}
export async function quickCreateCookie(data): Promise<any> {
    return request({
        url: `/${apiCookie}/quickCreate`,
        method: 'POST',
        data: data,
    });
}
export async function removeCookie(id: number): Promise<any> {
    return request({
        url: `/${apiCookie}/${id}`,
        method: 'DELETE',
    });
}

// checkpoint
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

// script
export async function getScript(id: number): Promise<any> {
    return request({
        url: `/${apiScript}/${id}`,
        method: 'GET',
    });
}
export async function saveScript(data): Promise<any> {
    return request({
        url: `/${apiScript}`,
        method: data.id ? 'PUT' : 'POST',
        data: data,
    });
}
export async function removeScript(id: number): Promise<any> {
    return request({
        url: `/${apiScript}/${id}`,
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
export async function getJslibs(): Promise<any> {
    return request({
        url: `/${apiSnippets}/getJslibs`,
        method: 'GET',
    });
}

export async function saveResponseDefine(data): Promise<any> {
    return request({
        url: `/${apiResponseDefine}`,
        method:'PUT',
        data: data,
    });
}

