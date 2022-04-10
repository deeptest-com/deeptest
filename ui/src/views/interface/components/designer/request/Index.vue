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
import {test} from "@/views/interface/service";
import {Interface} from "@/views/interface/data";

const useForm = Form.useForm;

export default defineComponent({
  name: 'InterfaceRequest',
  props: {
  },
  components: {
    RequestSender, RequestConfig,
  },
  setup(props) {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

    const sendRequest = (e) => {
      console.log('sendRequest')
      store.dispatch('Interface/test', interfaceData.value)
    };

    onMounted(() => {
      console.log('onMounted')
    })

    return {
      interfaceData,
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