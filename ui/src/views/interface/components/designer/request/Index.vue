<template>
  <div id="request-main">
    <RequestSender :onSend="sendRequest"></RequestSender>
    <RequestConfig></RequestConfig>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, onMounted, PropType, Ref} from "vue";
import {useI18n} from "vue-i18n";
import {Form, message} from 'ant-design-vue';
import {useStore} from "vuex";
import {StateType} from "@/views/interface/store";
import RequestSender from './Sender.vue';
import RequestConfig from './Config.vue';

const useForm = Form.useForm;

interface InterfaceRequestSetupData {
  requestData: ComputedRef;
  sendRequest: (e) => void;
}

export default defineComponent({
  name: 'InterfaceRequest',
  props: {
  },
  components: {
    RequestSender, RequestConfig,
  },
  setup(props): InterfaceRequestSetupData {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const requestData = computed<any>(() => store.state.Interface.requestData);

    const sendRequest = (e) => {
      console.log('sendRequest')
    };

    onMounted(() => {
      console.log('onMounted')
    })

    return {
      requestData,
      sendRequest,
    }
  }
})
</script>

<style lang="less" scoped>
#request-main {
  height: 100%;
}
</style>