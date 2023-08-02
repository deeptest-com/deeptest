/**
 * 是否运行在客户端 Electron 容器中
 * */

const win:any = window?.process;
export const isElectronEnv = win?.versions?.electron;


/**
 * 可选的 Agent 服务地址
 * */
export const agentUrlOpts = [
    {
        label: '本地',
        value: 'local',
        url: 'http://127.0.0.1:8086/api/v1',
    },
    {
        label: '远程',
        value: 'remote',
        url: 'https://leyanapi.nancalcloud.com/agent/api/v1',
    },
];


/**
 * 获取当前的 Agent 的 URL
 *
 * */
export function getAgentUrl() {
    // 运行在客户端 Electron 容器中
    if (isElectronEnv) {
        const localCacheAgentVal = window.localStorage.getItem('dp-cache-agent-value') || 'local';
        const selectedAgent = agentUrlOpts.find((item) => {
            return item.value === localCacheAgentVal;
        });
        if (selectedAgent?.url) {
            return selectedAgent.url;
        }
        // 如果是浏览器环境，只能通过远程服务调取
    } else {
        return process.env.VUE_APP_API_AGENT;
    }
}


/**
 * 获取当前的 Agent 的 URL
 *
 * */
export function getAgentLabel() {
    // 运行在客户端 Electron 容器中
    if (isElectronEnv) {
        const localCacheAgentVal = window.localStorage.getItem('dp-cache-agent-value') || 'local';
        const selectedAgent = agentUrlOpts.find((item) => {
            return item.value === localCacheAgentVal;
        });
        if (selectedAgent?.label) {
            return selectedAgent?.label;
        }
        // 如果是浏览器环境，只能通过远程服务调取
    } else {
        return '本地'
    }
}
