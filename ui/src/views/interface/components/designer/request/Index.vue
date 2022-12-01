<template>
  <div id="request-main">
    <RequestInvocation :onSend="invoke" :onSave="saveInterface"></RequestInvocation>
    <RequestConfig></RequestConfig>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, onMounted, PropType, Ref} from "vue";
import {useI18n} from "vue-i18n";
import {Form, message, notification} from 'ant-design-vue';
import {useStore} from "vuex";
import {StateType} from "@/views/interface/store";
import RequestInvocation from './Invocation.vue';
import RequestConfig from './Config.vue';
import {Interface} from "@/views/interface/data";
import {NotificationKeyCommon} from "@/utils/const";

const useForm = Form.useForm;

export default defineComponent({
  name: 'InterfaceRequest',
  props: {
  },
  components: {
    RequestInvocation, RequestConfig,
  },
  setup(props) {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

    const invoke = (data) => {
      console.log('invoke', data)
      store.dispatch('Interface/invoke', data)
    };

    const saveInterface = (data) => {
      console.log('saveInterface', data)
      store.dispatch('Interface/saveInterface', data).then((res) => {
        if (res === true) {
          notification.success({
            key: NotificationKeyCommon,
            message: `保存成功`,
          });
        } else {
          notification.success({
            key: NotificationKeyCommon,
            message: `保存失败`,
          });
        }
      })
    };

    onMounted(() => {
      console.log('onMounted')
    })

    return {
      interfaceData,
      invoke,
      saveInterface,
    }
  }
})
</script>

<style lang="less" scoped>
#request-main {
  height: 100%;
}
</style>