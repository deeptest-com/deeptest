<template>
  <div class="project-edit-main">
    <a-card :bordered="false">
      <template #title>
        <div>编辑项目</div>
      </template>
      <template #extra>
        <a-button type="link" @click="() => back()">返回</a-button>
      </template>

      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-form-item label="名称" v-bind="validateInfos.name">
            <a-input v-model:value="modelRef.name"
                     @blur="validate('name', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>
          <a-form-item label="描述" v-bind="validateInfos.desc">
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
import {defineComponent, computed, ref, reactive, ComputedRef} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import { useI18n } from "vue-i18n";
import { Props, validateInfos } from 'ant-design-vue/lib/form/useForm';
import {message, Form, notification} from 'ant-design-vue';
const useForm = Form.useForm;
import {StateType} from "../store";
import {Project} from "@/views/project/data";

interface EditProjectPageSetupData {
  formRef: any
  modelRef: ComputedRef<Partial<Project>>
  rulesRef: any
  labelCol: any
  wrapperCol: any
  validate: any
  validateInfos: validateInfos
  submitForm:  () => void;
  resetFields:  () => void;

  back: () => void;
}

export default defineComponent({
    name: 'ScriptEditPage',
    setup(): EditProjectPageSetupData {
      const router = useRouter();

      const { t } = useI18n();

      const formRef = ref();

      const rulesRef = reactive({
        name: [
          { required: true, message: '请输入名称', trigger: 'blur' },
        ],
      });

      const store = useStore<{ Project: StateType }>();
      const modelRef = computed<Partial<Project>>(() => store.state.Project.detailResult);
      const { resetFields, validate, validateInfos } = useForm(modelRef, rulesRef);

      const get = async (id: number): Promise<void> => {
        await store.dispatch('Project/getProject', id);
      }
      const id = +router.currentRoute.value.params.id
      get(id)

      const submitForm = async() => {
        validate()
            .then(() => {
              console.log(modelRef);

              store.dispatch('Project/saveProject', modelRef.value).then((res) => {
                console.log('res', res)
                if (res === true) {
                  notification.success({
                    message: `保存成功`,
                  });
                  router.replace('/project/index')
                } else {
                  notification.success({
                    message: `保存失败`,
                  });
                }
              })
            })
            .catch(err => {
              console.log('error', err);
            });
      };

      const back = ():void =>  {
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
      }
    }
})
</script>

<style lang="less" scoped>
.project-edit-main {

}
</style>
