const logger = require('electron-log');
import {DEBUG} from './consts';
logger.transports.file.resolvePath = () =>
    require("path").join(require("os").homedir(), 'deeptest', 'log', 'electron.log');

export function logDebug(...params) {
    if (DEBUG) {
        logger.debug(params);
    }
}
export function logInfo(...params) {
    logger.info(params.join(', '));
}
export function logErr(...params) {
    logger.error(params);
}
