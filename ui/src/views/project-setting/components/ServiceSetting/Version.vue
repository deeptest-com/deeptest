<template>
  <div class="content">
    <div class="header">
      <a-form layout="inline" :model="formState">
        <a-form-item name="name"  :rules="[{ required: true, message: '请设置版本号' }]">
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
          >
          </a-select>
        </a-form-item>
        <a-form-item>
          <a-input v-model:value="formState.description" style="width: 300px" placeholder="输入描述"/>
        </a-form-item>
        <a-button
            class="editable-add-btn"
            @click="handleAdd"
            type="primary"
            html-type="submit"
            style="margin-bottom: 8px">
          添加版本
        </a-button>
      </a-form>
    </div>
    <a-table bordered :data-source="dataSource" :columns="columns">
      <template #name="{ text }">
        <div class="editable-cell">
          <div class="editable-cell-text-wrapper">
            {{ text || ' ' }}
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
  defineEmits,
  defineProps,
  reactive,
  Ref,
  ref,
  UnwrapRef,
  watch,
  computed
} from 'vue';
import { useStore } from 'vuex';
import {message} from "ant-design-vue";
import { StateType as ProjectSettingStateType } from '../../store';

const props = defineProps({
  serveId: {
    type: String,
    required: true
  },
})

interface FormState {
  name: string;
  createUser: string | null | undefined,
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

const store = useStore<{ ProjectSetting: ProjectSettingStateType }>();
const dataSource = computed<DataItem[]>(() => store.state.ProjectSetting.serveVersionsList);
const userListOptions = computed<any[]>(() => store.state.ProjectSetting.userListOptions);


const formState: UnwrapRef<FormState> = reactive({
  name: '',
  description: '',
  createUser: null,
});

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

async function handleAdd() {
  if (!formState.name) {
    return;
  }
  const result = await store.dispatch('ProjectSetting/saveVersion', {
    "serveId": props.serveId,
    "value": formState.name,
    "createUser": formState.createUser,
    "description": formState.description
  });
  if (result) {
    message.success('添加版本号成功');
    // 清空表单中的数据
    formState.name = '';
    formState.createUser = null;
    formState.description = '';
    await getList();
  } else {
    message.error('保存失败');
  }
}


async function getList() {
  await store.dispatch('ProjectSetting/getVersionList');
}

async function onDelete(record: any) {
  const result = await store.dispatch('ProjectSetting/deleteVersion', record.id);
  if (result) {
    await getList();
  }
}

async function onDisabled(record: any) { 
  const result = await store.dispatch('ProjectSetting/disabledVersion', record.id);
  if (result) {
    await getList();
  }
}

async function selectUserFocus() {
  await store.dispatch('ProjectSetting/getUserOptionsList');
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
