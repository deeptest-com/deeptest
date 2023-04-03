<template>
  <div class="content">
    <!-- header -->
    <div class="header">
      <TableFilter />
      <!-- <Filter :form-schema-list="schemaList" :need-search="true" :search-place-holder="'输入服务名称搜索'" @handle-search="onSearch"/> -->
    </div>
    <!-- content -->
    <a-table  :data-source="dataSource" :columns="serviceColumns" rowKey="id">

      <template #name="{ text, record }">
        <div class="editable-cell">
          <div class="editable-cell-text-wrapper">
            {{ text || ' ' }}
            <edit-outlined class="editable-cell-icon" @click="edit(record)"/>
          </div>
        </div>
      </template>
      <template #customServers="{ record }">
        <span v-if="record?.servers.length > 0">
               <span  v-for="server in record.servers" :key="server.id">
          {{server.name || server.description}}
        </span>
        </span>
        <span v-else>暂无关联服务</span>
      </template>
      <template #customStatus="{ text,record }">
        <a-tag :color="record.statusTag">{{ text }}</a-tag>
      </template>
      <template #operation="{ record }">
        <a-dropdown>
          <MoreOutlined/>
          <template #overlay>
            <a-menu>
              <a-menu-item key="1">
                <a class="operation-a"  href="javascript:void (0)" @click="onDisabled(record)">禁用</a>
              </a-menu-item>
              <a-menu-item key="2">
                <a  class="operation-a"  href="javascript:void (0)" @click="onCopy(record)">复制</a>
              </a-menu-item>
              <a-menu-item key="3">
                <a  class="operation-a"  href="javascript:void (0)" @click="onDelete(record)">删除</a>
              </a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>
      </template>
    </a-table>
    <!-- 抽屉 -->
    <Drawer 
      :edit-key="editKey"
      :drawer-visible="drawerVisible" 
      @onClose="onClose" />
  </div>
</template>
<script setup lang="ts">

import {
  computed,
  reactive,
  ref,
  UnwrapRef,
  watch
} from 'vue';
import {useStore} from "vuex";
import {EditOutlined,MoreOutlined} from '@ant-design/icons-vue';
import TableFilter from '../commom/TableFilter.vue';
// import Filter from '../commom/Filter.vue'; //后续需要转移到这个组件上
import Drawer from './Drawer.vue';
import {StateType as ProjectStateType} from "@/store/project";
import {StateType as ProjectSettingStateType} from '../../store';
import { serviceColumns } from '../../config';

const store = useStore<{ ProjectGlobal: ProjectStateType, ProjectSetting: ProjectSettingStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const dataSource = computed<any>(() => store.state.ProjectSetting.serviceOptions);

const drawerVisible = ref(false);
const editKey = ref(0);
const keyword = ref('');

function onClose() {
  drawerVisible.value = false;
}

function edit(record: any) {
  store.dispatch('ProjectSetting/setServiceDetail', {
    name: record.name,
    description: record.description,
    id: record.id
  })
  editKey.value++;
  drawerVisible.value = true;
}

async function onDelete(record: any) {
  store.dispatch('ProjectSetting/deleteStoreServe', { id: record.id, projectId: currProject.value.id });
}

async function onDisabled(record: any) {
  store.dispatch('ProjectSetting/disabledStoreServe', { id: record.id, projectId: currProject.value.id });
}

async function onCopy(record: any) {
  store.dispatch('ProjectSetting/copyStoreServe', { id: record.id, projectId: currProject.value.id });
}

async function getList() {
    await store.dispatch('ProjectSetting/getServersList', {
        projectId: currProject.value.id,
        page: 0,
        pageSize: 100,
        name: keyword.value
    })
}

// 实时监听项目切换，如果项目切换了则重新请求数据
watch(() => {
  return currProject.value;
}, async (newVal) => {
  await getList()
}, {
  immediate: true
})

</script>

<style scoped lang="less">
.content {
  margin: 20px;

  .header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 8px;
  }
}

.drawer-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.operation-a{
  text-align: center;
  display: inline-block;
  width: 80px;
}

</style>
