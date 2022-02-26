<template>
  <div id="content">
    <div id="top-panel">

    </div>
    
    <div id="splitter-v"></div>

    <div id="bottom-panel">
      {{ model.name }}
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent, onMounted, PropType, Ref} from "vue";
import {useI18n} from "vue-i18n";
import {Form, message} from 'ant-design-vue';
import {resizeHeight} from "@/utils/dom";

const useForm = Form.useForm;

interface InterfaceDesignerSetupData {
  submit: (e) => void;
}

export default defineComponent({
  name: 'InterfaceDesigner',
  props: {
    model: {
      type: Object,
      required: true
    },
    onSubmit: {
      type: Function as PropType<(model: any) => void>,
      required: true
    }
  },
  components: {
  },
  setup(props): InterfaceDesignerSetupData {
    const {t} = useI18n();

    const submit = (e) => {
      console.log('submit')

      props.onSubmit(props.model);
    };

    onMounted(() => {
      console.log('onMounted')
      resizeHeight('content', 'top-panel', 'splitter-v', 'bottom-panel', 200, 200, 50)
    })

    return {
      submit
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