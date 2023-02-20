<template>
  <div class="content">
    <div class="header">
      <a-button class="editable-add-btn"
                @click="handleAdd"
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
            <edit-outlined class="editable-cell-icon" @click="edit(record.key)"/>
          </div>
        </div>
      </template>
      <template #operation="{ record }">
        <a-popconfirm
            v-if="dataSource.length"
            title="Sure to delete?"
            @confirm="onDelete(record.key)"
        >
          <a>Delete</a>
        </a-popconfirm>
      </template>
    </a-table>

    <a-modal v-model:visible="visible"
             @cancel="handleCancel"
             title="新建服务"
             @ok="handleOk">
      <a-form :model="formState" :label-col="{ span: 6 }" :wrapper-col=" { span: 15 }">
        <a-form-item label="服务名称">
          <a-input v-model:value="formState.name" placeholder="请输入内容"/>
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

    <a-drawer
        :closable="false"
        :width="1000"
        :visible="drawerVisible"
        @close="onClose"
    >
      <template #title>
        <div class="drawer-header">
          <div>服务编辑</div>
          <div
              class="btns"
              :style="{}">
            <a-button style="margin-right: 8px" @click="onClose">Cancel</a-button>
            <a-button type="primary" @click="onClose">Submit</a-button>
          </div>
        </div>

      </template>
      <div class="drawer-content">
        <a-form :model="formState" :label-col="{ span: 2 }" :wrapper-col=" { span: 15 }">
          <a-form-item label="服务名称">
            <a-input v-model:value="formState.name" placeholder="请输入内容"/>
          </a-form-item>
          <a-form-item label="描述">
            <a-textarea
                v-model:value="formState.desc"
                placeholder="请输入内容"
                :auto-size="{ minRows: 2, maxRows: 5 }"
            />
          </a-form-item>

          <a-tabs v-model:activeKey="activeKey">
            <a-tab-pane key="1" tab="服务版本">
              <ServiceVersion/>
            </a-tab-pane>
            <a-tab-pane key="2" tab="服务组件" force-render>
              <ServiceComponent/>
            </a-tab-pane>
          </a-tabs>

        </a-form>
      </div>


    </a-drawer>

  </div>

</template>
<script setup lang="ts">

import {computed, defineComponent, defineEmits, defineProps, reactive, Ref, ref, UnwrapRef} from 'vue';
import {CheckOutlined, EditOutlined} from '@ant-design/icons-vue';
import ServiceVersion from './ServiceVersion';
import ServiceComponent from './ServiceComponent';


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
    title: '服务名称',
    dataIndex: 'name',
    width: '30%',
    slots: {customRender: 'name'},
  },
  {
    title: '描述',
    dataIndex: 'age',
  },
  {
    title: '状态',
    dataIndex: 'address',
  },
  {
    title: '创建人',
    dataIndex: 'address',
  },
  {
    title: '创建时间',
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

.drawer-header{
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
