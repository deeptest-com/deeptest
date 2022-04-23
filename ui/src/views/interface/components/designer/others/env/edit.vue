<template>
  <a-modal
      :title="modelRef.id ? '编辑' : '创建'"
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
import {getEnvironment, saveEnvironment} from "@/views/interface/service";
import {useStore} from "vuex";
import {StateType} from "@/views/interface/store";
const useForm = Form.useForm;

export default defineComponent({
  name: 'EnvEdit',
  props: {
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
  },

  components: {},

  setup(props) {
    const { t } = useI18n();

    const store = useStore<{ Interface: StateType }>();
    const environmentsData = computed<any[]>(() => store.state.Interface.environmentsData);

    const rulesRef = reactive({
      name: [
        { required: true, message: '请输入变量名', trigger: 'blur' },
      ],
    });

    const modelRef = reactive<any>({name: ''})

    const { resetFields, validate, validateInfos } = useForm(modelRef, rulesRef);

    const getModel = async () => {
      if (props.modelId === 0) {
        modelRef.value = {name: '', interfaceId: props.interfaceId}
      } else {
        modelRef.value = await getEnvironment(props.modelId, 0)
      }
    }
    getModel()

    const onSubmit = async () => {
      console.log('onSubmit', modelRef)

      validate().then(async () => {
        store.dispatch('Interface/saveEnvironment', modelRef).then(() => {
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