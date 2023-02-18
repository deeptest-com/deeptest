import {app, dialog} from 'electron';
import {DEBUG} from './app/utils/consts';
import {DeepTestApp} from "./app/app";
import {logInfo, logErr, logger} from "./app/utils/log";

// Handle creating/removing shortcuts on Windows when installing/uninstalling.
if (require('electron-squirrel-startup')) { // eslint-disable-line global-require
  app.quit();
}

logInfo(`DEBUG=${DEBUG}`)

const isDev = require('electron-is-dev')
const mode = isDev ? 'development' : 'production'
logInfo(`Start DeepTest v${app.getVersion()} in ${mode} mode ...`)

const { autoUpdater } = require("electron-updater")
autoUpdater.logger = logger

const deeptestApp = new DeepTestApp();
app.on('ready', () => {
  deeptestApp.ready()

  if (!isDev) {
    checkForUpdates()
  }
});

function checkForUpdates() {
  logInfo('Set up event listeners...')
  autoUpdater.on('checking-for-update', () => {
    logInfo('Checking for update...')
  })
  autoUpdater.on('update-available', (info) => {
    logInfo('Update available.')
  })
  autoUpdater.on('update-not-available', (info) => {
    logInfo('Update not available.')
  })
  autoUpdater.on('error', (err) => {
    logErr('Error in auto-updater.' + err)
  })
  autoUpdater.on('download-progress', (progressObj) => {
    let msg = "Download speed: " + progressObj.bytesPerSecond
    msg = msg + ' - Downloaded ' + progressObj.percent + '%'
    msg = msg + ' (' + progressObj.transferred + "/" + progressObj.total + ')'
    logInfo(msg)
  })
  autoUpdater.on('update-downloaded', (info) => {
    logInfo('Update downloaded.')

    // The update will automatically be installed the next time the app launches.
    // If you want to, you can force the installation now:
    const dialogOpts = {
      type: 'info',
      buttons: ['重启', '稍后'],
      title: '自动更新',
      message: process.platform === 'win32' ? info.releaseNotes : info.releaseName,
      detail: `新版本（${info.version}）下载成功，请重启完成升级。`
    }

    dialog.showMessageBox(dialogOpts).then((returnValue) => {
      if (returnValue.response === 0) autoUpdater.quitAndInstall()
    })
  })

  // More properties on autoUpdater, see https://www.electron.build/auto-update#AppUpdater
  //autoUpdater.autoDownload = true
  //autoUpdater.autoInstallOnAppQuit = true

  // No debugging! Check log for details.
  // Ready? Go!
  logInfo('checkForUpdates() -- begin')
  try {
    const server = 'update_server'
    const url = `${server}/update/${process.platform}/ ${app.getVersion()}`
    autoUpdater.setFeedURL({ url })

    autoUpdater.checkForUpdates()
    //autoUpdater.checkForUpdatesAndNotify()
  } catch (error) {
    logErr(error)
  }
  logInfo('checkForUpdates() -- end')
}
