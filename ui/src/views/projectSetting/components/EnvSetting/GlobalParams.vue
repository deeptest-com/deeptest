<template>
    <div class="title">全局参数</div>

    <a-tabs :pagination="false" v-model:activeKey="globalParamsActiveKey">
        <a-tab-pane v-for="(tabItem) in tabPaneList" :key="tabItem.type" :tab="tabItem.name">

            <a-button class="envDetail-btn" @click="addGlobalParams">
                <template #icon>
                    <PlusOutlined />
                </template>
                添加
            </a-button>

            <a-table size="small" bordered :pagination="false" :columns="globalParamscolumns"
                :data-source="globalParamsData?.[tabItem.name] || []">
                <template #customName="{ text, index }">
                    <a-input :value="text" @change="(e) => {
                        handleGlobalParamsChange(tabItem.name, 'name', index, e);
                    }" placeholder="请输入参数名" />
                </template>
                <template #customType="{ text, index }">
                    <a-select class="custom-select" :value="text" style="width: 120px" @change="(e) => {
                        handleGlobalParamsChange(tabItem.name, 'type', index, e)
                    }">
                        <a-select-option value="string">string</a-select-option>
                        <a-select-option value="number">number</a-select-option>
                        <a-select-option value="integer">integer</a-select-option>
                    </a-select>
                </template>
                <template #customRequired="{ text, index }">
                    <a-switch :checked="text" @change="(checked) => {
                        handleGlobalParamsChange(tabItem.name, 'required', index, checked)
                    }" />
                </template>
                <template #customDefaultValue="{ text, index }">
                    <a-input :value="text" @change="(e) => {
                        handleGlobalParamsChange(tabItem.name, 'defaultValue', index, e);
                    }" placeholder="默认值" />
                </template>
                <template #customDescription="{ text, index }">
                    <a-input :value="text" @change="(e) => {
                        handleGlobalParamsChange(tabItem.name, 'description', index, e);
                    }" placeholder="说明" />
                </template>
                <template #customAction="{ index }">
                    <a-button danger type="text" @click="handleGlobalParamsChange(tabItem.name, '', index, '', 'delete');"
                        :size="'small'">删除
                    </a-button>
                </template>
            </a-table>
        </a-tab-pane>
    </a-tabs>

    <div class="envDetail-footer">
        <a-button class="save-btn" @click="handleSaveGlobalParams" type="primary">保存</a-button>
    </div>
</template>
<script setup lang="ts">
import { ref, computed, createVNode, watch } from 'vue';
import { useStore } from 'vuex';
import { Modal } from 'ant-design-vue';
import { ExclamationCircleOutlined } from '@ant-design/icons-vue';
import { globalParamscolumns, tabPaneList } from '../../config';
import { StateType as ProjectStateType } from "@/store/project";
import { StateType as ProjectSettingStateType } from "@/views/projectSetting/store";

const store = useStore<{ ProjectGlobal: ProjectStateType, ProjectSetting: ProjectSettingStateType }>();
const globalParamsData = computed<any>(() => store.state.ProjectSetting.globalParamsData);
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const globalParamsActiveKey = ref('header');

function getParamsData() {
    store.dispatch('ProjectSetting/getEnvironmentsParamList', { projectId: currProject.value.id });
}

function addGlobalParams() {
    store.dispatch('ProjectSetting/addGlobalParams', { globalParamsActiveKey });
}

function handleGlobalParamsChange(type: string, field: string, index: number, e: any, action?: string) {
    const confirmCallBack = () => store.dispatch('ProjectSetting/handleGlobalParamsChange', { type, field, index, e, action });
    if (action && action === 'delete') {
        Modal.confirm({
            title: '确认要删除该参数吗',
            icon: createVNode(ExclamationCircleOutlined),
            onOk() {
                confirmCallBack();
            },
        });
    } else {
        confirmCallBack();
    }
}

function handleSaveGlobalParams() {
    store.dispatch('ProjectSetting/saveEnvironmentsParam', { projectId: currProject.value.id });
}

watch(() => {
    return currProject.value
}, () => {
    getParamsData();
})
</script>
<style scoped lang="less">
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
