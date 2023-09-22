import {app, dialog} from 'electron';
import {DeepTestApp} from "./app/app";
import {logInfo, logErr, logger} from "./app/utils/log";
import {getCurrVersion} from "./app/utils/comm";

// Handle creating/removing shortcuts on Windows when installing/uninstalling.
if (require('electron-squirrel-startup')) {  // eslint-disable-line global-require
  app.quit();
}



const isDev = require('electron-is-dev')
const mode = isDev ? 'development' : 'production'
logInfo(`Start DeepTest v${app.getVersion()} in ${mode} mode ...`)

const deeptestApp = new DeepTestApp();
app.on('ready', () => {
  deeptestApp.ready()
});
