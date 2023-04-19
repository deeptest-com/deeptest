<template>
    <div class="login-form-main">
        <div class="menu-tab">
            <div class="menu-tab-item menu-tab-active">注册账号</div>
        </div>
        <a-form :wrapper-col="{span:24}">
            <a-form-item label="" v-bind="validateInfos.username">
                <div class="login-input-item">
                    <a-input v-model:value="modelRef.username" :placeholder="t('page.user.register.form-item-username')" @keyup.enter="handleSubmit" />
                </div>
            </a-form-item>

            <a-form-item label="" v-bind="validateInfos.email">
                <div class="login-input-item">
                    <a-input v-model:value="modelRef.email" :placeholder="t('page.user.register.form-item-email')" @keyup.enter="handleSubmit" />
                </div>
            </a-form-item>

            <a-form-item label="" v-bind="validateInfos.password">
                <div class="login-input-item">
                  <a-input-password v-model:value="modelRef.password" :placeholder="t('page.user.register.form-item-password')" @keyup.enter="handleSubmit" />
                </div>  
            </a-form-item>

            <a-form-item label="" v-bind="validateInfos.confirm">
                <div class="login-input-item">
                    <a-input-password v-model:value="modelRef.confirm" :placeholder="t('page.user.register.form-item-confirmpassword')" @keyup.enter="handleSubmit" />
                </div>
            </a-form-item>
            <div class="text-align-right">
                <router-link to="/user/login">
                    {{t('page.user.register.form.btn-jump')}}
                </router-link>
            </div>
            <a-form-item>
                <div class="login-input-button">
                    <a-button type="primary" class="submit" @click="handleSubmit" :loading="submitLoading">
                        {{t('page.user.register.form.btn-submit')}}
                    </a-button>
                </div>  
            </a-form-item>

            <a-alert v-if="errorMsg !== '' && typeof errorMsg !== 'undefined' &&  !submitLoading" :message="errorMsg" type="error" :show-icon="true" />

        </a-form>
    </div>
</template>
<script lang="ts">
import { computed, ComputedRef, defineComponent, reactive, Ref, ref } from "vue";
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';
import { useI18n } from "vue-i18n";

import { Props, validateInfos } from 'ant-design-vue/lib/form/useForm';
import {message, Form, notification} from 'ant-design-vue';
const useForm = Form.useForm;

import useI18nAntdFormVaildateInfos from '@/composables/useI18nAntdFormVaildateInfos';
import { RegisterParamsType } from "./data.d";
import { StateType as RegisterStateType } from "./store";
import {pattern} from "@/utils/const";

interface UserRegisterSetupData {
    t: (key: string | number) => string;
    validateInfos: ComputedRef<validateInfos>;
    modelRef: RegisterParamsType;
    submitLoading: Ref<boolean>;
    handleSubmit: (e: MouseEvent) => void;
    errorMsg: ComputedRef<string | undefined>;
}

export default defineComponent({
    name: 'UserRegister',
    setup(): UserRegisterSetupData {
        const router = useRouter();
        const store = useStore<{UserRegister: RegisterStateType}>();
        const { t } = useI18n();

        const modelRef = reactive<RegisterParamsType>({
          username: '',
          email: '',
          password: '',
          confirm: ''
        });

        const rulesRef = reactive({
          username: [
            {
                required: true,
                message: 'page.user.register.form-item-username.required',
            },
            {
              min: 4,
              message: '用户名最少4位'
            }
          ],
          email: [
            {
              type: 'string',
              required: true,
              pattern: pattern.email,
              message: 'page.user.register.form-item-email.required',
            },
          ],
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
                          return Promise.reject('page.user.register.form-item-password.required');
                      } else if (value !== modelRef.password) {
                          return Promise.reject("page.user.register.form-item-confirmpassword.compare");
                      } else {
                          return Promise.resolve();
                      }
                  }
              }
          ],
        });

        const { validate, validateInfos } = useForm(modelRef, rulesRef);
        const submitLoading = ref<boolean>(false);
        const handleSubmit = async (e: MouseEvent) => {
            e.preventDefault();
            submitLoading.value = true;
            try {
                const fieldsValue = await validate<RegisterParamsType>();
                const res: boolean = await store.dispatch('UserRegister/register',fieldsValue);
                if (res === true) {
                  notification.success({
                    message: t('page.user.register.form.register-success'),
                  });
                  router.replace('/user/login');
                }
            } catch (error) {
              notification.warn({
                message: t('page.user.register.form.register-fail'),
              });
            }
            submitLoading.value = false;
        };

        // 重置 validateInfos
        const validateInfosNew = useI18nAntdFormVaildateInfos(validateInfos);

         // 注册状态
        const errorMsg = computed<string | undefined>(()=> store.state.UserRegister.errorMsg);

        return {
            t,
            modelRef,
            validateInfos: validateInfosNew,
            submitLoading,
            handleSubmit,
            errorMsg
        }

    }
})
</script>
<style lang="less" scoped>
@import url('../assets/login.less');
</style>