import {message, Modal, notification} from "ant-design-vue";
import {createVNode} from 'vue';
import {ExclamationCircleOutlined} from "@ant-design/icons-vue";
import {NotificationKeyCommon} from "@/utils/const";

export function notifySuccess(title, desc = '', key = NotificationKeyCommon){
    notification.success({
        key: key,
        message: title,
        description: desc,
    });
}

export function notifyError(title, desc = '', key = NotificationKeyCommon){
    notification.error({
        key: key,
        message: title,
        description: desc,
    });
}

export function notifyWarn(title, desc = '', key = NotificationKeyCommon){
    notification.warning({
        key: key,
        message: title,
        description: desc ? desc : '',
    });
}