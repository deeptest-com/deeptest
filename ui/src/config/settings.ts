import { RoutesDataItem } from "@/utils/routes";

/**
 * 站点配置
 * @author LiQingSong
 */
export interface SettingsType {
    /**
     * 站点名称
     */
    siteTitle: string;

    /**
     * 顶部菜单开启
     */
    topNavEnable: boolean;

    /**
     * 头部固定开启
     */
    headFixed: boolean;

    /**
     * tab菜单开启
     */
     tabNavEnable: boolean;

     /**
     * 站点首页路由
     */
    homeRouteItem: RoutesDataItem;

    /**
     * 站点本地存储Token 的 Key值
     */
    siteTokenKey: string;

    /**
     * 站点本地存储当前项目 的 Key值
     */
    currProjectId: string;
    currServeId: string,

    settings: string;
    expandedKeys: string;
    selectedKey: string;
    skippedVersion: string;
    ignoreUtil: string,

    eventNotify,
    eventWebSocketConnStatus: string,
    eventWebSocketMsg: string,

    eventEditorAction: string,
    eventTypeContainerHeightChanged: string,
    eventTypeFormat: string,

    eventConditionSave: string,

    eventVariableSelectionStatus: string,
    eventVariableSelectionResult: any,
    webSocketRoom: string,
    electronMsg: string,
    electronMsgReplay: string,
    electronMsgUpdate: string,
    electronMsgDownloading: string,

    /**
     * Ajax请求头发送Token 的 Key值
     */
    ajaxHeadersTokenKey: string;

    /**
     * Ajax返回值不参加统一验证的api地址
     */
    ajaxResponseNoVerifyUrl: string[];

    /**
     * iconfont.cn 项目在线生成的 js 地址
     */
    iconfontUrl: string[];

    /**
     * 测试计划-执行列表
     */
    eventGetPlansReports: string;
    eventGetPlanDetail: string;

    /**
     * pane resize 
     */
    paneResizeTop: string;
}

const settings: SettingsType = {
    siteTitle: process.env.VUE_APP_DEPLOY_ENV === 'ly' ? 'leyanapi.com' : 'deeptest.com',
    topNavEnable: true,
    headFixed: true,
    tabNavEnable: false,
    homeRouteItem: {
        icon: 'interface',
        title: 'interface',
        path: '/endpoint/index',
        component: ()=> import('@/views/endpoint/index.vue')
    },
    siteTokenKey: 'admin_antd_vue_token',
    currProjectId: 'curr_project_id',
    currServeId: 'curr_serve_id',

    settings: 'settings',
    expandedKeys: 'deeptest-expandedKeys',
    selectedKey: 'deeptest-selectedKey',
    skippedVersion: 'skippedVersion',
    ignoreUtil: 'ignoreUtil',

    eventNotify: 'eventNotify',
    eventWebSocketConnStatus: 'eventWebSocketStatus',
    eventWebSocketMsg: 'eventWebSocketMsg',

    eventEditorAction: 'eventEditorAction',
    eventTypeContainerHeightChanged: 'heightChanged',
    eventTypeFormat: 'format',

    eventConditionSave: '',

    eventVariableSelectionStatus: 'eventVariableSelectionStatus',
    eventVariableSelectionResult: 'eventVariableSelectionResult',
    webSocketRoom: 'webSocketRoom',
    electronMsg: 'electronMsg',
    electronMsgReplay: 'electronMsgReplay',
    electronMsgUpdate: 'electronMsgUpdate',
    electronMsgDownloading: 'electronMsgDownloading',

    // ajaxHeadersTokenKey: 'x-token',
    ajaxHeadersTokenKey: 'Authorization',
    ajaxResponseNoVerifyUrl: [
        '/user/login', // 用户登录
    ],
    iconfontUrl: [],

    // 测试计划相关的
    eventGetPlanDetail: 'eventGetPlanDetail',
    eventGetPlansReports: 'eventGetPlansReports', 

    // pane Resizer
    paneResizeTop: 'paneResizeTop',
};

export default settings;
