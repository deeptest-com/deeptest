<template>
  <div id="main">
    <RequestSender></RequestSender>
  </div>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, onMounted, PropType, Ref} from "vue";
import {useI18n} from "vue-i18n";
import {Form, message} from 'ant-design-vue';
import {resizeHeight} from "@/utils/dom";
import {useStore} from "vuex";
import {StateType} from "@/views/interface/store";

const useForm = Form.useForm;

interface InterfaceRequestSetupData {
  modelData: ComputedRef;
  submit: (e) => void;
}

export default defineComponent({
  name: ' InterfaceRequest',
  props: {
    onSubmit: {
      type: Function as PropType<(model: any) => void>,
      required: true
    }
  },
  components: {
  },
  setup(props): InterfaceRequestSetupData {
    const {t} = useI18n();
    const store = useStore<{ Interface: StateType }>();
    const modelData = computed<any>(() => store.state.Interface.modelResult);

    const submit = (e) => {
      console.log('submit')
      props.onSubmit(modelData.value);
    };

    onMounted(() => {
      console.log('onMounted')
      resizeHeight('content', 'top-panel', 'splitter-v', 'bottom-panel', 200, 200, 50)
    })

    return {
      modelData,
      submit,
    }
  }
})
</script>

<style lang="less" scoped>
#main {
}
</style>