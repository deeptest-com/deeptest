/**
 * 是否运行在客户端 Electron 容器中
 * */

const win: any = window?.process;
export const isElectronEnv = win?.versions?.electron;


/**
 * 可选的 Agent 服务地址
 * */
export const agentUrlOpts = isElectronEnv ? [
    {
        label: '本地调试',
        value: 'local',
        url: 'http://127.0.0.1:8086/api/v1',
        desc:'可调用本地接口（localhost、127.0.0.1），所有请求通过本地Agent转发',
    },
    {
        label: '内网调试',
        value: 'test',
        url: 'https://leyanapi-test.nancalcloud.com/agent/api/v1',
        desc:'可调用内网接口，所有请求通过公司内网Agent转发',
    },
    {
        label: '外网调试',
        value: 'remote',
        url: 'https://leyanapi.nancalcloud.com/agent/api/v1',
        desc:'可调用外网接口，所有请求通过华为云外网Agent转发',
    },
] : [
    {
        label: '内网调试',
        value: 'test',
        url: 'https://leyanapi-test.nancalcloud.com/agent/api/v1',
        desc:'可调用内网接口，所有请求通过公司内网Agent转发',
    },
    {
        label: '外网调试',
        value: 'remote',
        url: 'https://leyanapi.nancalcloud.com/agent/api/v1',
        desc:'可调用外网接口，所有请求通过华为云外网Agent转发',
    },
];


/**
 * 获取当前的 Agent 的 URL
 *
 * */
export function getAgentUrl() {
    const localCacheAgentVal = window.localStorage.getItem('dp-cache-agent-value') || (isElectronEnv ? 'local' : 'test');
    const selectedAgent = agentUrlOpts.find((item) => {
        return item.value === localCacheAgentVal;
    });
    if (selectedAgent?.url) {
        return selectedAgent.url;
    }
    return process.env.VUE_APP_API_AGENT;
}


/**
 * 获取当前的 Agent 的 URL
 *
 * */
export function getAgentLabel() {
    const localCacheAgentVal = window.localStorage.getItem('dp-cache-agent-value') || (isElectronEnv ? 'local' : 'test');
    const selectedAgent = agentUrlOpts.find((item) => {
        return item.value === localCacheAgentVal;
    });
    if (selectedAgent?.label) {
        return selectedAgent?.label;
    }
    return '本地调试'
}
