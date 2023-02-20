<template>
  <div class="content">
    <div class="header">
      <a-button class="editable-add-btn"
                @click="handleAdd"
                style="margin-bottom: 8px">
        添加版本
      </a-button>
      <a-input-search
          v-model:value="keyword"
          placeholder="输入组件名称搜索"
          style="width: 300px"
          @search="onSearch"
      />
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
          <a-popconfirm
              v-if="dataSource.length"
              title="Sure to delete?"
              @confirm="onDelete(record.key)"
          >
            <a>过期</a>
          </a-popconfirm>
          <a-popconfirm
              v-if="dataSource.length"
              title="Sure to delete?"
              @confirm="onDelete(record.key)"
          >
            <a>删除</a>
          </a-popconfirm>
        </a-space>
      </template>
    </a-table>

    <a-modal v-model:visible="visible"
             @cancel="handleCancel"
             title="添加版本"
             @ok="handleOk">
      <a-form :model="formState" :label-col="{ span: 6 }" :wrapper-col=" { span: 15 }">
        <a-form-item label="版本号">
          <a-input v-model:value="formState.name" placeholder="请输入内容"/>
        </a-form-item>
        <a-form-item label="版本号">
          <a-select
              v-model:value="value"
              show-search
              placeholder="Select a person"
              style="width: 200px"
              :options="options"
              @focus="handleFocus"
              @blur="handleBlur"
              @change="handleChange"
          >
          </a-select>
        </a-form-item>
        <a-form-item label="描述">
          <a-textarea
              v-model:value="formState.desc"
              placeholder="请输入内容"
              :auto-size="{ minRows: 2, maxRows: 5 }"
          />
        </a-form-item>
      </a-form>
    </a-modal>

  </div>

</template>
<script setup lang="ts">

import {computed, defineComponent, defineEmits, defineProps, reactive, Ref, ref, UnwrapRef} from 'vue';
import {CheckOutlined, EditOutlined} from '@ant-design/icons-vue';
import { SelectTypes } from 'ant-design-vue/es/select';


const props = defineProps({})
const emit = defineEmits(['ok', 'close', 'refreshList']);

interface FormState {
  name: string;
  desc: string;
}

interface DataItem {
  key: string;
  name: string;
  age: number;
  address: string;
}


const formState: UnwrapRef<FormState> = reactive({
  name: '',
  desc: '',
});

const visible = ref(false);
const drawerVisible = ref(false);

const columns = [
  {
    title: '版本号',
    dataIndex: 'name',
    width: '30%',
    slots: {customRender: 'name'},
  },
  {
    title: '负责人',
    dataIndex: 'age',
  },
  {
    title: '描述',
    dataIndex: 'address',
  },
  {
    title: '操作',
    dataIndex: 'operation',
    slots: {customRender: 'operation'},
  },
];

const dataSource: Ref<DataItem[]> = ref([
  {
    key: '0',
    name: 'Edward King 0',
    age: 32,
    address: 'London, Park Lane no. 0',
  },
  {
    key: '1',
    name: 'Edward King 1',
    age: 32,
    address: 'London, Park Lane no. 1',
  },
]);
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

const options = ref<SelectTypes['options']>([
  { value: 'jack', label: 'Jack' },
  { value: 'lucy', label: 'Lucy' },
  { value: 'tom', label: 'Tom' },
]);
const handleChange = (value: string) => {
  console.log(`selected ${value}`);
};
const handleBlur = () => {
  console.log('blur');
};
const handleFocus = () => {
  console.log('focus');
};
const filterOption = (input: string, option: any) => {
  return option.value.toLowerCase().indexOf(input.toLowerCase()) >= 0;
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

const save = (key: string) => {
  // Object.assign(dataSource.value.filter(item => key === item.key)[0], editableData[key]);
  // delete editableData[key];
};

const onDelete = (key: string) => {
  dataSource.value = dataSource.value.filter(item => item.key !== key);
};

const handleAdd = () => {
  visible.value = true;
};

// 确定
function handleOk() {
  visible.value = false;
}

// 取消
function handleCancel() {
  visible.value = false;
}

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
