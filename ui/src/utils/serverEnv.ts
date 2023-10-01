/**
 * 是否运行在客户端 Electron 容器中
 * */
import {
    Cache_Key_Agent_Local_Port,
    Cache_Key_Agent_Url,
    Cache_Key_Agent_Value,
    Cache_Key_Server_Url
} from "@/utils/const";

const win: any = window?.process;
export const isElectronEnv = win?.versions?.electron;

/**
 * 获取当前ServerURL
 * */
export function getCachedServerUrl() {
    if(isElectronEnv) {
        const serverUrlFromCache = window.localStorage.getItem(Cache_Key_Server_Url)
        return serverUrlFromCache
    }

    return null;
}
