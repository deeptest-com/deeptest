<template>
  <a-spin tip="Loading..." :spinning="isImporting" style="z-index: 2000;">
    <div class="container">
      <div class="content">
        <div class="left tree" v-if="!collapsed">
          <Tree @select="selectNode" :serveId="currServe.id"/>
        </div>
        <CollapsedIcon
            :style="{left:'294px',top:'300px'}"
            :collapsedStyle="{left:'-9px', top:'300px'}"
            @click="collapsed = !collapsed" :collapsed="collapsed"/>
        <div :class="{'right': true, 'right-not-collapsed': !collapsed}">
          <div class="top-action">
            <div class="top-action-left">
              <PermissionButton
                  class="action-new"
                  text="新建接口"
                  code="ENDPOINT-ADD"
                  type="primary"
                  :loading="loading"
                  @handle-access="handleCreateEndPoint"/>
              <a-dropdown :trigger="['hover']" :placement="'bottomLeft'">
                <a class="ant-dropdown-link" @click.prevent>
                  <a-button>批量操作</a-button>
                </a>
                <template #overlay>
                  <a-menu style="margin-top: 8px;">
                    <a-menu-item key="0">
                      <a-button type="link" :size="'small'" href="javascript:void (0)" @click="inportApi">导入接口
                      </a-button>
                    </a-menu-item>
                    <a-menu-item key="1">
                      <a-button :disabled="!hasSelected" :size="'small'" type="link" @click="goDocs">查看文档</a-button>
                    </a-menu-item>
                    <a-menu-item key="1">
                      <a-button :disabled="!hasSelected" :size="'small'" type="link"
                                @click="showPublishDocsModal = true">发布文档
                      </a-button>
                    </a-menu-item>
                  </a-menu>
                </template>
              </a-dropdown>
            </div>
            <div class="top-search-filter">
              <TableFilter @filter="handleTableFilter"/>
            </div>
          </div>
          <EmptyCom>
            <template #content>
              <a-table :loading="fetching"
                       :row-selection="{
                selectedRowKeys: selectedRowKeys,
                onChange: onSelectChange
              }"
                       :pagination="{
                  ...pagination,
                  onChange: (page) => {
                    loadList(page,pagination.pageSize);
                  },
                  onShowSizeChange: (page, size) => {
                    loadList(page,size);
                  },
              }"
                       :scroll="{ x: '1280px' || true }"
                       :columns="columns"
                       :data-source="list">
                <template #colTitle="{text,record}">
                  <div class="customTitleColRender">
                    <span>
                      <a :title="record?.title" href="javascript:void (0)" @click="editEndpoint(record)">{{ text }}</a>
                    </span>
<!--                    <EditAndShowField :custom-class="'custom-endpoint show-on-hover'"-->
<!--                                      :value="text"-->
<!--                                      placeholder="请输入接口名称"-->
<!--                                      @update="(e: string) => handleUpdateEndpoint(e, record)"-->
<!--                                      @edit="editEndpoint(record)"/>-->
                  </div>
                </template>

                <template #colStatus="{record}">
                  <div class="customStatusColRender">
                    <EditAndShowSelect
                        :label="endpointStatus.get(record?.status || 0 )"
                        :value="record?.status"
                        :options="endpointStatusOpts"
                        @update="(val) => { handleChangeStatus(val,record);}"/>
                  </div>
                </template>

                <template #colPath="{text}">
                  <div class="customPathColRender">
                    <a-tag>{{ text }}</a-tag>
                  </div>
                </template>

                <template #action="{record}">
                  <a-dropdown>
                    <MoreOutlined/>
                    <template #overlay>
                      <a-menu>
                        <a-menu-item v-for="menuItem in MenuList" :key="menuItem.key">
                          <PermissionButton
                              style="width: 80px"
                              :text="menuItem.text"
                              size="small"
                              type="link"
                              :code="menuItem.code"
                              @handle-access="menuItem.action(record)"/>
                        </a-menu-item>
                      </a-menu>
                    </template>
                  </a-dropdown>
                </template>
              </a-table>
            </template>
          </EmptyCom>
        </div>
      </div>
      <CreateEndpointModal
          :visible="createApiModalVisible"
          :selectedCategoryId="selectedCategoryId"
          @cancel="createApiModalVisible = false;"
          @ok="handleCreateApi"/>
      <ImportEndpointModal
          :visible="showImportModal"
          :selectedCategoryId="selectedCategoryId"
          @cancal="showImportModal = false;"
          @ok="handleImport"/>
      <PubDocs
          :visible="showPublishDocsModal"
          :endpointIds='selectedRowIds'
          @cancal="showPublishDocsModal = false;"
          @ok="publishDocs"/>
      <!-- 编辑接口时，展开抽屉：外层再包一层 div, 保证每次打开弹框都重新渲染   -->
      <div v-if="drawerVisible">
        <Drawer
            :destroyOnClose="true"
            :visible="drawerVisible"
            @refreshList="refreshList"
            @close="drawerVisible = false;"/>
      </div>
    </div>
  </a-spin>
