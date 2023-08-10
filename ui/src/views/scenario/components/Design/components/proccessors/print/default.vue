<template>
  <div class="processor_print_default-main dp-proccessors-container">
    <ProcessorHeader/>
    <a-card :bordered="false">
      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
          <a-form-item label="输出" v-bind="validateInfos.rightValue">
            <a-input v-model:value="modelRef.rightValue"
                     @blur="validate('rightValue', { trigger: 'blur' }).catch(() => {})" />
            <div class="dp-input-tip">{{t('tips_expression', {name: '{name}'})}}</div>
          </a-form-item>

          <a-form-item label="备注" v-bind="validateInfos.comments">
            <a-textarea v-model:value="modelRef.comments" :rows="3"/>
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
import {computed, onMounted, onUnmounted, reactive, ref} from "vue";
import {Form, message} from "ant-design-vue";
import {useRouter} from "vue-router";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {StateType as ScenarioStateType} from "@/views/scenario/store";
import {getCompareOpts} from "@/utils/compare";
import ProcessorHeader from '../../common/ProcessorHeader.vue';
const useForm = Form.useForm;

const router = useRouter();

const {t} = useI18n();

const formRef = ref();

const rulesRef = reactive({
  rightValue: [
    {required: true, message: '请输入内容', trigger: 'blur'},
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
  console.log('onMounted')
  if (!modelRef.value.leftValue) modelRef.value.leftValue = ''
  if (!modelRef.value.rightValue) modelRef.value.rightValue = ''
  if (!modelRef.value.operator) modelRef.value.operator = 'equal'
})

onUnmounted(() => {
  console.log('onUnmounted')
})

const labelCol = { span: 4 }
const wrapperCol = { span: 16 }

</script>

