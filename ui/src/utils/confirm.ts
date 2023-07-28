import {message, Modal} from "ant-design-vue";
import {createVNode} from 'vue';
import {ExclamationCircleOutlined} from "@ant-design/icons-vue";

export function confirmToDelete(title, content, callback, confirmText?, cancelText?){
    Modal.confirm({
        okType: 'danger',
        title: title,
        icon: createVNode(ExclamationCircleOutlined),
        content: content,
        okText: () => confirmText?confirmText:'确定',
        cancelText: () => cancelText?cancelText:'取消',
        onOk: async () => {
            if (callback) callback()
        },
        onCancel() {
            console.log('Cancel');
        },
    });
}
