<template>
  <div class="scenario-list-main">
    <div class="filter-header">
      <div class="left">
        <a-space :size="16">
          <a-button type="primary" @click="() => edit(0)">新建测试场景</a-button>
          <!--          <a-button  :disabled="true" @click="() => edit(0)">批量操作</a-button>-->

        </a-space>
      </div>
      <div class="right">
        <a-form :layout="'inline'" class="filter-items">
          <a-form-item :label="'状态'">
            <a-select style="width:120px"  @change="onSearch" v-model:value="queryParams.enabled" :options="statusArr" class="status-select"/>
          </a-form-item>
          <a-form-item :label="'优先级'">
            <a-select style="width:120px"  @change="onSearch" v-model:value="queryParams.enabled" :options="statusArr" class="status-select"/>
          </a-form-item>
          <a-input-search @change="onSearch" @search="onSearch" v-model:value="queryParams.keywords"
                          placeholder="请输入你需要搜索的测试场景名称" style="width:270px;margin-left: 8px;"/>
        </a-form>
      </div>
    </div>
    <a-table
        v-if="list.length > 0"
        :row-selection="rowSelection"
        row-key="id"
        :columns="columns"
        :data-source="list"
        :loading="loading"
        :pagination="{
                ...pagination,
                onChange: (page) => {
                    getList(page, nodeDataCategory.id);
                },
                onShowSizeChange: (page, size) => {
                    pagination.pageSize = size
                    getList(page, nodeDataCategory.id);
                },
            }"
        class="dp-table"
    >
      <template #name="{ record ,text }">
        <EditAndShowField :custom-class="'custom-endpoint show-on-hover'"
                          :value="text"
                          placeholder="场景名称"
                          @update="(e: string) => handleUpdateName(e, record)"
                          @edit="editScenario(record)"/>

      </template>

      <template #desc="{ record  }">
        <span>{{record?.desc || '暂无描述'}}</span>
      </template>

      <template #updatedAt="{ record }">
        <span>{{momentUtc(record.updatedAt) }}</span>
      </template>

      <template #status="{ record }">
        <a-tag v-if="record.disabled" color="green">禁用</a-tag>
        <a-tag v-else color="cyan">启用</a-tag>
      </template>

      <template #action="{ record }">
        <a-dropdown>
          <MoreOutlined />
          <template #overlay>
            <a-menu>
              <a-menu-item key="0">
                <a class="operation-a" href="javascript:void (0)" @click="linkPlan(record.id)">关联测试计划</a>
              </a-menu-item>
              <a-menu-item key="1">
                <a class="operation-a" href="javascript:void (0)" @click="exec(record.id)">执行测试场景</a>
              </a-menu-item>
              <a-menu-item key="2">
                <a class="operation-a" href="javascript:void (0)" @click="design(record.id)">复制</a>
              </a-menu-item>
              <a-menu-item key="3">
                <a class="operation-a" href="javascript:void (0)" @click="design(record.id)">禁用</a>
              </a-menu-item>
              <a-menu-item key="4">
                <a class="operation-a" href="javascript:void (0)" @click="remove(record.id)">删除</a>
              </a-menu-item>

              <a-menu-item key="5">
                <a class="operation-a" href="javascript:void (0)" @click="design(record.id)">设计</a>
              </a-menu-item>
              <a-menu-item key="6">
                <a class="operation-a" href="javascript:void (0)" @click="edit(record.id)">编辑</a>
              </a-menu-item>

            </a-menu>
          </template>
        </a-dropdown>
      </template>

    </a-table>
    <a-empty v-if="list.length === 0" :image="simpleImage" />

  </div>
  <div v-if="isEditVisible">
    <a-modal :title="currModelId > 0 ? '编辑测试场景' : '新建测试场景'"
             :visible="true"
             :onCancel="onEditFinish"
             class="scenario-edit"
             :footer="null"
             width="600px">
      <ScenarioEdit
          :modelId="currModelId"
          :categoryId="nodeDataCategory.id"
          :onFinish="onEditFinish">
      </ScenarioEdit>
    </a-modal>
  </div>
  <a-modal v-model:visible="linkPlanVisible"
           :onCancel="onEditFinish"
           class="scenario-edit"
           :footer="null"
           width="600px">
    <template #title>
      <span>关联到测试计划 <span class="subTitle"> (已选择{{10}}个测试计划)</span></span>
    </template>
    <LinkPlan/>
  </a-modal>
  <DrawerDetail :destroyOnClose="true"
                :visible="drawerVisible"
                @refreshList="refreshList"
                @close="drawerVisible = false;"/>
</template>

<script setup lang="ts">
import {computed, onMounted, reactive, ref, watch} from "vue";
import { Empty } from 'ant-design-vue';
import { MoreOutlined } from "@ant-design/icons-vue";
import {SelectTypes} from 'ant-design-vue/es/select';
import {PaginationConfig, QueryParams, Scenario} from '../data.d';
import {useStore} from "vuex";
import {momentUtc} from "@/utils/datetime";
import {StateType} from "../store";
import debounce from "lodash.debounce";
import {useRouter} from "vue-router";
import {message, Modal, notification} from "ant-design-vue";
import {StateType as ProjectStateType} from "@/store/project";
import EditAndShowField from '@/components/EditAndShow/index.vue';
import ScenarioEdit from "../edit/index.vue";
import LinkPlan from "../edit/linkPlan.vue";
import DrawerDetail from "../components/Drawer/index.vue";
import { ColumnProps } from 'ant-design-vue/es/table/interface';

