<template>
  <div>
    <div style="margin-bottom: 16px">
      <a-button class="action-new" type="primary"  :loading="loading" @click="addInterface">
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
      <template #action>
        <div class="action-btn">
          <a-button type="link"  @click="copy">复制</a-button>
          <a-button type="link"  @click="del">删除</a-button>
          <a-button type="link"   @click="disabled">过时</a-button>
        </div>
      </template>

    </a-table>

    <a-drawer
        title="[接口编号] 接口名称"
        :placement="right"
        :width="1000"
        :closable="true"
        :visible="drawerVisible"
        @close="onCloseDrawer"
    >

      接口定义

    </a-drawer>

  </div>
</template>
<script setup lang="ts">
import {computed, reactive, toRefs,ref} from 'vue';
import {ColumnProps} from 'ant-design-vue/es/table/interface';

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


const data: DataType[] = [];
for (let i = 0; i < 46; i++) {
  data.push({
    key: i,
    title:`接口 ${i}`,
    name: `用户 ${i}`,
    index: i + 1,
    age: 32,
    address: `api/${i}`,
    updatedAt: new Date(),
  });
}



const selectedRowKeys:Key[] = ref([]);
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
 * 打开抽屉
 * */
function onCloseDrawer() {
    drawerVisible.value = false;
}
</script>

<style scoped lang="less">
.action-new{
  margin-right: 8px;
}
.action-btn {
  .ant-btn {
    display: inline-block;
  }


}
</style>
