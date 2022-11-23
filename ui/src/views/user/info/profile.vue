<template>
  <div class="profile-main">
    <a-card>
      <template #title>
        <div>用户信息</div>
      </template>
      <template #extra></template>

      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-form-item label="邮箱" v-bind="validateInfos.email">
            <span v-if="changing === 'email'" class="editor">
              <span class="input">
                <a-input v-model:value="modelRef.email"
                         @blur="validate('email', { trigger: 'blur' }).catch(() => {})"/>
              </span>

              <span class="btns">
                <CheckOutlined @click.stop="updateEmail()"
                               :class="{disabled: validateInfos.email.validateStatus === 'error'}"/>&nbsp;
                <CloseOutlined @click.stop="changing = ''"/>
              </span>
            </span>

            <span v-else>
              <span>{{ modelRef.email }}</span> &nbsp;
              <EditOutlined @click.stop="changeEmail()"/>
            </span>
          </a-form-item>

          <a-form-item label="用户名" v-bind="validateInfos.username">
            <span v-if="changing === 'username'" class="editor">
              <span class="input">
                <a-input v-model:value="modelRef.username"
                         @blur="validate('username', { trigger: 'blur' }).catch(() => {})"/>
              </span>

              <span class="btns">
                <CheckOutlined @click.stop="updateName()"
                               :class="{disabled: validateInfos.username.validateStatus === 'error'}"/>&nbsp;
                <CloseOutlined @click.stop="changing = ''"/>
              </span>
            </span>

            <span v-else>
              <span>{{ modelRef.username }}</span> &nbsp;
              <EditOutlined @click.stop="changeName()"/>
            </span>
          </a-form-item>

          <a-form-item label="密码" v-bind="validateInfos.password">
            <span>******</span> &nbsp;
            <EditOutlined @click.stop="changePassword()"/>
          </a-form-item>
        </a-form>
      </div>

    </a-card>

    <ChangePassword
        v-if="showChangePassword"
        :isVisible="showChangePassword"
        :submit="updatePassword"
        :cancel="cancelChangePassword"
    />

  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, reactive, ref} from "vue";
import {useRouter} from 'vue-router';
import {useI18n} from "vue-i18n";
import {EditOutlined, CloseOutlined, CheckOutlined} from "@ant-design/icons-vue";

import {Form, notification} from 'ant-design-vue';
import {useStore} from "vuex";
import {StateType as UserStateType} from "@/store/user";
import {NotificationKeyCommon, pattern} from "@/utils/const";
import ChangePassword from './changePassword.vue'

const useForm = Form.useForm;

const router = useRouter();
const {t} = useI18n();

const store = useStore<{User: UserStateType}>();
const modelRef = computed<any>(()=> store.state.User.currentUser);

const showChangePassword = ref(false)

const rulesRef = reactive({
  username: [
    {
      required: true,
      message: t('page.user.register.form-item-username.required'),
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

const changing = ref('')

const changeEmail = async () => {
  console.log('changeEmail')
  changing.value = 'email'
}

const changeName = async () => {
  console.log('changeName')
  changing.value = 'username'
}

const changePassword = async () => {
  console.log('changePassword')
  showChangePassword.value = true
  changing.value = ''
}
const cancelChangePassword = () => {
  console.log('cancelChangePassword')
  showChangePassword.value = false
}

const updateEmail = async () => {
  console.log('updateEmail')
  validate().then(() => {
    store.dispatch('User/updateEmail', modelRef.value.email).then((result) => {
      console.log('result', result)
      afterUpdate(result)
    })
  })
}
const updateName = async () => {
  console.log('updateName')
  validate().then(() => {
    store.dispatch('User/updateName', modelRef.value.username).then((result) => {
      console.log('result', result)
      afterUpdate(result)
    })
  })
}
const updatePassword = async (data) => {
  console.log('updatePassword')
  store.dispatch('User/updatePassword', data).then((result) => {
    console.log('result', result)
    afterUpdate(result)
  })
}
const afterUpdate = async (result) => {
  if (result === true) {
    notification.success({
      key: NotificationKeyCommon,
      message: `更新成功`,
    });
    changing.value = ''
    showChangePassword.value = false
  }
}

onMounted(()=>{
  store.dispatch("User/fetchCurrent");
})

const labelCol = {span: 4}
const wrapperCol = {span: 14}

</script>
<style lang="less" scoped>
.profile-main {
  .editor {
    display: flex;
    .input {
      width: 280px;
      padding: 0 10px 0 0;
    }
    .btns {
      flex: 1;
      line-height: 30px;

      .disabled {
        color: #00000040;
      }
    }
  }
}
</style>