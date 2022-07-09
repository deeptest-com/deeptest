<template>
  <div class="processor_login_if-main">
    <a-card :bordered="false">
      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">

          <a-form-item label="备注" v-bind="validateInfos.comments">
            <a-input v-model:value="modelRef.comments"/>
          </a-form-item>

          <a-form-item label="左值" v-bind="validateInfos.leftValue">
            <a-input v-model:value="modelRef.leftValue"
                     @blur="validate('leftValue', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>

          <a-form-item label="操作" v-bind="validateInfos.operator">
            <a-select v-model:value="modelRef.operator"
                      @blur="validate('operator', { trigger: 'change' }).catch(() => {})">
              <a-select-option v-for="(item, idx) in operators" :key="idx" :value="item.value">
                {{ t(item.label) }}
              </a-select-option>
            </a-select>
          </a-form-item>

          <a-form-item label="右值" v-bind="validateInfos.rightValue">
            <a-input v-model:value="modelRef.rightValue"
                     @blur="validate('rightValue', { trigger: 'blur' }).catch(() => {})"/>
          </a-form-item>

          <a-form-item :wrapper-col="{ offset: 2 }">
            左值、右值可以是常量或形如${name}的变量。
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
import {computed, reactive, ref} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {useI18n} from "vue-i18n";
import {Form, message} from 'ant-design-vue';
import {StateType as ScenarioStateType} from "../../../../store";
import {EditOutlined, CheckOutlined, CloseOutlined} from "@ant-design/icons-vue";
import {getCompareOpts, getCompareOptsForString} from "@/utils/compare";

const useForm = Form.useForm;

const router = useRouter();

const {t} = useI18n();

const formRef = ref();

const rulesRef = reactive({
  leftValue: [
    {required: true, message: '请输入左值', trigger: 'blur'},
  ],
  rightValue: [
    {required: true, message: '请输入右值', trigger: 'blur'},
  ],
  operator: [
    {required: true, message: '请选择操作', trigger: 'blur'},
  ],
});

const store = useStore<{ Scenario: ScenarioStateType; }>();
const modelRef = computed<boolean>(() => store.state.Scenario.nodeData);
const {resetFields, validate, validateInfos} = useForm(modelRef, rulesRef);

const operators = getCompareOpts()

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
.processor_login_if-main {
  .icons {
    text-align: right;
    line-height: 32px;
  }
}
</style>
