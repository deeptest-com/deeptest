<template>
  <div class="processor_cookie_set-main dp-processors-container">
    <ProcessorHeader/>
    <a-card :bordered="false">
      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">

          <a-form-item label="Cookie名称" v-bind="validateInfos.cookieName">
            <a-input v-model:value="modelRef.cookieName"
                     @blur="validate('cookieName', { trigger: 'blur' }).catch(() => {})"/>
          </a-form-item>

          <a-form-item label="Cookie域">
            <a-input v-model:value="modelRef.domain"/>
          </a-form-item>

          <a-form-item label="过期时间">
            <a-date-picker v-model:value="modelRef.expireTime"/>
          </a-form-item>

          <a-form-item label="取值" v-bind="validateInfos.rightValue">
            <a-input v-model:value="modelRef.rightValue"
                     @blur="validate('rightValue', { trigger: 'blur' }).catch(() => {})"/>
            <div class="dp-input-tip">{{t('tips_expression', {name: '{name}', number: '{+number}'})}}</div>
          </a-form-item>

          <a-form-item label="备注" v-bind="validateInfos.comments">
            <a-textarea v-model:value="modelRef.comments" :rows="3"/>
          </a-form-item>

          <a-form-item class="processor-btn" :wrapper-col="{ span: 16, offset: 4 }">
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
import debounce from "lodash.debounce";
import {Form, notification} from 'ant-design-vue';
import {StateType as ScenarioStateType} from "../../../../../store";
import ProcessorHeader from '../../common/ProcessorHeader.vue';
const useForm = Form.useForm;

const router = useRouter();

const {t} = useI18n();

const formRef = ref();

const rulesRef = reactive({
  cookieName: [
    {required: true, message: '请输入Cookie名称', trigger: 'blur'},
  ],
  rightValue: [
    {required: true, message: '请输入取值', trigger: 'blur'},
  ],
});

const store = useStore<{ Scenario: ScenarioStateType; }>();
const modelRef = computed<any>(() => store.state.Scenario.nodeData);
const {resetFields, validate, validateInfos} = useForm(modelRef, rulesRef);

const submitForm = debounce(async () => {
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
}, 300);

onMounted(() => {
  console.log('onMounted')
  if (!modelRef.value.cookieName) modelRef.value.cookieName = ''
  if (!modelRef.value.rightValue) modelRef.value.rightValue = ''
  if (!modelRef.value.expireTime) modelRef.value.expireTime = null
})

onUnmounted(() => {
  console.log('onUnmounted')
})

const labelCol = { span: 4 }
const wrapperCol = { span: 16 }

</script>

<style lang="less" scoped>
.processor_cookie_set-main {
  .icons {
    text-align: right;
    line-height: 32px;
  }
}
</style>
