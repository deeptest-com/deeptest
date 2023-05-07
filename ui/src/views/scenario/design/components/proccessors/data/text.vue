<template>
  <div class="processor_data_text-main">
    <a-card :bordered="false">
      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">

          <a-form-item label="备注" v-bind="validateInfos.comments">
            <a-input v-model:value="modelRef.comments"/>
          </a-form-item>

          <a-form-item label="变量名称" v-bind="validateInfos.variableName">
            <a-input v-model:value="modelRef.variableName"
                     @blur="validate('variableName', { trigger: 'blur' }).catch(() => {})"/>
          </a-form-item>

          <a-form-item label="上传文件" v-bind="validateInfos.url">
            <div v-if="isElectron" class="flow-file-input">
              <a-input v-model:value="modelRef.url" readonly="readonly" />
              <a-button @click="uploadFile()">
                <UploadOutlined />
              </a-button>
            </div>

            <div v-else class="flow-file-input2">
              <div class="input-container">
                <a-input v-model:value="modelRef.url" readonly="readonly" />
              </div>
              <div class="upload-container">
                <a-upload :action="uploadUrl"
                          :beforeUpload="upload"
                          :showUploadList="false">
                  <a-button>
                    <UploadOutlined />
                  </a-button>
                </a-upload>
              </div>
            </div>
          </a-form-item>

          <a-form-item label="分隔符" v-bind="validateInfos.separator">
            <a-input v-model:value="modelRef.separator"
                     @blur="validate('separator', { trigger: 'blur' }).catch(() => {})"/>
          </a-form-item>

          <a-form-item label="重复次数" v-bind="validateInfos.repeatTimes">
            <a-input-number v-model:value="modelRef.repeatTimes"
                            @blur="validate('repeatTimes', { trigger: 'blur' }).catch(() => {})"/>
          </a-form-item>

          <a-form-item label="是否随机" v-bind="validateInfos.isRand">
            <a-switch v-model:checked="modelRef.isRand"
                      @blur="validate('isRand', { trigger: 'blur' }).catch(() => {})"/>
          </a-form-item>
          <a-form-item label="是否唯一" v-bind="validateInfos.isOnce">
            <a-switch v-model:checked="modelRef.isOnce"
                      @blur="validate('isOnce', { trigger: 'blur' }).catch(() => {})"/>
          </a-form-item>
          <a-form-item label="是否循环" v-bind="validateInfos.isLoop">
            <a-switch :disabled="modelRef.isOnce" v-model:checked="modelRef.isLoop"
                      @blur="validate('isLoop', { trigger: 'blur' }).catch(() => {})"/>
          </a-form-item>

          <a-form-item :wrapper-col="{ span: 16, offset: 4 }">
            <a-button type="primary" @click.prevent="submitForm">保存</a-button>
            <a-button style="margin-left: 10px" @click="resetFields">重置</a-button>
          </a-form-item>
        </a-form>
      </div>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, onUnmounted, reactive, ref} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {useI18n} from "vue-i18n";
import {Form, message, notification} from 'ant-design-vue';
import {StateType as ScenarioStateType} from "../../../../store";
import {UploadOutlined} from "@ant-design/icons-vue";
import {NotificationKeyCommon} from "@/utils/const";
import {getServerUrl} from "@/utils/request";
import {getToken} from "@/utils/localToken";
import settings from "@/config/settings";
import {uploadRequest} from "@/utils/upload";

const useForm = Form.useForm;

const router = useRouter();

const {t} = useI18n();

const formRef = ref();

const rulesRef = reactive({
  variableName: [
    {required: true, message: '请输入变量名称', trigger: 'blur'},
  ],
  url: [
    {required: true, message: '请输入文件路径', trigger: 'blur'},
  ],
  separator: [
    {required: true, message: '请输入分隔符', trigger: 'blur'},
  ]
});

const store = useStore<{ Scenario: ScenarioStateType; }>();
const modelRef = computed<any>(() => store.state.Scenario.nodeData);
const {resetFields, validate, validateInfos} = useForm(modelRef, rulesRef);

const uploadUrl = ref(getServerUrl() + '/upload')

const submitForm = async () => {
  validate()
      .then(() => {
        store.dispatch('Scenario/saveProcessor', modelRef.value).then((res) => {
          if (res === true) {
            notification.success({
              key: NotificationKeyCommon,
              message: `保存成功`,
            });
          } else {
            notification.error({
              key: NotificationKeyCommon,
              message: `保存失败`,
            });
          }
        })
      })
};

const isElectron = ref(!!window.require)
let ipcRenderer = undefined as any
if (isElectron.value && !ipcRenderer) {
  ipcRenderer = window.require('electron').ipcRenderer

  ipcRenderer.on(settings.electronMsgReplay, (event, result) => {
    console.log('from electron: ', result)
    if (result.code === 0) {
      modelRef.value.url = result.data.path
    }
  })
}

const uploadFile = async () => {
  console.log('uploadFile')

  if (isElectron.value) {
    const data = {
      act: 'uploadFile',
      url: getServerUrl() + '/upload',
      token: await getToken(),
      filters: [
        {name: 'Text Files', extensions: ['txt']},
      ]
    }

    ipcRenderer.send(settings.electronMsg, data)

  } else {
    notification.warn({
      key: NotificationKeyCommon,
      message: `请使用客户端上传文件`,
    });
  }
}

const upload = async (file, fileList) => {
  console.log('upload', file, fileList)

  const path = await uploadRequest(file)
  modelRef.value.url = path

  return false
}

onMounted(() => {
  console.log('onMounted')
  if (!modelRef.value.url) modelRef.value.url = ''
  if (!modelRef.value.variableName) modelRef.value.variableName = ''
  if (!modelRef.value.separator) modelRef.value.separator = ','
  if (!modelRef.value.repeatTimes) modelRef.value.repeatTimes = 1
})

onUnmounted(() => {
  console.log('onUnmounted')
})

const labelCol = { span: 4 }
const wrapperCol = { span: 16 }

</script>

<style lang="less" scoped>
.processor_data_text-main {
  .icons {
    text-align: right;
    line-height: 32px;
  }
}
</style>
