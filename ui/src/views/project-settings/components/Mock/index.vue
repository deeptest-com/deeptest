<template>
  <div class="project-settings-mock-main">
    <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
      <a-form-item label="Mock引擎">
        <div>Mock.js</div>
      </a-form-item>

      <a-form-item label="Mock优先方式">
        <a-radio-group name="radioGroup"
                       v-model:value="modelRef.priority"
                       @blur="validate('name', { trigger: 'change' }).catch(() => {})">
          <a-radio value="smart">智能Mock优先</a-radio>
          <a-radio value="sample">响应示例优先</a-radio>
        </a-radio-group>

        <div class="dp-input-tip">Mock优先级：高级Mock > 智能Mock</div>
      </a-form-item>

      <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
        <a-button type="primary" @click="onSubmit" :disabled="disabled">保存</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup lang="ts">
import {computed, ref, watch} from 'vue';
import {useStore} from "vuex";
import {Form} from 'ant-design-vue';
import {notifyError, notifySuccess} from "@/utils/notify";

const useForm = Form.useForm;
const store = useStore<{ Endpoint, ProjectGlobal, ProjectSetting }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const treeDataCategory = computed<any>(() => store.state.Endpoint.treeDataCategory);

const modelRef = ref({priority: 'smart'} as any)
const rules = {
  priority: [
    {required: true, message: '请选择优先设置', trigger: 'blur'},
  ],
};

const {validate, validateInfos} = useForm(modelRef, rules);

const disabled = ref()
watch(() => modelRef.value, (val) => {
  disabled.value = false
}, {immediate: false, deep: true});

const onSubmit = () => {
  validate().then(async () => {
    const res = await store.dispatch('ProjectSetting/saveSwaggerSync', modelRef.value);
    if (res.code === 0) {
      notifySuccess('保存成功');
    } else {
      notifyError('保存失败')
    }
  })
}

const labelCol = {span: 4}
const wrapperCol = {span: 18, offset: 1}

</script>

<style lang="less" scoped>
  .project-settings-mock-main {
    padding: 36px;
  }
</style>
