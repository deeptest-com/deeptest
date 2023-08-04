/**
 * 场景编排配置相关
 * */


/**
 * 场景编排菜单配置
 * */
export const DESIGN_MENU_CONFIG = [
    {
        key: 'processor_interface',
        title: ' 添加请求',
        icon: 'arrange-interface',
        children: [
            {
                key: 'add-child-interface-fromDefine',
                title: '导入接口定义',
                icon: 'interface',
            },
            {
                key:'add-child-interface-fromTest',
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
                key: 'processor_loop',
                title: '循环',
                icon: 'arrange-var',
                children: [
                    {
                        title: '循环次数',
                        key: 'processor_loop_time',
                        icon: 'arrange-count',
                    },
                    // {
                    //     title: '数据送代',
                    //     key: 'loopData',
                    //     icon: 'arrange-database',
                    // },
                    {
                        title: '循环直到',
                        key: 'processor_loop_until',
                        icon: 'arrange-untils'
                    },
                    {
                        title: '循环区间',
                        key: 'processor_loop_range',
                        icon: 'arrange-range',
                    },
                    {
                        title: '跳出循环',
                        key: 'processor_loop_break',
                        icon: 'arrange-return',
                    },
                ]
            },
            // 条件
            {
                key: 'processor_logic',
                title: '条件',
                icon: 'arrange-if',
            },
            // 等待时间
            {
                key: 'processor_timer',
                title: '等待时间',
                icon: 'arrange-wait',
            },
            //    数据迭代
            {
                key: 'processor_data',
                title: '数据迭代',
                icon: 'arrange-data-loop',
            },
            {
                key: 'processor_cookie',
                title: 'Cookie',
                icon: 'arrange-cookie',
                children: [
                    {
                        title: '添加Cookie',
                        key: 'processor_cookie_get',
                        icon: 'arrange-add',
                    },
                    {
                        title: '删除Cookie',
                        key: 'processor_cookie_set',
                        icon: 'arrange-delete',
                    },
                    {
                        title: '清空Cookie',
                        key: 'processor_cookie_clear',
                        icon: 'arrange-clear',
                    },
                ]
            },
            //   变量
            {
                key: 'processor_variable',
                title: '变量',
                icon: 'arrange-var',
                children: [
                    {
                        title: '添加变量',
                        key: 'processor_variable_set',
                        icon: 'arrange-add',
                    },
                    {
                        title: '删除变量',
                        key: 'processor_variable_clear',
                        icon: 'arrange-delete',
                    }
                ]
            },
            //  输出
            {
                key: 'processor_print',
                title: '输出',
                icon: 'arrange-output',
            },
            //   断言
            {
                key: 'processor_assertion_default',
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
        key: 'disabled',
        title: '禁用',
        icon: 'arrange-disable',
    },
    // 删除
    {
        key: 'remove',
        title: '删除',
        icon: 'arrange-delete',
    },
]

/**
 * 场景编排类型对应的图标映射
 * */

export const DESIGN_TYPE_ICON_MAP = {
    'processor_group_default': 'arrange-group',
    'processor_loop_time': 'arrange-count',
    'processor_loop_until': 'arrange-untils',
    "processor_variable_set": "arrange-var",
}
