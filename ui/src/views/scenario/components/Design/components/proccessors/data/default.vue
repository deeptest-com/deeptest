<template>
  <div class="processor_data_excel-main">
    <a-card :bordered="false">
      <div class="top-header-tip">
<!--        <strong>说明：</strong>-->
<!--        <span>数据迭代处理器将循环读取文件中的行内容，并将读取的内容赋值给指定的变量</span>-->
        <a-alert message="说明：数据迭代处理器将循环读取文件中的行内容，并将读取的内容赋值给指定的变量" type="info" show-icon />
      </div>

      <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
        <a-form-item label="变量名称" v-bind="validateInfos.variableName">
          <a-input v-model:value="modelRef.variableName"
                   @blur="validate('variableName', { trigger: 'blur' }).catch(() => {})"/>
          <div v-if="modelRef.variableName" class="dp-input-tip">
            可使用 {{ '${' + modelRef.variableName + '.列名' + '}' }} 访问数据变量
          </div>
        </a-form-item>

        <a-form-item label="上传文件" v-bind="validateInfos.url">
          <div v-if="isElectron" class="upload-file-by-electron">
            <a-input v-model:value="modelRef.url" readonly="readonly"/>
            <a-button @click="uploadFile()">
              <UploadOutlined/>
            </a-button>
          </div>

          <div v-else class="upload-file">
            <div class="input-container">
              <a-input v-model:value="modelRef.url" readonly="readonly"/>
            </div>
            <div class="upload-container">
              <a-upload :beforeUpload="upload"
                        :showUploadList="false"
                        accept="application/vnd.ms-excel,application/vnd.openxmlformats-officedocument.spreadsheetml.sheet">
                <a-button>
                  <UploadOutlined/>
                </a-button>
              </a-upload>
            </div>
          </div>
          <span class="dp-input-tip">仅支持excel、csv、 text三种文件格式</span>
        </a-form-item>

        <a-form-item label="重复次数" v-bind="validateInfos.repeatTimes">
          <a-input-number
              style="width: 200px;"
              v-model:value="modelRef.repeatTimes"
              @blur="validate('repeatTimes', { trigger: 'blur' }).catch(() => {})"/>
          <div class="dp-input-tip">将按指定次数循环读取文件内容</div>
        </a-form-item>

        <a-form-item label="是否随机" v-bind="validateInfos.isRand">
          <a-switch v-model:checked="modelRef.isRand"
                    @blur="validate('isRand', { trigger: 'blur' }).catch(() => {})"/>
          <div class="dp-input-tip">开关打开，将按照随机顺序读取文件行内容</div>
        </a-form-item>

        <a-form-item label="备注" v-bind="validateInfos.comments">
          <a-input v-model:value="modelRef.comments"/>
        </a-form-item>


        <a-form-item :wrapper-col="{ span: 16, offset: 4 }">
          <a-button type="primary" @click.prevent="submitForm">保存</a-button>
        </a-form-item>
      </a-form>


    </a-card>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, onUnmounted, reactive, ref} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {useI18n} from "vue-i18n";
import {Form, message} from 'ant-design-vue';
import {StateType as ScenarioStateType} from "../../../../../store";
import {UploadOutlined} from "@ant-design/icons-vue";
import settings from "@/config/settings";
import {getUrls} from "@/utils/request";
import {getToken} from "@/utils/localToken";
import {uploadRequest} from "@/utils/upload";
import IconSvg from "@/components/IconSvg";

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
            message.success('保存成功');
          } else {
            message.error('保存失败');
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
      url: getUrls().serverUrl + '/upload',
      token: await getToken(),
      filters: [
        {name: 'Excel Files', extensions: ['xlsx']},
      ]
    }

    ipcRenderer.send(settings.electronMsg, data)
  }
}

const upload = async (file, fileList) => {
  console.log('upload', file, fileList)

  const res = await uploadRequest(file)
  modelRef.value.url = res.path

  return false
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

const labelCol = {span: 4}
const wrapperCol = {span: 16}

</script>

<style lang="less" scoped>
.top-header-tip {
  position: relative;
  //left: 63px;
  margin: 6px  auto 24px 60px;
  //font-size: 14px;
}
</style>
