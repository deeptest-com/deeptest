<!-- 迭代列表 -->
<template>
  <div class="processor_loop_in-main dp-proccessors-container">
    <ProcessorHeader/>
    <a-card :bordered="false">
      <div>
        <a-form :label-col="labelCol" :wrapper-col="wrapperCol">

          <a-form-item label="变量名称" v-bind="validateInfos.variableName">
            <a-input v-model:value="modelRef.variableName"
                     @blur="validate('variableName', { trigger: 'blur' }).catch(() => {})"/>
          </a-form-item>

          <a-form-item label="列表" v-bind="validateInfos.list">
            <a-input v-model:value="modelRef.list"
                     @blur="validate('list', { trigger: 'blur' }).catch(() => {})"/>
            <div class="dp-input-tip">列表以逗号分隔</div>
          </a-form-item>

          <a-form-item label="是否随机">
            <a-switch v-model:checked="modelRef.isRand" />
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
import ProcessorHeader from '../../common/ProcessorHeader.vue';
import {StateType as ScenarioStateType} from "../../../../../store";
const useForm = Form.useForm;

const router = useRouter();

const {t} = useI18n();

const formRef = ref();

const rulesRef = reactive({
  variableName: [
    {required: true, message: '请输入变量名称', trigger: 'blur'},
  ],
  list: [
    {required: true, message: '请输入列表', trigger: 'blur'},
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
  if (!modelRef.value.repeatTimes) modelRef.value.repeatTimes = 1
})

onUnmounted(() => {
  console.log('onUnmounted')
})

const labelCol = { span: 4 }
const wrapperCol = { span: 16 }

</script>

<style lang="less" scoped>
.processor_loop_in-main {
  .icons {
    text-align: right;
    line-height: 32px;
  }
}
</style>
