<template>
    <!-- ::::环境详情 -->
    <a-form :model="activeEnvDetail">
        <div class="title">{{ activeEnvDetail.displayName }}</div>
        <div class="envDetail-content">
            <a-form-item :labelCol="{ span: 2 }" :wrapperCol="{ span: 10 }" label="环境名称" name="name"
                :rules="[{ required: true, message: '环境名称不能为空' }]">
                <a-input class="env-name" :value="activeEnvDetail.name || ''" @change="handleEnvNameChange"
                    placeholder="请输入环境名称" />
            </a-form-item>
            <div class="serveServers">
                <div class="serveServers-header">服务 (前置URL)</div>
                <a-button class="envDetail-btn" @click="addService">
                    <template #icon>
                        <PlusOutlined />
                    </template>
                    关联服务
                </a-button>
                <a-table v-if="activeEnvDetail.serveServers.length > 0" size="small" bordered :pagination="false"
                    :columns="serveServersColumns" :data-source="activeEnvDetail.serveServers">
                    <template #customName="{ text, index }">
                        <a-input :value="text" @change="(e) => {
                            handleEnvChange('serveServers', 'serveName', index, e);
                        }" placeholder="请输入参数名" />
                    </template>

                    <template #customUrl="{ text, index }">
                        <a-input :value="text" @change="(e) => {
                            handleEnvChange('serveServers', 'url', index, e);
                        }" placeholder="http 或 https 起始的合法 URL" />
                    </template>
                </a-table>
            </div>
            <div class="vars">
                <div class="vars-header">环境变量</div>
                <a-button class="envDetail-btn" @click="addVar">
                    <template #icon>
                        <PlusOutlined />
                    </template>
                    添加
                </a-button>
                <a-table v-if="activeEnvDetail.vars.length > 0" bordered size="small" :pagination="false"
                    :columns="globalVarsColumns" :data-source="activeEnvDetail.vars">
                    <template #customName="{ text, index }">
                        <a-input @change="(e) => {
                            handleEnvChange('vars', 'name', index, e);
                        }" :value="text" placeholder="请输入参数名" />
                    </template>
                    <template #customLocalValue="{ text, index }">
                        <a-input :value="text" @change="(e) => {
                            handleEnvChange('vars', 'localValue', index, e);
                        }" placeholder="请输入本地值" />
                    </template>
                    <template #customRemoteValue="{ text, index }">
                        <a-input :value="text" @change="(e) => {
                            handleEnvChange('vars', 'remoteValue', index, e);
                        }" placeholder="请输入远程值" />
                    </template>
                    <template #customDescription="{ text, index }">
                        <a-input :value="text" @change="(e) => {
                            handleEnvChange('vars', 'description', index, e);
                        }" placeholder="请输入描述信息" />
                    </template>
                    <template #customAction="{ index }">
                        <a-button danger type="text" @click="handleEnvChange('vars', '', index, '', 'delete');"
                            :size="'small'">删除
                        </a-button>
                    </template>
                </a-table>
            </div>
        </div>
        <div class="envDetail-footer">
            <a-button v-if="activeEnvDetail.id" class="save-btn" @click="deleteEnvData" type="danger">删除</a-button>
            <a-button v-if="activeEnvDetail.id" class="save-btn" @click="copyEnvData" type="primary">复制</a-button>
            <a-button class="save-btn" @click="addEnvData" html-type="submit" type="primary">保存</a-button>
        </div>
    </a-form>
</template>
<script setup lang="ts">
import { ref, defineEmits, defineProps } from 'vue';
import { globalVarsColumns, serveServersColumns } from '../../config';

const props = defineProps<{
    activeEnvDetail: any
}>();
const emits = defineEmits(['deleteEnvData', 'copyEnvData', 'addEnvData', 'handleEnvChange', 'handleEnvNameChange', 'addVar', 'addService']);

function deleteEnvData() {
    emits('deleteEnvData');
}
function copyEnvData() {
    emits('copyEnvData');
}
function addEnvData() {
    emits('addEnvData');
}

function handleEnvChange(type: string, filed: string, index: number, e: any, action?: string) {
    emits('handleEnvChange', type, filed, index, e, action);
}
function handleEnvNameChange(e: any) {
    emits('handleEnvNameChange', e);
}

function addVar() {
    emits('addVar');
}
function addService() {
    emits('addService');
}
</script>

<style scoped lang="less">
.title {
    font-weight: bold;
    font-size: 18px;
    margin-bottom: 16px;
}

.vars-header,
.serveServers-header {
    padding: 0 0 8px;
    line-height: 1.4;
    white-space: normal;
    text-align: left;
    margin-bottom: 8px;
}

.serveServers-header,
.vars-header {
    font-weight: bold;
    margin-bottom: 0;
    margin-top: 16px;
}

.envDetail-content {
    position: relative;
}

.select-service {
    .ant-select-selector {
        border: 1px solid #d9d9d9;
    }
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
:deep(.ant-input:not(.env-name):hover),
:deep(.ant-input:active),
:deep(.ant-input:focus) {
    border: 1px solid #4096ff !important
}

:deep(.ant-input:not(.env-name)) {
    border: 1px solid transparent !important
}

:deep(.custom-select .ant-select-selector) {
    border: 1px solid transparent !important;
}

:deep(.custom-select .ant-select-selector:hover),
:deep(.custom-select .ant-select-selector:active),
:deep(.custom-select .ant-select-selector:focus) {
    border: 1px solid #4096ff !important
}</style>