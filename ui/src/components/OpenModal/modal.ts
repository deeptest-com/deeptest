import { createVNode, render as vueRender } from "vue";
import { Modal as AntModal } from "ant-design-vue";

export default function openModal(component, props, config, appContext) {
    props.onClose = async function () {
        close();
    };
    config.wrapClassName = 'dp-modal-confirm'

    const container = document.createDocumentFragment();
    const _contentVnode = createVNode(component, props);
    _contentVnode.appContext = {...appContext}

    const metadata = Object.create({
        footer: false,
        visible: true,
        ...props,
        ...config,
    });

    const vm = createVNode(AntModal, metadata, () => _contentVnode);
    vm.appContext = {...appContext}

    function close() {
        metadata.visible = false;
        update(metadata);
    }
    function update(config) {
        if (vm.component) {
            Object.assign(vm.component.props, config);
            vm.component.update();
        }
    }
    // function destroy() {
    //     if (vm) {
    //         vueRender(null, container);
    //         vm = null;
    //     }
    // }
    // metadata.onCancel = async function (...arg) {
    //     await config.onCancel?.(...arg);
    //     close();
    // };
    // metadata.onOk = async function () {
    //     if (!(config.onOk instanceof Function)) {
    //         close();
    //         return;
    //     }
    //     const result = config.onOk();
    //     if (!(result instanceof Promise)) {
    //         close();
    //         return;
    //     }
    //     update({ confirmLoading: true });
    //     return result
    //         .then(() => {
    //             update({ confirmLoading: false });
    //             close();
    //         })
    //         .catch(() => {
    //             update({ confirmLoading: false });
    //         });
    // };

    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    vueRender(vm, container);

    return {
        ..._contentVnode,
        close,
        update,
        // destroy,
    };
}