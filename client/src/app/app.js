import {app, BrowserWindow, ipcMain, Menu, shell, dialog, globalShortcut} from 'electron';
const path = require('path');
import {
    DEBUG,
    electronMsg,
    electronMsgReplay,
    electronMsgUpdate,
    minimumSizeHeight,
    minimumSizeWidth
} from './utils/consts';
import {IS_MAC_OSX, IS_LINUX, IS_WINDOWS_OS} from './utils/env';
import {logInfo, logErr} from './utils/log';
import Config from './utils/config';
import Lang, {initLang} from './core/lang';
import {startUIService} from "./core/ui";
import {startAgent, killAgent} from "./core/agent";
import { nanoid } from 'nanoid'
import {uploadFile} from "./utils/upload";
import {getCurrVersion, mkdir} from "./utils/comm";
import {checkUpdate, updateApp} from "./utils/hot-update";

const cp = require('child_process');
const fs = require('fs');

const bent = require('bent');
const getBuffer = bent('buffer')

app.commandLine.appendSwitch('disable-web-security');
let postmanToOpenApi = null

mkdir('converted')

export class DeepTestApp {
    constructor() {
        app.name = Lang.string('app.title', Config.pkg.displayName);

        app.commandLine.appendSwitch('disable-features', 'OutOfBlinkCors')
        app.commandLine.appendSwitch('disable-site-isolation-trials')
        app.commandLine.appendSwitch('disable-features','BlockInsecurePrivateNetworkRequests')

        this._windows = new Map();

        startAgent().then((agentUrl)=> {
            if (agentUrl) logInfo(`>> deeptest-agent service started successfully on : ${agentUrl}`);
            this.bindElectronEvents();
        }).catch((err) => {
            logErr('>> deeptest-agent service started failed, err: ' + err);
            process.exit(1);
        })
    }

    showAndFocus() {
        logInfo(`>> deeptest app: AppWindow[${this.name}]: show and focus`);

        const {browserWindow} = this;
        if (browserWindow.isMinimized()) {
            browserWindow.restore();
        } else {
            browserWindow.setOpacity(1);
            browserWindow.show();
        }
        browserWindow.focus();
    }

    async createWindow() {
        process.env['ELECTRON_DISABLE_SECURITY_WARNINGS'] = 'true';

        const mainWin = new BrowserWindow({
            show: false,
            frame: true, // 不创建无边框窗口，不然都没法拖动
            webPreferences: {
                nodeIntegration: true,
                contextIsolation: false,
                enableRemoteModule: true,
                webSecurity: false,
            }
        })
        if (IS_LINUX && !DEBUG) {
            const pth = path.join(__dirname, 'icon', 'favicon.png')
            mainWin.setIcon(pth);
        }

        require('@electron/remote/main').initialize()
        require('@electron/remote/main').enable(mainWin.webContents)
        const {currVersionStr} = getCurrVersion()
        global.sharedObj =  { version : currVersionStr};

        mainWin.setSize(minimumSizeWidth, minimumSizeHeight)
        mainWin.setMovable(true)
        mainWin.maximize()
        mainWin.show()

        this._windows.set('main', mainWin);

        const url = await startUIService()
        await mainWin.loadURL(url);

        // const temp = path.resolve(process.resourcesPath, 'ui_temp')
        // logInfo('===' + temp)
        // fs.mkdirSync(temp);
        // let url = 'http://110.42.146.127:8085/ui'
        // if (process.env.UI_SERVER_URL) {
        //     url = process.env.UI_SERVER_URL
        // }
        // logInfo('load ' + url)

        ipcMain.on(electronMsg, (event, arg) => {
            logInfo('msg from renderer', JSON.stringify(arg))

            if (arg.act == 'loadSpec') { // load openapi spec
                this.loadSpec(event, arg)
                return
            } else if (arg.act == 'chooseFile') { // choose file for interface form
                this.chooseFile(event, arg)
                return
            } else if (arg.act == 'uploadFile') { // upload file
                this.uploadFile(event, arg)
                return
            }

            switch (arg) {
                case 'selectDir':
                    this.showFolderSelection(event)
                    break;
                case 'selectFile':
                    this.showFileSelection(event)
                    break;

                case 'fullScreen':
                    mainWin.setFullScreen(!mainWin.isFullScreen());
                    break;
                case 'minimize':
                    mainWin.minimize();
                    break;
                case 'maximize':
                    mainWin.maximize();
                    break;
                case 'unmaximize':
                    mainWin.unmaximize();
                    break;

                case 'help':
                    shell.openExternal('https://deeptest.com');
                    break;
                case 'exit':
                    app.quit()
                    break;
                default:
                   logInfo('--', arg.action, arg.path)
                    if (arg.action == 'openInExplore')
                        this.openInExplore(arg.path)
                   else if (arg.action == 'openInTerminal')
                        this.openInTerminal(arg.path)
            }
        })
    }

