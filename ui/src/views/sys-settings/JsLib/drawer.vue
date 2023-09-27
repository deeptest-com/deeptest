<!-- 本页面是数据池编辑页面的抽屉 -->
<template>
  <div class="jslib-edit-main">
    <a-drawer :width="1000" :bodyStyle="{padding:'16px'}"
              :closable="true"
              :key="modelId"
              :visible="visible"
              @close="onCancel">

      <template #title>
        <div class="drawer-header">
          <div>{{model.id?'编辑':'新建'}}自定义类库</div>
        </div>
      </template>

      <div v-if="visible">
        <a-form :model="model" :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-form-item label="名称" v-bind="validateInfos.name" required>
            <a-input v-model:value="model.name"
                     @blur="validate('name', { trigger: 'blur' }).catch(() => {})"/>
          </a-form-item>

          <a-form-item label="上传脚本文件"
                       v-bind="validateInfos.scriptFile" required>
            <div v-if="isElectron" class="upload-file">
              <div class="input-container">
                <a-input v-model:value="model.path" readonly="readonly"/>
              </div>
              <div class="upload-container">
                <a-button @click="uploadFile()">
                  <UploadOutlined/>
                </a-button>
              </div>
            </div>
            <div v-if="isElectron" class="upload-file-by-electron">
              <a-input v-model:value="model.path" readonly="readonly"/>
              <a-button @click="uploadFile()">
                <UploadOutlined/>
              </a-button>
            </div>

            <div v-else class="upload-file">
              <div class="input-container">
                <a-input v-model:value="model.scriptFile" readonly="readonly"
                         @blur="validate('scriptFile', { trigger: 'blur' }).catch(() => {})"/>
              </div>
              <div class="upload-container">
                <a-upload :beforeUpload="uploadScript"
                          :showUploadList="false"
                          accept=".js">
                  <a-button>
                    <UploadOutlined/>
                  </a-button>
                </a-upload>
              </div>
            </div>
          </a-form-item>

          <a-form-item label="上传类型文件">
            <div v-if="isElectron" class="upload-file">
              <div class="input-container">
                <a-input v-model:value="model.path" readonly="readonly"/>
              </div>
              <div class="upload-container">
                <a-button @click="uploadFile()">
                  <UploadOutlined/>
                </a-button>
              </div>
            </div>
            <div v-if="isElectron" class="upload-file-by-electron">
              <a-input v-model:value="model.path" readonly="readonly"/>
              <a-button @click="uploadFile()">
                <UploadOutlined/>
              </a-button>
            </div>

            <div v-else class="upload-file">
              <div class="input-container">
                <a-input v-model:value="model.typesFile" readonly="readonly"/>
              </div>
              <div class="upload-container">
                <a-upload :beforeUpload="uploadTypes"
                          :showUploadList="false"
                          accept=".ts">
                  <a-button>
                    <UploadOutlined/>
                  </a-button>
                </a-upload>
              </div>
            </div>
          </a-form-item>

          <a-form-item :wrapperCol="{ span: wrapperCol.span, offset: labelCol.span }">
            <div class="dp-input-tip">
              代码中使用 {{model.name?model.name:'name'}}() 的形式来调用自定义库函数。<br />
              需要准备实现和声明两个JavaScript文件，具体请参照<a href="https://deeptest.com/jslib.html" target="_blank">这里</a>。
            </div>
          </a-form-item>

          <a-form-item :wrapper-col="{ span: wrapperCol.span, offset: labelCol.span }">
            <a-button type="primary" @click="onSubmit" class="dp-btn-gap">保存</a-button> &nbsp;
            <a-button @click="onCancel" class="dp-btn-gap">取消</a-button>
          </a-form-item>
        </a-form>
      </div>
    </a-drawer>
  </div>
</template>

<script setup lang="ts">
import {computed, defineEmits, defineProps, reactive, ref, watch} from 'vue';
import {Form, notification} from 'ant-design-vue';
import {useStore} from 'vuex';
import {UploadOutlined} from '@ant-design/icons-vue';

import settings from "@/config/settings";
import {getUrls} from "@/utils/request";
import {getToken} from "@/utils/localToken";

import {StateType as SysSettingStateType} from "../store";
import {uploadRequest} from "@/utils/upload";
import {notifyWarn} from "@/utils/notify";
import {pattern} from "@/utils/const";

const useForm = Form.useForm;

const store = useStore<{ SysSetting: SysSettingStateType }>();
const model = computed<any>(() => store.state.SysSetting.jslibModel);

const props = defineProps({
  visible: {
    type: Boolean,
    required: true,
  },
  modelId: {
    type: Number,
    required: true,
  },
  onClose: {
    type: Function,
    required: true,
  },
})

const onCancel = () => {
  props.onClose()
}

const rulesRef = reactive({
  name: [
    {required: true, message: '名称以字母开头包含字母和数字，且不能为空。', pattern: pattern.alphanumeric, trigger: 'blur'},
  ],
  scriptFile: [
    {required: true, message: '请上传脚本文件', trigger: 'blur'},
  ],
});

const {resetFields, validate, validateInfos} = useForm(model, rulesRef);

watch(props, () => {
  console.log('editId', props)

  if (props.modelId === 0) {
    store.commit('SysSetting/setJslib', {name: '', typesFile: '', scriptFile: ''});
  } else {
    store.dispatch('SysSetting/getJslib', props.modelId);
  }
}, {deep: true, immediate: true})

const isElectron = ref(!!window.require)
let ipcRenderer = undefined as any
if (isElectron.value && !ipcRenderer) {
  ipcRenderer = window.require('electron').ipcRenderer

  ipcRenderer.on(settings.electronMsgReplay, (event, result) => {
    console.log('from electron: ', result)
    if (result.code === 0) {
      // data.value = result.data.data
      model.value.data = JSON.stringify(result.data.data)
      model.value.path = result.data.path
    }
  })
}

const uploadFile = async () => {
  console.log('uploadFile')

  if (isElectron.value) {
    const data = {
      act: 'uploadFile',
      url: getUrls().serverUrl + '/upload',
      params: {isJslib: true},
      token: await getToken(),
      filters: [
        {name: 'Excel Files', extensions: ['xlsx']},
      ]
    }

    ipcRenderer.send(settings.electronMsg, data)
  }
}

const uploadTypes = (file, fileList) => {
  console.log('upload', file, fileList)

  uploadRequest(file, {isJslib: true}).then((res) => {
    model.value.typesFile = res.path
  })

  return false
}
const uploadScript = (file, fileList) => {
  console.log('upload', file, fileList)

  uploadRequest(file, {isJslib: true}).then((res) => {
    model.value.scriptFile = res.path
  })

  return false
}

const onSubmit = async () => {
  console.log('onSubmit', model.value)

  validate().then(async () => {
    store.dispatch('SysSetting/saveJslib', model.value).then(() => {
      props.onClose();
    })
  }).catch(err => {
    console.log(err)
  })
}

const labelCol = {span: 4}
const wrapperCol = {span: 18}

</script>

<style lang="less" scoped>
.jslib-edit-main {
}
</style>
