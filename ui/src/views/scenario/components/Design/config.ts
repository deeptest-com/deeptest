/**
 * 场景编排配置相关
 * */

/**
 * 循环相关的迭代器
 * */
export const loopIteratorTypes = ['processor_loop_time', 'processor_loop_in', 'processor_loop_until', 'processor_loop_range', 'processor_data_default'];
/**
 * 仅显示禁用和删除的操作的类型，即叶子节点
 * */
export const onlyShowDisableAndDeleteTypes = [
    'processor_time_default',
    // cookie 相关
    'processor_cookie_get',
    'processor_cookie_set',
    'processor_cookie_clear',
    "processor_assertion_default",
    'processor_custom_code',
    'processor_print_default',
    'processor_variable_set',
    "processor_variable_clear",
    // 提取器相关
    'processor_extractor_boundary',
    'processor_extractor_jsonquery',
    'processor_extractor_htmlquery',
    'processor_extractor_xmlquery',
    'processor_extractor_regex',
    // 跳出循环也是叶子结点
    'processor_loop_break',
    // 请求也是叶子结点
    'processor_interface_default',
];

/**
 * 场景编排菜单配置
 * */
export const DESIGN_MENU_CONFIG = [
    {
        key: 'addInterface',
        title: ' 添加请求',
        icon: 'arrange-interface',
        hideInNodeTypes: ['processor_interface_default', ...onlyShowDisableAndDeleteTypes],
        children: [
            {
                key: 'add-child-interface-define',
                title: '导入接口定义',
                icon: 'interface',
            },
            {
                key: 'add-child-interface-case',
                title: '导入接口用例',
                icon: 'arrange-case',
            },
            {
                key: 'add-child-interface-diagnose',
                title: '导入快捷调试',
                icon: 'arrange-debug',
            },
            {
                key: 'add-child-interface-custom',
                title: '自定义请求',
                icon: 'arrange-http',
            },
            {
                key: 'add-child-interface-curl',
                title: 'cURL导入',
                icon: 'arrange-url',
            }
        ]
    },
    {
        key: 'addProcessor',
        title: '添加处理器',
        icon: 'arrange-control',
        hideInNodeTypes: [...onlyShowDisableAndDeleteTypes],
        children: [
            {
                key: 'processor_loop',
                title: '迭代',
                icon: 'arrange-loop',
                hideInNodeTypes: ['processor_interface_default'],
                children: [
                    {
                        title: '迭代次数',
                        key: 'processor_loop_time',
                        icon: 'arrange-count',
                    },
                    {
                        title: '迭代列表',
                        key: 'processor_loop_in',
                        icon: 'arrange-loop-list',
                    },
                    {
                        title: '迭代直到',
                        key: 'processor_loop_until',
                        icon: 'arrange-untils'
                    },
                    {
                        title: '迭代区间',
                        key: 'processor_loop_range',
                        icon: 'arrange-range',
                    },
                    {
                        title: '迭代循环',
                        key: 'processor_loop_break',
                        icon: 'arrange-return',
                        showInNodeTypes: [...loopIteratorTypes],
                    },
                ]
            },
            // 条件
            {
                key: 'processor_logic_if',
                title: '条件分支',
                icon: 'arrange-if',
                hideInNodeTypes: ['processor_interface_default'],
                // children: [
                //     {
                //         title: '如果',
                //         key: 'processor_logic_if',
                //         icon: 'arrange-logic-if',
                //         hideInNodeTypes: null,
                //     },
                //     {
                //         title: '否则',
                //         key: 'processor_logic_else',
                //         icon: 'arrange-logic-if',
                //         hideInNodeTypes: ['processor_interface_default'],
                //         showInNodeTypes: ['processor_logic_if'],
                //     }
                // ]
            },
            // 等待时间
            {
                key: 'processor_time_default',
                title: '等待时间',
                icon: 'arrange-wait',
                hideInNodeTypes: ['processor_interface_default'],
            },
            //    数据迭代
            {
                key: 'processor_data_default',
                title: '数据迭代',
                icon: 'arrange-data-loop',
                hideInNodeTypes: ['processor_interface_default'],
            },
            {
                key: 'processor_cookie',
                title: 'Cookie',
                icon: 'arrange-cookie',
                hideInNodeTypes: null,
                children: [
                    // {
                    //     title: '获取Cookie',
                    //     key: 'processor_cookie_get',
                    //     icon: 'arrange-add',
                    //     // showInNodeTypes: ['processor_interface_default'],
                    // },
                    {
                        title: '设置Cookie',
                        key: 'processor_cookie_set',
                        icon: 'setting-dark',
                        // showInNodeTypes: ['processor_interface_default'],
                    },
                    {
                        title: '清空Cookie',
                        key: 'processor_cookie_clear',
                        icon: 'arrange-clear',
                        // showInNodeTypes: ['processor_interface_default'],
                    },
                ]
            },
            // {
            //     key: 'processor_extractor',
            //     title: '提取器',
            //     icon: 'arrange-extractor',
            //     showInNodeTypes: ['processor_interface_default'],
            //     children: [
            //         {
            //             title: '边界提取器',
            //             key: 'processor_extractor_boundary',
            //             icon: 'arrange-extractor-boundary',
            //         },
            //         {
            //             title: 'JSON提取器',
            //             key: 'processor_extractor_jsonquery',
            //             icon: 'arrange-extractor-json',
            //         },
            //         {
            //             title: 'HTML提取器',
            //             key: 'processor_extractor_htmlquery',
            //             icon: 'arrange-extractor-html',
            //         },
            //         {
            //             title: 'XML提取器',
            //             key: 'processor_extractor_xmlquery',
            //             icon: 'arrange-extractor-xml',
            //         },
            //         // {
            //         //     title: '正则提取器',
            //         //     key: 'processor_extractor_regex',
            //         //     icon: 'arrange-extractor-regex',
            //         // },
            //         // arrange-extractor-xml.svg
            //     ]
            // },
            //   变量
            {
                key: 'processor_variable',
                title: '变量',
                icon: 'arrange-var',
                hideInNodeTypes: null,
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
                key: 'processor_print_default',
                title: '输出',
                icon: 'arrange-output',
                hideInNodeTypes: null,
            },
            //   断言
            {
                key: 'processor_assertion_default',
                title: '断言',
                icon: 'arrange-assert',
                hideInNodeTypes: null,
            },
            // 定制代码
            {
                key: 'processor_custom_code',
                title: '自定义代码',
                icon: 'arrange-code',
                hideInNodeTypes: null,
            },
        ]
    },
    {
        key: 'processor_group_default',
        title: '添加分组',
        icon: 'arrange-group',
        hideInNodeTypes: [...onlyShowDisableAndDeleteTypes],
    },
    //  分割线
    {
        key: 'divider',
        title: '分割线',
        hideInNodeTypes: ['processor_root_default', ...onlyShowDisableAndDeleteTypes],
    },
    // //   修改名字
    // {
    //     key: 'edit',
    //     title: '编辑',
    //     icon: 'edit',
    //     hideInNodeTypes: ['processor_root_default'],
    // },
    //    禁用
    {
        key: 'disable',
        title: '禁用',
        icon: 'arrange-disabled',
        hideInNodeTypes: ['processor_root_default'],
    },
    //    禁用
    {
        key: 'enable',
        title: '启用',
        icon: 'arrange-enable',
        hideInNodeTypes: ['processor_root_default'],
    },
    // 删除
    {
        key: 'remove',
        title: '删除',
        icon: 'arrange-delete',
        hideInNodeTypes: ['processor_root_default'],
    },
]

