/**
 * 场景编排配置相关
 * */


/**
 * 场景编排菜单配置
 * */
export const DESIGN_MENU_CONFIG = [
    {
        key: 'addRequest',
        title: ' 添加请求',
        icon: 'arrange-interface',
        children: [
            {
                key: 'importInterface',
                title: '导入接口定义',
                icon: 'interface',
            },
            {
                key: 'importInterfaceCase',
                title: '导入接口用例',
                icon: 'arrange-case',
            },
            {
                key: 'importQuickDebug',
                title: '导入快捷调试',
                icon: 'arrange-debug',
            },
            {
                key: 'customRequest',
                title: '自定义请求',
                icon: 'arrange-http',
            },
            {
                key: 'importCurl',
                title: 'cURL导入',
                icon: 'arrange-url',
            }
        ]
    },
    {
        key: 'addHandler',
        title: '添加处理器',
        icon: 'arrange-control',
        children: [
            {
                key: 'addLoop',
                title: '循环',
                icon: 'arrange-var',
                children: [
                    {
                        title: '循环次数',
                        key: 'loopCount',
                        icon: 'arrange-count',
                    },
                    {
                        title: '数据送代',
                        key: 'loopData',
                        icon: 'arrange-database',
                    },
                    {
                        title: '循环直到',
                        key: 'loopUntil',
                        icon: 'arrange-untils'
                    },
                    {
                        title: '循环区间',
                        key: 'loopRange',
                        icon: 'arrange-range',
                    },
                    {
                        title: '跳出循环',
                        key: 'loopBreak',
                        icon: 'arrange-return',
                    },
                ]
            },
            // 条件
            {
                key: 'addCondition',
                title: '条件',
                icon: 'arrange-if',
            },
            // 等待时间
            {
                key: 'addWait',
                title: '等待时间',
                icon: 'arrange-wait',
            },
            //    数据迭代
            {
                key: 'addData',
                title: '数据迭代',
                icon: 'arrange-data-loop',
            },
            {
                key: 'addCookie',
                title: 'Cookie',
                icon: 'arrange-cookie',
                children: [
                    {
                        title: '添加Cookie',
                        key: 'addCookie',
                        icon: 'arrange-add',
                    },
                    {
                        title: '删除Cookie',
                        key: 'deleteCookie',
                        icon: 'arrange-delete',
                    },
                    {
                        title: '清空Cookie',
                        key: 'clearCookie',
                        icon: 'arrange-clear',
                    },
                ]
            },
            //   变量
            {
                key: 'addVariable',
                title: '变量',
                icon: 'arrange-var',
                children: [
                    {
                        title: '添加变量',
                        key: 'addVariable',
                        icon: 'arrange-add',
                    },
                    {
                        title: '删除变量',
                        key: 'deleteVariable',
                        icon: 'arrange-delete',
                    }
                ]
            },
            //  输出
            {
                key: 'addOutput',
                title: '输出',
                icon: 'arrange-output',
            },
            //   断言
            {
                key: 'addAssert',
                title: '断言',
                icon: 'arrange-assert',
            },
            //     定制代码
            {
                key: 'addCustomCode',
                title: '定制代码',
                icon: 'arrange-code',
            },
        ]
    },
    {
        key: 'addGroup',
        title: '添加分组',
        icon: 'arrange-group',
    },
//    分割线
    {
        key: 'divider',
        title: '分割线',
    },
    //    禁用
    {
        key: 'addDisable',
        title: '禁用',
        icon: 'arrange-disable',
    },
    // 删除
    {
        key: 'addDelete',
        title: '删除',
        icon: 'arrange-delete',
    },
]


