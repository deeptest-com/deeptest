<template>
  <div class="project-settings-mock-main">
    <a-form :label-col="labelCol" :wrapper-col="wrapperCol">
      <a-form-item label="Mock引擎">
        <div>Mock.js</div>
      </a-form-item>

      <a-form-item label="Mock优先方式">
        <a-radio-group name="radioGroup"
                       v-model:value="modelRef.priority"
                       @blur="validate('priority', { trigger: 'change' }).catch(() => {})">
          <a-radio value="smart">智能Mock优先</a-radio>
          <a-radio value="example">响应示例优先</a-radio>
        </a-radio-group>

        <div class="dp-input-tip">
          Mock优先级：高级Mock >
          <template v-if="modelRef.priority === 'example'"> 响应示例 > </template>
          智能Mock
        </div>
      </a-form-item>

      <a-form-item :wrapper-col="{ offset: 5 }">
        <a-button type="primary" @click="onSubmit" :disabled="!dataChanged && modelRef.id > 0">保存</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup lang="ts">
import {computed, onMounted, ref, watch} from 'vue';
import {useStore} from "vuex";
import {Form} from 'ant-design-vue';
import {notifyError, notifySuccess} from "@/utils/notify";

const useForm = Form.useForm;
const store = useStore<{ Endpoint, ProjectGlobal, ProjectSetting }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const modelRef = computed<any>(() => store.state.ProjectSetting.mockSettings);

const rules = ref({
  priority: [
    {required: true, message: '请选择优先设置', trigger: 'blur'},
  ],
})

const {validate, validateInfos} = useForm(modelRef, rules);

const dataLoaded = ref(false)
onMounted(async () => {
  await store.dispatch('ProjectSetting/getMock');
  dataLoaded.value = true
})

const dataChanged = ref(false)
watch(modelRef, (val) => {
  if (!dataLoaded.value) return
  dataChanged.value = true
}, {immediate: false, deep: true});

const onSubmit = () => {
  validate().then(async () => {
    const res = await store.dispatch('ProjectSetting/saveMock', modelRef.value)
    if (res) dataChanged.value = false
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
