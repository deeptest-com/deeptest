/**
 * @description 启动 agent 服务
 *
 * */

import {app} from 'electron';
import os from 'os';
import path from 'path';
import {execSync, spawn} from 'child_process';

import {DEBUG, uuid,agentProcessName} from '../utils/consts';
import {IS_WINDOWS_OS} from "../utils/env";
import {logErr, logInfo} from '../utils/log';
import {getBinPath} from "../utils/comm";

let _agentProcess;
let _uuid = uuid;

export async function startAgent(portAgent) {
    // uuid 和 portAgent 一起作为 agent 的唯一标识
    _uuid = uuid + '@' + portAgent;
    if (process.env.SKIP_AGENT_SERVER) {
        logInfo(`>> skip to start deeptest agent service by env "SKIP_AGENT_SERVER=${process.env.SKIP_AGENT_SERVER}".`);
        return Promise.resolve();
    }
    if (_agentProcess) {
        return Promise.resolve(_agentProcess);
    }
    let {SERVER_EXE_PATH: agentExePath} = process.env;
    if (!agentExePath && !DEBUG) {
        agentExePath = getBinPath(agentProcessName);
    }

    if (agentExePath) {
        if (!path.isAbsolute(agentExePath)) {
            agentExePath = path.resolve(app.getAppPath(), agentExePath);
            logInfo(3, agentExePath)
        }
        return new Promise((resolve, reject) => {
            const cwd = process.env.AGENT_CWD_PATH || path.dirname(agentExePath);
            logInfo(`>> starting deeptest-agent with ${agentExePath} -p ${portAgent} -uuid ${_uuid} in ${cwd} ...`);

            const cmd = spawn('"'+agentExePath+'"', ['-p', portAgent, '-uuid', _uuid], {
                cwd,
                shell: true,
            });

            _agentProcess = cmd;
            logInfo(`>> deeptest-agent service process = ${_agentProcess.pid}`)

            cmd.on('close', (code) => {
                logInfo(`>> deeptest-agent service closed with code ${code}`);
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
                        reject(new Error(`start agent service failed, error: ${line.substring('[ERRO]'.length)}`));
                        if (!DEBUG) {
                            break;
                        }
                    }
                }
            });
            cmd.on('error', spawnError => {
                logErr('>>> start agent service failed with error', spawnError);
                reject(spawnError)
            });
        });
    }

    return new Promise((resolve, reject) => {
        const cwd = process.env.AGENT_CWD_PATH || path.resolve(app.getAppPath(), '../');
        logInfo(`>> starting deeptest-agent development service from source with command "go run cmd/agent/main.go -p ${portAgent}" in "${cwd}"`);
        const cmd = spawn('go', ['run', 'cmd/agent/main.go', '-p', portAgent], {
            cwd,
            shell: true,
        });
        cmd.on('close', (code) => {
            logInfo(`>> deeptest-agent service closed with code ${code}`);
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
                    reject(new Error(`start deeptest-agent service failed, error: ${line.substring('[ERRO]'.length)}`));
                    if (!DEBUG) {
                        break;
                    }
                }
            }
        });
        cmd.on('error', spawnError => {
            console.error('>>> start deeptest-agent service failed with error', spawnError);
            reject(spawnError)
        });
        _agentProcess = cmd;
    });
}

export function killAgent() {
    try {
        if (!IS_WINDOWS_OS) {
            logInfo(`>> not windows`);

            const cmd = `ps -ef | grep ${_uuid} | grep -v "grep" | awk '{print $2}' | xargs kill -9`
            logInfo(`>> kill deeptest-agent cmd: ${cmd}`);

            const stdout  = execSync(cmd).toString().trim()
            logInfo(`>> kill deeptest-agent result: ${stdout}`);

        } else {
            const cmd = 'WMIC path win32_process  where "Commandline like \'%%' + _uuid + '%%\'" get Processid,Caption';
            logInfo(`>> list deeptest-agent process cmd: ${cmd}`);

            const stdout = execSync(cmd, {windowsHide: true}).toString().trim()
            logInfo(`>> list deeptest-agent process result: ${stdout}`)

            let pid = 0
            const lines = stdout.split('\n')
            lines.forEach(function(line){
                line = line.trim()
                console.log(`<${line}>`)
                logInfo(`<${line}>`)
                const cols = line.split(/\s+/)

                logInfo(`cols.length=${cols.length}`, JSON.stringify(cols));
                if (line.indexOf('deeptest') > -1 && cols.length > 1) {
                    const pidStr = cols[1].trim();
                    console.log(`>> deeptest-agent pid: ${pidStr}`);
                    logInfo(`>> deeptest-agent pid: ${pidStr}`)

                    if (pidStr && parseInt(pidStr, 10)) {
                        pid = parseInt(pidStr, 10)
                    }
                }
            });

            if (pid && pid > 0) {
                const killCmd = `taskkill /F /pid ${pid}`
                logInfo(`>> kill deeptest-agent cmd: exec ${killCmd}`)

                const out = execSync(`taskkill /F /pid ${pid}`, {windowsHide: true}).toString().trim()
                logInfo(`>> kill deeptest-agent result: ${out}`)
            }
        }
    }catch (e){
        logInfo(`killAgent error: ${e}`)
    }

}
