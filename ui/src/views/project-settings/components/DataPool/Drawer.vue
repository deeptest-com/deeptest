<!-- 本页面是数据池编辑页面的抽屉 -->
<template>
  <div class="datapool-main">
    <a-drawer :closable="true" :width="1000" :key="editId" :bodyStyle="{padding:'16px'}" :visible="drawerVisible"
              @close="onClose">
      <template #title>
        <div class="drawer-header">
          <div>数据池编辑</div>
        </div>
      </template>

      <div v-if="drawerVisible" class="">
        <a-form :model="formState" :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-form-item label="名称" v-bind="validateInfos.name">
            <a-input v-model:value="formState.name"
                     @blur="validate('name', { trigger: 'blur' }).catch(() => {})"/>
          </a-form-item>

          <a-form-item label="从文件导入" :wrapper-col="wrapperCol">
            <div v-if="isElectron" class="upload-file">
              <div class="input-container">
                <a-input v-model:value="formState.path" readonly="readonly"/>
              </div>
              <div class="upload-container">
                <a-button @click="uploadFile()">
                  <UploadOutlined/>
                </a-button>
              </div>
            </div>

            <div v-if="isElectron" class="upload-file-by-electron">
              <a-input v-model:value="formState.path" readonly="readonly"/>
              <a-button @click="uploadFile()">
                <UploadOutlined/>
              </a-button>
            </div>

            <div v-else class="upload-file">
              <div class="input-container">
                <a-input v-model:value="formState.path" readonly="readonly"/>
              </div>
              <div class="upload-container">
                <a-upload :beforeUpload="upload"
                          :showUploadList="false"
                          accept="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet">
                  <a-button>
                    <UploadOutlined/>
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
            请求中可使用形如${_dp(myData, A, 1)}的表达式引用数据池，追回一个参数支持数字（索引）、seq（顺序）和rand（随机）。
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
import {computed, defineEmits, defineProps, reactive, ref, watch} from 'vue';
import {Form, notification} from 'ant-design-vue';
import {useStore} from 'vuex';
import {UploadOutlined} from '@ant-design/icons-vue';
import HandsonTable from "@/components/sheet/handsontable.vue";

import settings from "@/config/settings";
import {NotificationKeyCommon} from "@/utils/const";
import {getUrls} from "@/utils/request";
import {getToken} from "@/utils/localToken";

import {StateType as ProjectStateType} from "@/store/project";
import {StateType as ProjectSettingStateType} from '../../store';
import {DatapoolDetail} from '../../data';
import {uploadRequest} from "@/utils/upload";

const useForm = Form.useForm;
const store = useStore<{ ProjectGlobal: ProjectStateType, ProjectSetting: ProjectSettingStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const formState = computed<DatapoolDetail>(() => store.state.ProjectSetting.datapoolDetail);

const props = defineProps<{
  drawerVisible: boolean
  editId?: number
}>();

const emits = defineEmits(['onClose', 'update:formState']);
const onClose = () => {
  emits('onClose');
}

const dataArr = [['A', 'B'], ['foo', 'bar']]
const data = ref<any[][]>(dataArr)

const rulesRef = reactive({
  name: [
    {required: true, message: '请输入变量名', trigger: 'blur'},
  ],
});

const {resetFields, validate, validateInfos} = useForm(formState, rulesRef);

watch(props, () => {
  console.log('editId', props)

  if (props.editId === 0) {
    store.commit('ProjectSetting/setDatapoolDetail', {name: '', path: '', data: ''});
  } else {
    store.dispatch('ProjectSetting/getDatapool', props.editId);
  }

}, {deep: true, immediate: true})

watch(formState, () => {
  console.log('formState', formState.value)
  if (!formState.value.id) {
    data.value = dataArr
    return
  }

  if (formState.value.data && formState.value.data.length > 0) {
    data.value = JSON.parse(formState.value.data)
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
      formState.value.data = JSON.stringify(result.data.data)
      formState.value.path = result.data.path
    }
  })
}

const uploadFile = async () => {
  console.log('uploadFile')

  if (isElectron.value) {
    const data = {
      act: 'uploadFile',
      url: getUrls().serverUrl + '/upload',
      params: {isDatapool: true},
      token: await getToken(),
      filters: [
        {name: 'Excel Files', extensions: ['xlsx']},
      ]
    }

    ipcRenderer.send(settings.electronMsg, data)
  }
}

const upload = (file, fileList) => {
  console.log('upload', file, fileList)

  uploadRequest(file, {isDatapool: true}).then((res) => {
    formState.value.path = res.path
    formState.value.data = JSON.stringify(res.data)
  })

  return false
}

const onSubmit = async () => {
  console.log('onSubmit', formState.value)

  validate().then(async () => {
    if (data.value.length < 2) {
      notification.warn({
        key: NotificationKeyCommon,
        message: `表格至少包含标题和数据两行。`,
      });
      return
    }

    formState.value.data = JSON.stringify(data.value)

    store.dispatch('ProjectSetting/saveDatapool', {
      "projectId": currProject.value.id,
      formState: formState.value,
      action: formState.value.id > 0 ? 'update':'create'
    }).then(() => {
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
