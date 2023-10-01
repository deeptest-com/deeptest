<template>
  <div class="agent-edit-main">
    <a-drawer :width="1000" :bodyStyle="{padding:'16px'}"
              :closable="true"
              :key="modelId"
              :visible="visible"
              @close="onCancel">

      <template #title>
        <div class="drawer-header">
          <div>{{model.id?'编辑':'新建'}}执行代理</div>
        </div>
      </template>

      <div v-if="visible">
        <a-form :model="model" :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-form-item label="名称" v-bind="validateInfos.name" required>
            <a-input v-model:value="model.name"
                     @blur="validate('name', { trigger: 'blur' }).catch(() => {})"/>
          </a-form-item>

          <a-form-item label="地址" v-bind="validateInfos.url" required>
            <a-input v-model:value="model.url"
                     @blur="validate('url', { trigger: 'blur' }).catch(() => {})"/>
          </a-form-item>

          <a-form-item label="描述" v-bind="validateInfos.desc">
            <a-textarea v-model:value="model.desc"
                        @blur="validate('desc', { trigger: 'blur' }).catch(() => {})"/>
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
import {urlValidator} from "@/utils/validate";

const useForm = Form.useForm;

const store = useStore<{ SysSetting: SysSettingStateType }>();
const model = computed<any>(() => store.state.SysSetting.agentModel);

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
    {required: true, message: '请输入名称', trigger: 'blur'},
  ],
  url: [
    {required: true, validator: urlValidator, message: '请输入合法的代理URL地址', trigger: 'blur'},
  ],
});

const {resetFields, validate, validateInfos} = useForm(model, rulesRef);

watch(props, () => {
  console.log('editId', props)

  if (props.modelId === 0) {
    store.commit('SysSetting/setAgent', {name: ''});
  } else {
    store.dispatch('SysSetting/getAgent', props.modelId);
  }
}, {deep: true, immediate: true})

const onSubmit = async () => {
  console.log('onSubmit', model.value)

  validate().then(async () => {
    store.dispatch('SysSetting/saveAgent', model.value).then(() => {
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
.agent-edit-main {
}
</style>
