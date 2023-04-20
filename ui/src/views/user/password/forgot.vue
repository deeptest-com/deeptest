<template>
  <div class="login-form-main">
    <div class="menu-tab">
      <div class="menu-tab-item menu-tab-active">忘记密码</div>
    </div>
    <a-form :wrapper-col="{ span: 24 }">
      <a-form-item label="" v-bind="validateInfos.usernameOrPassword">
        <div class="login-input-item">
          <a-input v-model:value="modelRef.usernameOrPassword" placeholder="输入用户名或邮箱" />
        </div>
      </a-form-item>
      <div class="text-align-right">
        <router-link to="/user/login">
          返回登录
        </router-link>
      </div>
      <a-form-item>
        <div class="login-input-button">
          <a-button type="primary" class="submit" @click="handleSubmit" :loading="submitLoading">
            提交
          </a-button>
        </div>
      </a-form-item>
    </a-form>
  </div>
</template>
<script setup lang="ts">
import { reactive, ref } from "vue";
import { useRouter } from 'vue-router';
import { useI18n } from "vue-i18n";

import { message, Form, notification } from 'ant-design-vue';
import { NotificationKeyCommon } from "@/utils/const";
import { forgotPassword } from "@/views/user/password/service";
const useForm = Form.useForm;

const router = useRouter();
const { t } = useI18n();

const modelRef = ref({ usernameOrPassword: '' });

const rulesRef = reactive({
  usernameOrPassword: [
    {
      required: true,
      message: '请输入用户名或邮箱',
    },
  ]
});

const { resetFields, validate, validateInfos } = useForm(modelRef, rulesRef);

const submitLoading = ref<boolean>(false);
const handleSubmit = async (e: MouseEvent) => {
  e.preventDefault();
  submitLoading.value = true;

  validate().then(() => {
    forgotPassword(modelRef.value.usernameOrPassword).then((json) => {
      if (json.code === 0) {
        notification.success({
          key: NotificationKeyCommon,
          message: `重置密码成功，请检查邮件开始进一步操作。`,
        });
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
@import '../assets/login.less';
</style>