<template>
  <div class="import-spec5555">
    <a-modal title="导入yapi项目接口"
             :visible="isVisible"
             :onSubmit="onSubmit"
             :onCancel="onCancel"
             class="import-yapi"
             width="700px">

      <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
        <a-form-item label="yapi域名">
          <a-input v-model:value="formState.domain"></a-input>
        </a-form-item>
        <a-form-item label="项目token">
          <a-input v-model:value="formState.projectToken"></a-input>
        </a-form-item>
      </a-form>

      <template #footer>
        <a-button  @click="onCancel" type="primary">取消</a-button>
        <a-button :disabled="!formState" @click="onSubmit" type="primary">导入</a-button>
      </template>

    </a-modal>
  </div>
</template>
<script lang="ts">
import {defineComponent, Ref, ref, PropType, onMounted, getCurrentInstance, onUnmounted, reactive, UnwrapRef } from "vue";
import settings from "@/config/settings";
import {Form} from "ant-design-vue";
const useForm = Form.useForm;

interface FormState {
  domain: string;
  projectToken: string;
}

export default defineComponent({
  name: 'ImportYapi',
  components: {},
  props: {
    model: {
      required: true
    },
    isVisible: {
      type: Boolean,
      required: true
    },
    submit: {
      type: Function,
      required: true,
    },
    cancel: {
      type: Function,
      required: true,
    },
  },

  setup(props) {
    const isElectron = ref(!!window.require)
    const formState: UnwrapRef<FormState> = reactive({
      domain: "",
      projectToken: "",
    });

    const onSubmit  = () => {
      console.log('onSubmit');
      props.submit(formState)
    }

    const onCancel = () => {
      console.log('onCancel')
      props.cancel()
    }

    onMounted(() => {
      console.log('onMounted')
    })

    onUnmounted(() => {
      console.log('onUnmounted')
    })

    return {
      isElectron,
      formState,
      onSubmit,
      onCancel,
      labelCol: { span: 4 },
      wrapperCol: { span: 18 },
    }
  }
})
</script>

<style lang="less">
</style>

<style lang="less" scoped>
.radio {
  display: block;
  height: 30px;
  lineHeight: 30px;
}
</style>