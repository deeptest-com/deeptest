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
                    <EditAndShowField :placeholder="'请输入服务名称'" :value="formState.name" @update="(e: string) => changeServiceInfo({ name: e })"/>
                </a-form-item>
                <a-form-item label="描述">
                    <EditAndShowField :placeholder="'请输入服务简介描述'" :value="formState.description" @update="(e: string) => changeServiceInfo({ description: e })"/>
                </a-form-item>
                <a-tabs v-model:activeKey="activeKey">
                    <a-tab-pane key="1" tab="服务版本">
                        <ServiceVersion :serveId="formState.id" />
                    </a-tab-pane>
                    <a-tab-pane key="2" tab="服务组件">
                        <ServiceComponent :serveId="formState.id" />
                    </a-tab-pane>
                    <a-tab-pane key="3" tab="Security">
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
    onMounted
} from 'vue';
import { useStore } from 'vuex';
import { message } from 'ant-design-vue';
import ServiceVersion from './Version.vue';
import ServiceComponent from './Component.vue';
import ServiceSecurity from './Security.vue';
import EditAndShowField from '@/components/EditAndShow/index.vue'; 
import { StateType as ProjectStateType } from "@/store/project";
import { StateType as ProjectSettingStateType } from '../../store';
import { ServeDetail } from '../../data';
import { placeholder } from '@babel/types';

const store = useStore<{ ProjectGlobal: ProjectStateType, ProjectSetting: ProjectSettingStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const formState = computed<ServeDetail>(() => store.state.ProjectSetting.selectServiceDetail);

const props = defineProps<{
    drawerVisible: boolean
    editKey?: number
    params:any
    
   
}>();

const emits = defineEmits(['onClose', 'update:formState']);

    

const activeKey = ref('1');
const isEditServiceDesc = ref(false);
const isEditServiceName = ref(false);

onMounted(()=>{
    console.log('~~~~~~~~~drawer props.params',props.params)
    if(props.params?.sectab){
        const sectab:string=props.params?.sectab 
        activeKey.value=sectab
    }  
    
})

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