<template>
  <div id="content" v-if="modelData.method">
    <div id="top-panel">
      <InterfaceRequest></InterfaceRequest>
    </div>
    
    <div id="splitter-v"></div>

    <div id="bottom-panel">
      <InterfaceResponse></InterfaceResponse>
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
import InterfaceRequest from './designer/request/Index.vue';
import InterfaceResponse from './designer/response/Index.vue';

const useForm = Form.useForm;

interface InterfaceDesignerSetupData {
  modelData: ComputedRef;
}

export default defineComponent({
  name: 'InterfaceDesigner',
  props: {
  },
  components: {
    InterfaceRequest, InterfaceResponse,
  },
  setup(props): InterfaceDesignerSetupData {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const modelData = computed<any>(() => store.state.Interface.modelResult);

    onMounted(() => {
      console.log('onMounted')
      resizeHeight('content', 'top-panel', 'splitter-v', 'bottom-panel',
          100, 100, 0)
    })

    return {
      modelData,
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
    padding: 4px 4px 0 4px;
    height: 300px;
    width: 100%;
  }

  #bottom-panel {
    flex: 1;
    padding: 4px;
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