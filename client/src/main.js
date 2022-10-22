import {app} from 'electron';
import {DEBUG} from './app/utils/consts';
import {DeepTestApp} from "./app/app";
import {logInfo} from "./app/utils/log";

// Handle creating/removing shortcuts on Windows when installing/uninstalling.
if (require('electron-squirrel-startup')) { // eslint-disable-line global-require
  app.quit();
}

logInfo(`DEBUG=${DEBUG}`)

const deeptestApp = new DeepTestApp();
app.on('ready', () => {
  deeptestApp.ready()
});
