import {app, dialog} from 'electron';
import {DEBUG} from './app/utils/consts';
import {DeepTestApp} from "./app/app";
import {logInfo, logErr, logger} from "./app/utils/log";

// Handle creating/removing shortcuts on Windows when installing/uninstalling.
if (require('electron-squirrel-startup')) { // eslint-disable-line global-require
  app.quit();
}

logInfo(`DEBUG=${DEBUG}`)
logInfo('===' + __dirname)

const isDev = require('electron-is-dev')
const mode = isDev ? 'development' : 'production'
logInfo(`Start DeepTest v${app.getVersion()} in ${mode} mode ...`)

const { autoUpdater } = require("electron-updater")
autoUpdater.logger = logger

const deeptestApp = new DeepTestApp();
app.on('ready', () => {
  deeptestApp.ready()
});
