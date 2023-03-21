<template>
  <div class="container">
    <div class="content">
      <div class="left tree">
        左侧区域 待添加，公共树形结构
      </div>
      <!--  头部搜索区域  -->
      <div class="right">
        <div class="top-action">
          <a-button
              class="action-new"
              type="primary"
              :loading="loading"
              @click="createApiModaVisible = true;">新建接口
          </a-button>
          <a-button class="action-import" type="primary" :disabled="!hasSelected" :loading="loading"
                    @click="importApi">批量操作
          </a-button>
        </div>
        <div class="top-search">
          <TableFilter @filter="handleTableFilter"/>
        </div>
        <a-table
            style="margin: 0 16px;"
            :row-selection="{
              selectedRowKeys: selectedRowKeys,
              onChange: onSelectChange
            }"
            :pagination="{
                ...pagination,
                onChange:async (page) => {
                  await loadList(currProject.id,page,pagination.pageSize);
                },
                onShowSizeChange: async (page, size) => {
                   await loadList(currProject.id,page,size);
                },
            }"
            :columns="columns"
            :data-source="list">
          <template #colTitle="{text,record}">
            <div class="customTitleColRender">
              <span>{{ text }}</span>
              <span class="edit" @click="editInterface(record)"><EditOutlined/></span>
            </div>
          </template>
          <template #colStatus="{record}">
            <div class="customTitleColRender">
              <a-tag :color="interfaceStatusColor.get(record.status)">{{ interfaceStatus.get(record.status) }}</a-tag>
            </div>
          </template>
          <template #colPath="{text}">
            <div class="customTitleColRender">
              <a-tag>{{ text }}</a-tag>
            </div>
          </template>
          <template #action="{record}">
            <a-dropdown>
              <MoreOutlined/>
              <template #overlay>
                <a-menu>
                  <a-menu-item key="1">
                    <a-button style="width: 80px" type="link" size="small" @click="copy(record)">复制</a-button>
                  </a-menu-item>
                  <a-menu-item key="2">
                    <a-button style="width: 80px" type="link" size="small" @click="del(record)">删除</a-button>
                  </a-menu-item>
                  <a-menu-item key="3">
                    <a-button style="width: 80px" type="link" size="small" @click="disabled(record)">过时</a-button>
                  </a-menu-item>
                </a-menu>
              </template>
            </a-dropdown>
          </template>
        </a-table>
      </div>
    </div>
    <!--  创建新接口弹框  -->
    <CreateApiModal
        :visible="createApiModaVisible"
        @cancal="createApiModaVisible = false;"
        @ok="handleCreateApi"/>
    <!-- 编辑接口时，展开抽屉   -->
    <Drawer
        :destroyOnClose="true"
        :interfaceId="editInterfaceId"
        :visible="drawerVisible"
        @refreshList="refreshList"
        @close="drawerVisible = false;"/>
  </div>
</template>
<script setup lang="ts">
import {
  computed, reactive, toRefs, ref, onMounted,
  watch
} from 'vue';
import {ColumnProps} from 'ant-design-vue/es/table/interface';
import {
  EditOutlined,
  MoreOutlined
} from '@ant-design/icons-vue';
import {interfaceStatus, interfaceStatusColor} from '@/config/constant';
import CreateApiModal from './components/CreateApiModal.vue';
import TableFilter from './components/TableFilter.vue';
import Drawer from './components/Drawer/index.vue'
import {useStore} from "vuex";
import {Interface, PaginationConfig} from "@/views/interface/data";
import {filterFormState} from "@/views/interfaceV2/data";

const store = useStore<{ InterfaceV2, ProjectGlobal }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const list = computed<Interface[]>(() => store.state.InterfaceV2.listResult.list);
let pagination = computed<PaginationConfig>(() => store.state.InterfaceV2.listResult.pagination);
const createApiModaVisible = ref(false);
type Key = ColumnProps['key'];

/**
 * 表格数据
 * */
