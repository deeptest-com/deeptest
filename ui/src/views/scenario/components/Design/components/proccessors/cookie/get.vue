<template>
  <div class="processor_cookie_get-main dp-processors-container">
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

          <a-form-item label="赋予变量名" v-bind="validateInfos.variableName">
            <a-input v-model:value="modelRef.variableName"
                     @blur="validate('variableName', { trigger: 'blur' }).catch(() => {})"/>
            <div class="dp-input-tip">不存在会自动新建，已有的会被覆盖。</div>
          </a-form-item>

          <a-form-item label="默认值">
            <a-input v-model:value="modelRef.default"/>
            <div class="dp-input-tip">Cookie不存时的默认值</div>
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
import {computed, onMounted, onUnmounted, reactive, ref, watch} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {useI18n} from "vue-i18n";
import debounce from "lodash.debounce";
import {Form, notification} from 'ant-design-vue';
import {StateType as ScenarioStateType} from "../../../../../store";
import {NotificationKeyCommon} from "@/utils/const";
import ProcessorHeader from '../../common/ProcessorHeader.vue';
const useForm = Form.useForm;

const router = useRouter();

const {t} = useI18n();

const formRef = ref();

const rulesRef = reactive({
  cookieName: [
    {required: true, message: '请输入名称', trigger: 'blur'},
  ],
  variableName: [
    {required: true, message: '请输入接受变量的名称', trigger: 'blur'},
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
  if (!modelRef.value.variable) modelRef.value.variable = ''
})

onUnmounted(() => {
  console.log('onUnmounted')
})

const labelCol = { span: 4 }
const wrapperCol = { span: 16 }

</script>

<style lang="less" scoped>
.processor_cookie_get-main {
  .icons {
    text-align: right;
    line-height: 32px;
  }
}
</style>
