<template>
  <div class="user-edit-main">
    <a-card :bordered="false">
      <template #title>
        <div>编辑用户</div>
      </template>
      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-form-item label="用户名" v-bind="validateInfos.username">
            <a-input v-model:value="modelRef.username"
                     @blur="validate('username', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>
          <a-form-item label="邮箱" v-bind="validateInfos.email">
            <a-input v-model:value="modelRef.email"
                     @blur="validate('email', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>
          <a-form-item label="姓名" v-bind="validateInfos.name">
            <a-input v-model:value="modelRef.name"/>
          </a-form-item>
          <a-form-item label="角色" v-bind="validateInfos.role_ids">
            <a-select
                v-model:value="modelRef.role_ids"
                mode="multiple"
                placeholder="请选择角色"
            >
              <a-select-option
                  v-for="(option, key) in sysRoles"
                  :key="key"
                  :value="option.id"
              >{{ option.displayName }}</a-select-option
              >
            </a-select>
          </a-form-item>
          <a-form-item label="介绍" v-bind="validateInfos.intro">
            <a-input v-model:value="modelRef.intro"/>
          </a-form-item>
          <a-form-item label="密码" v-bind="validateInfos.password">
            <a-input v-model:value="modelRef.password"/>
          </a-form-item>

          <a-form-item :wrapper-col="{ offset: labelCol.span, span: wrapperCol.span}">
            <a-button type="primary" @click.prevent="submitForm">保存</a-button>
            <a-button style="margin-left: 10px" @click="resetFields">重置</a-button>
          </a-form-item>
        </a-form>
      </div>
    </a-card>
  </div>
</template>
<script lang="ts">
import {computed, defineComponent, reactive, ref, watchEffect} from "vue";
import {useRouter} from "vue-router";
import {pattern} from "@/utils/const";
import {useStore} from "vuex";
import {StateType} from "../store";
import {User} from "@/views/user/data";
import {Form, message} from "ant-design-vue";

const useForm = Form.useForm;

export default defineComponent({
  name: 'UserEditPage',
  props:{ currentUserId: Number,getList: Function,closeModal: Function, sysRoles:Array},
  setup(props:any) {
    const router = useRouter();
    const formRef = ref();

    const rulesRef = reactive({
      username: [
        { required: true, message: '请输入用户名', trigger: 'blur' },
      ],
      email: [
        { required: true, pattern: pattern.email, message: '请输入邮箱', trigger: 'blur' },
      ],
      adminId: [
        { required: true, message: '请选择管理员'},
      ],
    });

    const store = useStore<{ UserInternal: StateType }>();
    const modelRef = computed<Partial<User>>(() => store.state.UserInternal.detailResult);
    const { resetFields, validate, validateInfos } = useForm(modelRef, rulesRef);

    const get = async (id: number): Promise<void> => {
      await store.dispatch('UserInternal/getUser', id);
      if (props.currentUserId != 0) {
        modelRef.value.password = '******'
      }
    }
    watchEffect(() => {
      get(props.currentUserId)
    })

    const submitForm = async() => {
      validate().then(() =>{
        if (modelRef.value.password == '******') {
          modelRef.value.password = ''
        }
        store.dispatch('UserInternal/saveUser', modelRef.value).then((res) => {
          if (res === true) {
            message.success("保存成功")
            props.getList(1)
          } else {
            message.error("保存失败")
          }
          props.closeModal()
        })
      }).catch(err => {
        console.log('error', err);
      });
    }

    const back = (): void => {
      router.replace(`/user/index`)
    }

    return {
      labelCol: { span: 6 },
      wrapperCol: { span: 14 },
      formRef,
      modelRef,
      rulesRef,
      resetFields,
      validate,
      validateInfos,
      submitForm,
      back,
    }
  }
})
</script>
<style lang="less" scoped>
</style>