    async openOrCreateWindow() {
        const mainWin = this._windows.get('main');
        if (mainWin) {
            this.showAndFocus(mainWin)
        } else {
            await this.createWindow();
        }
    }

    showAndFocus(mainWin) {
        if (mainWin.isMinimized()) {
            mainWin.restore();
        } else {
            mainWin.setOpacity(1);
            mainWin.show();
        }
        mainWin.focus();
    }

     ready() {
        logInfo('>> deeptest app ready.');

        initLang()
        this.buildAppMenu();
        this.openOrCreateWindow()
        this.setAboutPanel();

         globalShortcut.register('CommandOrControl+D', () => {
             const mainWin = this._windows.get('main');
             mainWin.toggleDevTools()
         })

         ipcMain.on(electronMsgUpdate, (event, arg) => {
             logInfo('update confirm from renderer', arg)

             const mainWin = this._windows.get('main');
             updateApp(arg.newVersion, mainWin)
         });

         setInterval(async () => {
             await checkUpdate(this._windows.get('main'))
         }, 6000);
    }

    quit() {
        killAgent();
    }

    bindElectronEvents() {
        app.on('activate', () => {
            logInfo('>> event: app activate');

            this.buildAppMenu();

            // 在 OS X 系统上，可能存在所有应用窗口关闭了，但是程序还没关闭，此时如果收到激活应用请求，需要重新打开应用窗口并创建应用菜单。
            this.openOrCreateWindow()
        });

        app.on('window-all-closed', () => {
            logInfo(`>> event: app window-all-closed`)
            app.quit();
        });

        app.on('will-quit', (e) => {
            e.preventDefault()
            logInfo(`>> event: app will-quit`)

            logInfo(`>> start to kill child process`)
            this.quit();
            logInfo(`>> end to kill child process`)

            app.exit()
        });

        app.on('quit', () => {
            logInfo(`>> event: app quit`)
        });
    }

    // choose file
    async chooseFile(event, arg) {
        const result = await dialog.showOpenDialog({properties: ['openFile']})

        if (result.filePaths && result.filePaths.length > 0) {
            event.reply(electronMsgReplay, {
                filepath: result.filePaths[0],
            });
        }
    }

    // upload file
    async uploadFile(event, arg) {
        const result = await dialog.showOpenDialog({
            properties: ['openFile'],
            filters: arg.filters,
        })

        if (result.filePaths && result.filePaths.length > 0) {
            const resp = await uploadFile(arg.url, arg.token, result.filePaths[0], arg.params)
            console.log('uploadFile resp: ', resp)

            event.reply(electronMsgReplay, resp);
        }
    }

    // load spec
    loadSpec(event, arg) {
        if (arg.src === 'url') {
            this.loadFromUrl(event, arg)
        } else {
            this.showSpecSelection(event, arg)
        }
    }

    async loadFromUrl(event, arg) {
        logInfo('loadFromUrl')
        const buffer = await getBuffer(arg.url)

        const index = arg.url.lastIndexOf('/')
        const fileName = arg.url.substring(index)
        const file = path.join(convertedDir, fileName);

        fs.writeFileSync(file, buffer)

        this.replyLoadApiSpec(event, arg.type, 'url', file, arg.url)
    }

    async showSpecSelection(event, arg) {
        const result = await dialog.showOpenDialog({properties: ['openFile']})

        if (result.filePaths && result.filePaths.length > 0) {
            this.replyLoadApiSpec(event, arg.type, 'file', result.filePaths[0], null)
        }
    }

    async replyLoadApiSpec(event, spaceType, spaceSrc, file, url) {
        console.log(`replyLoadApiSpec`, spaceType, spaceSrc, file, url)

        // load the content to check that if it is a postman file
        let content = fs.readFileSync(file).toString()

        const isPostMan = spaceType === 'postman' || content.indexOf('"_postman_id"') > -1

        if (isPostMan) {
            if (!postmanToOpenApi) postmanToOpenApi = require('postman-to-openapi')

            const oldFile = file
            file = path.join(convertedDir, 'postman-' + nanoid() + '.json')
            await postmanToOpenApi(oldFile, file, {defaultTag: 'General'})
        }

        event.reply(electronMsgReplay, {src: spaceSrc, type: spaceType, file: file, url: url});
    }

    // common file selection
    showFileSelection(event) {
        dialog.showOpenDialog({
            properties: ['openFile']
        }).then(result => {
            this.replyFile(event, result)
        }).catch(err => {
            logErr(err)
        })
    }

