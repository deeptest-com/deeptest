export var CONSTANT: any = {
    _SERVICE_URL_DEV: 'http://localhost:8080/events/',
    _SERVICE_URL_PRODUCTION: 'http://localhost:8080/events/',

    SERVICE_URL: undefined,
    API_URL: undefined,

    API_PATH: 'api/admin/v1/',
    UPLOAD_URI: 'uploadSingle',

    TOKEN: '',
    COOKIE_KEY: 'cn.linkr.events.token',

    EventStatus: [{not_start: '未开始'}, {register: '报名'}, {sign: '签到'},
                             {in_progress: '进行中'}, {end: '已结束'}, {cancel: '取消'}]
};

