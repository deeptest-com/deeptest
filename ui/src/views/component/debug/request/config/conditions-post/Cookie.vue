<template>
  <div class="response-cookie-main">
    <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
      <a-form-item label="Cookie名称" v-bind="validateInfos.cookieName">
        <a-input v-model:value="model.cookieName"
                 @blur="validate('cookieName', { trigger: 'blur' }).catch(() => {})" />
      </a-form-item>

      <a-form-item label="Cookie域">
        <a-input v-model:value="model.cookieDomain" />
      </a-form-item>

      <a-form-item label="赋予变量名" v-bind="validateInfos.variableName">
        <a-input v-model:value="model.variableName"
                 @blur="validate('variableName', { trigger: 'blur' }).catch(() => {})"/>
      </a-form-item>

      <a-form-item label="默认值">
        <a-input v-model:value="model.default" />
        <div class="dp-input-tip">Cookie不存时的默认值</div>
      </a-form-item>

    </a-form>
  </div>
</template>

<script setup lang="ts">
import {computed, defineProps, inject, onBeforeUnmount, onMounted, PropType, reactive, Ref, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {message, Form, notification} from 'ant-design-vue';

import {UsedBy} from "@/utils/enum";

import {StateType as Debug} from "@/views/component/debug/store";
import {NotificationKeyCommon} from "@/utils/const";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";

const useForm = Form.useForm;
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

const store = useStore<{  Debug: Debug }>();

const debugInfo = computed<any>(() => store.state.Debug.debugInfo);
const debugData = computed<any>(() => store.state.Debug.debugData);
const model = computed<any>(() => store.state.Debug.cookieData);

const props = defineProps({
  condition: {
    type: Object,
    required: true,
  },
  finish: {
    type: Function,
    required: false,
  },
})

const load = () => {
  console.log('load', props.condition)
  store.dispatch('Debug/getCookie', props.condition.entityId)
}
load()

const variables = ref([])

const rulesRef = computed(() => { return {
  cookieName: [
    { required: true, message: '请输入Cookie名称', trigger: 'blur' },
  ],
  variableName: [
    { required: true, message: '请输入变量名称', trigger: 'blur' },
  ],
}})

let { resetFields, validate, validateInfos } = useForm(model, rulesRef);

const save = () => {
  console.log('save', model.value)
  validate().then(() => {
    model.value.debugInterfaceId = debugInfo.value.debugInterfaceId
    model.value.endpointInterfaceId = debugInfo.value.endpointInterfaceId
    model.value.projectId = debugData.value.projectId

    store.dispatch('Debug/saveCookie', model.value).then((result) => {
      if (result) {
        notification.success({
          key: NotificationKeyCommon,
          message: `保存成功`,
        });
        if (props.finish) {
          props.finish()
        }
      } else {
        notification.error({
          key: NotificationKeyCommon,
          message: `保存失败`,
        });
      }
    })
  })
}
const cancel = () => {
  console.log('cancel')
  if (props.finish) {
    props.finish()
  }
}

onMounted(() => {
  console.log('onMounted')
  bus.on(settings.eventConditionSave, save);
})
onBeforeUnmount( () => {
  console.log('onBeforeUnmount')
  bus.off(settings.eventConditionSave, save);
})

const labelCol = { span: 4 }
const wrapperCol = { span: 18 }

</script>

<style lang="less" scoped>
.response-cookie-main {
  height: 100%;
  width: 100%;
}
</style>