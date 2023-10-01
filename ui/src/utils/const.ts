import {ref} from "vue";

// export const NotificationKeyRequest = 'key_request'
export const NotificationKeyCommon = 'key_common'

export const EventNodeIdJson = 'deeptest-event-node-json'
export const EventNameJson = 'deeptest-event-from-chrome-json'

export const EventNodeIdImg = 'deeptest-event-node-img'
export const EventNameImg = 'deeptest-event-from-chrome-img'

export const ScopeDeeptest = 'scope-com-deeptest'
export const ActionRecordStart = 'recordStart'
export const ActionRecordedMsg = 'recordMsg'

export const Cache_Key_Server_Url = 'dp-cache-server-url'
export const Cache_Key_Agent_Url = 'dp-cache-agent-url'
export const Cache_Key_Agent_Local_Port = 'dp-cache-agent-local-port'
export const Cache_Key_Agent_Value = 'dp-cache-agent-value'

export const MonacoOptions = {
    colorDecorators: true,
    lineHeight: 24,
    tabSize: 2,
    autoIndent: true,
    formatOnPaste: true,
    formatOnType: true
}

export const pattern = {
    // http://emailregex.com/
    email: /^(([^<>()\\[\]\\.,;:\s@"]+(\.[^<>()\\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]+\.)+[a-zA-Z\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]{2,}))$/,
    // url: new RegExp(
    //   '^(?!mailto:)(?:(?:http|https|ftp)://|//)(?:\\S+(?::\\S*)?@)?(?:(?:(?:[1-9]\\d?|1\\d\\d|2[01]\\d|22[0-3])(?:\\.(?:1?\\d{1,2}|2[0-4]\\d|25[0-5])){2}(?:\\.(?:[0-9]\\d?|1\\d\\d|2[0-4]\\d|25[0-4]))|(?:(?:[a-z\\u00a1-\\uffff0-9]+-*)*[a-z\\u00a1-\\uffff0-9]+)(?:\\.(?:[a-z\\u00a1-\\uffff0-9]+-*)*[a-z\\u00a1-\\uffff0-9]+)*(?:\\.(?:[a-z\\u00a1-\\uffff]{2,})))|localhost)(?::\\d{2,5})?(?:(/|\\?|#)[^\\s]*)?$',
    //   'i',
    // ),
    hex: /^#?([a-f0-9]{6}|[a-f0-9]{3})$/i,

    cron:/(@(annually|yearly|monthly|weekly|daily|hourly|reboot))|(@every (\d+(ns|us|µs|ms|s|m|h))+)|((((\d+,)+\d+|(\d+(\/|-)\d+)|\d+|\*) ?){5,7})/,

    alphanumeric: /^[a-z][a-z0-9]*$/i
};