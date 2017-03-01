export var CONSTANT: any = {
    _SERVICE_URL_DEV: 'http://localhost:8080/testspace/',
    _SERVICE_URL_PRODUCTION: 'http://localhost:8080/testspace/',

    SERVICE_URL: undefined,
    API_URL: undefined,

    API_PATH: 'api/client/v1/',
    UPLOAD_URI: 'uploadSingle',

    TOKEN: '',
    PROFILE: {},

    PROFILE_KEY: 'cn.linkr.events.profile',
    PROFILE_EXPIRE: 'cn.linkr.events.expire',

    EventStatus: [{not_start: '未开始'}, {register: '报名'}, {sign: '签到'},
                             {in_progress: '进行中'}, {end: '已结束'}, {cancel: '取消'}]
};

