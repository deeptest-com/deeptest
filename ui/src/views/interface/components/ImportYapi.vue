<template>
  <div class="import-spec">
    <a-modal title="导入yapi项目接口"
             :visible="isVisible"
             :onCancel="onCancel"
             class="import-yapi"
             width="700px">

      <a-form ref="formRef" :rules="rules" :model="formYapi" :label-col="labelCol" :wrapper-col="wrapperCol" >
        <a-form-item ref="name" label="yapi域名" name="yapiHost">
          <a-input v-model:value="formYapi.yapiHost"></a-input>
        </a-form-item>
        <a-form-item label="项目token" name="token">
          <a-input v-model:value="formYapi.token"></a-input>
        </a-form-item>

      </a-form>

      <template #footer>
        <a-button @click="onSubmit" type="primary">导入</a-button>
      </template>
    </a-modal>
  </div>
</template>
<script lang="ts">
import { ValidateErrorEntity } from 'ant-design-vue/es/form/interface';
import {defineComponent, Ref, ref, PropType, onMounted, getCurrentInstance, onUnmounted, reactive,UnwrapRef } from "vue";
import settings from "@/config/settings";
import {Form} from "ant-design-vue";
const useForm = Form.useForm;
interface FormYapi {
  // target: string;
  yapiHost: string;
  token: string;
}
export default defineComponent({
  name: 'ImportYapi',
  components: {},
  props: {
    isVisible: {
      type: Boolean,
      required: true
    },
    submit: {
      type: Function,
      required: true,
    },
    cancel: {
      type: Function,
      required: true,
    },
  },

  setup(props) {
    const formRef = ref();
    const formYapi: UnwrapRef<FormYapi> = reactive({
      // target: '',
      yapiHost: '',
      token: '',
    });

    const rules = {
      yapiHost:[
        { required: true, message: '请输入yapiHost', trigger: 'blur' }
      ],
      token: [
        { required: true, message: '请输入项目token', trigger: 'blur' }
      ],
    };

    const { resetFields, validate, validateInfos } = useForm(formRef, rules);
    const onSubmit  = async () => {
      console.log('onSubmit')
      formRef.value.validate()
          .then(async () => {
            props.submit(formYapi);
          })
          .catch((error: ValidateErrorEntity<FormYapi>) => {
            console.log('error', error);
          });
    }

    const onCancel = () => {
      console.log('onSubmit')
      props.cancel()
    }

    onMounted(() => {
      console.log('onMounted')
    })

    onUnmounted(() => {
      console.log('onUnmounted')
    })

    return {
      onSubmit,
      onCancel,
      formYapi,
      rules,
      formRef,
      validate,
      validateInfos,
      resetFields,


      labelCol: { span: 4 },
      wrapperCol: { span: 18 },
    }
  }
})
</script>

<style lang="less">
</style>

<style lang="less" scoped>
.radio {
  display: block;
  height: 30px;
  lineHeight: 30px;
}
</style>