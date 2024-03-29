import {app} from 'electron';
import {
    changeVersion, checkFileExist, checkMd5, getAppUrl,
    getCurrVersion,
    getDownloadPath,
    getRemoteVersion,
    getResPath, mkdir,
    restart
} from "./comm";
import {agentProcessName, electronMsgDownloading, electronMsgUpdate, WorkDir} from "./consts";
import path from "path";
import {execSync} from 'child_process';
import {IS_WINDOWS_OS} from "../utils/env";
import fse from 'fs-extra'
import {logErr, logInfo} from "./log";



/**
 * 版本号大小比较
 * @param {string} v1 版本号1
 * @param {string} v2 版本号2
 * @return {number} 1: v1 > v2, 0: v1 = v2, -1: v1 < v2
 * */
function compareVersion(v1, v2) {
    const [x1, y1, z1] = v1.split('.');
    const [x2, y2, z2] = v2.split('.');

    if (x1 !== x2) {
        return parseInt(x1) > parseInt(x2) ? 1 : -1;
    }
    if (y1 !== y2) {
        return parseInt(y1) > parseInt(y2) ? 1 : -1;
    }
    if (z1 !== z2) {
        return parseInt(z1) > parseInt(z2) ? 1 : -1;
    }

    return 0;
}

/**
 * 检查更新，如果有更新，则通知渲染进程，即 UI 服务
 * */
export async function checkUpdate(mainWin) {
    logInfo('checkUpdate ...')

    const {currVersion, currVersionStr} = getCurrVersion()
    const {newVersion, newVersionStr, forceUpdate} = await getRemoteVersion()
    logInfo(`currVersion=${currVersion}(${currVersionStr}), newVersion=${newVersion}(${newVersionStr}), forceUpdate=${forceUpdate}`)
    logInfo(currVersion < newVersion)

    // 需要更新
    if (compareVersion(newVersionStr, currVersionStr) === 1) {
        if (forceUpdate) {
            // logInfo('forceUpdate')
        } else {
            mainWin.webContents.send(electronMsgUpdate, {
                currVersionStr, newVersionStr, forceUpdate
            })
        }
    }
}

// 更新应用
export const updateApp = (version, mainWin) => {
    downLoadApp(version, mainWin, doUpdate)
}

// 更新应用后，复制静态文件，即 ui目录和 Agent 目录，重启应用
const doUpdate = async (downloadPath, version) => {
    let ok = copyFiles(downloadPath);
    if (!ok) return
    ok = changeVersion(version);
    if (!ok) return
    // 重启应用
    logInfo(`restart app ...`)
    restart();
}

const admZip = require('adm-zip');

import {promisify} from 'node:util';
import stream from 'node:stream';
import fs from 'node:fs';
import got from 'got';
import {cpSync} from "fs";
import os from "os";
import {killAgent} from "../core/agent";
import {isBoolean} from "@vercel/webpack-asset-relocator-loader";
const pipeline = promisify(stream.pipeline);

mkdir(path.join('tmp', 'download'))

// 下载应用
const downLoadApp = (version, mainWin, cb) => {

    // 通过 Node.js 内置的 http 模块发送请求，获取文件，然后写入到本地文件
    const downloadUrl = getAppUrl(version)
    const downloadPath = getDownloadPath(version)

    const downloadStream = got.stream(downloadUrl);
    const fileWriterStream = fs.createWriteStream(downloadPath);

    logInfo(`start download ${downloadUrl} ...`)

    // 更新下载进度，然后通知渲染进程
    downloadStream.on("downloadProgress", ({ transferred, total, percent }) => {
        mainWin.webContents.send(electronMsgDownloading, {percent})
    });

    pipeline(downloadStream, fileWriterStream).then(async () => {
        logInfo(`success to downloaded to ${downloadPath}`)
        const md5Pass = await checkMd5(version, downloadPath);
        if (md5Pass) {
            cb(downloadPath, version)
        } else {
            logInfo('check md5 failed')
        }

    }).catch((err) => {
        logErr(`update failed: ${err}`)
    });
}


/**
 * 复制然后解压，杀 Agent 进程等操作
 * */
const copyFiles = (downloadPath) => {
    const downloadDir = path.dirname(downloadPath)
    const extractedPath = path.resolve(downloadDir, 'extracted')
    logInfo(`downloadPath=${downloadPath}, extractedPath=${extractedPath}`);
    const unzip = new admZip(downloadPath, {});
    let pass = ''
    unzip.extractAllTo(extractedPath, true, true, pass);
    logInfo(`success to extract ${downloadPath}`)
    logInfo(`${extractedPath} ${pass}`)
    const {uiPath, agentPath} = getResPath()
    logInfo(`uiPath=${uiPath}, agentPath=${agentPath}`)
    killAgent();
    fs.rmSync(uiPath, {recursive: true})
    fs.rmSync(agentPath)

    if (!checkFileExist(uiPath) && !checkFileExist(agentPath)) {
        logInfo(`success to remove old resources`)
    } else {
        logInfo(`failed to remove old resources`)
        return false
    }

    const agentFileName = `${agentProcessName}${os.platform() === 'win32' ? '.exe' : ''}`

    logInfo(`agentFileName=${agentFileName}`)

    fse.copySync(path.resolve(downloadDir, 'extracted', 'ui'),          uiPath, {recursive: true})
    fse.copySync(path.resolve(downloadDir, 'extracted', agentFileName), agentPath)

    logInfo(`agentFileName=${IS_WINDOWS_OS}`)

    if (!IS_WINDOWS_OS) {
        const cmd = `chmod +x ${agentPath}`
        logInfo(`cmd=${cmd}`)
        try {
            execSync(cmd, {windowsHide: true})
            logInfo(`success to chmod +x ${agentPath}`)
        }catch (e){
            logInfo(`failed to chmod +x ${agentPath}`, e)
        }

    }

    logInfo(`success to copy new resources`)

    return true
}
