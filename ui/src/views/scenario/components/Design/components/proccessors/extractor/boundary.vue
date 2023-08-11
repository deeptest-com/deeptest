<template>
  <div class="processor_extractor_boundary-main dp-processors-container">
    <ProcessorHeader/>
    <a-card :bordered="false">
      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">

          <a-form-item label="变量名称" v-bind="validateInfos.variable">
            <a-input v-model:value="modelRef.variable"
                     @blur="validate('variable', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>

          <a-form-item label="边界开始" v-bind="validateInfos.boundaryStart">
            <a-input v-model:value="modelRef.boundaryStart"
                     @blur="validate('boundaryStart', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>
          <a-form-item  label="边界结束" v-bind="validateInfos.boundaryEnd">
            <a-input v-model:value="modelRef.boundaryEnd"
                     @blur="validate('boundaryEnd', { trigger: 'blur' }).catch(() => {})" />
          </a-form-item>
          <a-form-item  label="索引值">
            <a-input-number v-model:value="modelRef.boundaryIndex" />
          </a-form-item>
          <a-form-item  label="是否包含边界">
            <a-switch v-model:checked="modelRef.boundaryIncluded" />
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
import {computed, onMounted, reactive, ref} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {useI18n} from "vue-i18n";
import {Form, message, notification} from 'ant-design-vue';
import {StateType as ScenarioStateType} from "../../../../../store";
import ProcessorHeader from '../../common/ProcessorHeader.vue';
const useForm = Form.useForm;

const router = useRouter();

const {t} = useI18n();

const formRef = ref();

const rulesRef = reactive({
  variable: [
    { required: true, message: '请输入变量名', trigger: 'blur' },
  ],
  boundaryStart:  [
      { required: true, message: '请输入边界开始字符串', trigger: 'blur' }
  ],
  boundaryEnd: [
      { required: true, message: '请输入边界结束字符串', trigger: 'blur' }
  ]
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
  if (!modelRef.value.variableName) modelRef.value.variableName = ''
  if (!modelRef.value.boundaryStart) modelRef.value.boundaryStart = ''
  if (!modelRef.value.boundaryEnd) modelRef.value.boundaryEnd = ''
  if (!modelRef.value.boundaryIndex) modelRef.value.boundaryIndex = 0
})



const labelCol = { span: 4 }
const wrapperCol = { span: 16 }

</script>
