<template>
  <a-modal
      :title="modelRef.id ? '编辑' : '创建' + '变量'"
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
                   @blur="validate('name', { trigger: 'blur' }).catch(() => {})" />
        </a-form-item>
        <a-form-item label="取值" v-bind="validateInfos.rightValue">
          <a-input v-model:value="modelRef.rightValue"
                   @blur="validate('rightValue', { trigger: 'blur' }).catch(() => {})" />
        </a-form-item>

        <a-form-item :wrapper-col="{ span: wrapperCol.span, offset: labelCol.span }">
          <a-button type="primary" @click="onSubmit" class="dp-btn-gap">保存</a-button> &nbsp;
          <a-button @click="() => onCancel()" class="dp-btn-gap">取消</a-button>
        </a-form-item>

      </a-form>

    </div>

  </a-modal>
</template>

<script lang="ts">
import {computed, ComputedRef, defineComponent, onMounted, PropType, reactive, ref, Ref} from "vue";
import {message, Form} from 'ant-design-vue';
import {useI18n} from "vue-i18n";
import {getEnvironment, saveEnvironment} from "@/services/environment";
import {useStore} from "vuex";
import {StateType as InterfaceStateType} from "@/views/interface1/store";
import {StateType as EnvironmentStateType} from "@/store/environment";
const useForm = Form.useForm;

export default defineComponent({
  name: 'EnvVarEdit',
  props: {
    model: {
      required: true
    },
    environmentId: {
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
  },

  components: {},

  setup(props) {
    const { t } = useI18n();

    const store = useStore<{ Interface1: InterfaceStateType, Environment: EnvironmentStateType }>();

    const rulesRef = reactive({
      name: [
        { required: true, message: '请输入变量名', trigger: 'blur' },
      ],
      rightValue: [
        { required: true, message: '请输入取值', trigger: 'blur' },
      ],
    });

    const model = props.model as any
    const modelRef = reactive<any>({
      id: model.id,
      name: model.name,
      rightValue: model.rightValue,
      environmentId: props.environmentId,
    })

    const { resetFields, validate, validateInfos } = useForm(modelRef, rulesRef);

    const onSubmit = async () => {
      console.log('onSubmit', modelRef)

      validate().then(async () => {
        store.dispatch('Environment/saveEnvironmentVar', modelRef).then(() => {
          props.onFinish();
        })
      }).catch(err => { console.log('') })
    }

    onMounted(()=> {
      console.log('onMounted')
    })

    return {
      t,
      labelCol: { span: 6 },
      wrapperCol: { span: 16 },
      rulesRef,
      validate,
      validateInfos,
      resetFields,

      modelRef,
      onSubmit,
    }
  }
})
</script>