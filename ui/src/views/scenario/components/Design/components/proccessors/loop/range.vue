<!-- ::::迭代区间 -->
<template>
  <div class="processor_loop_range-main dp-processors-container">
    <ProcessorHeader/>
    <a-card :bordered="false">
      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">

          <a-form-item label="变量名称" v-bind="validateInfos.variableName">
            <a-input v-model:value="modelRef.variableName"
                     @blur="validate('variableName', { trigger: 'blur' }).catch(() => {})"/>
          </a-form-item>

          <a-form-item label="区间" v-bind="validateInfos.range">
            <a-input v-model:value="modelRef.range"
                     @blur="validate('range', { trigger: 'blur' }).catch(() => {})"/>
            <div class="dp-input-tip">类似1-9或a-z的取值区间。</div>
          </a-form-item>
          <a-form-item label="间隔">
            <a-input-number v-model:value="modelRef.step"/>
          </a-form-item>
          <a-form-item label="是否随机">
            <a-switch v-model:checked="modelRef.isRand" />
          </a-form-item>

          <a-form-item label="备注" v-bind="validateInfos.comments">
            <a-textarea v-model:value="modelRef.comments" :rows="3"/>
          </a-form-item>

          <a-form-item class="processor-btn" :wrapper-col="{ span: 16, offset: 4 }">
            <a-button type="primary" @click.prevent="submitForm">保存</a-button>
<!--            <a-button style="margin-left: 10px" @click="resetFields">重置</a-button>-->
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
import ProcessorHeader from '../../common/ProcessorHeader.vue';
import {NotificationKeyCommon} from "@/utils/const";
import debounce from "lodash.debounce";

const useForm = Form.useForm;

const router = useRouter();

const {t} = useI18n();

const formRef = ref();

const rulesRef = reactive({
  variableName: [
    {required: true, message: '请输入变量名称', trigger: 'blur'},
  ],
  range: [
    {required: true, message: '请输入区间', trigger: 'blur'},
  ],
});

const store = useStore<{ Scenario: ScenarioStateType; }>();
const modelRef:any = computed<boolean>(() => store.state.Scenario.nodeData);
const {resetFields, validate, validateInfos} = useForm(modelRef, rulesRef);

const submitForm = debounce(async () => {
  validate()
      .then(() => {
        modelRef.value.step = modelRef.value.step + ''
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
  if (!modelRef.value.step) modelRef.value.step = 1
})

onUnmounted(() => {
  console.log('onUnmounted')
})

const labelCol = { span: 4 }
const wrapperCol = { span: 16 }

</script>

<style lang="less" scoped>
.processor_loop_range-main {
  .icons {
    text-align: right;
    line-height: 32px;
  }
}
</style>
