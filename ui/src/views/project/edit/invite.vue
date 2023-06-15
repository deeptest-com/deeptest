<template>
  <a-modal
      :visible="visible"
      @cancel="close"
      width="700px"
      :footer="null">
  <div class="invite-main">
    <a-card>
      <template #title>
      </template>
      <!--
      <template #extra>
        <a-button type="link" @click="() => back()">返回</a-button>
      </template>
    -->

      <a-form :label-col="labelCol" :wrapper-col="wrapperCol" ref="formRef">


        <a-form-item label="用户名" v-bind="validateInfos.userId">
                <a-select v-model:value="modelRef.userId" show-search @change="selectUser" :options="options" optionFilterProp="label"
                          @blur="validate('userId', { trigger: 'blur' }).catch(() => {})">
                </a-select>
              </a-form-item>

        <a-form-item label="角色" v-bind="validateInfos.roleName">
                <a-select v-model:value="modelRef.roleName" show-search
                 @blur="validate('roleName', { trigger: 'blur' }).catch(() => {})">
                 <a-select-option  v-for="(option,key) in roles" :key=key :value="option.value">{{option.label}}</a-select-option>
                </a-select>
        </a-form-item>

        <a-form-item label="邮箱" v-bind="validateInfos.email" >
          <a-input v-model:value="modelRef.email"/>
        </a-form-item>

        <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
          <a-button @click="submit" type="primary" class="submit">
            确认
          </a-button>
        </a-form-item>
      </a-form>

    </a-card>

  </div>
  </a-modal>
</template>

<script setup lang="ts">
import {ref, reactive, computed} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {Form, notification, } from 'ant-design-vue';
import {NotificationKeyCommon} from "@/utils/const";
import {inviteUser} from "@/views/user/info/service";
import {useRouter} from "vue-router";
import { SelectTypes } from "ant-design-vue/lib/select";
import {StateType as UserStateType} from "@/store/user";
import {StateType} from "../store";
import { defineProps,defineEmits,} from 'vue';
const props = defineProps({
  visible:Boolean
})

const emit = defineEmits(['ok', 'cancel']);

const router = useRouter();

const useForm = Form.useForm;

const {t} = useI18n();


const modelRef = reactive({
  userId:"",
  email:"",
  roleName:"user",
  username:""
});


const projectId = Number(window.localStorage.getItem('currentProjectId'));


const rulesRef = reactive({
  userId: [
    {
      required: true,
      message: '请选择用户',
    }
  ],

  email: [
    {
      required: true,
      message: '邮箱地址不能为空',
    }
  ],
  roleName :[
    {
      required: true,
      message: '角色不能为空',
    },
  ]
});

const { validate, validateInfos} = useForm(modelRef, rulesRef);

const submit = async (e: MouseEvent) => {
  validate().then(() => {
    console.log(modelRef);

    inviteUser(modelRef, projectId).then((json) => {
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
      close()
    })
  })
}

const back = () => {
  console.log('back')
  router.push(`/project/members/${projectId}`)
}

const labelCol = {span: 4}
const wrapperCol = {span: 14}
const store = useStore<{ Project: StateType, User: UserStateType }>();
const options = computed<SelectTypes["options"]>(()=>store.state.Project.notExistedUserList);
const roles = computed<SelectTypes["options"]>(()=>store.state.Project.roles);
const formRef = ref();

const close = ()=>{
  emit("cancel")
}



const selectUser  = (value:any) => {
      store.state.Project.notExistedUserList?.forEach((item)=>{
        if (item.id == value){
          modelRef.email = item.email
        }
      })
    };

</script>
<style lang="less" scoped>
.invite-main {

}
</style>
