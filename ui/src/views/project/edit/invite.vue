<template>
  <div class="invite-main">
    <a-card>
      <template #title>
        <div>邀请用户</div>
      </template>

      <template #extra>
        <a-button type="link" @click="() => back()">返回</a-button>
      </template>

      <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
        <a-form-item label="用户名" v-bind="validateInfos.username">
          <a-input v-model:value="modelRef.username"/>
        </a-form-item>

        <a-form-item label="邮箱" v-bind="validateInfos.email">
          <a-input v-model:value="modelRef.email"/>
        </a-form-item>

        <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
          <a-button @click="submit" type="primary" class="submit">
            邀请
          </a-button>
        </a-form-item>
      </a-form>

    </a-card>

  </div>
</template>

<script setup lang="ts">
import {ref, reactive} from "vue";
import {useI18n} from "vue-i18n";

import {Form, notification} from 'ant-design-vue';
import {NotificationKeyCommon, pattern} from "@/utils/const";
import {inviteUser} from "@/views/user/info/service";
import {useRouter} from "vue-router";
const router = useRouter();

const useForm = Form.useForm;

const {t} = useI18n();

const modelRef = ref<any>({
  username: '',
  email: '',
});

const projectId = +router.currentRoute.value.params.id

const rulesRef = reactive({
  username: [
    {
      required: true,
      message: '请输入用户名',
    },
    {
      min: 4,
      message: '用户名最少4位'
    }
  ],
  email: [
    {
      required: true,
      message: '邮箱地址不能为空',
    },
    {
      pattern: pattern.email,
      message: '不正确的邮箱地址',
    },
  ],
});

const {resetFields, validate, validateInfos} = useForm(modelRef, rulesRef);

const submit = async (e: MouseEvent) => {
  validate().then(() => {
    console.log(modelRef);

    inviteUser(modelRef.value, projectId).then((json) => {
      if (json.code === 0) {
        notification.success({
          key: NotificationKeyCommon,
          message: `保存成功`,
        });
      } else {
        notification.success({
          key: NotificationKeyCommon,
          message: `保存失败`,
        });
      }
    })
  })
}

const back = () => {
  console.log('back')
  router.push(`/project/members/${projectId}`)
}

const labelCol = {span: 4}
const wrapperCol = {span: 14}

</script>
<style lang="less" scoped>
.invite-main {

}
</style>