/**
 * 场景编排类型对应的图标映射
 * */

export const DESIGN_TYPE_ICON_MAP = {
    'processor_interface_default': 'interface',
    'add-child-interface-define': 'interface',
    'add-child-interface-case': 'arrange-case',
    'add-child-interface-diagnose': 'arrange-debug',
    'add-child-interface-custom': 'arrange-http',
    'add-child-interface-curl': 'arrange-url',
    'case': 'arrange-case',
    'diagnose': 'arrange-debug',
    'custom': 'arrange-http',
    'curl': 'arrange-url',

    'processor_group_default': 'arrange-group',

    "processor_variable_set": "arrange-var",
    "processor_variable_clear": "arrange-var",

    "processor_cookie_get": "arrange-cookie",
    'processor_cookie_set': 'arrange-cookie',
    'processor_cookie_clear': 'arrange-cookie',

    'processor_extractor_boundary': 'arrange-extractor',
    'processor_extractor_jsonquery': 'arrange-extractor',
    'processor_extractor_htmlquery': 'arrange-extractor',
    'processor_extractor_xmlquery': 'arrange-extractor',
    'processor_extractor_regx': 'arrange-extractor',


    "processor_loop_break": "arrange-return",
    "processor_loop_range": "arrange-range",
    "processor_loop_in": "arrange-loop-list",
    'processor_loop_time': 'arrange-count',
    'processor_loop_until': 'arrange-untils',

    'processor_logic_if': 'arrange-logic-if',
    'processor_logic_else': 'arrange-logic-if',

    'processor_time_default': 'arrange-wait',

    'processor_print_default': 'arrange-output',
    'processor_assertion_default': 'arrange-assert',
    'processor_custom_code': 'arrange-code',

    'processor_data_default': 'arrange-data-loop',
}


