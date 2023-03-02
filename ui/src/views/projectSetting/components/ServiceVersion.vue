<template>
  <div class="content">
    <div class="header">
      <a-form layout="inline" :model="formState">
        <a-form-item>
          <a-input v-model:value="formState.name" placeholder="版本号，格式如: 1.2.1"/>
        </a-form-item>
        <a-form-item>
          <a-select
              v-model:value="formState.createUser"
              show-search
              placeholder="请选择负责人"
              style="width: 200px"
              :options="userListOptions"
              @focus="selectUserFocus"
              @blur="handleBlur"
              @change="handleChange"
          >
          </a-select>
        </a-form-item>
        <a-form-item>
          <a-input v-model:value="formState.description" style="width: 300px" placeholder="输入描述"/>
        </a-form-item>
        <a-button class="editable-add-btn"
                  @click="handleAdd"
                  style="margin-bottom: 8px">
          添加版本
        </a-button>
      </a-form>
    </div>
    <a-table bordered :data-source="dataSource" :columns="columns">
      <template #name="{ text, record }">
        <div class="editable-cell">
          <div class="editable-cell-text-wrapper">
            {{ text || ' ' }}
            <edit-outlined class="editable-cell-icon" @click="edit(record.key)"/>
          </div>
        </div>
      </template>
      <template #operation="{ record }">
        <a-space>
          <a href="javascript:void (0)" @click="onDisabled(record)">过期</a>
          <a href="javascript:void (0)" @click="onDelete(record)">删除</a>
        </a-space>
      </template>
    </a-table>
  </div>

</template>
<script setup lang="ts">

import {
  computed,
  defineComponent,
  defineEmits,
  defineProps,
  onMounted,
  reactive,
  Ref,
  ref,
  UnwrapRef,
  watch
} from 'vue';
import {CheckOutlined, EditOutlined} from '@ant-design/icons-vue';
import {SelectTypes} from 'ant-design-vue/es/select';
import {
  getServeVersionList,
  deleteServeVersion,
  disableServeVersions,
  saveServeVersion,
  deleteServe,
  getUserList,
  disableServe
} from '../service';
import {message} from "ant-design-vue";

const props = defineProps({
  serveId: {
    type: String,
    required: true
  },
})
const emit = defineEmits(['ok', 'close', 'refreshList']);

interface FormState {
  name: string;
  createUser:string,
  description: string;
}


interface DataItem {
  updatedAt: string;
  value: string;
  createUser: number;
  createdAt: string;
  id: string;
  serveId: string;
}


const formState: UnwrapRef<FormState> = reactive({
  name: '',
  description: '',
  createUser: '',
});

const visible = ref(false);
const drawerVisible = ref(false);

const columns = [
  {
    title: '版本号',
    dataIndex: 'value',
    width: '30%',
    slots: {customRender: 'value'},
  },
  {
    title: '负责人',
    dataIndex: 'createUser',
  },
  {
    title: '描述',
    dataIndex: 'description',
  },
  {
    title: '操作',
    dataIndex: 'operation',
    slots: {customRender: 'operation'},
  },
];


const dataSource: Ref<DataItem[]> = ref();
const count = computed(() => dataSource.value.length + 1);
const editableData: UnwrapRef<Record<string, DataItem>> = reactive({});

function onSearch(e) {
  console.log(e.target.value)
}

const keyword = ref('');

const activeKey = ref('1');

function onClose() {
  console.log('xxx')
  drawerVisible.value = false;

}

const userListOptions = ref<SelectTypes['options']>([]);


const handleChange = (value: string) => {
  console.log(`selected ${value}`);
};
const handleBlur = () => {
  console.log('blur');
};
const handleFocus = () => {
  console.log('focus');
};

const edit = (key: string) => {
  drawerVisible.value = true;
};

const isEditServiceDesc = ref(false);
const isEditServiceName = ref(false);

function editServiceDesc() {
  isEditServiceDesc.value = true;
}

function editServiceName() {
  isEditServiceName.value = true;
}

async function handleAdd() {
  const res = await saveServeVersion({
    "serveId": props.serveId,
    "value": formState.name,
    "createUser": formState.createUser,
    "description": formState.description
  });
  if (res.code === 0) {
    message.success('保存成功');
    await getList();
  } else {
    message.error('保存失败');
  }
}


async function getList() {
  const res = await getServeVersionList({
    "serveId": props.serveId,
    "createUser": "",
    "version": "",
    "page": 1,
    "pageSize": 20
  });
  if (res.code === 0) {
    dataSource.value = res.data.result;
  }
}

async function onDelete(record: any) {
  const res = await deleteServeVersion(record.id);
  if (res.code === 0) {
    message.success('删除版本成功');
    await getList();
  } else {
    message.error('删除版本失败');
  }
}

async function onDisabled(record: any) {
  const res = await disableServe(record.id);
  if (res.code === 0) {
    message.success('禁用版本成功');
    await getList();
  } else {
    message.error('禁用版本失败');
  }
}

async function setUserListOptions() {
  const res = await getUserList('');
  if (res.code === 0) {
    res.data.result.forEach((item) => {
      item.label = item.name;
      item.value = item.username
    })
    userListOptions.value = res.data.result;
  }
}

async function selectUserFocus(e) {
  await setUserListOptions();
}

watch(() => {
  return props.serveId
}, async () => {
  await getList();
}, {
  immediate: true
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
</style>
