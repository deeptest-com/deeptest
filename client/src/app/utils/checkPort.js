import {logInfo} from "./log";

const portfinder = require('portfinder');

/**
 * @description 检测端口是否被占用
 * */
export async function checkPortIsUsed(port) {
    const newPort = await portfinder.getPortPromise({
        port: port,    // minimum port
    });
    logInfo(`>> check port ${port} is used: ${newPort !== port}`);
    return newPort === port;
}

/**
 * @description 获取可用端口
 * @param port {number} 起始端口
 * @param maxPort {number} 最大端口
 * */
export const getUsefulPort = async function (port, maxPort) {
    const newPort = await portfinder.getPortPromise({
        port: port,    // minimum port
        maxPort: maxPort, // maximum port
    });
    logInfo(`>> getUsefulPort port ${port} is used ${newPort} : ${newPort !== port}`);
    return newPort;
}



