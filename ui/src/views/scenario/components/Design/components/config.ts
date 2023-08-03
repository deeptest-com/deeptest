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
                key: 'addVariable',
                title: '循环',
                icon: 'arrange-variable',
                //循环次数
                // 数据送代
                // 循环直到
                // 循环区间
                // 跳出循环
                children: [
                    {
                        title: '循环次数',
                        key: 'loopCount',
                        icon: 'arrange-count',
                    },
                    {
                        title: '数据送代',
                        key: 'loopData',
                        icon: 'arrange-data',
                    },
                    {
                        title: '循环直到',
                        key: 'loopUntil',
                        icon: 'arrange-until',
                    },
                    {
                        title: '循环区间',
                        key: 'loopRange',
                        icon: 'arrange-range',
                    },
                    {
                        title: '跳出循环',
                        key: 'loopBreak',
                        icon: 'arrange-break',
                    },
                ]
            },
            // 条件
            {
                key: 'addCondition',
                title: '条件',
                icon: 'arrange-condition',
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
                icon: 'arrange-data',
            },
            // Cookie
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
            //

        ]
    },
    {
        key: 'addGroup',
        title: '添加分组',
        icon: 'arrange-group',
    }
]


