<template>
  <div class="content">
    <div class="header">
      <CustomForm :form-config="formConfig" :rules="rules" @handle-ok="handleAdd" />
    </div>
    <EmptyCom>
      <template #content>
        <a-table bordered :data-source="dataSource" :columns="versionColumns" :rowKey="(_record, index) => index">
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
      </template>
    </EmptyCom>
  </div>
</template>
<script setup lang="ts">

import {
  defineProps,
  watch,
  computed,
  createVNode
} from 'vue';
import { useStore } from 'vuex';
import { message, Modal } from "ant-design-vue";
import { ExclamationCircleOutlined } from '@ant-design/icons-vue';
import CustomForm from '../common/CustomForm.vue';
import EmptyCom from '@/components/TableEmpty/index.vue';
import { StateType as ProjectSettingStateType } from '../../store';
import { versionColumns } from '../../config';

const props = defineProps({
  serveId: {
    required: true
  },
})

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

let validateVersion = async (_rule: any, value: string) => {
  if (value === '') {
    return Promise.reject('版本号不能为空');
  } else {
    if (!/^(\d+\.){2}\d+/.test(value)) {
      return Promise.reject('请输入正确格式的版本号');
    }
    return Promise.resolve();
  }
};

const rules = {
  name: [{
    required: true,
    validator: validateVersion
  }]
}

const formConfig = [
  {
    type: 'input',
    modelName: 'name',
    valueType: 'string',
    placeholder: '版本号，格式如: 1.2.1',
    attrs: "width: 200px"
  },
  {
    type: 'select',
    modelName: 'createUser',
    valueType: 'string',
    options: userListOptions.value,
    placeholder: '请选择负责人',
    attrs: "width: 200px"
  },
  {
    type: 'input',
    modelName: 'description',
    valueType: 'string',
    placeholder: '输入描述',
    attrs: "width: 300px"
  },
  {
    type: 'button',
    text: '添加版本'
  },
]

async function handleAdd(formState: any) {
  if (!/^(\d+\.){2}\d+/.test(formState.name)) {
    message.error('请输入正确格式的版本号');
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
  const confirmCallBack = async () => {
    const result = await store.dispatch('ProjectSetting/deleteVersion', record.id);
    if (result) {
      await getList();
    }
  }
  Modal.confirm({
    title: '确认要删除版本号吗',
    icon: createVNode(ExclamationCircleOutlined),
    onOk() {
      confirmCallBack();
    }
  })
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