const columns = [
  {
    title: '序号',
    dataIndex: 'index',
  },
  {
    title: '接口名称',
    dataIndex: 'title',
    slots: {customRender: 'colTitle'},
  },
  {
    title: '状态',
    dataIndex: 'status',
    slots: {customRender: 'colStatus'},
  },
  {
    title: '创建人',
    dataIndex: 'createUser',

  },
  {
    title: '接口路径',
    dataIndex: 'path',
    slots: {customRender: 'colPath'},
  },
  {
    title: '所属服务',
    dataIndex: 'serveName',
  },
  {
    title: '最近更新',
    dataIndex: 'updatedAt',
  },
  {
    title: '操作',
    key: 'operation',
    fixed: 'right',
    width: 100,
    slots: {customRender: 'action'},
  },
];
const selectedRowKeys = ref<Key[]>([]);
const loading = false;
// 是否批量选中了
// const hasSelected = computed(() => state.selectedRowKeys.length > 0);
const hasSelected = false;
// 抽屉是否打开
const drawerVisible = ref<boolean>(false);
const onSelectChange = (keys: Key[], rows: any) => {
  selectedRowKeys.value = [...keys];
};
const editInterfaceId = ref('');

async function editInterface(record) {
  editInterfaceId.value = record.id;
  await store.dispatch('InterfaceV2/getInterfaceDetail', {id:record.id});
  drawerVisible.value = true;
}

async function copy(record: any) {
  await store.dispatch('InterfaceV2/copy', record);
}

async function disabled(record: any) {
  await store.dispatch('InterfaceV2/disabled', record);
}

/**
 * 删除接口
 * */
async function del(record: any) {
  await store.dispatch('InterfaceV2/del', record);
}


function importApi() {
  console.log('导入')
}

async function handleCreateApi(data) {
  await store.dispatch('InterfaceV2/createApi', {
    "title": data.title,
    "projectId": currProject.value.id,
    "description": data.description || null,
    "parentId": data.parentId || null,
  });
  createApiModaVisible.value = false;
}

async function loadList(currProjectId, page, size, opts?: filterFormState) {
  await store.dispatch('InterfaceV2/loadList', {
    currProjectId,
    "page": page,
    "pageSize": size,
    opts,
  });
}

async function handleTableFilter(filterState) {
  await loadList(currProject.value.id, pagination.value.current, pagination.value.pageSize, filterState);
}

// 实时监听项目切换，如果项目切换了则重新请求数据
watch(() => {
  return currProject.value;
}, async (newVal) => {
  if (newVal.id) {
    await loadList(newVal.id, pagination.value.current, pagination.value.pageSize);
    await store.dispatch('InterfaceV2/loadCategory');
  }
}, {
  immediate: true
})

async function refreshList() {
  await loadList(currProject.value.id, pagination.value.current, pagination.value.pageSize);
}


</script>
<style scoped lang="less">
.container {
  margin: 16px;
  background: #ffffff;
  min-height: calc(100vh - 80px);
}

.tag-filter-form {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 60px;

  .search-input {
    margin-left: 8px;
  }

  .add-btn {
    margin-left: 12px;
    margin-right: 16px;
    cursor: pointer;
  }
}

.content {
  display: flex;
  width: 100%;

  .left {
    width: 300px;
    border-right: 1px solid #f0f0f0;
    height: calc(100vh - 80px);
  }

  .right {
    flex: 1
  }
}

.action-new {
  margin-right: 8px;
}

.top-search {
  height: 60px;
  display: flex;
  align-items: center;
  margin-left: 16px;
  margin-bottom: 8px;
}

.top-action {
  height: 60px;
  display: flex;
  align-items: center;
  margin-left: 16px;
  margin-top: 8px;

  .ant-btn {
    margin-right: 16px;
  }
}

.action-btns {
  display: flex;
}

.customTitleColRender {
  display: flex;

  .edit {
    margin-left: 8px;
    cursor: pointer;
  }
}

.form-item-con {
  display: flex;
  justify-content: center;
  align-items: center;
}

.more-icon {
  position: absolute;
  right: 8px;
}


::v-deep {
  .ant-alert-info {
    padding: 12px;
  }

  .ant-alert-icon {
    font-size: 14px;
    position: relative;
    top: 4px;
    left: 8px;
  }

  .ant-alert-message {
    font-size: 14px;
  }

  .ant-alert-description {
    font-size: 12px;
  }
}

</style>
