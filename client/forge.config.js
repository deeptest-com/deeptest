module.exports = {
    packagerConfig: {
        "name": "deeptest",
        asar: true,
        "icon": "./icon/favicon",
        extraResource: [
            './bin',
            './ui',
            './lang',
        ]
    },
    electronPackagerConfig: {
        "name": "deeptest",
        "icon": "./ui/favicon.ico"
    },
    makers: [
        {
            name: '@electron-forge/maker-squirrel',
            config: {
                name: 'deeptest' //todo 从package.json中获取
            }
        },
        {
            name: '@electron-forge/maker-zip',
            platforms: [
                'darwin'
            ]
        },
        {
            name: '@electron-forge/maker-deb',
            config: {}
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
                    // 其实以下配置可以不用写，因为默认都是自己远程加载的，或者本地启动 Express 服务加载的
                    entryPoints: [
                        {
                            html: './src/index.html',
                            js: './src/renderer.js',
                            name: 'main_window',
                            preload: {
                                js: './src/preload.js',
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
