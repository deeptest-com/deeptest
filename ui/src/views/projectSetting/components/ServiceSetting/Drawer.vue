<!-- 本页面是服务编辑页面的抽屉 -->
<template>
    <a-drawer :closable="true" :width="1000" :key="editKey" :visible="drawerVisible" @close="onClose">
        <template #title>
            <div class="drawer-header">
                <div>服务编辑</div>
            </div>
        </template>
        <div class="drawer-content">
            <a-form :model="formState" :label-col="{ span: 2 }" :wrapper-col="{ span: 15 }">
                <a-form-item label="服务名称">
                    <EditAndShowField :value="formState.name" @update="(e: string) => changeServiceInfo({ name: e })"/>
                </a-form-item>
                <a-form-item label="描述">
                    <EditAndShowField :value="formState.description" @update="(e: string) => changeServiceInfo({ description: e })"/>
                </a-form-item>
                <a-tabs v-model:activeKey="activeKey">
                    <a-tab-pane key="1" tab="服务版本">
                        <ServiceVersion :serveId="formState.id" />
                    </a-tab-pane>
                    <a-tab-pane key="2" tab="服务组件">
                        <ServiceComponent :serveId="formState.id" />
                    </a-tab-pane>
                </a-tabs>
            </a-form>
        </div>
    </a-drawer>
</template>
<script setup lang="ts">
import {
    ref,
    computed,
    defineEmits,
    defineProps,
} from 'vue';
import { useStore } from 'vuex';
import { message } from 'ant-design-vue';
import ServiceVersion from './Version.vue';
import ServiceComponent from './Component.vue';
import EditAndShowField from '@/components/EditAndShow/index.vue'; 
import { StateType as ProjectStateType } from "@/store/project";
import { StateType as ProjectSettingStateType } from '../../store';
import { ServeDetail } from '../../data';

const store = useStore<{ ProjectGlobal: ProjectStateType, ProjectSetting: ProjectSettingStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const formState = computed<ServeDetail>(() => store.state.ProjectSetting.selectServiceDetail);

defineProps<{
    drawerVisible: boolean
    editKey?: number
}>();

const emits = defineEmits(['onClose', 'update:formState']);


const activeKey = ref('1');
const isEditServiceDesc = ref(false);
const isEditServiceName = ref(false);


function onClose() {
    emits('onClose');
}

async function changeServiceInfo(updateFieldInfo: any) {
    isEditServiceDesc.value = false;
    isEditServiceName.value = false;
    const serviceInfo = { ...formState.value, ...updateFieldInfo };
    if (!serviceInfo.name) {
        message.error('服务名称不能为空');
    }
    await store.dispatch('ProjectSetting/saveStoreServe', {
        "projectId": currProject.value.id,
        formState: { ...serviceInfo },
        action: 'update'
    });
}
</script>