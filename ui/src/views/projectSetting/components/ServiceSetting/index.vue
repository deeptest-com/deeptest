<template>
  <div class="content">
    <div class="header">
      <a-button class="editable-add-btn"
                @click="handleAdd"
                type="primary"
                style="margin-bottom: 8px">新建服务
      </a-button>
      <a-input-search
          v-model:value="keyword"
          placeholder="输入服务名称搜索"
          style="width: 300px"
          @search="onSearch"
      />
    </div>
    <a-table bordered :data-source="dataSource" :columns="columns">
      <template #name="{ text, record }">
        <div class="editable-cell">
          <div class="editable-cell-text-wrapper">
            {{ text || ' ' }}
            <edit-outlined class="editable-cell-icon" @click="edit(record)"/>
          </div>
        </div>
      </template>
      <template #operation="{ record }">
        <a-space>
          <a href="javascript:void (0)" @click="onDisabled(record)">禁用</a>
          <a href="javascript:void (0)" @click="onCopy(record)">复制</a>
          <a href="javascript:void (0)" @click="onDelete(record)">删除</a>
        </a-space>
      </template>
    </a-table>
    <!-- ::::新建服务弹框 -->
    <a-modal v-model:visible="visible"
             @cancel="handleCancel"
             title="新建服务"
             @ok="handleOk">
      <a-form :model="formState" :label-col="{ span: 6 }" :wrapper-col=" { span: 15 }">
        <a-form-item label="服务名称">
          <a-input v-model:value="formState.name" placeholder="请输入内容"/>
        </a-form-item>
        <a-form-item label="描述">
          <a-input v-if="editServiceDesc" v-model:value="formState.description" placeholder="请输入内容"/>
        </a-form-item>
      </a-form>
    </a-modal>

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
          </a-tabs>
        </a-form>
      </div>
    </a-drawer>

  </div>

</template>
<script setup lang="ts">

import {computed, defineComponent, defineEmits, defineProps, onMounted, reactive, Ref, ref, UnwrapRef} from 'vue';
import {CheckOutlined, EditOutlined} from '@ant-design/icons-vue';
import ServiceVersion from './Version.vue';
import ServiceComponent from './Component.vue';

import {getServeList, deleteServe, copyServe, disableServe, saveServe} from '../../service';
import {momentUtc} from '@/utils/datetime';
import {message} from "ant-design-vue";
import {serveStatus} from "@/config/constant";

const props = defineProps({})
const emit = defineEmits(['ok', 'close', 'refreshList']);

interface FormState {
  name: string;
  description: string;
  serveId?: string,

}

interface DataItem {
  key: string;
  name: string;
  age: number;
  address: string;
}


const formState: UnwrapRef<FormState> = reactive({
  name: '',
  description: '',
});

const editFormState: UnwrapRef<FormState> = reactive({
  name: '',
  description: '',
  serveId: '',
});


const visible = ref(false);
const drawerVisible = ref(false);

const columns = [
  {
    title: '服务名称',
    dataIndex: 'name',
    width: '30%',
    slots: {customRender: 'name'},
  },
  {
    title: '描述',
    dataIndex: 'description',
  },
  {
    title: '状态',
    dataIndex: 'statusDesc',
  },
  {
    title: '创建人',
    dataIndex: 'createUser',
  },
  {
    title: '创建时间',
    dataIndex: 'createdAt',
  },
  {
    title: '操作',
    dataIndex: 'operation',
    slots: {customRender: 'operation'},
  },
];
const dataSource: Ref<DataItem[]> = ref([]);
const count = computed(() => dataSource.value.length + 1);
const editKey = ref(0);

const keyword = ref('');

const activeKey = ref('1');

function onClose() {
  drawerVisible.value = false;

}


async function onSearch(e) {
  await getList();
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
    const res = await saveServe({
      "projectId": 1,
      "name": editFormState.name,
      "description": editFormState.description,
      "id": editFormState.serveId,
    });
    if (res.code === 0) {
      // message.success('修改服务描述成功');
      await getList();
    } else {
      // message.error('修改服务描述失败');
    }
  }
}

async function onDelete(record: any) {
  const res = await deleteServe(record.id);
  if (res.code === 0) {
    message.success('删除成功');
    await getList();
  } else {
    message.error('删除失败');
  }
}

async function onDisabled(record: any) {
  const res = await disableServe(record.id);
  if (res.code === 0) {
    message.success('禁用服务成功');
    await getList();
  } else {
    message.error('禁用服务失败');
  }
}

async function onCopy(record: any) {
  const res = await copyServe(record.id);
  if (res.code === 0) {
    message.success('复制服务成功');
    await getList();
  } else {
    message.error('复制服务失败');
  }
}

const handleAdd = () => {
  visible.value = true;
};

// 确定
async function handleOk() {
  visible.value = false;
  // :::: todo 需要更换数据
  const res = await saveServe({
    "projectId": 1,
    "name": formState.name,
    "description": formState.description
  });
  if (res.code === 0) {
    message.success('新建服务成功');
    await getList();
  } else {
    message.error('新建服务失败');
  }
}

// 取消
function handleCancel() {
  visible.value = false;
}


async function getList() {
  let res = await getServeList({
    "projectId": 1,
    "page": 0,
    "pageSize": 100,
    "name": keyword.value,
  });
  if (res.code === 0) {
    res.data.result.forEach((item) => {
      item.statusDesc = serveStatus.get(item.status);
      item.createdAt = momentUtc(item.createdAt)
    })
    dataSource.value = res.data.result;
  }
}

onMounted(async () => {
  await getList()
})

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

  .btns {
    //border-bottom: 1px solid #e9e9e9;
    //padding: 10px 16px;
    //background: #fff;
    //margin-bottom: 8px;
  }
}

</style>
