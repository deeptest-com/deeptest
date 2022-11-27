<template>
  <div class="resetPassword-main">
    <h1 class="title">
      设置密码
    </h1>
    <a-form :wrapper-col="{span:24}">
      <a-form-item label="">
        用户 {{username}}
      </a-form-item>

      <a-form-item label="" v-bind="validateInfos.password">
        <a-input v-model:value="modelRef.password" placeholder="输入新密码"
                 @blur="validate('password', { trigger: 'blur' }).catch(() => {})" />
      </a-form-item>

      <a-form-item label="" v-bind="validateInfos.confirm">
        <a-input v-model:value="modelRef.confirm" placeholder="重复新密码"
                 @blur="validate('confirm', { trigger: 'blur' }).catch(() => {})" />
      </a-form-item>

      <a-form-item>
        <a-button type="primary" class="submit" @click="handleSubmit" :loading="submitLoading">
          提交
        </a-button>
        <div class="text-align-right">
          <router-link to="/user/login">
            返回登录
          </router-link>
        </div>
      </a-form-item>
    </a-form>
  </div>
</template>
<script setup lang="ts">
import { reactive, ref } from "vue";
import { useRouter } from 'vue-router';
import { useI18n } from "vue-i18n";

import {message, Form, notification} from 'ant-design-vue';
import {NotificationKeyCommon} from "@/utils/const";
import {resetPassword} from "@/views/user/password/service";
const useForm = Form.useForm;

const router = useRouter();
const { t } = useI18n();

const username = router.currentRoute.value.params.username
const vcode = router.currentRoute.value.params.vcode
const modelRef = ref({usernameOrPassword: '', username:username, vcode: vcode});

const rulesRef = reactive({
  password: [
    {
      required: true,
      message: 'page.user.register.form-item-password.required',
      trigger: 'blur'
    },
    {
      min: 6,
      message: '密码长度最少6位',
    }
  ],
  confirm: [
    {
      validator: (rule: any, value: string, callback: any) => {
        if (value === '') {
          return Promise.reject('密码不能为空');
        } else if (value !== modelRef.value.password) {
          return Promise.reject("两次密码不一样");
        } else {
          return Promise.resolve();
        }
      }
    }
  ],
});

const { resetFields, validate, validateInfos } = useForm(modelRef, rulesRef);

const submitLoading = ref<boolean>(false);
const handleSubmit = async (e: MouseEvent) => {
  e.preventDefault();
  submitLoading.value = true;

  validate().then(() => {
    resetPassword(modelRef.value).then((json) => {
      if (json.code === 0) {
        notification.success({
          key: NotificationKeyCommon,
          message: `修改密码成功，请登录。`,
        });

        router.replace('/user/login')
      } else {
        notification.error({
          key: NotificationKeyCommon,
          message: `重置密码失败。`,
        });
      }
    })
  })

  submitLoading.value = false;
};

</script>
<style lang="less" scoped>
.resetPassword-main {
  flex: none;
  width: 380px;
  padding: 36px;
  margin: 0 auto;
  border-radius: 4px;
  background-color: rgba(255, 255, 255, 0.2);
  .title {
    font-size: 28px;
    margin-top: 0;
    margin-bottom: 30px;
    text-align: center;
    color: #ffffff;
  }
  .submit {
    width: 100%;
  }
}
</style>