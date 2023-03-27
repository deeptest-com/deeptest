<template>
  <div class="content">
    <div class="header">
      <TableFilter />
    </div>
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

    <a-drawer
        :closable="true"
        :width="1000"
        :key="editKey"
        :visible="drawerVisible"
        @close="onClose"
    >
      <template #title>
        <div class="drawer-header">
          <div>服务编辑</div>
        </div>
      </template>
      <div class="drawer-content">
        <a-form :model="formState" :label-col="{ span: 2 }" :wrapper-col=" { span: 15 }">
          <a-form-item label="服务名称">
            <a-input
                v-if="isEditServiceName"
                @focusout="changeServiceInfo"
                @pressEnter="changeServiceInfo"
                v-model:value="editFormState.name"
                placeholder="请输入内容"/>
            <span v-else>{{ editFormState.name }}
              <edit-outlined class="editable-cell-icon" @click="editServiceName"/></span>
          </a-form-item>
          <a-form-item label="描述">
            <a-input
                @focusout="changeServiceInfo"
                @pressEnter="changeServiceInfo"
                v-if="isEditServiceDesc"
                v-model:value="editFormState.description"
                placeholder="请输入内容"/>
            <span v-if="!isEditServiceDesc">{{ editFormState.description }}
              <edit-outlined class="editable-cell-icon" @click="editServiceDesc"/> </span>
          </a-form-item>
          <a-tabs v-model:activeKey="activeKey">
            <a-tab-pane key="1" tab="服务版本">
              <ServiceVersion :serveId="editFormState.serveId"/>
            </a-tab-pane>
            <a-tab-pane key="2" tab="服务组件">
              <ServiceComponent :serveId="editFormState.serveId"/>
            </a-tab-pane>
          </a-tabs>
        </a-form>
      </div>
    </a-drawer>

  </div>
</template>
<script setup lang="ts">

import {
  computed,
  defineEmits,
  defineProps,
  reactive,
  ref,
  UnwrapRef,
  watch
} from 'vue';
import {EditOutlined,MoreOutlined} from '@ant-design/icons-vue';
import ServiceVersion from './Version.vue';
import TableFilter from '../commom/TableFilter.vue';
import ServiceComponent from './Component.vue';

import {deleteServe, copyServe, disableServe, saveServe,} from '../../service';
import {message} from "ant-design-vue";
import {StateType as ProjectStateType} from "@/store/project";
import {StateType as ProjectSettingStateType} from '../../store';
import {useStore} from "vuex";
import { serviceColumns } from '../../config';

const store = useStore<{ ProjectGlobal: ProjectStateType, ProjectSettingV2: ProjectSettingStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const dataSource = computed<any>(() => store.state.ProjectSettingV2.serviceOptions);
const props = defineProps({})
const emit = defineEmits(['ok', 'close', 'refreshList']);

const formState: UnwrapRef<any> = reactive({
  name: '',
  description: '',
  userId: null,
});

const editFormState: UnwrapRef<any> = reactive({
  name: '',
  description: '',
  serveId: '',
});

const drawerVisible = ref(false);
const editKey = ref(0);
const activeKey = ref('1');

function onClose() {
  drawerVisible.value = false;

}

const edit = (record: any) => {
  editKey.value++;
  drawerVisible.value = true;
  editFormState.name = record.name;
  editFormState.description = record.description;
  editFormState.serveId = record.id;
};

const isEditServiceDesc = ref(false);
const isEditServiceName = ref(false);

function editServiceDesc() {
  isEditServiceDesc.value = true;
}

function editServiceName() {
  isEditServiceName.value = true;
}

async function changeServiceInfo(e) {
  isEditServiceDesc.value = false;
  isEditServiceName.value = false;
  if (editFormState.name && editFormState.description) {
    const res = await saveServe({
      "projectId": currProject.value.id,
      "name": editFormState.name,
      "description": editFormState.description,
      "id": editFormState.serveId,
    });
    if (res.code === 0) {
      // message.success('修改服务描述成功');
      await getList();
    } else {
      // message.error('修改服务描述失败');
    }
  }
}

async function onDelete(record: any) {
  const res = await deleteServe(record.id);
  if (res.code === 0) {
    message.success('删除成功');
    await getList();
  } else {
    message.error('删除失败');
  }
}

async function onDisabled(record: any) {
  const res = await disableServe(record.id);
  if (res.code === 0) {
    message.success('禁用服务成功');
    await getList();
  } else {
    message.error('禁用服务失败');
  }
}

async function onCopy(record: any) {
  const res = await copyServe(record.id);
  if (res.code === 0) {
    message.success('复制服务成功');
    await getList();
  } else {
    message.error('复制服务失败');
  }
}

async function getList() {
  await store.dispatch('ProjectSettingV2/getServersList', {
    projectId: currProject.value.id,
    page: 0,
    pageSize: 100,
    name: ''
  })
}

// onMounted(async () => {
//   await getList()
// })

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

  .btns {
    //border-bottom: 1px solid #e9e9e9;
    //padding: 10px 16px;
    //background: #fff;
    //margin-bottom: 8px;
  }
}

.operation-a{
  text-align: center;
  display: inline-block;
  width: 80px;
}

</style>
