<template>
    <a-form layout="inline" :model="formState">
        <a-form-item>
            <span style="cursor: pointer;font-weight: bold">新建服务
                <a-tooltip placement="topLeft" arrow-point-at-center
                    title="一个产品服务端通常对应一个或多个服务(微服务)，服务可以有多个版本并行，新的服务默认起始版本为v0.1.0。">
                    <QuestionCircleOutlined class="icon"
                        style="position: relative;top:-8px; font-size: 12px;transform: scale(0.9)" />
                </a-tooltip>
            </span>
        </a-form-item>
        <a-form-item name="name" :rules="[{ required: true, message: '服务名称不能为空' }]">
            <a-input v-model:value="formState.name" placeholder="服务名称" />
        </a-form-item>
        <a-form-item>
            <a-select v-model:value="formState.username" show-search placeholder="负责人(默认创建人)" style="width: 200px"
                :options="userListOptions" @focus="selectUserFocus">
            </a-select>
        </a-form-item>
        <a-form-item>
            <a-input v-model:value="formState.description" placeholder="输入描述" />
        </a-form-item>
        <a-form-item>
            <a-button class="editable-add-btn" @click="handleOk" type="primary" html-type="submit" style="margin-bottom: 8px">
                确定
            </a-button>
        </a-form-item>
    </a-form>


    <a-input-search v-model:value="keyword" placeholder="输入服务名称搜索" style="width: 300px" @search="onSearch" />
</template>
<script setup lang="ts">
import { reactive, UnwrapRef, ref, computed } from 'vue';
import {StateType as ProjectStateType} from "@/store/project";
import {StateType as ProjectSettingStateType} from '../../store';
import { useStore } from 'vuex';

const store = useStore<{ ProjectGlobal: ProjectStateType, ProjectSetting: ProjectSettingStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const userListOptions = computed<any>(() => store.state.ProjectSetting.userListOptions);
const keyword = ref('');
const formState: UnwrapRef<any> = reactive({
  name: '',
  description: '',
  username: null,
});

async function getList() {
    await store.dispatch('ProjectSetting/getServersList', {
        projectId: currProject.value.id,
        page: 0,
        pageSize: 100,
        name: keyword.value
    })
}

async function selectUserFocus() {
  await store.dispatch('ProjectSetting/getUserOptionsList')
}

// 确定
async function handleOk() {
    const { name, username, description } = formState;
    if (!name) {
        return;
    }
    const result = userListOptions.value.filter((e: any) => e.value === username);
    await store.dispatch('ProjectSetting/saveStoreServe', {
        projectId: currProject.value.id,
        formState: {
            userId: result && result[0] && result[0].id,
            name,
            description
        },
        action: 'create'
    })
}

async function onSearch(e) {
  await getList();
}

</script>