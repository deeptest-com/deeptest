<template>
  <a-spin tip="Loading..." :spinning="isImporting" style="z-index: 2000;">
    <ContentPane>
      <template #left>
        <Tree @select="selectNode" :serveId="currServe.id"/>
      </template>
      <template #right>
        <div style="min-width: 1080px;overflow-x:scroll ">
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
                    <a-menu-item key="0">
                      <a-button :disabled="!hasSelected" type="link" :size="'small'" @click="batchUpdate">批量修改
                      </a-button>
                    </a-menu-item>
                  </a-menu>
                </template>
              </a-dropdown>
            </div>
            <div class="top-search-filter">
              <TableFilter @filter="handleTableFilter" ref="filter"/>
            </div>
          </div>
          <EmptyCom>
            <template #content>
              <a-table :loading="fetching"
                       :rowKey="'id'"
                       :row-selection="{
                      selectedRowKeys: selectedRowKeys,
                      onChange: onSelectChange
              }"
                       :pagination="{
                  ...pagination,
                  onChange: (page) => {
                    loadList(page,pagination.pageSize);
                  },
                  onShowSizeChange: (_page, size) => {
                    loadList(1,size);
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

                <template #colTags="{record}">
                  <div class="customTagsColRender">
                    <Tags
                        :values = "record?.tags"
                        :options = "tagList"
                        @updateTags = "(values:[])=>{
                      updateTags(values,record.id)
                    }"
                    />
                  </div>
                </template>
                <template #colCreateUser="{record}">
                  <div class="customTagsColRender">
                    {{username(record.createUser)}}
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
      </template>
    </ContentPane>
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
    <BatchUpdateFieldModal
        :visible="showBatchUpdateModal"
        :selectedCategoryId="selectedCategoryId"
        :selectedEndpointNum="selectedEndpointNum"
        @cancel="showBatchUpdateModal = false;"
        @ok="handleBatchUpdate"/>
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
  </a-spin>
</template>
<script setup lang="ts">
import {
  computed, reactive, toRefs, ref, onMounted,
  watch, createVNode, onUnmounted
} from 'vue';
import {useRouter} from 'vue-router';
import debounce from "lodash.debounce";
import {ColumnProps} from 'ant-design-vue/es/table/interface';
import {ExclamationCircleOutlined, MoreOutlined} from '@ant-design/icons-vue';
import {endpointStatusOpts, endpointStatus} from '@/config/constant';
import EditAndShowField from '@/components/EditAndShow/index.vue';
import ContentPane from '@/views/component/ContentPane/index.vue';
import CreateEndpointModal from './components/CreateEndpointModal.vue';
import PubDocs from './components/PubDocs.vue';
import ImportEndpointModal from './components/ImportEndpointModal.vue';
import TableFilter from './components/TableFilter.vue';
import Drawer from './components/Drawer/index.vue'
import EditAndShowSelect from '@/components/EditAndShowSelect/index.vue';
import EmptyCom from '@/components/TableEmpty/index.vue';
import PermissionButton from "@/components/PermissionButton/index.vue";
import {useStore} from "vuex";
import {Endpoint, PaginationConfig} from "@/views/endpoint/data";
import {StateType as ServeStateType} from "@/store/serve";
import {StateType as Debug} from "@/views/component/debug/store";
import {message, Modal, notification} from 'ant-design-vue';
import Tree from './components/Tree.vue'
import BatchUpdateFieldModal from './components/BatchUpdateFieldModal.vue';
import Tags from './components/Tags/index.vue';
const store = useStore<{ Endpoint, ProjectGlobal, Debug: Debug, ServeGlobal: ServeStateType,Project }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const currServe = computed<any>(() => store.state.ServeGlobal.currServe);
const serves = computed<any>(() => store.state.ServeGlobal.serves);
const list = computed<Endpoint[]>(() => store.state.Endpoint.listResult.list);
let pagination = computed<PaginationConfig>(() => store.state.Endpoint.listResult.pagination);
const createApiModalVisible = ref(false);
const router = useRouter();
type Key = ColumnProps['key'];
const tagList: any = computed(()=>store.state.Endpoint.tagList);
const userList = computed<any>(() => store.state.Project.userList);

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
    ellipsis: true,
    width: 150,
  },
  {
    title: '状态',
    dataIndex: 'status',
    slots: {customRender: 'colStatus'},
    width: 150,
  },
  {
    title: '标签',
    dataIndex: 'tags',
    slots: {customRender: 'colTags'},
    width: 200,
  },
  {
    title: '创建人',
    dataIndex: 'createUser',
    slots: {customRender: 'colCreateUser'},
    width: 100,
    ellipsis: true
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
    text: '过期',
    action: (record: any) => disabled(record)
  },
]

const selectedRowKeys = ref<Key[]>([]);