</template>
<script setup lang="ts">
import {
  computed, reactive, toRefs, ref, onMounted,
  watch, createVNode, onUnmounted
} from 'vue';
import {useRouter} from 'vue-router';
import debounce from "lodash.debounce";
import EndpointTree from './list/tree.vue';
import {ColumnProps} from 'ant-design-vue/es/table/interface';
import {ExclamationCircleOutlined, MoreOutlined} from '@ant-design/icons-vue';
import {endpointStatusOpts, endpointStatus} from '@/config/constant';
import EditAndShowField from '@/components/EditAndShow/index.vue';
import CreateEndpointModal from './components/CreateEndpointModal.vue';
import PubDocs from './components/PubDocs.vue';
import ImportEndpointModal from './components/ImportEndpointModal.vue';
import TableFilter from './components/TableFilter.vue';
import Drawer from './components/Drawer/index.vue'
import EditAndShowSelect from '@/components/EditAndShowSelect/index.vue';
import EmptyCom from '@/components/Empty/index.vue';
import PermissionButton from "@/components/PermissionButton/index.vue";
import {useStore} from "vuex";
import {Endpoint, PaginationConfig} from "@/views/endpoint/data";
import CollapsedIcon from "@/components/CollapsedIcon/index.vue"
import {StateType as ServeStateType} from "@/store/serve";
import {StateType as Debug} from "@/views/component/debug/store";
import {message, Modal, notification} from 'ant-design-vue';
import Tree from './components/Tree.vue'

const store = useStore<{ Endpoint, ProjectGlobal, Debug: Debug, ServeGlobal: ServeStateType }>();
const collapsed = ref(false);
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const currServe = computed<any>(() => store.state.ServeGlobal.currServe);
const serves = computed<any>(() => store.state.ServeGlobal.serves);
const list = computed<Endpoint[]>(() => store.state.Endpoint.listResult.list);
let pagination = computed<PaginationConfig>(() => store.state.Endpoint.listResult.pagination);
const createApiModalVisible = ref(false);
const router = useRouter();
type Key = ColumnProps['key'];

/**
 * 表格数据
 * */
const columns = [
  {
    title: '编号',
    dataIndex: 'serialNumber',
    width: 150,
  },
  {
    title: '接口名称',
    dataIndex: 'title',
    slots: {customRender: 'colTitle'},
    ellipsis: true
  },
  {
    title: '状态',
    dataIndex: 'status',
    slots: {customRender: 'colStatus'},
    width: 150,
  },
  {
    title: '创建人',
    dataIndex: 'createUser',
    width: 100,
  },
  {
    title: '接口路径',
    dataIndex: 'path',
    width: 300,
    slots: {customRender: 'colPath'},
    ellipsis: true
  },
  {
    title: '所属服务',
    dataIndex: 'serveName',
    ellipsis: true,
    width: 100,
  },
  {
    title: '最近更新',
    dataIndex: 'updatedAt',
    width: 200,
  },
  {
    title: '操作',
    key: 'operation',
    fixed: 'right',
    width: 80,
    slots: {customRender: 'action'},
  },
];
const MenuList = [
  {
    key: '1',
    code: 'ENDPOINT-COPY',
    text: '复制',
    action: (record: any) => copy(record)
  },
  {
    key: '2',
    code: 'ENDPOINT-DELETEE',
    text: '删除',
    action: (record: any) => del(record)
  },
  {
    key: '3',
    code: 'ENDPOINT-OUTDATED',
    text: '过时',
    action: (record: any) => disabled(record)
  },
]
const selectedRowKeys = ref<Key[]>([]);
const selectedRowIds = ref<Key[]>([]);
const loading = false;
// 抽屉是否打开
const drawerVisible = ref<boolean>(false);
const selectedCategoryId = ref<string | number>('');
const onSelectChange = (keys: Key[], rows: any) => {
  console.log('onSelectChange', keys, rows)
  selectedRowKeys.value = [...keys];
  selectedRowIds.value = rows.map((item: any) => item.id);
};
const hasSelected = computed(() => selectedRowKeys.value.length > 0);

const fetching = ref(false);

/*查看选中的接口文档*/
function goDocs() {
  window.open(`/#/docs/view?endpointIds=${selectedRowIds.value.join(',')}`);
}

const showPublishDocsModal: any = ref(false)

// 发布文档版本
async function publishDocs() {
  showPublishDocsModal.value = false;
  selectedRowIds.value = [];
}

/**
 * 导入接口
 * */
const showImportModal = ref(false);

function inportApi() {
  showImportModal.value = true;
}


function handleCreateEndPoint() {
  if (serves.value.length === 0) {
    Modal.confirm({
      title: '请先创建服务',
      content: '创建接口前需先创建服务才能使用,确认将跳转服务页面',
      onOk: () => {
        router.push('/project-setting/service-setting');
      }
    })
    return;
  }
  createApiModalVisible.value = true;
}

async function handleChangeStatus(value: any, record: any,) {
  await store.dispatch('Endpoint/updateStatus', {
    id: record.id,
    status: value
  });
}

