
/**
 * @description 启动 ui 服务，
 *
 *  如果传入了 UI_SERVER_URL 环境变量，
 *      如果，有效的 http 地址，则直接返该地址
 *      如果，是本地的静态资源路径，则启动 express 服务
 *  如果，没有传入 UI_SERVER_URL 环境变量，
 *     则直接返回该地址，否则从本地启动 【UI 服务】 即执行命令： npm run serve
 *
 * */
import {app} from 'electron';
import path from 'path';
import {spawn} from 'child_process';
import express from 'express';

import {DEBUG} from '../utils/consts';
import {logInfo, logErr} from '../utils/log';
const history = require('connect-history-api-fallback');

let _uiService;

export function startUIService(portClient) {
    // ui 已经启动
    if (_uiService) {
        return Promise.resolve();
    }

    // 从环境变量中获取 ui 服务地址
    let {UI_SERVER_URL: uiServerUrl} = process.env;

    if (!uiServerUrl && !DEBUG) {
        uiServerUrl = path.resolve(process.resourcesPath, 'ui');
    }

    uiServerUrl = path.resolve(process.resourcesPath, 'ui');


    if (uiServerUrl) {
        // 有效的 http 地址
        if (/^https?:\/\//.test(uiServerUrl)) {
            return Promise.resolve(uiServerUrl);
        }
        // 返回本地的静态资源路径，启动 express 服务
        return new Promise((resolve, reject) => {
            if (!path.isAbsolute(uiServerUrl)) {
                uiServerUrl = path.resolve(app.getAppPath(), uiServerUrl);
            }

            const port = portClient;
            logInfo(`>> starting ui serer at ${uiServerUrl} with port ${port}`);

            const uiServer = express();
            uiServer.use(history());
            uiServer.use(express.static(uiServerUrl));
            const server = uiServer.listen(port, serverError => {
                if (serverError) {
                    console.error('>>> start ui server failed with error', serverError);
                    _uiService = null;
                    reject(serverError);
                } else {
                    logInfo(`>> ui server started successfully on http://localhost:${port}.`);
                    resolve(`http://localhost:${port}`);
                }
            });
            // express 服务关闭时，清空 _uiService
            server.on('close', () => {
                _uiService = null;
            });
            _uiService = uiServer;
        })
    }

    // 从本地启动 npm 启动 ui 服务，即 npm run serve
    return new Promise((resolve, reject) => {
        const cwd = path.resolve(app.getAppPath(), '../ui');
        logInfo(`>> starting ui development server with command "npm run serve" in "${cwd}"...`);

        let resolved = false;
        const cmd = spawn('npm', ['run', 'serve'], {
            cwd,
            shell: true,
        });
        cmd.on('close', (code) => {
            logInfo(`>> ui server closed with code ${code}`);
            _uiService = null;
        });
        cmd.stdout.on('data', data => {
            if (resolved) {
                return;
            }
            const dataString = String(data);
            const lines = dataString.split('\n');
            for (let i = 0; i < lines.length; i++) {
                const line = lines[i];
                if (DEBUG) {
                    logInfo('\t' + line);
                }
                if (line.includes('App running at:')) {
                    const nextLine = lines[i + 1] || lines[i + 2];
                    if (DEBUG) {
                        logInfo('\t' + nextLine);
                    }
                    if (!nextLine) {
                        logErr('\t' + `cannot grabing running address after line "${line}".`);
                        throw new Error(`cannot grabing running address after line "${line}".`);
                    }
                    const url = nextLine.split('Local:   ')[1];
                    if (url) {
                        resolved = true;
                        resolve(url);
                    }
                    if (!DEBUG) {
                        break;
                    }
                }
            }
        });
        cmd.on('error', spawnError => {
            logErr(`>>> start ui server failed, error ${spawnError}`);
            reject(spawnError)
        });
        _uiService = cmd;
    });
}
