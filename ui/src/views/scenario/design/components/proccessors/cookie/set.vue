<template>
  <div class="processor_cookie_set-main">
    <a-card :bordered="false">
      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">

          <a-form-item label="备注" v-bind="validateInfos.comments">
            <a-input v-model:value="modelRef.comments"/>
          </a-form-item>

          <a-form-item label="Cookie名称" v-bind="validateInfos.cookieName">
            <a-input v-model:value="modelRef.cookieName"
                     @blur="validate('cookieName', { trigger: 'blur' }).catch(() => {})"/>
          </a-form-item>

          <a-form-item label="取值" v-bind="validateInfos.rightValue">
            <a-input v-model:value="modelRef.rightValue"
                     @blur="validate('rightValue', { trigger: 'blur' }).catch(() => {})"/>
            <div class="dp-tip-small">常量或用${name}表示的变量</div>
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
import {computed, reactive, ref} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {useI18n} from "vue-i18n";
import {Form, message} from 'ant-design-vue';
import {StateType as ScenarioStateType} from "../../../../store";
import {EditOutlined, CheckOutlined, CloseOutlined} from "@ant-design/icons-vue";

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
const modelRef = computed<boolean>(() => store.state.Scenario.nodeData);
const {resetFields, validate, validateInfos} = useForm(modelRef, rulesRef);

const submitForm = async () => {
  validate()
      .then(() => {
        store.dispatch('Scenario/saveProcessor', modelRef.value).then((res) => {
          if (res === true) {
            message.success(`保存成功`);
          } else {
            message.error(`保存失败`);
          }
        })
      })
};

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
