<template>
  <div class="plan-create-main">
    <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
      <a-form-item label="名称" v-bind="validateInfos.name">
        <a-input v-model:value="modelRef.name"
                 @blur="validate('name', { trigger: 'blur' }).catch(() => {})" />
      </a-form-item>

      <a-form-item label="描述">
        <a-input v-model:value="modelRef.desc" />
      </a-form-item>

      <a-form-item :wrapper-col="{ offset: labelCol.span }">
        <a-button type="primary" @click.prevent="save">保存</a-button>
        <a-button @click="resetFields">重置</a-button>
      </a-form-item>

    </a-form>
  </div>
</template>

<script setup lang="ts">
import {defineProps, PropType, reactive, ref} from "vue";
import {EditOutlined, DeleteOutlined} from '@ant-design/icons-vue';
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {useI18n} from "vue-i18n";
import {Form, notification} from 'ant-design-vue';
import {StateType} from "../store";
import {addScenarios, get, getDetail, removeScenarioFromPlan} from "@/views/plan/service";
import SelectScenario from "./select-scenario.vue"
import {NotificationKeyCommon} from "@/utils/const";

const useForm = Form.useForm;

const router = useRouter();
const {t} = useI18n();

const props = defineProps({
  categoryId: {
    type: Number,
    required: true
  },
  onSaved: {
    type: Function as PropType<() => void>,
    required: true
  }
})

const store = useStore<{ Plan: StateType }>();

const modelRef = ref({name: '', categoryId: props.categoryId} as any)
const formRef = ref();
const rulesRef = reactive({
  name: [
    { required: true, message: '请输入名称', trigger: 'blur' },
  ],
});

const { resetFields, validate, validateInfos } = useForm(modelRef, rulesRef);

const save = async () => {
  console.log('save');
  validate().then(() => {
    console.log(modelRef);
    store.dispatch('Plan/savePlan', modelRef.value).then((res) => {
      if (res === true) {
        props.onSaved()
      }
    })
  })
}

const labelCol = {span: 3}
const wrapperCol = {span: 21}

</script>

<style lang="less" scoped>
.plan-create-main {

}
</style>
