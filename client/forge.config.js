/**
 *  打包ly的包时，需要修改下面的引用 package 路径
 * */
// import pkg from './src/app/package.json';
const pkg = require('./src/app/package-ly.json');

module.exports = {
    // Windows and macOS only 设置图标
    packagerConfig: {
        "name": pkg?.name || "deeptest",
        asar: true,
        "icon": pkg?.icon || "./icon/favicon",  // no file extension required
        extraResource: [
            './bin',
            './ui',
            './lang',
        ]
    },
    // electronPackagerConfig: {
    //     "name": "deeptest",
    //     "icon": "./ui/favicon.ico"
    // },
    makers: [
        // 使用 Electron Forge 为你的 Electron 应用程序创建 Windows 安装程序
        {
            name: '@electron-forge/maker-squirrel',
            config: {
                name: pkg?.name || 'deeptest'
            }
        },
        {
            name: '@electron-forge/maker-zip',
            platforms: [
                'darwin'
            ]
        },
        // Linux 下  需要显式 设置 icon
        {
            name: '@electron-forge/maker-deb',
            config: {
                options: {
                    icon: pkg?.linuxIcon || './icon/favicon.png'
                }
            }
        },
        {
            name: '@electron-forge/maker-rpm',
            config: {}
        }
    ],
    plugins: [
        {
            name: '@electron-forge/plugin-auto-unpack-natives',
            config: {},
        },
        {
            name: '@electron-forge/plugin-webpack',
            config: {
                mainConfig: './webpack.main.config.js',
                renderer: {
                    config: './webpack.renderer.config.js',
                    // 其实以下配置没有用，因为默认都是自己远程加载的，或者本地启动 Express 服务加载的
                    // 但不配置，打包时会一直 pendding，所以这里配置一下
                    // TODO  找时间研究一下
                    entryPoints: [
                        {
                            html: './src/entry/index.html',
                            js: './src/entry/renderer.js',
                            name: 'main_window',
                            preload: {
                                js: './src/entry/preload.js',
                            },
                        },
                    ],
                },
            },
        },
        /**
         * 在 Webpack 中使用 Electron 时，支持原生模块的最简单方法是将它们添加到 Webpack 的外部配置中。
         * 这样，Webpack 就会通过 require() 从 node_modules 中加载这些模块
         * */
        {
            name: '@timfish/forge-externals-plugin',
            config: {
                externals: ['@electron/remote'],
                includeDeps: true,
            },
        }
        // [
        //     "@timfish/forge-externals-plugin",
        //     {
        //         "externals": ["@electron/remote"],
        //         "includeDeps": true
        //     }
        // ]

    ]
}
