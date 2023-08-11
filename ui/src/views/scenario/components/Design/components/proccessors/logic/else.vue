<template>
  <div class="processor_login_else-main  dp-processors-container">
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
import {Form, message, notification} from "ant-design-vue";
import ProcessorHeader from '../../common/ProcessorHeader.vue';
import debounce from "lodash.debounce";
const useForm = Form.useForm;

const store = useStore<{ Scenario: ScenarioStateType; }>();
const nodeData: any = computed<boolean>(() => store.state.Scenario.nodeData);
const formState: any = ref({
  comments: '',
});

watch(() => {
  return nodeData.value;
}, (val: any) => {
  if (!val) return;
  formState.value.comments = val.comments || '';
},{
  immediate: true,
});

const rulesRef = reactive({
  name: [
    {required: true, message: '请输入名称', trigger: 'blur'},
  ],
})
const {resetFields, validate, validateInfos} = useForm(formState, rulesRef);

const submit = debounce(async () => {
  validate()
      .then(async () => {
        // 下面代码改成 await 的方式
        const res = await store.dispatch('Scenario/saveProcessor', {
          ...nodeData.value,
          comments: formState.value.comments,
        });
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
      .catch(error => {
        console.log('error', error);
      });
}, 300);

</script>

