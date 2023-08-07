<template>
  <div class="processor_extractor_jsonquery-main">
    <a-card :bordered="false">
      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">

          <a-form-item label="变量名称" v-bind="validateInfos.variable">
            <a-input v-model:value="modelRef.variable"
                     @blur="validate('variable', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>

          <a-form-item label="元素路径" v-bind="validateInfos.expression">
            <a-input v-model:value="modelRef.expression"
                     @blur="validate('expression', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>

          <a-form-item label="备注" v-bind="validateInfos.comments">
            <a-input v-model:value="modelRef.comments"/>
          </a-form-item>

          <a-form-item :wrapper-col="{ span: 16, offset: 4 }">
            <a-button type="primary" @click.prevent="submitForm">保存</a-button>
            <a-button style="margin-left: 10px" @click="resetFields">重置</a-button>
          </a-form-item>
        </a-form>
      </div>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, reactive, ref} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {useI18n} from "vue-i18n";
import {Form, message, notification} from 'ant-design-vue';
import {StateType as ScenarioStateType} from "../../../../../store";

const useForm = Form.useForm;

const router = useRouter();

const {t} = useI18n();

const formRef = ref();

const rulesRef = reactive({
  variable: [
    { required: true, message: '请输入变量名', trigger: 'blur' },
  ],
  expression: [
    { required: true, message: '请输入路径表达式', trigger: 'blur' },
  ],
});

const store = useStore<{ Scenario: ScenarioStateType; }>();
const modelRef = computed<any>(() => store.state.Scenario.nodeData);
const {resetFields, validate, validateInfos} = useForm(modelRef, rulesRef);

const submitForm = async () => {
  validate()
      .then(() => {
        store.dispatch('Scenario/saveProcessor', modelRef.value).then((res) => {
          if (res === true) {
            message.success('保存成功');
          } else {
            message.error('保存失败');
          }
        })
      })
};
onMounted(() => {
  if (!modelRef.value.variable) modelRef.value.variable = ''
  if (!modelRef.value.expression) modelRef.value.expression = ''
})

const labelCol = { span: 4 }
const wrapperCol = { span: 16 }

</script>


