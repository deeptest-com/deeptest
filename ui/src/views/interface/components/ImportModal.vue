<template>
  <div class="import-spec">
    <a-modal title="导入定义"
             :visible="isVisible"
             :onCancel="onCancel"
             class="import-modal">

      <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-form-item label="文件">
            <a-input-search
                v-model:value="modelRef.path"
                @search="selectFile"
            >
              <template #enterButton>
                <a-button v-if="isElectron" >选择文件</a-button>

                <a-upload v-if="!isElectron"
                          v-model:file-list="fileList"
                          :before-upload="beforeUpload"
                          :showUploadList="false"
                          accept=".json,.yml,.yaml"
                >
                  <div>选择文件</div>
                </a-upload>
              </template>

            </a-input-search>

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
import {defineComponent, Ref, ref, PropType} from "vue";
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

    const selectFile = () => {
      console.log('selectFile')

      if (!isElectron.value) return

      const { ipcRenderer } = window.require('electron')
      ipcRenderer.send(settings.electronMsg, {act: 'selectSpec', type: modelRef.value.type})

      ipcRenderer.on(settings.electronMsgReplay, (event, data) => {
        console.log('from electron: ', data)
        modelRef.value.path = data.path
        modelRef.value.content = data.content
      })
    }

    const fileList = ref([]);

    const beforeUpload = (file) => {
      console.log('beforeUpload', file)
      modelRef.value.path = file.name

      modelRef.value.formData = new FormData()
      modelRef.value.formData.append('file', file)

      return false
    };

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
      beforeUpload,
      onSubmit,
      onCancel,

      labelCol: { span: 4 },
      wrapperCol: { span: 18 },
    }
  }
})
</script>

<style lang="less">
.import-modal {
  .ant-input-search {
    .ant-btn-primary {
      color: #000000d9 !important;
      background: #fff !important;
      border-color: #d9d9d9 !important;
    }
  }
}
</style>

<style lang="less" scoped>
.radio {
  display: block;
  height: 30px;
  lineHeight: 30px;
}
</style>