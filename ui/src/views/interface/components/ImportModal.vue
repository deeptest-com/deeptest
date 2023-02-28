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
                v-model:value="modelRef.file"
                @search="selectFile">
              <template #enterButton>
                <a-button>选择文件</a-button>
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
        <a-button :disabled="!modelRef.file" @click="onSubmit" type="primary">导入</a-button>
      </template>
    </a-modal>
  </div>
</template>
<script lang="ts">
import {defineComponent, Ref, ref, PropType, onMounted, getCurrentInstance, onUnmounted} from "vue";
import settings from "@/config/settings";

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
        modelRef.value.file = data.file
        modelRef.value.src = data.src
      })
    }

    const selectFile = () => {
      console.log('selectFile')

      if (isElectron.value) {
        const data = {act: 'loadSpec', type: modelRef.value.type, src: 'file'} as any
        ipcRenderer.send(settings.electronMsg, data)
      }
    }

    const fileList = ref([]);

    const onSubmit  = () => {
      console.log('onSubmit')
      props.submit(modelRef.value)
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