async function handleUpdateEndpoint(value: string, record: any) {
  await store.dispatch('Endpoint/updateEndpointDetail',
      {...record, title: value}
  );
}

async function editEndpoint(record) {
  await store.dispatch('Endpoint/getEndpointDetail', {id: record.id});
  drawerVisible.value = true;
}

async function copy(record: any) {
  await store.dispatch('Endpoint/copy', record);
}

async function disabled(record: any) {
  await store.dispatch('Endpoint/disabled', record);
}

async function del(record: any) {
  Modal.confirm({
    title: () => '确定删除该接口吗？',
    icon: createVNode(ExclamationCircleOutlined),
    okText: () => '确定',
    okType: 'danger',
    cancelText: () => '取消',
    onOk: async () => {
      const res = await store.dispatch('Endpoint/del', record);
      if (res) {
        message.success('删除成功');
      } else {
        message.error('删除失败');
      }
    },
  });
}

async function handleCreateApi(data) {
  await store.dispatch('Endpoint/createApi', {
    "title": data.title,
    "projectId": currProject.value.id,
    "serveId": currServe.value.id,
    "description": data.description || null,
    "categoryId": data.categoryId || null,
  });
  await refreshList();
  createApiModalVisible.value = false;
}


const isImporting = ref(false);

async function handleImport(data, callback) {

  isImporting.value = true;

  const res = await store.dispatch('Endpoint/importEndpointData', {
    ...data,
    "serveId": currServe.value.id,
  });

  // 导入成功，重新拉取列表 ，并且关闭弹窗
  if (res) {
    await refreshList();
    if (callback) {
      callback();
    }
    showImportModal.value = false;
  }
  isImporting.value = false

}

async function selectNode(id) {
  selectedCategoryId.value = id;
  await loadList(pagination.value.current, pagination.value.pageSize, {
    categoryId: id,
    serveId: currServe.value.id,
  });
}

const loadList = debounce(async (page, size, opts?: any) => {
  fetching.value = true;
  await store.dispatch('Endpoint/loadList', {
    "projectId": currProject.value.id,
    "page": page,
    "pageSize": size,
    opts,
  });
  fetching.value = false;
}, 300)

async function handleTableFilter(filterState) {
  await loadList(pagination.value.current, pagination.value.pageSize, filterState);
}

// 实时监听项目/服务 ID，如果项目切换了则重新请求数据
watch(() => [currProject.value.id, currServe.value.id], async (newVal) => {
  const [newProjectId, newServeId] = newVal;
  if (newProjectId !== undefined) {
    await loadList(pagination.value.current, pagination.value.pageSize, {
      serveId: newServeId || 0,
    });
    if (newServeId) {
      await store.dispatch('Endpoint/getServerList', {id: newServeId});
      // 获取授权列表
      await store.dispatch('Endpoint/getSecurityList', {id: newServeId});
    }
  }
}, {
  immediate: true
})

async function refreshList() {
  await loadList(pagination.value.current, pagination.value.pageSize);
  //await store.dispatch('Endpoint/loadCategory');
}

watch(
    () => [createApiModalVisible.value, showImportModal.value, drawerVisible.value],
    async (newValue) => {
      if (!newValue[0] || !newValue[1] || !newValue[2]) {
        await store.dispatch('Endpoint/loadCategory');
      }
    },
    {immediate: true}
);

// 页面路由卸载时，清空搜索条件
onUnmounted(async () => {
  await store.commit('Endpoint/clearFilterState');
})


</script>
<style scoped lang="less">
.container {
  margin: 16px;
  background: #ffffff;
  height: calc(100vh - 96px);
  overflow: hidden;



  :deep(.ant-pagination) {
    margin-right: 24px;
  }
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
    margin-left: 8px;
    margin-right: 16px;
    cursor: pointer;
  }
}

.content {
  display: flex;
  width: 100%;
  position: relative;
  height: 100%;

  .left {
    width: 300px;
    border-right: 1px solid #f0f0f0;
    height: calc(100vh - 80px);
    overflow-y: scroll;
  }

  .right {
    flex: 1;
    overflow: scroll;

    &.right-not-collapsed {
      width: calc(100% - 300px);
    }
  }
}

.action-new {
  margin-right: 8px;
}

.top-action {
  width: 100%;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px;
  box-sizing: border-box;

  .ant-btn {
    margin-right: 16px;
  }

  .top-action-left {
    min-width: 220px;
  }
}

:deep(.top-action .ant-row.ant-form-item) {
  margin: 0;
}

.action-btns {
  display: flex;
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

:deep(.ant-alert-info) {
  padding: 12px;
}

:deep(.ant-alert-icon) {
  font-size: 14px;
  position: relative;
  top: 4px;
  left: 8px;
}

:deep(.ant-alert-message) {
  font-size: 14px;
}

:deep(.ant-alert-description) {
  font-size: 12px;
}

@media screen and (max-width: 1440px) {
  .right {
    width: 1180px;

    &.right-not-collapsed {
      width: calc(100% - 300px);
    }
  }
}


.customTitleColRender{
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  color: #447DFD;
}

</style>
