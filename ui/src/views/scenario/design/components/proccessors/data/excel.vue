<template>
  <div class="processor_data_excel-main">
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

          <a-form-item label="选择文件" v-bind="validateInfos.url">
            <div class="flow-file-input">
              <a-input v-model:value="modelRef.url" readonly="readonly"
                       @blur="validate('url', { trigger: 'blur' }).catch(() => {})"/>
              <a-button @click="uploadFile()">
                <UploadOutlined />
              </a-button>
            </div>
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
import {EditOutlined, CheckOutlined, CloseOutlined} from "@ant-design/icons-vue";
import {NotificationKeyCommon} from "@/utils/const";
import settings from "@/config/settings";
import {getServerUrl} from "@/utils/request";
import {getToken} from "@/utils/localToken";

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
});

const store = useStore<{ Scenario: ScenarioStateType; }>();
const modelRef = computed<any>(() => store.state.Scenario.nodeData);
const {resetFields, validate, validateInfos} = useForm(modelRef, rulesRef);

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

  ipcRenderer.on(settings.electronMsgReplay, (event, data) => {
    console.log('from electron: ', data.data)
  })
}

const uploadFile = () => {
  console.log('uploadFile')

  if (isElectron.value) {
    const data = {
      act: 'uploadFile',
      url: getServerUrl() + '/processors/data/upload',
      token: getToken(),
      filters: [
        {name: 'Excel Files', extensions: ['xlsx']},
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

onMounted(() => {
  console.log('onMounted')
  if (!modelRef.value.url) modelRef.value.url = ''
  if (!modelRef.value.variableName) modelRef.value.variableName = ''
  if (!modelRef.value.repeatTimes) modelRef.value.repeatTimes = 1
})

onUnmounted(() => {
  console.log('onUnmounted')
})

const labelCol = { span: 4 }
const wrapperCol = { span: 16 }

</script>

<style lang="less" scoped>
.processor_data_excel-main {
  .icons {
    text-align: right;
    line-height: 32px;
  }
}
</style>
