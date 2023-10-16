import pth from "path";
import path from "path";
import os from "os";
import {App} from "./constant";
export const DEBUG = process.env.NODE_ENV === 'development';
export const WORK_DIR = process.cwd()

export const portClient = 55111
export const portAgent = 56111
export const uuid = '1CF17A46-B136-4AEB-96B4-F21C8200EF5A@DEEPTEST.COM'

export const electronMsg = 'electronMsg'
export const electronMsgUsePort = 'electronMsgUsePort'
export const electronMsgReplay = 'electronMsgReplay'
export const electronMsgUpdate = 'electronMsgUpdate'
export const electronMsgDownloading = 'electronMsgDownloading'
export const electronMsgServerUrl = 'electronMsgServerUrl'

export const minimumSizeWidth = 1024
export const minimumSizeHeight = 640


export const WorkDir = path.join(os.homedir(), App);
export const ResDir = process.resourcesPath;
// Agent进程名称
export const agentProcessName = 'deeptest-agent'