/**
 * 根据 菜单的key 对应的分类，用于保存场景编排时，根据分类保存到不同的字段
 * */
export const menuKeyMapToProcessorCategory = {
    'processor_group_default': 'processor_group',


    'processor_cookie_get': 'processor_cookie',
    'processor_cookie_set': 'processor_cookie',
    'processor_cookie_clear': 'processor_cookie',

    "add-child-interface-define": "processor_interface",
    "add-child-interface-case": "processor_interface",
    "add-child-interface-diagnose": "processor_interface",
    "add-child-interface-custom": "processor_interface",
    "add-child-interface-curl": "processor_interface",

    "processor_loop_time": "processor_loop",
    "processor_loop_until": "processor_loop",
    "processor_loop_range": "processor_loop",
    "processor_loop_break": "processor_loop",
    "processor_loop_in": "processor_loop",


    'processor_logic_if': 'processor_logic',
    'processor_logic_else': 'processor_logic',
    'processor_time_default': 'processor_timer',

    'processor_extractor_boundary': 'processor_extractor',
    'processor_extractor_jsonquery': 'processor_extractor',
    'processor_extractor_htmlquery': 'processor_extractor',
    'processor_extractor_xmlquery': 'processor_extractor',
    'processor_extractor_regx': 'processor_extractor',

    'processor_variable_set': 'processor_variable',
    'processor_variable_clear': 'processor_variable',

    'processor_print_default': 'processor_print',
    'processor_assertion_default': 'processor_assertion',

    // todo 定制代码对应的 类型确定
    'processor_custom_code': 'processor_custom_code',


    // 数据处理
    'processor_data_default': 'processor_data',

}


/**
 * 展示目录竖线的场景编排类型
 * */
export const showLineScenarioType = [
    // 接口
    'processor_interface_default',
    'processor_group_default',
    'processor_loop_time',
    'processor_loop_until',
    'processor_loop_range',
    'processor_loop_in',
    'processor_logic_if',
    // 数据处理
    'processor_data_default',
]

/**
 * 各节点类型对应的文案
 * */
export const scenarioTypeMapToText = {
    'processor_interface_default': '接口定义',

    'define': '接口定义导入',
    'case': '接口用例导入',
    'diagnose': '接口调试导入',
    'custom': '自定义请求导入',
    'curl': 'cURL导入',

    'processor_group_default': '分组',
    'processor_time_default': '定时器',
    'processor_print_default': '输出',
    'processor_assertion_default': '断言',
    'processor_data_default': '数据迭代',
    'add-child-interface-define': '接口',
    'add-child-interface-case': '接口',
    'add-child-interface-diagnose': '接口',
    'add-child-interface-custom': '接口',
    'add-child-interface-curl': '接口',
    'processor_cookie_get': '获取 Cookie',
    'processor_cookie_set': '设置 Cookie',
    'processor_cookie_clear': '清除 Cookie',
    'processor_loop_break': '跳出迭代',
    'processor_loop_time': '迭代次数',
    'processor_loop_until': '迭代直到',
    'processor_loop_range': '迭代区间',
    'processor_loop_in': '迭代列表',
    'processor_logic_if': 'if',
    'processor_logic_else': 'else',

    'processor_extractor_boundary': '边界提取',
    'processor_extractor_jsonquery': 'JSON 提取',
    'processor_extractor_htmlquery': 'HTML 提取',
    'processor_extractor_xmlquery': 'XML 提取',
    'processor_variable_set': '设置变量',
    'processor_variable_clear': '清除变量',
    'processor_custom_code': '自定义代码',
}


/**
 * 场景绑定接口的类型对应的文案
 * */
export const scenarioTypeMapToBindText = {

    'define': '绑定接口',
    'case': '绑定接口用例',
    'diagnose': '绑定快捷调试',
    // 'custom': '自定义请求导入',
    // 'curl': 'cURL导入',

}
