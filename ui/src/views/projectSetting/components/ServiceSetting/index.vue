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
            <a-tab-pane key="3" tab="Security">
              <ServiceSecurity :serveId="editFormState.serveId"/>
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
  watch,
} from 'vue';
import { useRouter } from "vue-router";
import {EditOutlined,MoreOutlined} from '@ant-design/icons-vue';
import ServiceVersion from './Version.vue';
import TableFilter from '../commom/TableFilter.vue';
import Filter from '../commom/Filter.vue';
import ServiceComponent from './Component.vue';
import ServiceSecurity from './Security.vue';

import {StateType as ProjectStateType} from "@/store/project";
import {StateType as ProjectSettingStateType} from '../../store';
import {useStore} from "vuex";
import { serviceColumns } from '../../config';
import { Schema } from '../../data';
const props = defineProps({
  params: {
    type: Object,
  },
})
const router = useRouter();
const store = useStore<{ ProjectGlobal: ProjectStateType, ProjectSetting: ProjectSettingStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const dataSource = computed<any>(() => store.state.ProjectSetting.serviceOptions);

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

const schemaList: Schema[] = [
  {
    type: 'tooltip',
    text: '新建组件',
    title: '一个产品服务端通常对应一个或多个服务(微服务)，服务可以有多个版本并行，新的服务默认起始版本为v0.1.0。'
  },
  {
    type: 'input',
    stateName: 'name',
    placeholder: '服务名称',
    valueType: 'string'
  },
  {
    type: 'select',
    stateName: 'serveId',
    placeholder: '负责人(默认创建人)',
    options: [],
    valueType: 'string'
  },
  {
    type: 'input',
    stateName: 'description',
    placeholder: '输入描述',
    valueType: 'string'
  },
  {
    type: 'button',
    text: '确定',
  },
] 

const drawerVisible = ref(false);
const editKey = ref(0);
const activeKey = ref('1');

function onClose() {
  drawerVisible.value = false;
  router.currentRoute.value.query={}

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
    await store.dispatch('ProjectSetting/saveStoreServe', {
      "projectId": currProject.value.id,
      "name": editFormState.name,
      "description": editFormState.description,
      "id": editFormState.serveId,
    });
  }
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
    name: ''
  })
}



// 实时监听项目切换，如果项目切换了则重新请求数据
watch(() => {
  return currProject.value;
}, async (newVal) => {
  await getList()
  await isHasProps()

}, {
  immediate: true
})

// 判断是否携带参数，用于security模块
async function isHasProps(){
  if(JSON.stringify(props.params) !=='{}'){
    let record={}  
    dataSource?.value?.map((item)=>{
      if(item.id==props.params?.serveId*1){
          record=  item
      }
    })
    await edit(record)
    activeKey.value=props.params?.sectab       
  }
}

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
