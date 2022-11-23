<template>
  <div class="change-password">
    <a-modal title="修改密码"
             :visible="isVisible"
             :onCancel="onCancel"
             class="change-password-modal">

      <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
        <a-form-item label="当前密码" v-bind="validateInfos.password">
          <a-input v-model:value="modelRef.password"
                   @blur="validate('password', { trigger: 'blur' }).catch(() => {})"/>
        </a-form-item>

        <a-form-item label="新密码" v-bind="validateInfos.newPassword">
          <a-input v-model:value="modelRef.newPassword"
                   @blur="validate('newPassword', { trigger: 'blur' }).catch(() => {})"/>
        </a-form-item>

        <a-form-item label="重复密码" v-bind="validateInfos.confirm">
          <a-input v-model:value="modelRef.confirm"
                   @blur="validate('confirm', { trigger: 'blur' }).catch(() => {})"/>
        </a-form-item>
      </a-form>

      <template #footer>
        <a-button @click="onSubmit" type="primary">提交</a-button>
      </template>
    </a-modal>
  </div>
</template>
<script lang="ts">
import {defineComponent, Ref, ref, PropType, onMounted, getCurrentInstance, onUnmounted, reactive} from "vue";
import {message, Form, notification} from 'ant-design-vue';
import settings from "@/config/settings";
import {pattern} from "@/utils/const";
const useForm = Form.useForm;

export default defineComponent({
  name: 'ImportModal',
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
    const modelRef = ref({
      password: '',
      newPassword: '',
      confirm: ''
    } as any)

    const rulesRef = reactive({
      password: [
        {
          type: 'string',
          required: true,
          message: '原有密码不能为空',
        },
      ],
      newPassword: [
        {
          required: true,
          message: '新密码不能为空',
        },
        {
          min: 6,
          message: '新密码长度最少6位'
        }
      ],
      confirm: [
        {
          validator: (rule: any, value: string, callback: any) => {
            if (value === '') {
              return Promise.reject('重复密码不能为空');
            } else if (value !== modelRef.value.newPassword) {
              return Promise.reject("两次密码不相等");
            } else {
              return Promise.resolve();
            }
          }
        }
      ],
    });

    const { resetFields, validate, validateInfos } = useForm(modelRef, rulesRef);

    const onSubmit  = () => {
      console.log('onSubmit')

      validate().then(() => {
        props.submit(modelRef.value)
      })
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
      modelRef,
      onSubmit,
      onCancel,
      resetFields, validate, validateInfos,
      labelCol: { span: 4 },
      wrapperCol: { span: 18 },
    }
  }
})
</script>

<style lang="less" scoped>
.change-password {

}
</style>