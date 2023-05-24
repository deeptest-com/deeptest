<template>
  <div class="content">
    <!-- header -->
    <div class="header">
      <CustomForm :form-config="formConfig" :rules="rules" :show-search="true" :search-placeholder="'输入服务名称搜索'"
        @handle-ok="handleAdd" @handle-search="handleSearch" />
    </div>
    <!-- content -->
    <EmptyCom>
      <template #content>
        <a-table :data-source="dataSource" :columns="serviceColumns" :rowKey="(_record, index) => index">
          <template #name="{ text, record }">
            <div class="serve-name">
              <EditAndShowField :custom-class="'custom-serve show-on-hover'" placeholder="请输入服务名称" :value="text || ''"
                @update="(e: string) => handleUpdateName(e, record)" @edit="edit(record)" />
            </div>
          </template>
          <template #description="{ text }">
            <div class="serve-description">
              {{ text || '' }}
            </div>
          </template>
          <template #customServers="{ record }">
            <span v-if="record?.servers.length > 0">
              <a-tag v-for="server in record.servers" :key="server.id">
                {{ server.name || server.description }}
              </a-tag>
            </span>
            <span v-else>---</span>
          </template>
          <template #customStatus="{ text, record }">
            <a-tag :color="record.statusTag">{{ text }}</a-tag>
          </template>
          <template #operation="{ record }">
            <a-dropdown>
              <MoreOutlined />
              <template #overlay>
                <a-menu>
                  <a-menu-item key="1">
                    <a class="operation-a" href="javascript:void (0)" @click="onOpenComponent(record)">服务组件</a>
                  </a-menu-item>
                  <a-menu-item key="2">
                    <a class="operation-a" href="javascript:void (0)" @click="onOpenVersion(record)">服务版本</a>
                  </a-menu-item>
                  <a-menu-item key="3">
                    <a class="operation-a" href="javascript:void (0)" @click="onOpenSecurity(record)">security</a>
                  </a-menu-item>
                  <a-menu-item key="4">
                    <a class="operation-a" href="javascript:void (0)" @click="onDisabled(record)">禁用</a>
                  </a-menu-item>
                  <a-menu-item key="5">
                    <a class="operation-a" href="javascript:void (0)" @click="onCopy(record)">复制</a>
                  </a-menu-item>
                  <a-menu-item key="6">
                    <a class="operation-a" href="javascript:void (0)" @click="onDelete(record)">删除</a>
                  </a-menu-item>
                </a-menu>
              </template>
            </a-dropdown>
          </template>
        </a-table>
      </template>
    </EmptyCom>

    <!-- 抽屉 -->
    <Drawer :edit-key="editKey" :drawer-visible="drawerVisible" :tab-key="currentTabKey"
      @update:tab-key="handleUpdateTabKey" @onClose="onClose" />
  </div>
</template>
<script setup lang="ts">

import {
  computed,
  ref,
  watch,
  createVNode,
} from 'vue';
import { useStore } from "vuex";
import { useRouter } from 'vue-router';
import { Modal } from 'ant-design-vue';
import { ExclamationCircleOutlined, MoreOutlined } from '@ant-design/icons-vue';
import CustomForm from '../common/CustomForm.vue';
import EditAndShowField from '@/components/EditAndShow/index.vue';
import EmptyCom from '@/components/Empty/index.vue';
import Drawer from './Drawer.vue';
import { StateType as ProjectStateType } from "@/store/project";
import { StateType as ProjectSettingStateType } from '../../store';
import { serviceColumns } from '../../config';

const store = useStore<{ ProjectGlobal: ProjectStateType, ProjectSetting: ProjectSettingStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const dataSource = computed<any>(() => store.state.ProjectSetting.serviceOptions);
const userListOptions = computed<any>(() => store.state.ProjectSetting.userListOptions);
const route = useRouter();

const drawerVisible = ref(false);
const editKey = ref(0);
const currentTabKey = ref('');

let formConfig = ref([
  {
    type: 'tooltip',
    title: '一个产品服务端通常对应一个或多个服务(微服务)，服务可以有多个版本并行，新的服务默认起始版本为v0.1.0。',
    text: '新建服务'
  },
  {
    type: 'input',
    modelName: 'name',
    placeholder: '服务名称',
    valueType: 'string'
  },
  {
    type: 'select',
    modelName: 'username',
    placeholder: '负责人(默认创建人)',
    options: [],
    valueType: 'string',
  },
  {
    type: 'input',
    modelName: 'description',
    placeholder: '输入描述',
    valueType: 'string'
  },
  {
    type: 'button',
    text: '新建'
  }
]);

const rules = {
  name: [
    {
      required: true,
      message: '服务名称不能为空'
    }
  ]
};

async function handleAdd(formData: any) {
  const { name, username, description } = formData;
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

function onClose() {
  drawerVisible.value = false;
}

function handleSearch(value: any) {
  getList(value);
}

function handleUpdateName(value: string, record: any) {
  const serviceInfo = { name: value, description: record.description, id: record.id };
  store.dispatch('ProjectSetting/saveStoreServe', {
    "projectId": currProject.value.id,
    formState: { ...serviceInfo },
    action: 'update'
  });
}

async function edit(record: any) {
  if (!record || (record && Object.keys(record).length === 0)) {
    return;
  }
  await store.dispatch('ProjectSetting/setServiceDetail', {
    name: record.name,
    description: record.description,
    id: record.id
  })
  editKey.value++;
  drawerVisible.value = true;
}

async function onOpenComponent(record: any) {
  await edit(record);
  currentTabKey.value = 'service-component';
}

async function onOpenSecurity(record: any) {
  await edit(record);
  currentTabKey.value = 'service-security';
}

async function onOpenVersion(record: any) {
  await edit(record);
  currentTabKey.value = 'service-version';
}

function handleUpdateTabKey(val: string) {
  currentTabKey.value = val;
}

async function onDelete(record: any) {
  Modal.confirm({
    title: '确认要删除该服务吗',
    icon: createVNode(ExclamationCircleOutlined),
    onOk() {
      store.dispatch('ProjectSetting/deleteStoreServe', { id: record.id, projectId: currProject.value.id });
    }
  })
}

async function onDisabled(record: any) {
  store.dispatch('ProjectSetting/disabledStoreServe', { id: record.id, projectId: currProject.value.id });
}

async function onCopy(record: any) {
  store.dispatch('ProjectSetting/copyStoreServe', { id: record.id, projectId: currProject.value.id });
}

async function getList(name = '') {
  await store.dispatch('ProjectSetting/getServersList', {
    projectId: currProject.value.id,
    page: 0,
    pageSize: 100,
    name
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

watch(() => {
  return userListOptions.value;
}, (val: any) => {
  if (val && val.length > 0) {
    const config = JSON.parse(JSON.stringify(formConfig.value));
    config.forEach((e: any) => {
      if (e.type === 'select') {
        e.options = [...val];
      }
    })
    formConfig.value = config;
  }
}, {
  immediate: true
})

// 判断是否携带参数，用于security模块
async function isHasProps() {
  const { query: { serveId = '', sectab = '' } = {} }: any = route.currentRoute.value;
  if (serveId) {
    let record = {}
    dataSource?.value?.map((item) => {
      if (item.id == serveId * 1) {
        record = item
      }
    })
    await edit(record)
    currentTabKey.value = sectab;
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

.operation-a {
  text-align: center;
  display: inline-block;
  width: 80px;
}

.serve-name {
  width: 120px;
}

.serve-description {
  max-width: 140px;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
}
</style>
