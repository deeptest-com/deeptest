<!-- 本页面是服务编辑页面的抽屉 -->
<template>
    <a-drawer :closable="true" :width="1000" :key="editKey" :bodyStyle="{padding:'16px'}" :visible="drawerVisible" @close="onClose">
        <template #title>
            <div class="drawer-header">
                <div>服务编辑</div>
            </div>
        </template>
        <div class="drawer-content">
            <a-form :model="formState" :label-col="{ span: 2 }" :wrapper-col="{ span: 15 }"> 
                <ConBoxTitle :backgroundStyle="'background: #FBFBFB;'" :title="'基本信息'" />
                <div class="drawer-basicinfo">
                    <a-form-item label="服务名称">
                        <EditAndShowField :placeholder="'请输入服务名称'" :value="formState.name" @update="(e: string) => changeServiceInfo({ name: e })"/>
                    </a-form-item>
                    <a-form-item label="描述">
                        <EditAndShowField :placeholder="'请输入服务简介描述'" :value="formState.description" @update="(e: string) => changeServiceInfo({ description: e })"/>
                    </a-form-item>
                </div>
                <ConBoxTitle :backgroundStyle="'background: #FBFBFB;'" :title="'服务管理'" />
                <a-tabs :activeKey="activeKey" @change="handleTabChange">
                    <a-tab-pane key="service-version" tab="服务版本">
                        <ServiceVersion :serveId="formState.id" />
                    </a-tab-pane>
                    <a-tab-pane key="service-component" tab="服务组件">
                        <ServiceComponent :serveId="formState.id" />
                    </a-tab-pane>
                    <a-tab-pane key="service-security" tab="Security">
                        <ServiceSecurity :serveId="formState.id"/>
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
    watch
} from 'vue';
import { useStore } from 'vuex';
import { message } from 'ant-design-vue';
import ServiceVersion from './Version.vue';
import ServiceComponent from './Component.vue';
import ServiceSecurity from './Security.vue';
import EditAndShowField from '@/components/EditAndShow/index.vue';
import ConBoxTitle from '@/components/ConBoxTitle/index.vue';
import { StateType as ProjectStateType } from "@/store/project";
import { StateType as ProjectSettingStateType } from '../../store';
import { ServeDetail } from '../../data';

const store = useStore<{ ProjectGlobal: ProjectStateType, ProjectSetting: ProjectSettingStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const formState = computed<ServeDetail>(() => store.state.ProjectSetting.selectServiceDetail);

const props = defineProps<{
    drawerVisible: boolean
    tabKey: string
    editKey?: number
}>();

const emits = defineEmits(['onClose', 'update:formState', 'update:tabKey']);
console.log('props-------- tabKey', props.tabKey);
const activeKey = ref('service-version');

function onClose() {
    emits('onClose');
}

async function changeServiceInfo(updateFieldInfo: any) {
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

function handleTabChange(value: string) {
    emits('update:tabKey', value);
}

watch(() => {
    return props.tabKey;
}, (val) => {
    activeKey.value = val || 'service-version';
}, {
    immediate: true
})
</script>
<style lang="less" scoped>
.drawer-basicinfo {
    padding: 9px 0 33px 0;
}
</style>
