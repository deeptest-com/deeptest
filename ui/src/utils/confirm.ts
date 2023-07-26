import {message, Modal} from "ant-design-vue";

export function confirmToDelete(title, content, callback, confirmText?, cancelText?){
    Modal.confirm({
        okType: 'danger',
        title: title,
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

export function confirmToSave(callback, title?, content?, confirmText?, cancelText?){
    Modal.confirm({
        okType: 'danger',
        title: title || '有修改内容未保存',
        content: content || '是否放弃未保存的修改？',
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