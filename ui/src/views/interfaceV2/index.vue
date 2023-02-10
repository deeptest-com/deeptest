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
            :tree-data="treeData"
        >
          <template #title="{ title, key }">
            <span v-if="key === '0-0-0'" style="color: #1890ff">{{ title }}</span>

            <template v-else>{{ title }}</template>
          </template>
        </a-tree>
      </div>
      <div class="right">
        <!--  头部区域  -->
        <div class="top-action">
          <a-button class="action-new" type="primary" :loading="loading" @click="addApi">
            新建接口
          </a-button>
          <a-button class="action-import" type="primary" :disabled="!hasSelected" :loading="loading" @click="importApi">
            导入
          </a-button>
        </div>
        <a-table
            :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
            :columns="columns"
            :data-source="data"
        >
          <template #title="{text}">
            <div class="customTitleColRender">
              <span>{{ text }}</span>
              <span class="edit" @click="addInterface"><EditOutlined/></span>
            </div>
          </template>

          <template #action>
            <div class="action-btns">
              <a-button type="link" @click="copy">复制</a-button>
              <a-button type="link" @click="del">删除</a-button>
              <a-button type="link" @click="disabled">过时</a-button>
            </div>
          </template>

        </a-table>
      </div>
    </div>
    <!-- 编辑接口时，展开抽屉   -->
    <EditInterfaceDrawer :visible="drawerVisible" @close="onCloseDrawer"/>
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
import {computed, reactive, toRefs, ref} from 'vue';
import {ColumnProps} from 'ant-design-vue/es/table/interface';
import {PlusOutlined, EditOutlined} from '@ant-design/icons-vue';
import {requestMethodOpts} from '@/config/constant';

import CreateApiModal from './components/CreateApiModal.vue';
import CreateTagModal from './components/CreateTagModal.vue'
import EditInterfaceDrawer from './components/EditInterfaceDrawer.vue'


type Key = ColumnProps['key'];

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
    slots: {customRender: 'title'},
  },
  {
    title: '状态',
    dataIndex: 'state',
  },
  {
    title: '创建人',
    dataIndex: 'name',
  },
  {
    title: '接口路径',
    dataIndex: 'address',
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


const treeData = [
  {
    title: '接口分类1',
    key: '0-0',
    children: [
      {
        title: '接口0-0-0',
        key: '0-0-0',
        children: [
          {title: '接口0-0-0-0', key: '0-0-0-0'},
          {title: '接口0-0-0-1', key: '0-0-0-1'},
          {title: '接口0-0-0-2', key: '0-0-0-2'},
        ],
      },
      {
        title: '接口0-0-1',
        key: '0-0-1',
        children: [
          {title: '接口0-0-1-0', key: '0-0-1-0'},
          {title: '接口0-0-1-1', key: '0-0-1-1'},
          {title: '接口0-0-1-2', key: '0-0-1-2'},
        ],
      },
    ],
  },
  {
    title: '接口分类2',
    key: '0-1',
  },
  {
    title: '接口分类3',
    key: '0-2',
  },
  {
    title: '接口分类4',
    key: '0-3',
  },
];

const data: DataType[] = [];
for (let i = 0; i < 46; i++) {
  data.push({
    key: i,
    title: `接口 ${i}`,
    name: `用户 ${i}`,
    index: i + 1,
    age: 32,
    address: `api/${i}`,
    updatedAt: new Date(),
  });
}


const selectedRowKeys: Key[] = ref([]);
const loading = false;

// 是否批量选中了
// const hasSelected = computed(() => state.selectedRowKeys.length > 0);
const hasSelected = false;


// 抽屉是否打开
const drawerVisible = ref<boolean>(false);

const onSelectChange = (selectedRowKeys: Key[]) => {
  console.log('selectedRowKeys changed: ', selectedRowKeys);
  selectedRowKeys.value = selectedRowKeys;
};


/**
 * 新建接口
 * */
function addInterface() {
  console.log('addInterface')
  drawerVisible.value = true;
}

/**
 * 批量操作
 * */
function batchHandle() {
  console.log('batchHandle')
}

function copy() {
  console.log()
}

function disabled() {
  console.log(1)
}

/**
 * 删除接口
 * */
function del() {
  console.log('删除接口')
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


function handleCreateApi() {
  createApiModalvisible.value = false;
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