    showFolderSelection(event) {
        dialog.showOpenDialog({
            properties: ['openDirectory']
        }).then(result => {
            this.replyFile(event, result)
        }).catch(err => {
            logErr(err)
        })
    }

    replyFile(event, result)  {
        if (result.filePaths && result.filePaths.length > 0) {
            event.reply(electronMsgReplay, result.filePaths[0]);
        }
    }

    openInExplore(pth) {
        shell.showItemInFolder(pth);
    }
    openInTerminal(pth) {
        logInfo('openInTerminal')

        const stats = fs.statSync(pth);
        if (stats.isFile()) {
            pth = pth.resolve(pth, '..')
        }

        if (IS_WINDOWS_OS) {
            cp.exec('start cmd.exe /K cd /D ' + pth);
        } else if (IS_LINUX) {
            // support other terminal types
            cp.spawn ('gnome-terminal', [], { cwd: pth });
        } else if (IS_MAC_OSX) {
            cp.exec('open -a Terminal ' + pth);
        }
    }

    get windows() {

        return this._windows;
    }

    setAboutPanel() {
        if (!app.setAboutPanelOptions) {
            return;
        }

        let version = Config.pkg.buildTime ? 'build at ' + new Date(Config.pkg.buildTime).toLocaleString() : ''
        version +=  DEBUG ? '[debug]' : ''
        app.setAboutPanelOptions({
            applicationName: Lang.string(Config.pkg.name) || Config.pkg.displayName,
            applicationVersion: Config.pkg.version,
            copyright: `${Config.pkg.copyright} ${Config.pkg.company}`,
            credits: `Licence: ${Config.pkg.license}`,
            version: version
        });
    }

    buildAppMenu() {
        logInfo('>> deeptest app: build application menu.');

        if (IS_MAC_OSX) {
            const template = [
                {
                    label: Lang.string('app.title', Config.pkg.displayName),
                    submenu: [
                        {
                            label: Lang.string('app.about'),
                            selector: 'orderFrontStandardAboutPanel:'
                        }, {
                            label: Lang.string('app.exit'),
                            accelerator: 'Command+Q',
                            click: () => {
                                app.quit();
                            }
                        }
                    ]
                },
                {
                    label: Lang.string('app.edit'),
                    submenu: [{
                        label: Lang.string('app.undo'),
                        accelerator: 'Command+Z',
                        selector: 'undo:'
                    }, {
                        label: Lang.string('app.redo'),
                        accelerator: 'Shift+Command+Z',
                        selector: 'redo:'
                    }, {
                        type: 'separator'
                    }, {
                        label: Lang.string('app.cut'),
                        accelerator: 'Command+X',
                        selector: 'cut:'
                    }, {
                        label: Lang.string('app.copy'),
                        accelerator: 'Command+C',
                        selector: 'copy:'
                    }, {
                        label: Lang.string('app.paste'),
                        accelerator: 'Command+V',
                        selector: 'paste:'
                    }, {
                        label: Lang.string('app.select_all'),
                        accelerator: 'Command+A',
                        selector: 'selectAll:'
                    }]
                },
                {
                    label: Lang.string('app.view'),
                    submenu:  [
                        {
                            label: Lang.string('app.switch_to_full_screen'),
                            accelerator: 'Ctrl+Command+F',
                            click: () => {
                                const mainWin = this._windows.get('main');
                                mainWin.setFullScreen(!mainWin.isFullScreen());
                            }
                        }
                    ]
                },
                {
                    label: Lang.string('app.window'),
                    submenu: [
                        {
                            label: Lang.string('app.minimize'),
                            accelerator: 'Command+M',
                            selector: 'performMiniaturize:'
                        },
                        {
                            label: Lang.string('app.close'),
                            accelerator: 'Command+W',
                            selector: 'performClose:'
                        },
                        {
                            label: 'Reload',
                            accelerator: 'Command+R',
                            click: () => {
                                this._windows.get('main').webContents.reload();
                            }
                        },
                        {
                            type: 'separator'
                        },
                        {
                            label: Lang.string('app.bring_all_to_front'),
                            selector: 'arrangeInFront:'
                        }
                    ]
                },
                {
                    label: Lang.string('app.help'),
                    submenu: [{
                        label: Lang.string('app.website'),
                        click: () => {
                            shell.openExternal('https://deeptest.com');
                        }
                    }]
                }
            ];

            const menu = Menu.buildFromTemplate(template);
            Menu.setApplicationMenu(menu);
        } else {
            Menu.setApplicationMenu(null);
        }
    }
}
