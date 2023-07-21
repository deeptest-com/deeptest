<template>
    <!-- ::::环境详情 -->
    <a-form :model="activeEnvDetail" ref="formRef">
        <div class="title">
            <ConBoxTitle :title="activeEnvDetail.displayName"/>
        </div>
        <div class="envDetail-content">
            <a-form-item :labelCol="{ span: 2 }" :wrapperCol="{ span: 10 }" label="环境名称" name="name"
                :rules="rules.name">
                <a-input class="env-name" :value="activeEnvDetail.name || ''" @change="handleEnvNameChange"
                    placeholder="请输入环境名称" />
            </a-form-item>
            <div class="serveServers">
                <div class="serveServers-header">服务访问地址</div>
                <PermissionButton
                    code="LINK-SERVICE"
                    class="envDetail-btn"
                    text="关联服务"
                    @handle-access="addService">
                    <template #before>
                        <PlusOutlined />
                    </template>
                </PermissionButton>
                <a-table v-if="activeEnvDetail.serveServers.length > 0" size="small" bordered :pagination="false"
                    :columns="serveServersColumns" :data-source="activeEnvDetail.serveServers" :rowKey="(_record, index) => index">
                    <template #customUrl="{ text, index }">
                        <a-form-item :name="['serveServers', index, 'url']"
                            :rules="rules.serveUrl">
                            <a-input :value="text" @change="(e) => {
                                handleEnvChange('serveServers', 'url', index, e);
                            }" placeholder="http 或 https 起始的合法 URL" />
                        </a-form-item>
                    </template>
                    <template #customAction="{ index }">
                        <PermissionButton
                            code="UNLINK-SERVICE"
                            type="text"
                            size="small"
                            :danger="true"
                            text="解除关联"
                            @handle-access="handleEnvChange('serveServers', '', index, '', 'delete')" />
                    </template>
                </a-table>
            </div>
            <div class="vars">
                <div class="vars-header">环境变量</div>
                <PermissionButton
                    class="envDetail-btn"
                    code="ADD-ENVIRONMENT-VARIABLE"
                    text="添加"
                    @handle-access="addVar">
                    <template #before>
                        <PlusOutlined />
                    </template>
                </PermissionButton>
                <a-table v-if="activeEnvDetail.vars.length > 0" bordered size="small" :pagination="false"
                    :columns="globalVarsColumns" :data-source="activeEnvDetail.vars" :rowKey="(_record, index) => index">
                    <template #customName="{ text, index }">
                        <a-form-item :name="['vars', index, 'name']"
                            :rules="rules.var">
                            <a-input @change="(e) => {
                                handleEnvChange('vars', 'name', index, e);
                            }" :value="text" placeholder="请输入参数名" />
                        </a-form-item>
                    </template>
                    <template #customLocalValue="{ text, index }">
                        <a-form-item :name="['vars', index, 'localValue']"
                            :rules="rules.localValue">
                            <a-input :value="text" @change="(e) => {
                                handleEnvChange('vars', 'localValue', index, e);
                            }" placeholder="请输入本地值" />
                        </a-form-item>
                    </template>
                    <template #customRemoteValue="{ text, index }">
                        <a-form-item :name="['vars', index, 'remoteValue']"
                            :rules="rules.remoteValue">
                            <a-input :value="text" @change="(e) => {
                                handleEnvChange('vars', 'remoteValue', index, e);
                            }" placeholder="请输入远程值" />
                        </a-form-item>
                    </template>
                    <template #customDescription="{ text, index }">
                        <a-input :value="text" @change="(e) => {
                            handleEnvChange('vars', 'description', index, e);
                        }" placeholder="请输入描述信息" />
                    </template>
                    <template #customAction="{ index }">
                        <PermissionButton
                            code="DELETE-ENVIRONMENT-VARIABLE"
                            type="text"
                            size="small"
                            :danger="true"
                            text="删除"
                            @handle-access="handleEnvChange('vars', '', index, '', 'delete')" />
                    </template>
                </a-table>
            </div>
        </div>
        <div class="envDetail-footer">
            <PermissionButton
                class="save-btn"
                code="SAVE-ENVIRONMENT"
                type="primary"
                text="保存"
                html-type="submit"
                @handle-access="addEnvData" />
        </div>
    </a-form>
    <a-modal v-model:visible="addServiceModalVisible" title="关联服务" @ok="handleAddServiceOk">
        <a-form-item class="select-service" :labelCol="{ span: 6 }" :wrapperCol="{ span: 16 }" label="请选择服务">
            <a-select v-model:value="selectedService" :options="serviceOptions" placeholder="请选择服务" style="width: 200px" />
        </a-form-item>

    </a-modal>
