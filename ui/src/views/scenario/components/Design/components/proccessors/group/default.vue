<template>
  <div class="processor_group_default-main  dp-processors-container">
    <ProcessorHeader/>
    <a-card :bordered="false">
      <a-form
          :model="formState"
          :label-col="{ span: 4 }"
          :wrapper-col="{ span: 16 }">

        <a-form-item label="备注" name="comments">
          <a-textarea v-model:value="formState.comments" :rows="3"/>
        </a-form-item>

        <a-form-item class="processor-btn" :wrapper-col="{ span: 16, offset: 4 }">
          <a-button type="primary" @click.prevent="submit">保存</a-button>
        </a-form-item>
      </a-form>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import {computed, reactive, ref, watch} from "vue";
import {useStore} from "vuex";
import {StateType as ScenarioStateType} from "../../../../../store";
import {Form, message} from "ant-design-vue";
import ProcessorHeader from '../../common/ProcessorHeader.vue';
const useForm = Form.useForm;

const store = useStore<{ Scenario: ScenarioStateType; }>();
const nodeData: any = computed<boolean>(() => store.state.Scenario.nodeData);
const formState: any = ref({
  name: '',
  comments: '',
});

watch(nodeData, (val: any) => {
  if (!val) return;
  formState.value.name = val.name || '分组';
  formState.value.comments = val.comments;
},{
  immediate: true
});

const rulesRef = reactive({

})
const {resetFields, validate, validateInfos} = useForm(formState, rulesRef);

const submit = async () => {
  validate()
      .then(async () => {
        // 下面代码改成 await 的方式
        const res = await store.dispatch('Scenario/saveProcessor', {
          ...nodeData.value,
          name: formState.value.name,
          comments: formState.value.comments,
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

const reset = () => {
  resetFields();
};

</script>