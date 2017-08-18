export var CONSTANT: CONSTANT_INTERFACE = {
  _SERVICE_URL_DEV: 'http://localhost:8080/platform/',
  _SERVICE_URL_PRODUCTION: 'http://172.16.90.66:8080/platform/',

  SERVICE_URL: undefined,
  API_URL: undefined,

  API_PATH: 'api/client/v1/',
  UPLOAD_URI: 'uploadSingle',

  TOKEN: undefined,
  ORG_ID: undefined,

  PROFILE: undefined,
  MY_ORGS: undefined,
  CURRENT_PROJECT: {id: null, name: null},
  RECENT_PROJECTS: undefined,
  CUSTOM_FIELD_FOR_PROJECT: undefined,

  TOKEN_KEY: 'com.ngtesting.token',
  TOKEN_EXPIRE: 'com.ngtesting.expire',

  ExeStatus: {'not_start': '未开始', 'in_progress': '执行中', 'end': '已结束'},
  EntityDisabled: {'false': '启用', 'true': '归档', '': '所有'},

  ScreenSize: {h: 0, w: 0},
  DebounceTime: 500
};

export interface CURRENT_PROJECT_INTERFACE {
  id: number,
  name: string
}

export interface CONSTANT_INTERFACE {
  _SERVICE_URL_DEV: string,
  _SERVICE_URL_PRODUCTION: string,

  SERVICE_URL: string,
  API_URL: string,

  API_PATH: string,
  UPLOAD_URI: string,

  TOKEN: string,
  ORG_ID: number,

  PROFILE: string,
  MY_ORGS: any[],
  CURRENT_PROJECT: CURRENT_PROJECT_INTERFACE,
  RECENT_PROJECTS: any[],
  CUSTOM_FIELD_FOR_PROJECT: any,

  TOKEN_KEY: string,
  TOKEN_EXPIRE: string,

  ExeStatus: any,
  EntityDisabled: any,

  ScreenSize: any,
  DebounceTime: number
}

