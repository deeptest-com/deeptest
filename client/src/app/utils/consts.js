import pth from "path";
import path from "path";
import os from "os";

export const DEBUG = process.env.NODE_ENV === 'development';
export const WORK_DIR = process.cwd()

export const portClient = 55111
export const portAgent = 56111
export const uuid = '1CF17A46-B136-4AEB-96B4-F21C8200EF5A@DEEPTEST.COM'

export const electronMsg = 'electronMsg'
export const electronMsgReplay = 'electronMsgReplay'
export const electronMsgUpdate = 'electronMsgUpdate'
export const electronMsgDownloading = 'electronMsgDownloading'

export const minimumSizeWidth = 1024
export const minimumSizeHeight = 640

export const App = 'LeyanAPI';
// export const App = 'deeptest';
export const WorkDir = path.join(os.homedir(), App);
export const ResDir = process.resourcesPath;
// export const downloadUrl = 'http://127.0.0.1:8085/upload/';
// export const downloadUrl = 'http://111.231.16.35:8085/upload/';
export const downloadUrl = 'http://192.168.159.163:8080/';
// Agent进程名称
export const agentProcessName = 'deeptest-agent'
