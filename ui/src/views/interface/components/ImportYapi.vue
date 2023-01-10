<template>
  <div class="import-spec5555">
    <a-modal title="导入yapi项目接口"
             :visible="isVisible"
             :onCancel="onCancel"
             class="import-yapi"
             width="700px">

      <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
        <a-form-item label="yapi域名" v-bind="validateInfos.name">
          <a-input v-model:value="yapiDomian"></a-input>
        </a-form-item>
        <a-form-item label="项目token" v-bind="validateInfos.rightValue">
          <a-input v-model:value="modelRef.rightValue"
                   @blur="validate('rightValue', { trigger: 'blur' }).catch(() => {})" />
        </a-form-item>

      </a-form>

      <template #footer>
        <a-button :disabled="!modelRef.file" @click="onSubmit" type="primary">导入</a-button>
      </template>
    </a-modal>
  </div>
</template>
<script lang="ts">
import {defineComponent, Ref, ref, PropType, onMounted, getCurrentInstance, onUnmounted, reactive} from "vue";
import settings from "@/config/settings";
import {Form} from "ant-design-vue";
const useForm = Form.useForm;

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
    const rulesRef = reactive({
      name: [
        { required: true, message: '请输入yapi域名', trigger: 'blur' },
      ],
      rightValue: [
        { required: true, message: '请输入项目token', trigger: 'blur' },
      ],
    });
    // const model = props.model as any
    // const modelRef = ref({type: 'openapi3'} as any)
    const modelRef = ref(name: string,rightValue: model.rightValue,)

    const { resetFields, validate, validateInfos } = useForm(modelRef, rulesRef);

    let ipcRenderer = undefined as any
    if (isElectron.value && !ipcRenderer) {
      ipcRenderer = window.require('electron').ipcRenderer

      ipcRenderer.on(settings.electronMsgReplay, (event, data) => {
        console.log('from electron: ', data)
        modelRef.value.file = data.file
      })
    }


    const onSubmit  = () => {
      console.log('onSubmit')
      props.submit(modelRef.value)
    }

    const onCancel = () => {
      console.log('onSubmit')
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
      modelRef,
      onSubmit,
      onCancel,
      resetFields,
      validateInfos,

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