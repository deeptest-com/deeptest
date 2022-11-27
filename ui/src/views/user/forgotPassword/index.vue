<template>
  <div class="forgotPassword-main">
    <h1 class="title">
      忘记密码
    </h1>
    <a-form :wrapper-col="{span:24}">
      <a-form-item label="" v-bind="validateInfos.usernameOrPassword">
        <a-input v-model:value="modelRef.usernameOrPassword" placeholder="输入用户名或邮箱" />
      </a-form-item>

      <a-form-item>
        <a-button type="primary" class="submit" @click="handleSubmit" :loading="submitLoading">
          重置密码
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
import {forgotPassword} from "@/views/user/forgotPassword/service";
const useForm = Form.useForm;

const router = useRouter();
const { t } = useI18n();

const modelRef = ref({usernameOrPassword: ''});

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
.forgotPassword-main {
  flex: none;
  width: 320px;
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