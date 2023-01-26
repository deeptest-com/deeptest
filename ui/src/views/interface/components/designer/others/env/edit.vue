<template>
  <a-modal
      :title="modelRef.id ? '编辑' : '创建' + '环境'"
      :destroy-on-close="true"
      :mask-closable="false"
      :visible="true"
      :onCancel="onCancel"
      :footer="null"
  >
    <div>
      <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
        <a-form-item label="名称" v-bind="validateInfos.name">
          <a-input v-model:value="modelRef.name"
                   @blur="validate('name', { trigger: 'blur' }).catch(() => {})"/>
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
import {Form} from 'ant-design-vue';
import {useI18n} from "vue-i18n";
import {getEnvironment} from "@/services/environment";
import {useStore} from "vuex";
import {StateType as InterfaceStateType} from "@/views/interface/store";
import {StateType as EnvironmentStateType} from "@/store/environment";

const useForm = Form.useForm;

const props = defineProps({
  modelId: {
    type: Number,
    required: true
  },
  interfaceId: {
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
const store = useStore<{ Interface: InterfaceStateType, Environment: EnvironmentStateType }>();

const rulesRef = reactive({
  name: [
    {required: true, message: '请输入变量名', trigger: 'blur'},
  ],
});

const modelRef = ref<any>({name: ''})

const {resetFields, validate, validateInfos} = useForm(modelRef, rulesRef);

const getModel = async () => {
  if (props.modelId === 0) {
    modelRef.value = {name: '', interfaceId: props.interfaceId}
  } else {
    getEnvironment(props.modelId, 0).then((json) => {
      console.log('json', json)
      modelRef.value = json.data
    })
  }
}
getModel()

const onSubmit = async () => {
  console.log('onSubmit', modelRef)

  validate().then(async () => {
    store.dispatch('Environment/saveEnvironment', modelRef.value).then(() => {
      props.onFinish();
    })
  }).catch(err => {
    console.log('')
  })
}

onMounted(() => {
  console.log('onMounted')
})

const labelCol = {span: 6}
const wrapperCol = {span: 16}

</script>