</template>
<script setup lang="ts">
import { ref, computed } from 'vue';
import { useStore } from "vuex";
import { message } from 'ant-design-vue';
import { PlusOutlined } from '@ant-design/icons-vue';
import ConBoxTitle from '@/components/ConBoxTitle/index.vue';
import PermissionButton from "@/components/PermissionButton/index.vue";
import { globalVarsColumns, serveServersColumns } from '../../config';
import { useGlobalEnv } from '../../hooks/useGlobalEnv';
import { StateType as ProjectSettingStateType } from "@/views/ProjectSetting/store";

const store = useStore<{ ProjectSetting: ProjectSettingStateType }>();
const serviceOptions = computed<any>(() => store.state.ProjectSetting.serviceOptions);
const addServiceModalVisible = ref(false);
const selectedService = ref(null);
const formRef = ref<any>();


const {
    activeEnvDetail,
    addVar,
    addEnvData,
    handleEnvChange,
    handleEnvNameChange
} = useGlobalEnv(formRef);

const rules = {
    name: [{ required: true, message: '环境名称不能为空' }],
    serveUrl: [{ required: true, validator: urlValidator }],
    var: [{ required: true, message: '参数名不可为空' }],
    localValue: [{ required: true, message: '本地值不可为空' }],
    remoteValue: [{ required: true, message: '远程值不可为空' }]
}

// 添加服务弹窗操作
async function addService() {
    addServiceModalVisible.value = true;
}

function urlValidator(...arg: any[]) {
    const value = arg[1];
    // 需要兼容 https://localhost
    const urlReg = /http(s)?:\/\/[\w-]+/; // eslint-disable-line
    if (value === '') {
        return Promise.reject('url不能为空');
    } else {
        if (!urlReg.test(value)) {
            return Promise.reject('url格式错误，请参考 http(s)://www.test.com ');
        }
        return Promise.resolve();
    }
}

function handleAddServiceOk() {
    const selectServe: any = serviceOptions.value.find((item: any) => {
        return selectedService.value === item.id;
    })
    const envDetail = JSON.parse(JSON.stringify(activeEnvDetail.value));
    const isExsitServe = envDetail.serveServers.find((item: any) => {
        return item.serveId === selectServe.id;
    });
    if (!isExsitServe) {
        store.dispatch('ProjectSetting/addEnvServe', {
            "url": "",
            "serveName": selectServe.name,
            "serveId": selectServe.id,
        })
        addServiceModalVisible.value = false;
    } else {
        message.error('不可添加重复的服务,请重新选择~');
    }

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
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: flex-end;
    margin-top: 60px;
    box-shadow: 0px -1px 0px rgba(0, 0, 0, 0.06);

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
}

:deep(.serveServers .ant-row.ant-form-item),
:deep(.vars .ant-row.ant-form-item) {
    margin-bottom: 0 !important;
}

:deep(.vars .ant-row.ant-form-item-has-error .ant-form-item-control-input),
:deep(.serveServers .ant-row.ant-form-item-has-error .ant-form-item-control-input) {
    border: 1px solid #f5222d;
}
</style>
