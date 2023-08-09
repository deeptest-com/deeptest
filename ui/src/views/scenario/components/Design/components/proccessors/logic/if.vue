<template>
  <div class="processor_login_if-main dp-proccessors-container">
    <ProcessorHeader/>

    <a-card :bordered="false">
      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">

          <a-form-item label="判断表达式" v-bind="validateInfos.expression">
            <a-input v-model:value="modelRef.expression"
                     @blur="validate('expression', { trigger: 'blur' }).catch(() => {})" />
            <div class="dp-input-tip">{{t('tips_expression_bool', {name: '{name}'})}}</div>
          </a-form-item>

          <a-form-item label="备注" v-bind="validateInfos.comments">
            <a-input v-model:value="modelRef.comments"/>
          </a-form-item>

          <a-form-item :wrapper-col="{ span: 16, offset: 4 }">
            <a-button type="primary" @click.prevent="submitForm">保存</a-button>
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
import {Form, message} from 'ant-design-vue';
import {StateType as ScenarioStateType} from "../../../../../store";
import ProcessorHeader from '../../common/ProcessorHeader.vue';

import {getCompareOpts} from "@/utils/compare";

const useForm = Form.useForm;

const router = useRouter();

const {t} = useI18n();

const formRef = ref();

const rulesRef = reactive({
  expression: [
    {required: true, message: '请输入表达式', trigger: 'blur'},
  ],
});

const store = useStore<{ Scenario: ScenarioStateType; }>();
const modelRef = computed<any>(() => store.state.Scenario.nodeData);
const {resetFields, validate, validateInfos} = useForm(modelRef, rulesRef);

const operators = getCompareOpts()

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
  if (!modelRef.value.leftValue) modelRef.value.leftValue = ''
  if (!modelRef.value.expression) modelRef.value.expression = ''
  if (!modelRef.value.operator) modelRef.value.operator = ''
})


const labelCol = { span: 4 }
const wrapperCol = { span: 16 }

</script>
