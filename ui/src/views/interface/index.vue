<template>
  <div class="container">
    <div class="content">
      <div class="left tree">
        <InterfaceTree @select="selectNode"/>
      </div>
      <div class="right">
        <!--  头部搜索区域  -->
        <div class="top-action">
          <a-button class="action-new" type="primary" :loading="loading"
                    @click="createApiModaVisible = true;">新建接口
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
                onChange: (page) => {
                  loadList(currProject.id,page,pagination.pageSize);
                },
                onShowSizeChange: (page, size) => {
                  loadList(currProject.id,page,size);
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
              <a-select
                  :value="record?.status"
                  style="width: 100px"
                  :size="'small'"
                  placeholder="请修改接口状态"
                  :options="interfaceStatusOpts"
                  @change="(val) => {
                  handleChangeStatus(val,record);
                  }"
              >
              </a-select>
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
        :selectedCategoryId="selectedCategoryId"
        @cancal="createApiModaVisible = false;"
        @ok="handleCreateApi"/>
    <!-- 编辑接口时，展开抽屉   -->
    <Drawer
        :destroyOnClose="true"
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
import InterfaceTree from './list/tree.vue';
import {ColumnProps} from 'ant-design-vue/es/table/interface';
import {
  EditOutlined,
  MoreOutlined
} from '@ant-design/icons-vue';
import {interfaceStatusOpts} from '@/config/constant';
import CreateApiModal from './components/CreateApiModal.vue';
import TableFilter from './components/TableFilter.vue';
import Drawer from './components/Drawer/index.vue'
import {useStore} from "vuex";
import {Interface, PaginationConfig} from "@/views/interface/data";
import {StateType as ServeStateType} from "@/store/serve";
const store = useStore<{ Interface, ProjectGlobal, ServeGlobal: ServeStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const list = computed<Interface[]>(() => store.state.Interface.listResult.list);
let pagination = computed<PaginationConfig>(() => store.state.Interface.listResult.pagination);
const currServe = computed<any>(() => store.state.ServeGlobal.currServe);
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
// 抽屉是否打开
const drawerVisible = ref<boolean>(false);
const selectedCategoryId = ref<string>('');
const onSelectChange = (keys: Key[], rows: any) => {
  selectedRowKeys.value = [...keys];
};


async function handleChangeStatus(value: any, record: any,) {
  await store.dispatch('Interface/updateStatus', {
    id: record.id,
    status: value
  });
}

async function editInterface(record) {
  await store.dispatch('Interface/getInterfaceDetail', {id: record.id});
  drawerVisible.value = true;
}

async function copy(record: any) {
  await store.dispatch('Interface/copy', record);
}

async function disabled(record: any) {
  await store.dispatch('Interface/disabled', record);
}

async function del(record: any) {
  await store.dispatch('Interface/del', record);
}

async function handleCreateApi(data) {
  await store.dispatch('Interface/createApi', {
    "title": data.title,
    "projectId": currProject.value.id,
    "serveId": currServe.value.id,
    "description": data.description || null,
    "categoryId": data.categoryId || null,
  });
  createApiModaVisible.value = false;
}

async function selectNode(id) {
  selectedCategoryId.value = id;
  await loadList(currProject.value.id, pagination.value.current, pagination.value.pageSize, {
    categoryId: id,
    serveId: currServe.value.id,
  });
}

async function loadList(currProjectId, page, size, opts?: any) {
  await store.dispatch('Interface/loadList', {
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
  }
}, {
  immediate: true
})

// 实时监听服务 ID，如果项目切换了则重新请求数据
watch(() => {
  return currServe.value.id;
}, async (newVal) => {
  if (newVal) {
    await loadList(newVal.id, pagination.value.current, pagination.value.pageSize, {
      serveId: newVal,
    });
    await store.dispatch('Interface/getServerList', {id: currServe.value.id});
    // 获取授权列表
    await store.dispatch('Interface/getSecurityList', {id: currServe.value.id});
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
