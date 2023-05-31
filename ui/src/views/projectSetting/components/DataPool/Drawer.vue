<!-- 本页面是数据池编辑页面的抽屉 -->
<template>
  <div class="datapool-main">
    <a-drawer :closable="true" :width="1000" :key="editKey" :bodyStyle="{padding:'16px'}" :visible="drawerVisible"
              @close="onClose">
        <template #title>
            <div class="drawer-header">
                <div>数据池编辑</div>
            </div>
        </template>
        <div class="drawer-content">
            <a-form :model="formState" :label-col="{ span: 2 }" :wrapper-col="{ span: 15 }">

            </a-form>
        </div>
    </a-drawer>
  </div>
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
import { StateType as ProjectStateType } from "@/store/project";
import { StateType as ProjectSettingStateType } from '../../store';
import { ServeDetail } from '../../data';

const store = useStore<{ ProjectGlobal: ProjectStateType, ProjectSetting: ProjectSettingStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const formState = computed<ServeDetail>(() => store.state.ProjectSetting.selectServiceDetail);

const props = defineProps<{
    drawerVisible: boolean
    editKey?: number
}>();

const emits = defineEmits(['onClose', 'update:formState']);

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

</script>

<style lang="less" scoped>
.datapool-main {

}
</style>
