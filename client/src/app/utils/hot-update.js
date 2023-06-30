import {app} from 'electron';
import {
    changeVersion, checkFileExist, checkMd5, getAppUrl,
    getCurrVersion,
    getDownloadPath,
    getRemoteVersion,
    getResPath, mkdir,
    restart
} from "./comm";
import {electronMsgDownloading, electronMsgUpdate, WorkDir} from "./consts";
import path from "path";
import {execSync} from 'child_process';
import {IS_WINDOWS_OS} from "../utils/env";
import fse from 'fs-extra'
import {logErr, logInfo} from "./log";

export async function checkUpdate(mainWin) {
    logInfo('checkUpdate ...')

    const {currVersion, currVersionStr} = getCurrVersion()
    const {newVersion, newVersionStr, forceUpdate} = await getRemoteVersion()
    logInfo(`currVersion=${currVersion}(${currVersionStr}), newVersion=${newVersion}(${newVersionStr}), forceUpdate=${forceUpdate}`)
    logInfo(currVersion < newVersion)

    if (currVersion < newVersion) {
        if (forceUpdate) {
            // logInfo('forceUpdate')
        } else {
            mainWin.webContents.send(electronMsgUpdate, {
                currVersionStr, newVersionStr, forceUpdate
            })
        }
    }
}

export const updateApp = (version, mainWin) => {
    downLoadApp(version, mainWin, doUpdate)
}

const doUpdate = async (downloadPath, version) => {
    let ok = copyFiles(downloadPath);
    if (!ok) return

    ok = changeVersion(version);
    if (!ok) return

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

const downLoadApp = (version, mainWin, cb) => {
    const downloadUrl = getAppUrl(version)
    const downloadPath = getDownloadPath(version)

    const downloadStream = got.stream(downloadUrl);
    const fileWriterStream = fs.createWriteStream(downloadPath);

    logInfo(`start download ${downloadUrl} ...`)

    downloadStream.on("downloadProgress", ({ transferred, total, percent }) => {
        mainWin.webContents.send(electronMsgDownloading, {percent})
    });

    pipeline(downloadStream, fileWriterStream).then(async () => {
        logInfo(`success to downloaded to ${downloadPath}`)

        const md5Pass = await checkMd5(version, downloadPath)
        if (md5Pass) {
            cb(downloadPath, version)
        } else {
            logInfo('check md5 failed')
        }

    }).catch((err) => {
        logErr(`update failed: ${err}`)
    });
}

const copyFiles = (downloadPath) => {
    const downloadDir = path.dirname(downloadPath)

    const extractedPath = path.resolve(downloadDir, 'extracted')
    logInfo(`downloadPath=${downloadPath}, extractedPath=${extractedPath}`)

    const unzip = new admZip(downloadPath, {});
    let pass = ''
    unzip.extractAllTo(extractedPath, true, true, pass);
    logInfo(pass)

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

    const agentFileName = `agent${os.platform() === 'win32' ? '.exe' : ''}`

    fse.copySync(path.resolve(downloadDir, 'extracted', 'ui'),          uiPath, {recursive: true})
    fse.copySync(path.resolve(downloadDir, 'extracted', agentFileName), agentPath)

    if (!IS_WINDOWS_OS) {
        const cmd = `chmod +x ${agentPath}`
        execSync(cmd, {windowsHide: true})
    }

    logInfo(`success to copy new resources`)

    return true
}