<template>
  <div class="container">
    <div class="content">
      <div class="left tree">
        <div class="tag-filter-form">
          <a-input-search
              class="search-input"
              v-model:value="searchValue"
              placeholder="搜索接口分类"/>
          <div class="add-btn" @click="addApiTag">
            <PlusOutlined style="font-size: 16px;"/>
          </div>
        </div>
        <a-tree
            draggable
            @dragenter="onDragEnter"
            @drop="onDrop"
            :tree-data="treeData">
          <template #title="{ title, key }">
            <span v-if="key === '0-0-0'" style="color: #1890ff">{{ title }}</span>
            <template v-else>{{ title }}</template>
          </template>
        </a-tree>
      </div>
      <div class="right">
        <!--  头部区域  -->
        <div class="top-action">
          <a-row type="flex" style="align-items: center;width: 100%">
            <a-col :span="4">
              <a-button class="action-new" type="primary" :loading="loading" @click="addApi">新建接口</a-button>
              <a-button class="action-import" type="primary" :disabled="!hasSelected" :loading="loading"
                        @click="importApi">
                导入
              </a-button>
            </a-col>
            <a-col :span="1"/>
            <a-col :span="4">
              <a-form-item label="创建人" style="margin-bottom: 0;">
                <a-select placeholder="请选择创建人">
                  <a-select-option value="admin"> admin</a-select-option>
                  <a-select-option value="superAdmin">super admin</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :span="1"/>
            <a-col :span="4">
              <a-form-item label="状态" style="margin-bottom: 0;">
                <a-select placeholder="请选择创建人" :options="interfaceStatusOpts"/>
              </a-form-item>
            </a-col>
            <a-col :span="2"/>
            <a-col :span="7" style="margin-right: 8px;">
              <a-input-search
                  placeholder="请输入你需要搜索的接口文档"
                  enter-button
                  @search="() => {

                  }"
              />
            </a-col>

          </a-row>


        </div>
        <a-table
            :row-selection="{
          selectedRowKeys: selectedRowKeys,
          onChange: onSelectChange
        }"
            :columns="columns"
            :data-source="data">
          <template #colTitle="{text,record}">
            <div class="customTitleColRender">
              <span>{{ text }}</span>
              <span class="edit" @click="editInterface(record)"><EditOutlined/></span>
            </div>
          </template>
          <template #colStatus="{record}">
            <div class="customTitleColRender">
              <span>{{ interfaceStatus.get(record.status) }}</span>
            </div>
          </template>
          <template #action="{record}">
            <div class="action-btns">
              <a-button type="link" @click="copy(record)">复制</a-button>
              <a-button type="link" @click="del(record)">删除</a-button>
              <a-button type="link" @click="disabled(record)">过时</a-button>
            </div>
          </template>
        </a-table>
      </div>
    </div>
    <!-- 编辑接口时，展开抽屉   -->
    <EditInterfaceDrawer
        :destroyOnClose="true"
        :interfaceId="editInterfaceId"
        :visible="drawerVisible"
        :key="clickTag"
        @refreshList="refreshList"
        @close="onCloseDrawer"/>
    <!--  创建接口 Tag  -->
    <CreateTagModal
        :visible="createTagModalvisible"
        @cancal="handleCancalCreateTag"
        @ok="handleCreateTag"/>
    <!--  创建新接口弹框  -->
    <CreateApiModal
        :visible="createApiModalvisible"
        @cancal="handleCancalCreateApi"
        @ok="handleCreateApi"/>

  </div>
</template>
<script setup lang="ts">
import {computed, reactive, toRefs, ref, onMounted} from 'vue';
import {ColumnProps} from 'ant-design-vue/es/table/interface';
import {PlusOutlined, EditOutlined} from '@ant-design/icons-vue';
import {requestMethodOpts, interfaceStatus, interfaceStatusOpts} from '@/config/constant';
import {momentUtc} from '@/utils/datetime';
import {message} from 'ant-design-vue';
import {
  getInterfaceList,
  saveInterface,
  expireInterface,
  deleteInterface,
  copyInterface,
  getYaml,
  moveCategories,
  getCategories,
  deleteCategories,
  editCategories,
  newCategories
} from './service';
import CreateApiModal from './components/CreateApiModal.vue';
import CreateTagModal from './components/CreateTagModal.vue'
import EditInterfaceDrawer from './components/EditInterfaceDrawer.vue'


import {TreeDataItem, TreeDragEvent, DropEvent} from 'ant-design-vue/es/tree/Tree';

type Key = ColumnProps['key'];

// todo 待处理接口类型定义
interface DataType {
  key: Key;
  name: string;
  age: number;
  address: string;
}

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

