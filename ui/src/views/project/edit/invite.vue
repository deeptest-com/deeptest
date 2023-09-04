<template>
    <a-modal
      title="邀请用户"
      :visible="visible"
      @cancel="close"
      @ok="ok"
      width="700px">
  <div class="invite-main">
    <a-card>

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
          <a-input v-model:value="modelRef.email" />
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
import {Form} from 'ant-design-vue';
import {useRouter} from "vue-router";
import { SelectTypes } from "ant-design-vue/lib/select";
import {StateType as UserStateType} from "@/store/user";
import {StateType} from "../store";
import { defineProps,defineEmits,} from 'vue';
import { Member } from "../data.d";
const props = defineProps({
  visible: {
    required: true,
    type: Boolean,
  },})

const emit = defineEmits(['ok', 'cancel']);

const router = useRouter();

const useForm = Form.useForm;

const {t} = useI18n();

const modelRef = ref<Member>({userId:"",username:"",email:"",roleName:""});


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

const { resetFields,validate, validateInfos} = useForm(modelRef, rulesRef);

const labelCol = {span: 4}
const wrapperCol = {span: 14}
const store = useStore<{ Project: StateType, User: UserStateType }>();
const options = computed<SelectTypes["options"]>(()=>store.state.Project.notExistedUserList);
const roles = computed<SelectTypes["options"]>(()=>store.state.Project.roles);
const formRef = ref();

const close = ()=>{
  emit("cancel")
  reset()
}

const ok  = ()=>{
  validate().then(() => {
    emit('ok', modelRef.value, () => {
          reset();
        });
  })
}

function reset() {
  resetFields()
}

const selectUser  = (value:any) => {
      store.state.Project.notExistedUserList?.forEach((item)=>{
        if (item.id == value){
          modelRef.value.email = item.email
        }
      })
    };

</script>
<style lang="less" scoped>
.invite-main {

}
</style>
