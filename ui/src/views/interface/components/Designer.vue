<template>
  <div id="content" v-if="modelData.method">
    <div id="top-panel">
      <RequestSender :onSend="sendRequest"></RequestSender>
    </div>
    
    <div id="splitter-v"></div>

    <div id="bottom-panel">
      {{ modelData.name }}
    </div>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, onMounted, PropType, Ref} from "vue";
import {useI18n} from "vue-i18n";
import {Form, message} from 'ant-design-vue';
import {resizeHeight} from "@/utils/dom";
import {useStore} from "vuex";
import {StateType} from "@/views/interface/store";
import RequestSender from './designer/request/Sender.vue';

const useForm = Form.useForm;

interface InterfaceDesignerSetupData {
  modelData: ComputedRef;
  sendRequest: (e) => void;
}

export default defineComponent({
  name: 'InterfaceDesigner',
  props: {
    onSubmit: {
      type: Function as PropType<() => void>,
      required: true
    }
  },
  components: {
    RequestSender,
  },
  setup(props): InterfaceDesignerSetupData {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const modelData = computed<any>(() => store.state.Interface.modelResult);

    const sendRequest = (e) => {
      console.log('sendRequest')
    };

    onMounted(() => {
      console.log('onMounted')
      resizeHeight('content', 'top-panel', 'splitter-v', 'bottom-panel', 200, 200, 50)
    })

    return {
      modelData,
      sendRequest,
    }
  }
})
</script>

<style lang="less" scoped>
#content {
  display: flex;
  flex-direction: column;
  position: relative;
  height: 100%;

  #top-panel {
    height: 200px;
    width: 100%;
    padding: 6px;
  }

  #bottom-panel {
    flex: 1;
    width: 100%;
    overflow: auto;
  }

  #splitter-v {
    width: 100%;
    height: 3px;
    background-color: #e6e9ec;
    cursor: ns-resize;

    &:hover {
      height: 3px;
      background-color: #D0D7DE;
    }

    &.active {
      height: 3px;
      background-color: #a9aeb4;
    }
  }
}

</style>