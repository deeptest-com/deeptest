import {ProcessorExtractor} from "@/utils/enum";

export default {
    '': '',
    'empty': '',
    'lang': 'zh-CN',
    'app.global.menu.notfound': 'Not Found',
    'app.global.form.validatefields.catch': '验证不通过，请检查输入',

    'start': '开始',
    'in_progress': '执行中',
    'end': '结束',

    'user-info': '个人信息',
    'logout': '退出',

    'home': '首页',
    'workbench': '工作台',
    'docs': '使用文档',
    'menu': '菜单',

    'project': '项目',
    'project.management': '项目管理',
    'project.list': '项目列表',
    'project.edit': '项目编辑',
    'project.members': '项目成员',
    'project.invite': '邀请成员',

    'endpoint-management': '接口定义',
    'endpoint-docs': '接口文档',
    'diagnose': '接口调试',
    'projectSetting': '项目设置', // 新版 接口模块
    'projectSetting.enviroment': '环境管理',
    'projectSetting.datapool': '数据池',
    'projectSetting.service': '服务管理',
    'envsetting.var': '全局变量',
    'envsetting.params': '全局参数',
    'envsetting.envdetail': '环境详情',
    'scenario': '测试开发',
    'scenario.edit': '场景编辑',
    'scenario.design': '场景设计',
    'scenario.exec': '执行测试场景',

    'plan': '测试计划',
    'plan.edit': '计划编辑',
    'plan.exec': '执行测试计划',

    'report': '测试报告',
    'reportNew': '报告New',
    'report.detail': '报告细节',
    'reportNew.detail': '报告细节New',

    'profile': '个人信息',
    'invite': '邀请用户',
    'message': '消息',

    'user': '用户',
    'user.management': '用户管理',
    'user.list': '用户列表',

    'mock.oauth2.callback': 'OAuth2回调模拟',

    'json': 'JSON',
    'xml': 'XML',
    'html': 'HTML',
    'text': 'TEXT',
    'plaintext': 'TEXT',

    'header':  '响应头',
    'body':  '响应体',

    'header_en':  'Header',
    'cookie_en':  'Cookie',
    'query_en':  'Query',
    'body_en':  'Body',

    'responseStatus': '响应码',
    'responseHeader': '响应头',
    'responseBody': '响应体',
    'extractor': '提取器',
    'judgement': '表达式',

    'fulltext': '全文本',
    'jsonquery':  'JSON查询',
    'htmlquery':  'HTML查询',
    'xmlquery':  'XML查询',
    'regx':  '正则表达式',
    'boundary':  '边界选择器',

    'equal':  '=',
    'notEqual':  '!=',
    'greaterThan':  '>',
    'lessThan':  '<',
    'greaterThanOrEqual':  '>=',
    'lessThanOrEqual':  '<=',

    'contain':  '包含',
    'notContain': '不包含',

    // 处理器分类
    'processor_thread':  '线程',
    'processor_group':  '分组',
    'processor_logic':  '逻辑',
    'processor_loop':  '循环',
    'processor_timer':  '计时器',
    'processor_variable': '变量',
    'processor_assertion':  '断言',
    'processor_extractor':  '提取器',
    'processor_cookie':  'Cookie',
    'processor_data':  '数据',

    // 处理器类型
    'processor_thread_default':  '线程',
    'processor_group_default':  '分组',
    'processor_time_default':  '计时器',
    'processor_print_default':  '输出',

    'processor_logic_if':  '如果',
    'processor_logic_else':  '否则',

    'processor_loop_time':  '迭代次数',
    'processor_loop_until':  '迭代直到',
    'processor_loop_in':  '迭代列表',
    'processor_loop_range':  '迭代区间',
    'processor_loop_break':  '跳出迭代',

    'processor_variable_get': '获取变量',
    'processor_variable_set': '设置变量',
    'processor_variable_clear': '清除变量',

    'processor_assertion_default': '断言',
    // 'processor_assertion_equal': '等于',
    // 'processor_assertion_not_equal': '不等于',
    // 'processor_assertion_contain': '包含',
    // 'processor_assertion_not_contain': '不包含',
    // 'processor_assertion_greater_than': '大于',
    // 'processor_assertion_less_than': '小于',
    // 'processor_assertion_greater_than_or_equal': '大于等于',
    // 'processor_assertion_less_than_or_equal': '小于等于',

    'processor_extractor_boundary':  '边界提取器',
    'processor_extractor_jsonquery':  'JSON提取器',
    'processor_extractor_htmlquery':  'HTML提取器',
    'processor_extractor_xmlquery':  'XML提取器',
    'processor_extractor_regx':  '正则表达式提取器',

    'processor_cookie_get':  '获取Cookie',
    'processor_cookie_set':  '设置Cookie',
    'processor_cookie_clear':  '清除Cookie',

    'processor_data_text':  '文本数据',
    'processor_data_excel':  'Excel数据',
    'processor_data_zendata':  'ZenData数据',

    'pass': '通过',
    'fail': '失败',
    'skip': '跳过',
    'block': '阻塞',
    'unknown': '位置',

    'tips_expression': '可以引用形如${name}的变量。',
    'tips_expression_bool': '可以引用形如${name}的变量，需返回一个布尔值。',

    'extractor_err': '提取失败',
    'extractor_err_short': '失败',
    'content_err': '内容错误',

    'http_1000': '无法连接到服务',

    'biz_401': '请重新登录',
    'biz_403': '权限不足',
    'biz_undefined': '未知错误',
    'biz_2000': '参数错误',
    'biz_3000': '请求失败',
    'biz_4000': '系统错误',
    'biz_10100': '同名记录已存在',
    'biz_10200': '用户名已存在',
    'biz_10300': '邮箱已存在',
    'biz_10500': '两次密码必须一样',

}
