<template>
  <a-modal
      :title="modelRef.id ? '编辑' : '创建' + '数据池'"
      :destroy-on-close="true"
      :mask-closable="false"
      :visible="true"
      :onCancel="onCancel"
      :footer="null"
      width="800px"
      height="600px"
  >
    <div class="data-pool-main">
      <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
        <a-form-item label="名称" v-bind="validateInfos.name">
          <a-input v-model:value="modelRef.name"
                   @blur="validate('name', { trigger: 'blur' }).catch(() => {})"/>
        </a-form-item>

        <a-form-item label="上传文件" :wrapper-col="wrapperCol">
          <div class="flow-file-input">
            <a-input v-model:value="modelRef.path" readonly="readonly" />
            <a-button @click="uploadFile()">
              <UploadOutlined />
            </a-button>
          </div>
        </a-form-item>

        <a-form-item :wrapper-col="{ span: wrapperCol.span, offset: labelCol.span }">
          <div class="handson-table-wrapper">
            <HandsonTable :data="data"></HandsonTable>
          </div>
        </a-form-item>

        <a-form-item :wrapper-col="{ span: wrapperCol.span, offset: labelCol.span }">
          <a-button type="primary" @click="onSubmit" class="dp-btn-gap">保存</a-button> &nbsp;
          <a-button @click="() => onCancel()" class="dp-btn-gap">取消</a-button>
        </a-form-item>
      </a-form>

    </div>

  </a-modal>
</template>

<script setup lang="ts">
import {defineProps, onMounted, PropType, reactive, ref} from "vue";
import {Form, notification} from 'ant-design-vue';
import {useI18n} from "vue-i18n";
import {getDatapool} from "@/services/datapool";
import {useStore} from "vuex";
import { UploadOutlined} from '@ant-design/icons-vue';

import {StateType as InterfaceStateType} from "@/views/interface/store";
import {StateType as DatapoolStateType} from "@/store/environment";
import settings from "@/config/settings";
import {NotificationKeyCommon} from "@/utils/const";
import {getServerUrl} from "@/utils/request";
import {getToken} from "@/utils/localToken";

import HandsonTable from "@/components/sheet/handsontable.vue";

const useForm = Form.useForm;

const props = defineProps({
  modelId: {
    type: Number,
    required: true
  },

  onCancel: {
    type: Function,
    required: true
  },
  onFinish: {
    type: Function as PropType<() => void>,
    required: true
  }
})

const {t} = useI18n();
const store = useStore<{ Interface: InterfaceStateType, Datapool: DatapoolStateType }>();

const rulesRef = reactive({
  name: [
    {required: true, message: '请输入变量名', trigger: 'blur'},
  ],
});

const modelRef = ref<any>({name: '', path: ''})

const data = ref<any[][]>([['A','B','C'], [1,2,3]])

const {resetFields, validate, validateInfos} = useForm(modelRef, rulesRef);

const getModel = async () => {
  if (props.modelId === 0) {
    modelRef.value = {name: '', path: ''}
  } else {
    getDatapool(props.modelId, 0).then((json) => {
      console.log('json', json)
      modelRef.value = json.data
    })
  }
}
getModel()

const isElectron = ref(!!window.require)
let ipcRenderer = undefined as any
if (isElectron.value && !ipcRenderer) {
  ipcRenderer = window.require('electron').ipcRenderer

  ipcRenderer.on(settings.electronMsgReplay, (event, result) => {
    console.log('from electron: ', result)
    if (result.code === 0) {
      data.value = result.data
      modelRef.value.path = result.data.path
    }
  })
}

const uploadFile = async () => {
  console.log('uploadFile')

  if (isElectron.value) {
    const data = {
      act: 'uploadFile',
      url: getServerUrl() + '/datapools/upload',
      token: await getToken(),
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

const onSubmit = async () => {
  console.log('onSubmit', modelRef.value)

  validate().then(async () => {
    if (data.value.length < 2) {
      notification.warn({
        key: NotificationKeyCommon,
        message: `表格至少包含标题和数据两行。`,
      });
      return
    }

    modelRef.value.data = JSON.stringify(data.value)

    store.dispatch('Datapool/saveDatapool', modelRef.value).then(() => {
      props.onFinish();
    })
  }).catch(err => {
    console.log(err)
  })
}

onMounted(() => {
  console.log('onMounted')
})

const labelCol = {span: 4}
const wrapperCol = {span: 20}

</script>

<style lang="less" scoped>
.data-pool-main {
  .handson-table-wrapper {
    height: 200px;
  }
}
</style>

<style src="@/../node_modules/handsontable/dist/handsontable.css"></style>