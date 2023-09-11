import path from "path";
import os from "os";
import {agentProcessName, App, downloadUrl, ResDir, WorkDir} from "./consts";
import {app} from "electron";
import {logInfo} from "./log";
import {killAgent} from "../core/agent";
import fs from 'node:fs';
import got from 'got';
import crypto from "crypto";

export function mkdir(dir) {
    const pth = path.join(WorkDir, dir);
    fs.mkdirSync(pth, {recursive:true})

    return pth
}

// export function getResDir(dir) {
//     const pth = path.resolve(process.resourcesPath, dir);
//     return pth
// }

export function getDownloadPath(version) {
    const pth = path.join(WorkDir, 'tmp', 'download', `${version}.zip`);
    return pth
}


export function getCurrVersion() {
    let currVersionStr = '0';

    const {versionPath} = getResPath();
    logInfo(`83222 versionPath=${versionPath}`)

    logInfo(`83222 versionPath=${app.getVersion()}`)

    if (fs.existsSync(versionPath)) {
        const content = fs.readFileSync(versionPath)
        let json = JSON.parse(content);
        currVersionStr = json.version;
    } else {
        currVersionStr = app.getVersion();
    }

    const currVersion = parseFloat(currVersionStr);

    return {currVersion, currVersionStr};
}

export async function getRemoteVersion() {
    const versionUrl = getVersionUrl();

    logInfo(`83222${versionUrl}`)

    const json = await got.get(versionUrl).json();
    const newVersionStr = json.version;
    const newVersion = parseFloat(newVersionStr);
    const forceUpdate = json.force?json.force:false;

    return {
        newVersion,
        newVersionStr,
        forceUpdate,
    }
}

export function changeVersion(newVersion) {
    let res = false;
    try {

        const versionPath = getResPath().versionPath;
        logInfo(`ResDir=${ResDir}, versionPath=${versionPath}`)

        let json = {}
        if (fs.existsSync(versionPath)) {
            const content = fs.readFileSync(versionPath)
            json = JSON.parse(content);
        }

        json.version = newVersion;
        fs.writeFileSync(versionPath, JSON.stringify(json));
        logInfo(`success to write new version`)
        res = true;
    }catch (e){
        logInfo(`83222 changeVersion error: ${e}`)
    }

    return  res;
}

export function restart() {
    try {
        killAgent();
        app.relaunch({
            args: process.argv.slice(1)
        });
        app.exit(0);
    }catch (e) {
        logInfo(`83222 restart error: ${e}`)
    }
}
export function getResPath() {
    // const versionPath = path.resolve(ResDir, 'version.json');
    const versionPath =`${ResDir}/version.json`;
    const uiPath =  path.resolve(ResDir, 'ui');
    const agentPath = getBinPath(agentProcessName);

    return {
        versionPath, uiPath, agentPath
    }
}

/**
 * @param name: bin name
 * */
export function getBinPath(name) {
    const platform = os.platform(); // 'darwin', 'linux', 'win32'
    const execPath = `bin/${platform}/${name}${platform === 'win32' ? '.exe' : ''}`;
    const pth = path.join(ResDir, execPath);

    return pth
}

export function computerFileMd5(pth) {
    const buffer = fs.readFileSync(pth);
    const hash = crypto.createHash('md5');
    hash.update(buffer, 'utf8');
    const md5 = hash.digest('hex') + '';
    return md5.trim()
}

export function getVersionUrl() {
    const url = new URL(`${App}/version.json`, downloadUrl) + '?ts=' + Date.now();
    logInfo(`versionUrl=${url}`)
    return url
}
export function getAppUrl(version) {
    const platform = getPlatform(); // 'darwin', 'linux', 'win32', 'win64'
    logInfo(`platform=${platform}`)
    const url = new URL(`${App}/${version}/${platform}/${App}-upgrade.zip`, downloadUrl) + '?ts=' + Date.now();
    logInfo(`appUrl=${url}`)
    return url
}

export function getPlatform() {
    let platform = os.platform(); // 'darwin', 'linux', 'win32'

    if (platform === 'win32' && ['arm64', 'ppc64', 'x64', 's390x'].includes(os.arch())) {
        platform = 'win64'
    }

    return platform
}

export async function checkMd5(version, file) {
    const platform = getPlatform(); // 'darwin', 'linux', 'win32'
    const url = new URL(`${App}/${version}/${platform}/${App}-upgrade.zip.md5`, downloadUrl) + '?ts=' + Date.now();

    logInfo(`md5Url=${url}, file=${file}`)

    const md5Remote = (await got.get(url).text() + '').trim();
    const md5File = computerFileMd5(file)
    const pass = md5Remote === md5File

    logInfo(`md5Remote=${md5Remote}, md5File=${md5File}, pass=${pass}`)

    return pass
}

export function checkFileExist(path) {
    const ret = fs.existsSync(path)
    return ret
}
