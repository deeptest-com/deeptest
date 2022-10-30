<template>
  <div class="import-spec">
    <a-modal title="导入定义"
             :visible="isVisible"
             :onCancel="onCancel"
             class="import-modal"
             width="700px">

      <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-form-item label="文件">

            <a-input-search
                v-if="isElectron"
                v-model:value="modelRef.path"
                @search="selectFile"
            >
              <template #enterButton>
                <a-button  >选择文件</a-button>
              </template>

            </a-input-search>

            <div v-if="!isElectron">请使用客户端导入</div>

          </a-form-item>

        <a-form-item label="文件类型">
          <a-radio-group v-model:value="modelRef.type">
            <a-radio class="radio" value="openapi3">OpenAPI 3</a-radio>
            <a-radio class="radio" value="openapi2">Swagger (OpenAPI2)</a-radio>
            <a-radio class="radio" value="postman">PostMan</a-radio>
          </a-radio-group>
        </a-form-item>

        </a-form>

      <template #footer>
        <a-button :disabled="!modelRef.path" @click="onSubmit" type="primary">导入</a-button>
      </template>
    </a-modal>
  </div>
</template>
<script lang="ts">
import {defineComponent, Ref, ref, PropType, onMounted, getCurrentInstance, onUnmounted} from "vue";
import settings from "@/config/settings";
import {loadSpecFromAgent} from "@/views/interface/service";

export default defineComponent({
  name: 'ImportModal',
  components: {},
  props: {
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
    const modelRef = ref({type: 'openapi3'} as any)

    let ipcRenderer = undefined as any

    if (isElectron.value && !ipcRenderer) {
      ipcRenderer = window.require('electron').ipcRenderer

      ipcRenderer.on(settings.electronMsgReplay, (event, data) => {
        console.log('from electron: ', data)
        modelRef.value.path = data.path

        if (modelRef.value.type === 'postman') { // already converted by postman
          modelRef.value.content = data.content
        } else { // call agent to convert
          loadSpecFromAgent(modelRef.value.path, modelRef.value.type).then((json) => {
            if (json.code === 0) {
              console.log(json.data)
              modelRef.value.content = json.data.content
            }
          })
        }
      })
    }

    onMounted(() => {
      console.log('onMounted')
    })

    onUnmounted(() => {
      console.log('onUnmounted')
    })

    const selectFile = () => {
      console.log('selectFile')

      if (!isElectron.value) return

      ipcRenderer.send(settings.electronMsg, {act: 'selectSpec', type: modelRef.value.type})
    }

    const fileList = ref([]);

    const onSubmit  = () => {
      console.log('onSubmit')
      props.submit(modelRef.value)
    }

    const onCancel = () => {
      console.log('onSubmit')
      props.cancel()
    }

    return {
      isElectron,
      modelRef,
      selectFile,
      fileList,
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