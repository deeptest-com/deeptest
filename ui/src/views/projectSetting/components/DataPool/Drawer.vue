<!-- 本页面是数据池编辑页面的抽屉 -->
<template>
  <div class="datapool-main">
    <a-drawer :closable="true" :width="1000" :key="editKey" :bodyStyle="{padding:'16px'}" :visible="drawerVisible"
              @close="onClose">
      <template #title>
        <div class="drawer-header">
          <div>数据池编辑</div>
        </div>
      </template>

      <div class="">
        <a-form :model="formState" :label-col="labelCol" :wrapper-col="wrapperCol">

          <a-form-item label="名称" v-bind="validateInfos.name">
            <a-input v-model:value="modelRef.name"
                     @blur="validate('name', { trigger: 'blur' }).catch(() => {})"/>
          </a-form-item>

          <a-form-item label="从文件导入" :wrapper-col="wrapperCol">
            <div v-if="isElectron" class="upload-file">
              <div class="input-container">
                <a-input v-model:value="modelRef.path" readonly="readonly"/>
              </div>
              <div class="upload-container">
                <a-button @click="uploadFile()">
                  <UploadOutlined/>
                </a-button>
              </div>
            </div>

            <div v-if="isElectron" class="upload-file">
            </div>

            <div v-if="isElectron" class="upload-file-by-electron">
              <a-input v-model:value="modelRef.path" readonly="readonly" />
              <a-button @click="uploadFile()">
                <UploadOutlined />
              </a-button>
            </div>

            <div v-else class="upload-file">
              <div class="input-container">
                <a-input v-model:value="modelRef.path" readonly="readonly" />
              </div>
              <div class="upload-container">
                <a-upload :beforeUpload="upload"
                          :showUploadList="false">
                  <a-button>
                    <UploadOutlined />
                  </a-button>
                </a-upload>
              </div>
            </div>
          </a-form-item>

          <a-form-item :wrapper-col="{ span: wrapperCol.span, offset: labelCol.span }">
            <div class="handson-table-wrapper">
              <HandsonTable :data="data"></HandsonTable>
            </div>
          </a-form-item>

          <a-form-item :wrapper-col="{ span: wrapperCol.span, offset: labelCol.span }">
            <a-button type="primary" @click="onSubmit" class="dp-btn-gap">保存</a-button> &nbsp;
            <a-button @click="onClose" class="dp-btn-gap">取消</a-button>
          </a-form-item>
        </a-form>
      </div>
    </a-drawer>
  </div>
</template>

<script setup lang="ts">
import {computed, defineEmits, defineProps, ref, reactive, watch} from 'vue';
import {Form, notification} from 'ant-design-vue';
import {useStore} from 'vuex';
import { UploadOutlined} from '@ant-design/icons-vue';
import HandsonTable from "@/components/sheet/handsontable.vue";

import settings from "@/config/settings";
import {NotificationKeyCommon} from "@/utils/const";
import {getServerUrl} from "@/utils/request";
import {getToken} from "@/utils/localToken";

import {StateType as ProjectStateType} from "@/store/project";
import {StateType as ProjectSettingStateType} from '../../store';
import {ServeDetail} from '../../data';
import {uploadRequest} from "@/utils/upload";

const useForm = Form.useForm;
const store = useStore<{ ProjectGlobal: ProjectStateType, ProjectSetting: ProjectSettingStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const formState = computed<ServeDetail>(() => store.state.ProjectSetting.datapoolDetail);

const props = defineProps<{
  drawerVisible: boolean
  editKey?: number
}>();

const emits = defineEmits(['onClose', 'update:formState']);
const onClose = () => {
  emits('onClose');
}

const modelRef = ref<any>({name: '', path: ''})
const data = ref<any[][]>([['A','B','C'], [1,2,3]])

const rulesRef = reactive({
  name: [
    {required: true, message: '请输入变量名', trigger: 'blur'},
  ],
});

const {resetFields, validate, validateInfos} = useForm(modelRef, rulesRef);

watch(props, () => {
  console.log('modelRef', props)

  if (props.editKey === 0) {
    modelRef.value = {name: '', path: ''}
  } else {
    store.dispatch('ProjectSetting/getDatapool', props.editKey);
  }

}, {deep: true, immediate: true})

watch(modelRef, () => {
  console.log('modelRef', modelRef.value)
  if (!modelRef.value.id) return

  data.value = JSON.parse(modelRef.value.data)
}, {deep: true, immediate: true})

const isElectron = ref(!!window.require)
let ipcRenderer = undefined as any
if (isElectron.value && !ipcRenderer) {
  ipcRenderer = window.require('electron').ipcRenderer

  ipcRenderer.on(settings.electronMsgReplay, (event, result) => {
    console.log('from electron: ', result)
    if (result.code === 0) {
      data.value = result.data.data
      modelRef.value.path = result.data.path
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
  }
}

const upload =  (file, fileList) => {
  console.log('upload', file, fileList)

  uploadRequest(file, {isDatapool: true}).then((res) => {
    modelRef.value.path = res.path
    data.value = res.data
  })

  return false
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
      // onClose();
    })
  }).catch(err => {
    console.log(err)
  })
}

const labelCol = {span: 4}
const wrapperCol = {span: 20}

</script>

<style lang="less" scoped>
.datapool-main {

}
</style>