const selectedRowIds = computed(() => {
  const ids: any[] = [];
  Object.keys(selectedRow.value).forEach((key: string) => {
    ids.push(...selectedRow.value[key]);
  });
  return ids;
});

const selectedRow = ref<any>({});
const currentPage = ref(1);
const loading = false;
// 抽屉是否打开
const drawerVisible = ref<boolean>(false);
const selectedCategoryId = ref<string | number>('');
const onSelectChange = (keys: Key[], rows: any) => {
  selectedRowKeys.value = [...keys];
  selectedRow.value[currentPage.value] = rows.map((item: any) => item.id);
};
const hasSelected = computed(() => selectedRowKeys.value.length > 0);
const selectedEndpointNum = computed(() => selectedRowIds.value.length);

const fetching = ref(false);

/*查看选中的接口文档*/
function goDocs() {
  window.open(`/#/docs/view?endpointIds=${selectedRowIds.value.join(',')}`);
}

const showPublishDocsModal: any = ref(false)

// 发布文档版本
async function publishDocs() {
  showPublishDocsModal.value = false;
  selectedRowKeys.value = [];
  selectedRow.value = {};
  // selectedRowIds.value = [];
}

/**
 * 导入接口
 * */
const showImportModal = ref(false);

function inportApi() {
  showImportModal.value = true;
}

/**
 * 接口批量修改
 * */
const showBatchUpdateModal = ref(false);

function batchUpdate() {
  showBatchUpdateModal.value = true;
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
      // // 删除后重新拉取列表，根据当前页面和当前筛选条件
      // await loadList(pagination.value.current, pagination.value.pageSize, filterState.value);
      // // 重新拉取目录树
      // await store.dispatch('Endpoint/loadCategory');
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
    "curl": data.curl || null,
  });
  await refreshList('reset');
  createApiModalVisible.value = false;
}

async function handleBatchUpdate(data) {
  await store.dispatch('Endpoint/batchUpdateField', {
    "fieldName": data.value.fieldName,
    "value": data.value.value,
    "endpointIds":selectedRowIds.value
  });
  await refreshList();
  showBatchUpdateModal.value = false;
  selectedRow.value = {};
  selectedRowKeys.value = [];
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
    await refreshList('reset');
    if (callback) {
      callback();
    }
    showImportModal.value = false;
  }
  isImporting.value = false

}

// 当前筛选条件，包括分类、服务、状态
const filterState: any = ref({});

async function selectNode(id) {
  selectedCategoryId.value = id;
  selectedRowKeys. value = [];
  selectedRow.value = {};
  // 选中节点时，重置分页为第一页
  await loadList(1, pagination.value.pageSize, {
    categoryId: id,
    serveId: currServe.value.id,
  });
}


const loadList = debounce(async (page, size, opts?: any) => {
  fetching.value = true;
  currentPage.value = page;
  await store.dispatch('Endpoint/loadList', {
    "projectId": currProject.value.id,
    "page": page,
    "pageSize": size,
    ...opts,
  });
  // await store.dispatch('Endpoint/loadCategory');
  fetching.value = false;
}, 300)


async function handleTableFilter(state) {
  filterState.value = state;
  await loadList(1, pagination.value.pageSize, state);
}

const filter = ref()

// 实时监听项目/服务 ID，如果项目切换了则重新请求数据
watch(() => [currProject.value.id, currServe.value.id], async (newVal) => {
  const [newProjectId, newServeId] = newVal;
  if (newProjectId !== undefined) {
    await loadList(1, pagination.value.pageSize, {
      serveId: newServeId || 0,
    });
    await store.dispatch('Endpoint/getEndpointTagList');
    if (newServeId) {
      await store.dispatch('Debug/listServes', {serveId: newServeId});
      // 获取授权列表
      await store.dispatch('Endpoint/getSecurityList', {id: newServeId});
    }
    store.commit('Endpoint/clearFilterState');
    filter.value.resetFields()
  }
}, {
  immediate: true
})

async function refreshList(resetPage?: string) {
  await loadList(resetPage ? 1 : pagination.value.current, pagination.value.pageSize);
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
  store.commit('Endpoint/clearFilterState');
})


function paneResizeStop(pane, resizer, size) {
  console.log(pane.className, resizer.className, size.split('px')[0])
  if (pane?.className?.includes('left')) {
    const leftWidth = size.split('px')[0];
    // 当左侧宽度小于 100 时，折叠左侧
  }
}

const updateTags = async (tags :[],id:number)=>{
   await store.dispatch('Endpoint/updateEndpointTag', {
      id:id,tagNames:tags
    });

}

const username = (user:string)=>{
  let result = userList.value.find(arrItem => arrItem.value == user);
  return result?.label || '-'
}

</script>
<style scoped lang="less">
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
  overflow: hidden;

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

  }
}

.customTitleColRender {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  color: #447DFD;
}

</style>
