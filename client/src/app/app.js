import {app, BrowserWindow, ipcMain, Menu, shell, dialog, globalShortcut, createWindow} from 'electron';
import {
    DEBUG,
    electronMsg,
    electronMsgReplay, electronMsgServerUrl,
    electronMsgUpdate,
    electronMsgUsePort,
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
import {portAgent,portClient} from "./utils/consts";
import {getUsefulPort} from "./utils/checkPort";
import {App} from "./utils/constant";
import logger from "electron-log";

const cp = require('child_process');
const path = require('path');
const os = require("os")
const fs = require('fs');
const yaml = require('js-yaml');
const bent = require('bent');
const getBuffer = bent('buffer')

app.commandLine.appendSwitch('disable-web-security');
let postmanToOpenApi = null

// agent 服务端口
let port = '';

mkdir('converted')

export class DeepTestApp {
    constructor() {
        app.name = Lang.string('app.title', Config.pkg.displayName);
        app.commandLine.appendSwitch('disable-features', 'OutOfBlinkCors')
        app.commandLine.appendSwitch('disable-site-isolation-trials')
        app.commandLine.appendSwitch('disable-features','BlockInsecurePrivateNetworkRequests')
        this._windows = new Map();

        // 需要启动本地 Agent 服务，之后再会启动 UI 服务
        (async () => {
            port = await getUsefulPort(portAgent,56999);

            logInfo(`>> starting deeptest agent on port ${port}`);

            startAgent(port).then((agentUrl)=> {
                if (agentUrl) logInfo(`>> deeptest agent started successfully at ${agentUrl}`);
                this.bindElectronEvents();
            }).catch((err) => {
                logErr('>> deeptest agent started failed on port ${port}, err: ' + err);
                process.exit(1);
            })
        })()

    }

    /**
     * @description 创建主窗口
     * */
    async createWindow() {
        process.env['ELECTRON_DISABLE_SECURITY_WARNINGS'] = 'true';
        const mainWin = new BrowserWindow({
            show: false,
            frame: true, // 不创建无边框窗口，不然都没法拖动
            autoHideMenuBar: true,
            webPreferences: {
                nodeIntegration: true,
                contextIsolation: false,
                enableRemoteModule: true,
                webSecurity: false,
            }
        })

        if (IS_LINUX && !DEBUG) {
            // const pth = path.join(__dirname, 'icon', 'favicon.png')
            const pth = path.join(__dirname, '../../',Config.pkg.linuxIcon);
            mainWin.setIcon(pth);
        }

        require('@electron/remote/main').initialize()
        require('@electron/remote/main').enable(mainWin.webContents);

        const {currVersionStr} = getCurrVersion()
        global.sharedObj =  { version : currVersionStr};

        mainWin.setSize(minimumSizeWidth, minimumSizeHeight)
        mainWin.setMovable(true)
        mainWin.maximize()
        mainWin.show()

        this._windows.set('main', mainWin);

        // 最终都是返回 http地址，远端 或者 本地http服务
        const uiPort = process.env.UI_SERVER_PORT || await getUsefulPort(portClient,55999);
        const url = await startUIService(uiPort);
        await mainWin.loadURL(url);

        // 通知渲染进程，agent服务端口
        const data = {uiPort, agentPort: port}
        logInfo(`send event to webpage, uiPort=${data.uiPort}, agentPort=${data.agentPort}`)
        mainWin.webContents.send(electronMsgUsePort, data);

        // 读取服务配置
        const confPath = path.join(os.homedir(), App, 'conf.yaml');
        if (!fs.existsSync(confPath)) {
            const data = {
                serverUrl: '',
            }
            const yamlStr = yaml.safeDump(data);
            fs.writeFileSync(confPath, yamlStr, 'utf8');
        }
        try {
            const doc = yaml.load(fs.readFileSync(confPath, 'utf8'));
            logInfo('read conf.yaml', doc.serverUrl);
            if (doc.serverUrl && doc.serverUrl.trim() != '') {
                mainWin.webContents.send(electronMsgServerUrl, {serverUrl: doc.serverUrl});
            }
        } catch (e) {
            logErr(e);
        }

        // 进程间通信逻辑
        ipcMain.on(electronMsg, (event, arg) => {
            logInfo('::::msg from renderer', JSON.stringify(arg))
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
                    logInfo('--', arg.action, arg.path);
                    if (arg.action == 'openInExplore')
                        this.openInExplore(arg.path)
                    else if (arg.action == 'openInTerminal')
                        this.openInTerminal(arg.path)
            }
        })
    }

    // 打开或者重新创建主窗口
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
        // this.buildAppMenu();
        this.openOrCreateWindow()
        this.setAboutPanel();

        // 使用默认的快捷键，和常用的快捷键有冲突
        // globalShortcut.register('CommandOrControl+D', () => {
        //     const mainWin = this._windows.get('main');
        //     mainWin.toggleDevTools()
        // })
        // check update,检查是否需要更新应用
        ipcMain.on(electronMsgUpdate, (event, arg) => {
            logInfo('update confirm from renderer', arg)
            const mainWin = this._windows.get('main');
            updateApp(arg.newVersion, mainWin)
        });

        // 一个小时检查一次更新
        setInterval(async () => {
            await checkUpdate(this._windows.get('main'))
        }, 1000*60*60);
    }

    async quit() {
        killAgent();
    }

    // 绑定 electron 事件
    bindElectronEvents() {
        app.on('activate', () => {
            logInfo('>> event: app activate');

            // this.buildAppMenu();

            // 在 OS X 系统上，可能存在所有应用窗口关闭了，但是程序还没关闭，此时如果收到激活应用请求，需要重新打开应用窗口并创建应用菜单。
            this.openOrCreateWindow()

            // On OS X it's common to re-create a window in the app when the
            // dock icon is clicked and there are no other windows open.
            // if (BrowserWindow.getAllWindows().length === 0) {
            //     createWindow();
            // }
        });

        // Quit when all windows are closed, except on macOS. There, it's common
        // for applications and their menu bar to stay active until the user quits
        // explicitly with Cmd + Q.
        app.on('window-all-closed', () => {
            logInfo(`>> event: app window-all-closed`)
            app.quit();
            // if (process.platform !== 'darwin') {
            //     app.quit();
            // }
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

    // 打开文件夹
    openInExplore(pth) {
        shell.showItemInFolder(pth);
    }

    // 打开终端
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

    // // 获取当前应用的主窗口
    // get windows() {
    //     return this._windows;
    // }

    // 设置关于应用面板
    setAboutPanel() {
        if (!app.setAboutPanelOptions) {
            return;
        }
        // todo 需要更新字段，待转成内部仓库后更新
        let version = Config.pkg.buildTime ? 'build at ' + new Date(Config.pkg.buildTime).toLocaleString() : ''
        version +=  DEBUG ? '[debug]' : ''
        app.setAboutPanelOptions({
            applicationName: Lang.string(Config.pkg.name) || Config.pkg.displayName,
            applicationVersion: Config.pkg.version,
            copyright: `${Config.pkg.copyright} ${Config.pkg.company}`,
            // credits: `Licence: ${Config.pkg.license}`,
            version: version
        });
    }

    // 构建应用菜单
    buildAppMenu() {
        logInfo('>> deeptest app: build application menu.');

        if(IS_WINDOWS_OS){
            Menu.setApplicationMenu(null);
        }
        return;

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