type Key = ColumnProps['key'];

interface DataType {
  key: Key;
  name: string;
  age: number;
  address: string;
}

const simpleImage = Empty.PRESENTED_IMAGE_SIMPLE

const statusArr = ref<SelectTypes['options']>([
  {
    label: '所有状态',
    value: '',
  },
  {
    label: '启用',
    value: '1',
  },
  {
    label: '禁用',
    value: '0',
  },
]);

const router = useRouter();
const store = useStore<{ Scenario: StateType, ProjectGlobal: ProjectStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const nodeDataCategory = computed<any>(()=> store.state.Scenario.nodeDataCategory);

const list = computed<Scenario[]>(() => store.state.Scenario.listResult.list);
let pagination = computed<PaginationConfig>(() => store.state.Scenario.listResult.pagination);
let queryParams = reactive<QueryParams>({
  keywords: '', enabled: '1',
  page: pagination.value.current, pageSize: pagination.value.pageSize
});

const currModelId = ref(0)

watch(nodeDataCategory, () => {
  console.log('watch nodeDataCategory', nodeDataCategory.value.id)
  getList(1, nodeDataCategory.value.id);
}, {deep: false})

watch(currProject, () => {
  console.log('watch currProject', currProject.value.id)
  getList(1, nodeDataCategory.value.id);
}, {deep: false})

onMounted(async () => {
  getList(1, nodeDataCategory.value.id);
})
const loading = ref<boolean>(true);

const getList = debounce(async (current: number, categoryId: number): Promise<void> => {
  console.log('getList')

  loading.value = true;

  await store.dispatch('Scenario/listScenario', {
    categoryId,
    keywords: queryParams.keywords,
    enabled: queryParams.enabled,
    pageSize: pagination.value.pageSize,
    page: current,
  });
  loading.value = false
}, 300)

const exec = (id: number) => {
  console.log('exec')
  router.push(`/scenario/exec/${id}`)
}

const linkPlanVisible = ref(false)
const linkPlan = () => {
  linkPlanVisible.value = true;
}

const design = (id: number) => {
  console.log('edit')
  router.push(`/scenario/design/${id}`)
}

const isEditVisible = ref(false)

const edit = (id: number) => {
  console.log('edit')
  currModelId.value = id
  isEditVisible.value = true
}

const onEditFinish = () => {
  console.log('onEditFinish')
  isEditVisible.value = false

  getList(pagination.value.current, nodeDataCategory.value.id)
}

const remove = (id: number) => {
  console.log('remove')

  Modal.confirm({
    title: '删除场景',
    content: '确定删除指定的场景？',
    okText: '确认',
    cancelText: '取消',
    onOk: async () => {
      store.dispatch('Scenario/removeScenario', id).then((res) => {
        console.log('res', res)
        if (res === true) {
          getList(1, nodeDataCategory.value.id)

          notification.success({
            message: `删除成功`,
          });
        } else {
          notification.error({
            message: `删除失败`,
          });
        }
      })
    }
  });
}


async function handleUpdateName(value: string, record: any) {
  console.log('handleUpdateName', value, record)
  // await store.dispatch('Scenario/updateEndpointDetail',
  //     {...record, title: value}
  // );
}

async function editScenario(record: any) {
  drawerVisible.value = true;
  // console.log('handleUpdateName', value, record)
  await store.dispatch('Scenario/getScenario', record.id);
}
// 抽屉是否打开
const drawerVisible = ref<boolean>(false);
// 关闭抽屉时，重新拉取列表数据
async function refreshList() {
  console.log('refreshList');
}
const rowSelection = {
  onChange: (selectedRowKeys: Key[], selectedRows: DataType[]) => {
    console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows);
  },
  getCheckboxProps: (record: DataType) => ({
    // disabled: record.name === 'Disabled User', // Column configuration not to be checked
    // name: record.name,
  }),
};


const onSearch = debounce(() => {
  getList(1, nodeDataCategory.value.id)
}, 500);

const columns = [
  {
    title: '编号',
    dataIndex: 'serialNumber',
  },
  {
    title: '测试场景名称',
    dataIndex: 'name',
    slots: {customRender: 'name'},
    ellipsis: true,
  },
  {
    title: '描述',
    dataIndex: 'desc',
    ellipsis: true,
    slots: {customRender: 'desc'},
  },
  {
    title: '状态',
    dataIndex: 'status',
    slots: {customRender: 'status'},
  },
  {
    title: '优先级',
    dataIndex: 'desc',
    ellipsis: true,
  },
  {
    title: '创建人',
    dataIndex: 'createUserName',
    width: 100,
    ellipsis: true,
  },
  {
    title: '最新更新',
    dataIndex: 'updatedAt',
    slots: {customRender: 'updatedAt'},
    ellipsis: true,
  },
  {
    title: '操作',
    key: 'action',
    width: 80,
    slots: {customRender: 'action'},
  },
];

onMounted(() => {
  console.log('onMounted')
})

</script>

<style lang="less" scoped>
.filter-header{
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 16px;
  height: 60px;
  .left{
    display: flex;
    align-items: center;
  }
  .right{
    display: flex;
    align-items: center;
  }
}
.filter-items{
  font-weight: normal;
}
.operation-a {
  text-align: center;
  display: inline-block;
  width: 80px;
}
.subTitle{
  font-size: 12px;
  color: #999;
  font-weight: normal;
}
</style>
