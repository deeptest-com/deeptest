<template>
  <div class="project-edit-main">
    <a-card :bordered="false">
          <template #title>
            <div>编辑项目</div>
          </template>
          <template #extra>
           <!-- <a-button type="link" @click="() => back()">返回</a-button>  --> 
          </template>

          <div>
            <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
              <a-form-item label="项目名称" v-bind="validateInfos.name">
                <a-input v-model:value="modelRef.name"
                         @blur="validate('name', { trigger: 'blur' }).catch(() => {})" />
              </a-form-item>
              <a-form-item label="项目logo" v-bind="validateInfos.logo">
                <a-input v-model:value="modelRef.logo"
                         @blur="validate('logo', { trigger: 'blur' }).catch(() => {})" />
              </a-form-item>
              <a-form-item label="英文缩写" v-bind="validateInfos.shortName">
                <a-input v-model:value="modelRef.shortName"
                         @blur="validate('shortName', { trigger: 'blur' }).catch(() => {})" />
              </a-form-item>
              <a-form-item label="管理员" v-bind="validateInfos.adminId">
                <a-select v-model:value="modelRef.adminId" show-search style="width: 250px" 
                 @blur="validate('adminId', { trigger: 'blur' }).catch(() => {})">
                 <a-select-option  v-for="(option,key) in options" :key=key :value="option.value">{{option.label}}</a-select-option>
                </a-select>     
              </a-form-item>
              <a-form-item label="项目简介" v-bind="validateInfos.desc">
                <a-input v-model:value="modelRef.desc"
                         @blur="validate('desc', { trigger: 'blur' }).catch(() => {})" />
              </a-form-item>

              <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
                <a-button type="primary" @click.prevent="submitForm">保存</a-button>
                <a-button style="margin-left: 10px" @click="resetFields">重置</a-button>
              </a-form-item>
            </a-form>
          </div>
        </a-card>
  </div>
</template>

<script lang="ts">
import {defineComponent, computed, ref, reactive, ComputedRef,watchEffect} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import { useI18n } from "vue-i18n";
import {message, Form, notification} from 'ant-design-vue';
const useForm = Form.useForm;
import {StateType as UserStateType} from "@/store/user";
import {StateType} from "../store";
import {Project} from "@/views/project/data";
import {SelectTypes} from 'ant-design-vue/es/select';

export default defineComponent({
    name: 'ScriptEditPage',
    props:{ currentProjectId:  Number,getList:Function,closeModal:Function},
    setup(props:any) {
      const router = useRouter();

      const { t } = useI18n();

      const formRef = ref();

      const rulesRef = reactive({
        name: [
          { required: true, message: '请输入名称', trigger: 'blur' },
        ],
        logo: [
          { required: true, message: '请输入logo', trigger: 'blur' },
        ],
        shortName: [
          { required: true, message: '大写英文字母开头,仅限字母和数字,<=10位,不可修改', trigger: 'blur' },
        ],
        adminId: [
          { required: true, message: '请选择管理员',trigger: 'blur'},
        ],
      });

      const store = useStore<{ Project: StateType, User: UserStateType }>();
      const modelRef = computed<Partial<Project>>(() => store.state.Project.detailResult);
      const { resetFields, validate, validateInfos } = useForm(modelRef, rulesRef);

      const get = async (id: number): Promise<void> => {
        await store.dispatch('Project/getProject', id);
      }
      //const id = +router.currentRoute.value.params.id
      watchEffect(() => {
      get(props.currentProjectId)
      //store.dispatch('Project/getUserList')
    })
     // get(id.value)

     const options = computed<SelectTypes["options"]>(()=>store.state.Project.userList);
    
      const submitForm = async() => {
        validate().then(() => {
          console.log(modelRef);

          store.dispatch('Project/saveProject', modelRef.value).then((res) => {
            console.log('res', res)
            if (res === true) {
              store.dispatch('User/fetchCurrent');
              message.success("保存成功")
              //notification.success({
               // key: NotificationKeyCommon,
              //  message: `保存成功`,
             // });
              props.getList(1)
              //router.replace('/project/index')
            } else {
              message.error("保存失败")
              //notification.success({
              //  key: NotificationKeyCommon,
               // message: `保存失败`,
              //});
            }
            props.closeModal()
          })
        })
        .catch(err => {
          console.log('error', err);
        });
      };

      const back = (): void => {
        router.replace(`/project/index`)
      }

      return {
        labelCol: { span: 4 },
        wrapperCol: { span: 14 },
        formRef,
        modelRef,
        rulesRef,
        resetFields,
        validate,
        validateInfos,
        submitForm,
        back,
        options,
      }
    }
})
</script>

<style lang="less" scoped>
.project-edit-main {

}
</style>