// const treeData = [
//   {
//     title: '接口分类1',
//     key: '0-0',
//     children: [
//       {
//         title: '接口0-0-0',
//         key: '0-0-0',
//         children: [
//           {title: '接口0-0-0-0', key: '0-0-0-0'},
//           {title: '接口0-0-0-1', key: '0-0-0-1'},
//           {title: '接口0-0-0-2', key: '0-0-0-2'},
//         ],
//       },
//       {
//         title: '接口0-0-1',
//         key: '0-0-1',
//         children: [
//           {title: '接口0-0-1-0', key: '0-0-1-0'},
//           {title: '接口0-0-1-1', key: '0-0-1-1'},
//           {title: '接口0-0-1-2', key: '0-0-1-2'},
//         ],
//       },
//     ],
//   },
//   {
//     title: '接口分类2',
//     key: '0-1',
//   },
//   {
//     title: '接口分类3',
//     key: '0-2',
//   },
//   {
//     title: '接口分类4',
//     key: '0-3',
//   },
// ];


const treeData = ref([])

const data = ref([]);


async function reloadList() {
  let res = await getInterfaceList({
    "prjectId": 0,
    "page": 1,
    "pageSize": 100,
    "status": 0,
    "userId": 0,
    // "title": "接口名称"
  });
  const {result, total} = res.data;

  result.forEach((item, index) => {
    item.index = index + 1;
    item.key = `${index + 1}`;
    item.updatedAt = momentUtc(item.updatedAt);
  })
  data.value = result;
  // TODO 待处理分页逻辑
}

async function loadCategories() {
  let res = await getCategories({
    currProjectId: 1,
    // serveId, moduleId=2
    moduleId: 2
  });
  const {result, total} = res.data;

  result.forEach((item, index) => {
    item.index = index + 1;
    item.key = `${index + 1}`;
    item.updatedAt = momentUtc(item.updatedAt);
  })
}


function onDragEnter(info: TreeDragEvent) {
  console.log(info);

}

function onDrop(info: DropEvent) {
  console.log(info);
  const dropKey = info.node.eventKey;
  const dragKey = info.dragNode.eventKey;
  const dropPos = info.node.pos.split('-');
  const dropPosition = info.dropPosition - Number(dropPos[dropPos.length - 1]);
  const loop = (data: TreeDataItem[], key: string, callback: any) => {
    data.forEach((item, index, arr) => {
      if (item.key === key) {
        return callback(item, index, arr);
      }
      if (item.children) {
        return loop(item.children, key, callback);
      }
    });


  };
}


onMounted(async () => {
  await reloadList();
  await loadCategories();
})


async function refreshList() {
  await reloadList();

}

// const selectedRowKeys: Key[] = ref([]);
const selectedRowKeys = ref<Key[]>([]);
const loading = false;

// 是否批量选中了
// const hasSelected = computed(() => state.selectedRowKeys.length > 0);
const hasSelected = false;


// 抽屉是否打开
const drawerVisible = ref<boolean>(false);

const onSelectChange = (keys: Key[], rows: any) => {
  console.log('selectedRowKeys changed: ', keys, rows);
  selectedRowKeys.value = [...keys];
};


const editInterfaceId = ref('');


const clickTag = ref(0);

/**
 * 接口编辑
 * */
function editInterface(record) {
  console.log('editInterface');
  editInterfaceId.value = record.id;
  drawerVisible.value = true;
  clickTag.value++;
}

/**
 * 批量操作
 * */
function batchHandle() {
  console.log('batchHandle')
}

async function copy(record: any) {
  let res = await copyInterface(record.id);
  if (res.code === 0) {
    message.success('复制成功');
    await reloadList();

  }
}

async function disabled(record: any) {
  let res = await expireInterface(record.id);
  if (res.code === 0) {
    message.success('置为无效成功');
    await reloadList();
  }
}

/**
 * 删除接口
 * */
async function del(record: any) {
  let res = await deleteInterface(record.id);
  if (res.code === 0) {
    message.success('删除成功');
    await reloadList();
  }
}

/**
 * 接口导入逻辑
 * */
function importApi() {
  console.log('导入')
}

/**
 * 关闭抽屉
 * */
function onCloseDrawer() {
  drawerVisible.value = false;
}


const createTagModalvisible = ref(false);
const createApiModalvisible = ref(false);

/**
 * 添加接口分类
 * */
function addApiTag() {
  createTagModalvisible.value = true;
}

/**
 * 添加接口
 * */
function addApi() {
  createApiModalvisible.value = true;
}


async function handleCreateApi(data) {
  let res = await saveInterface({
    "serveId": 1,
    "path": data.path,
    "title": data.title,
  });
  createApiModalvisible.value = false;
  if (res.code === 0) {
    message.success('新建接口成功');
    await reloadList();
  }
}

function handleCancalCreateApi() {
  createApiModalvisible.value = false;
}

function handleCreateTag() {
  createTagModalvisible.value = false;
}

function handleCancalCreateTag() {
  createTagModalvisible.value = false;
}
</script>

<style scoped lang="less">
.container {
  margin: 16px;
  background: #ffffff;
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
  }

  .right {
    flex: 1
  }
}

.action-new {
  margin-right: 8px;
}

.top-action {
  height: 60px;
  display: flex;
  align-items: center;
  margin-left: 16px;

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
</style>
