<template>
  <div class="content">
    <div class="header">
      <a-button class="editable-add-btn"
                @click="handleAdd"
                style="margin-bottom: 8px">新建组件
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
            <a>克隆</a>
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
             title="新建组件"
             @ok="handleOk">
      <a-form :model="formState" :label-col="{ span: 6 }" :wrapper-col=" { span: 15 }">
        <a-form-item label="组件名称">
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

    <a-modal v-model:visible="schemeVisible"
             @cancel="handleCancel"
             title="组件编辑"
             width="1000px"
             @ok="handleOk">

      <div class="editModal-content">
        <div class="btns">
          <a-button :type="showMode === 'form' ? 'primary' : 'default'" @click="switchMode('form')">
            <template #icon>
              <BarsOutlined/>
            </template>
            图形
          </a-button>
          <a-button :type="showMode === 'code' ? 'primary' : 'default'" @click="switchMode('code')">
            <template #icon>
              <CodeOutlined/>
            </template>
            YAML
          </a-button>
        </div>

      </div>


    </a-modal>

  </div>
</template>
<script setup lang="ts">

import {computed, defineComponent, defineEmits, defineProps, reactive, Ref, ref, UnwrapRef} from 'vue';
import {CheckOutlined, EditOutlined} from '@ant-design/icons-vue';
import {getYaml} from "@/views/interfaceV2/service";

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
const schemeVisible = ref(false);

const columns = [
  {
    title: '组件名称',
    dataIndex: 'name',
    width: '30%',
    slots: {customRender: 'name'},
  },
  {
    title: '标签',
    dataIndex: 'age',
  },
  {
    title: '应用范围',
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


const showMode = ref('form');
const yamlCode = ref('');
async function switchMode(val) {
  showMode.value = val;
  // 需求去请求YAML格式
  // if (val === 'code') {
  //   let res = await getYaml(interfaceDetail.value);
  //   yamlCode.value = res.data;
  // }
}

const keyword = ref('');

const activeKey = ref('1');

function onClose() {
  console.log('xxx')
  schemeVisible.value = false;

}

const edit = (key: string) => {
  schemeVisible.value = true;
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

.editModal-content{
  min-height: 200px;
  position: relative;
  .btns{
    position: absolute;
    top:8px;
    right: 8px;
  }

}
</style>
