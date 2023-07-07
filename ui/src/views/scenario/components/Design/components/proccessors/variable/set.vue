<template>
  <div class="processor_variable_set-main">
    <a-card :bordered="false">
      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">

          <a-form-item label="变量名称" v-bind="validateInfos.variableName">
            <a-input v-model:value="modelRef.variableName"
                     @blur="validate('variableName', { trigger: 'blur' }).catch(() => {})"/>
          </a-form-item>

          <a-form-item label="表达式" v-bind="validateInfos.expression">
            <a-input v-model:value="modelRef.expression"
                     @blur="validate('expression', { trigger: 'blur' }).catch(() => {})"/>
            <div class="dp-input-tip">{{t('tips_expression', {name: '{name}'})}}</div>
          </a-form-item>

          <a-form-item label="备注" v-bind="validateInfos.comments">
            <a-input v-model:value="modelRef.comments"/>
          </a-form-item>

          <a-form-item :wrapper-col="{ span: 16, offset: 2 }">
            <a-button type="primary" @click.prevent="submitForm">保存</a-button>
            <a-button style="margin-left: 10px" @click="resetFields">重置</a-button>
          </a-form-item>
        </a-form>
      </div>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, onUnmounted, reactive, ref} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {useI18n} from "vue-i18n";
import {Form, message, notification} from 'ant-design-vue';
import {StateType as ScenarioStateType} from "../../../../../store";
import {EditOutlined, CheckOutlined, CloseOutlined} from "@ant-design/icons-vue";

const useForm = Form.useForm;

const router = useRouter();

const {t} = useI18n();

const formRef = ref();

const rulesRef = reactive({
  variableName: [
    {required: true, message: '请输入变量名', trigger: 'blur'},
  ],
  expression: [
    {required: true, message: '请输入表达式', trigger: 'blur'},
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
            notification.success({
              message: `保存成功`,
            });
          } else {
            notification.error({
              message: `保存失败`,
            });
          }
        })
      })
};

onMounted(() => {
  console.log('onMounted')
  if (!modelRef.value.variableName) modelRef.value.variableName = ''
  if (!modelRef.value.expression) modelRef.value.expression = ''
})

onUnmounted(() => {
  console.log('onUnmounted')
})

const labelCol = { span: 4 }
const wrapperCol = { span: 16 }

</script>

<style lang="less" scoped>
.processor_variable_set-main {
  .icons {
    text-align: right;
    line-height: 32px;
  }
}
</style>
