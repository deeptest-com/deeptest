<template>
  <a-modal
      title="选择执行环境"
      :visible="envSelectDrawerVisible"
      class="env-selector"
      :closable="true"
      @cancel="onCancel"
      @ok="save"
      width="600px">
    <div class="env-selector-main">
      <a-form :label-col="labelCol" :wrapper-col="wrapperCol" :ref="formRef">
        <a-form-item
            label="执行环境"
            has-feedback
            :rules="[{ required: true, message: '请选择执行环境' }]">
          <a-select @change="changeEnv" v-model:value="currEnvId" placeholder="请选择" :options="envList"/>
        </a-form-item>
      </a-form>
    </div>
  </a-modal>
</template>
<script setup lang="ts">
import {defineProps, defineEmits, ref, reactive, computed, watch} from 'vue';
import {useStore} from 'vuex';
import {Form} from 'ant-design-vue';
import {StateType as ProjectStateType} from "@/store/project";
import {StateType as ProjectSettingStateType} from "@/views/project-settings/store";
import {Scenario} from "@/views/scenario/data";

const props = defineProps([
  'envSelectDrawerVisible',
  'execEnvId'
])

const useForm = Form.useForm;
const emits = defineEmits(['onCancel', 'onOk']);
const store = useStore<{ Plan, ProjectSetting: ProjectSettingStateType, ProjectGlobal: ProjectStateType, Scenario }>();
const envList = computed<any>(() => store.state.ProjectSetting.envList);
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const selectEnvId = computed<any>(() => store.state.ProjectSetting.selectEnvId);
// const envId: any = computed<Scenario>(() => store.state.Scenario.detailResult?.currEnvId);
const currEnvId = ref(null);

const labelCol = {span: 5};
const wrapperCol = {span: 17};
const formRef = ref();

function onCancel() {
  emits('onCancel');
}

async function changeEnv(value) {
  currEnvId.value = value;
  await store.dispatch('ProjectSetting/saveEnvId', value);
}

async function save() {
  emits('onOk');
}

watch(() => {
  return [props.envSelectDrawerVisible, props.execEnvId];
}, async (val) => {
  const [visible, envId] = val;
  if (visible) {
    await store.dispatch('ProjectSetting/getEnvsList', {
      projectId: currProject.value.id
    });

    if (envId) {
      currEnvId.value = envId;
      store.dispatch('ProjectSetting/saveEnvId',envId); 
    }
  }
}, {
  immediate: true
})

</script>
