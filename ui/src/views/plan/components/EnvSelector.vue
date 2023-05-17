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
                    :rules="[{ required: true, message: '请选择执行环境' }]"
                    >
                    <a-select @change="changeEnv" v-model:value="currEnvId" placeholder="请选择" :options="envList" />
                </a-form-item>
            </a-form>
        </div>
    </a-modal>
</template>
<script setup lang="ts">
import { defineProps, defineEmits, ref, reactive, computed, watch } from 'vue';
import { useStore } from 'vuex';
import { Form } from 'ant-design-vue';
import { StateType as PlanStateType } from "../store";
import { StateType as ProjectStateType } from "@/store/project";
import { StateType as ProjectSettingStateType } from "@/views/projectSetting/store";

const props = defineProps<{
    envSelectDrawerVisible: Boolean
}>();

const useForm = Form.useForm;
const emits = defineEmits(['onCancel', 'onOk']);
const store = useStore<{ Plan: PlanStateType, ProjectSetting: ProjectSettingStateType, ProjectGlobal: ProjectStateType }>();
const envList = computed<any>(() => store.state.ProjectSetting.envList);
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const selectEnvId = computed<any>(() => store.state.ProjectSetting.selectEnvId);
// activity.esm-bundler.js:1188 Write operation failed: computed value is readonly
const currEnvId = computed({
    get() {
        return selectEnvId.value;
    },
    set(val) {
        return val;
    }
});  
const labelCol = { span: 5 };
const wrapperCol = { span: 17 };
const formRef = ref();

function onCancel() {
    emits('onCancel');
}

async function changeEnv(value) {
    await store.dispatch('ProjectSetting/saveEnvId', value);
}

async function save() {
    console.log(selectEnvId.value);
    emits('onOk');
}

watch(() => {
    return props;
}, (val) => {
    console.log(val);
    if (props.envSelectDrawerVisible) {
        store.dispatch('ProjectSetting/getEnvsList', {
            projectId: currProject.value.id
        });
    }
}, {
    immediate: true,
    deep: true
})
</script>