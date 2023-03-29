<template>
    <div class="title">全局变量</div>
    <a-button class="envDetail-btn" @click="addGlobalVar">
        <template #icon>
            <PlusOutlined />
        </template>
        添加
    </a-button>

    <a-table bordered size="small" :pagination="false" :columns="globalVarsColumns" :data-source="globalVarsData">
        <template #customName="{ text, index }">
            <a-input @change="(e) => {
                handleGlobalVarsChange('name', index, e);
            }" :value="text" placeholder="请输入参数名" />
        </template>
        <template #customLocalValue="{ text, index }">
            <a-input :value="text" @change="(e) => {
                handleGlobalVarsChange('localValue', index, e);
            }" placeholder="请输入本地值" />
        </template>
        <template #customRemoteValue="{ text, index }">
            <a-input :value="text" @change="(e) => {
                handleGlobalVarsChange('remoteValue', index, e);
            }" placeholder="请输入远程值" />
        </template>
        <template #customDescription="{ text, index }">
            <a-input :value="text" @change="(e) => {
                handleGlobalVarsChange('description', index, e);
            }" placeholder="请输入描述信息" />
        </template>
        <template #customAction="{ index }">
            <a-button danger type="text" @click="handleGlobalVarsChange('description', index, '', 'delete');"
                :size="'small'">删除
            </a-button>
        </template>
    </a-table>

    <div class="envDetail-footer">
        <a-button class="save-btn" @click="handleSaveGlobalVars" type="primary">保存</a-button>
    </div>
</template>
<script setup lang="ts">
import { computed } from "vue";
import { useStore } from "vuex";
import { globalVarsColumns } from '../../config';
import { StateType as ProjectStateType } from "@/store/project";
import { StateType as ProjectSettingStateType } from "@/views/ProjectSetting/store";

// store 相关
const store = useStore<{ ProjectGlobal: ProjectStateType, ProjectSetting: ProjectSettingStateType }>();
const globalVarsData = computed<any>(() => store.state.ProjectSetting.globalVarsData);

const emits = defineEmits(['addGlobalVar', 'handleSaveGlobalVars', 'handleGlobalVarsChange']);

function addGlobalVar() {
    emits('addGlobalVar');
}

function handleSaveGlobalVars() {
    emits('handleSaveGlobalVars');
}

function handleGlobalVarsChange(field: string, index: number, e: any, action?: string) {
    emits('handleGlobalVarsChange', field, index, e, action);

}
</script>
<style lang="less" scoped>
.title {
    font-weight: bold;
    font-size: 18px;
    margin-bottom: 16px;
}

.envDetail-btn {
    margin-top: 16px;
    margin-bottom: 16px;
}

.envDetail-footer {
    height: 60px;
    position: absolute;
    top: 8px;
    right: 8px;
    width: 300px;
    display: flex;
    align-items: center;
    justify-content: flex-end;

    .save-btn {
        margin-right: 16px;
    }
}
</style>