<template>
  <div class="processor_group_default-main dp-processors-container">
    <ProcessorHeader/>
    <a-card :bordered="false">
      <a-form
          :model="formState"
          :label-col="{ span: 4 }"
          :wrapper-col="{ span: 16 }">
        <a-form-item label="跳出如果满足" name="breakIfExpression">
          <a-input  v-model:value="formState.breakIfExpression"/>
          <div class="dp-input-tip">{{t('tips_expression_bool', {name: '{name}'})}}</div>
        </a-form-item>

        <a-form-item label="备注" name="comments">
          <a-textarea v-model:value="formState.comments" :rows="3"/>
        </a-form-item>

        <a-form-item :wrapper-col="{ span: 16, offset: 4 }">
          <a-button type="primary" @click.prevent="submit">保存</a-button>
        </a-form-item>
      </a-form>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import {computed, ref, watch} from "vue";
import {useStore} from "vuex";
import {StateType as ScenarioStateType} from "../../../../../store";
import {Form, message} from "ant-design-vue";
import {useI18n} from "vue-i18n";
import ProcessorHeader from '../../common/ProcessorHeader.vue';
const {t} = useI18n();
const useForm = Form.useForm;
const store = useStore<{ Scenario: ScenarioStateType; }>();
const nodeData: any = computed<boolean>(() => store.state.Scenario.nodeData);

const formState: any = ref({
  breakIfExpression: '',
  comments: '',
});

watch(() => {
  return nodeData.value;
}, (val: any) => {
  if (!val) return;
  formState.value.breakIfExpression = val.breakIfExpression;
  formState.value.comments = val.comments;
},{
  immediate: true,
});

const rulesRef = {
  breakIfExpression: [
    {required: true, message: '请输入跳出循环满足的条件表达式', trigger: 'blur'},
  ],
}
const {resetFields, validate, validateInfos} = useForm(formState, rulesRef);

const submit = async () => {
  validate()
      .then(async () => {
        // 下面代码改成 await 的方式
        const res = await store.dispatch('Scenario/saveProcessor', {
          ...nodeData.value,
          breakIfExpression: formState.value.breakIfExpression,
        });
        if (res === true) {
          message.success('保存成功');
        } else {
          message.error('保存失败');
        }
      })
      .catch(error => {
        console.log('error', error);
      });
};

</script>
