import {app} from 'electron';
import os from 'os';
import path from 'path';
import {execSync, spawn} from 'child_process';

import {DEBUG, portAgent, uuid} from '../utils/consts';
import {IS_WINDOWS_OS} from "../utils/env";
import {logErr, logInfo} from '../utils/log';

let _agentProcess;

export async function startAgent() {
    if (process.env.SKIP_AGENT_SERVER) {
        logInfo(`>> skip to start deeptest agent service by env "SKIP_AGENT_SERVER=${process.env.SKIP_AGENT_SERVER}".`);
        return Promise.resolve();
    }
    if (_agentProcess) {
        return Promise.resolve(_agentProcess);
    }

    let {SERVER_EXE_PATH: agentExePath} = process.env;
    if (!agentExePath && !DEBUG) {
        const platform = os.platform(); // 'darwin', 'linux', 'win32'
        const exePath = `bin/${platform}/deeptest${platform === 'win32' ? '.exe' : ''}`;
        agentExePath = path.join(process.resourcesPath, exePath);
    }
    if (agentExePath) {
        if (!path.isAbsolute(agentExePath)) {
            agentExePath = path.resolve(app.getAppPath(), agentExePath);
        }
        return new Promise((resolve, reject) => {
            const cwd = process.env.AGENT_CWD_PATH || path.dirname(agentExePath);
            logInfo(`>> starting agent server with command ` +
                `"${agentExePath} -p ${portAgent} -uuid ${uuid}" in "${cwd}"...`);

            const cmd = spawn('"'+agentExePath+'"', ['-p', portAgent, '-uuid', uuid], {
                cwd,
                shell: true,
            });

            _agentProcess = cmd;
            logInfo(`>> agent server process = ${_agentProcess.pid}`)

            cmd.on('close', (code) => {
                logInfo(`>> agent server closed with code ${code}`);
                _agentProcess = null;
                cmd.kill()
            });
            cmd.stdout.on('data', data => {
                const dataString = String(data);
                const lines = dataString.split('\n');
                for (let line of lines) {
                    if (DEBUG) {
                        logInfo('\t' + line);
                    }
                    if (line.includes('Now listening on: http')) {
                        resolve(line.split('Now listening on:')[1].trim());
                        if (!DEBUG) {
                            break;
                        }
                    } else if (line.includes('启动HTTP服务于')) {
                        resolve(line.split(/启动HTTP服务于|，/)[1].trim());
                        if (!DEBUG) {
                            break;
                        }
                    } else if (line.startsWith('[ERRO]')) {
                        reject(new Error(`start agent server failed, error: ${line.substring('[ERRO]'.length)}`));
                        if (!DEBUG) {
                            break;
                        }
                    }
                }
            });
            cmd.on('error', spawnError => {
                logErr('>>> start agent server failed with error', spawnError);
                reject(spawnError)
            });
        });
    }

    return new Promise((resolve, reject) => {
        const cwd = process.env.AGENT_CWD_PATH || path.resolve(app.getAppPath(), '../');
        logInfo(`>> starting agent development server from source with command "go run cmd/agent/main.go -p ${portAgent}" in "${cwd}"`);
        const cmd = spawn('go', ['run', 'cmd/agent/main.go', '-p', portAgent], {
            cwd,
            shell: true,
        });
        cmd.on('close', (code) => {
            logInfo(`>> agent server closed with code ${code}`);
            _agentProcess = null;
        });
        cmd.stdout.on('data', data => {
            const dataString = String(data);
            const lines = dataString.split('\n');
            for (let line of lines) {
                if (DEBUG) {
                    logInfo('\t' + line);
                }
                if (line.includes('Now listening on: http')) {
                    resolve(line.split('Now listening on:')[1].trim());
                    if (!DEBUG) {
                        break;
                    }
                } else if (line.startsWith('[ERRO]')) {
                    reject(new Error(`start agent server failed, error: ${line.substring('[ERRO]'.length)}`));
                    if (!DEBUG) {
                        break;
                    }
                }
            }
        });
        cmd.on('error', spawnError => {
            console.error('>>> start agent server failed with error', spawnError);
            reject(spawnError)
        });
        _agentProcess = cmd;
    });
}

export function killAgent() {
    if (!IS_WINDOWS_OS) {
        logInfo(`>> not windows`);

        const cmd = `ps -ef | grep ${uuid} | grep -v "grep" | awk '{print $2}' | xargs kill -9`
        logInfo(`>> exit cmd: ${cmd}`);

        const cp = require('child_process');
        cp.exec(cmd, function (error, stdout, stderr) {
            logInfo(`>> exit result: stdout: ${stdout}; stderr: ${stderr}; error: ${error}`);
        });
    } else {
        const cmd = 'WMIC path win32_process  where "Commandline like \'%%' + uuid + '%%\'" get Processid,Caption';
        logInfo(`>> list process cmd: ${cmd}`);

        const stdout = execSync(cmd, {windowsHide: true}).toString().trim()
        logInfo(`>> list process result: exec ${cmd}, stdout: ${stdout}`)

        let pid = 0
        const lines = stdout.split('\n')
        lines.forEach(function(line){
            line = line.trim()
            console.log(`<${line}>`)
            logInfo(`<${line}>`)
            const cols = line.split(/\s/)

            if (line.indexOf('deeptest') > -1 && cols.length > 3) {
                const col3 = cols[3].trim()
                console.log(`col3=${col3}`);
                logInfo(`col3=${col3}`)

                if (col3 && parseInt(col3, 10)) {
                    pid = parseInt(col3, 10)
                }
            }
        });

        if (pid && pid > 0) {
            const killCmd = `taskkill /F /pid ${pid}`
            logInfo(`>> exit cmd: exec ${killCmd}`)

            const out = execSync(`taskkill /F /pid ${pid}`, {windowsHide: true}).toString().trim()
            logInfo(`>> exit result: ${out}`)
        }
    }
}