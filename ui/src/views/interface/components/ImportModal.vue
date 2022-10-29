<template>
  <div>
    <a-modal title="导入定义"
             :visible="isVisible"
             :onCancel="cancel">

      <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-form-item label="文件">
            <a-input v-if="isElectron" v-model:value="modelRef.path" />

            <a-button v-if="isElectron" @click="selectFile()">选择文件</a-button>
            <a-input type="text" v-if="!isElectron" v-model="modelRef.path" />
          </a-form-item>
        </a-form>

      <template #footer>
        <a-button @click="submit" type="primary">导入</a-button>
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
    const modelRef = ref({} as any)

    const selectFile = () => {
      console.log('selectFile')

      const { ipcRenderer } = window.require('electron')
      ipcRenderer.send(settings.electronMsg, 'importSpec')

      ipcRenderer.on(settings.electronMsgReplay, (event, data) => {
        console.log('from electron: ' + data.path + ', ' + data.spec)
      })
    }

    return {
      isElectron,
      modelRef,
      selectFile,

      labelCol: { span: 4 },
      wrapperCol: { span: 18 },
    }
  }
})
</script>
<style lang="less" scoped>

</style>