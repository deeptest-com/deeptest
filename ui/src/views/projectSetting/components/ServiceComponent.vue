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
            <a href="javascript:void (0)" @click="edit(record)">{{ text }}</a>
          </div>
        </div>
      </template>
      <template #operation="{ record }">
        <a-space>
          <a href="javascript:void (0)" @click="onCopy(record)">复制</a>
          <a href="javascript:void (0)" @click="onDelete(record)">删除</a>
        </a-space>
      </template>
    </a-table>

    <!-- ::::新建组件 -->
    <a-modal v-model:visible="visible"
             @cancel="handleCancel"
             title="新建组件"
             @ok="handleOk">
      <a-form :model="formState" :label-col="{ span: 6 }" :wrapper-col=" { span: 15 }">
        <a-form-item label="组件名称">
          <a-input v-model:value="formState.name" placeholder="请输入内容"/>
        </a-form-item>
        <a-form-item label="描述">
          <a-select
              v-model:value="formState.tags"
              mode="tags"
              style="width: 100%"
              placeholder="Tags Mode"
              :options="[]"
              @change="handleTagChange"
          >
          </a-select>
        </a-form-item>
      </a-form>
    </a-modal>

    <!-- ::::编辑scheme组件 -->
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
import {getYaml} from "@/views/interfaceV2/service";
import {
  saveSchema,
  deleteSchema,
  disableSchema,
  copySchema,
  getSchemaList,
} from '../service';
import {message} from "ant-design-vue";

const props = defineProps({
  serveId:{
    type: String,
    required: true
  },
})
const emit = defineEmits(['ok', 'close', 'refreshList']);

interface FormState {
  name: string;
  description: string;
  tags: Array<string>,
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
  tags: [],
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
    dataIndex: 'type',
  },
  {
    title: '操作',
    dataIndex: 'operation',
    slots: {customRender: 'operation'},
  },
];

const dataSource: Ref<DataItem[]> = ref([]);
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

const edit = (record: string) => {
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


const handleAdd = () => {
  visible.value = true;
};

// 保存组件
async function handleOk() {
  const res = await saveSchema({
    "name": formState.name,
    "serveId": props.serveId,
    "tags":formState.tags.join(','),
  });
  if (res.code === 0) {
    message.success('保存成功');
    visible.value = false;
    await getList();
  } else {
    message.error('保存失败');
  }

}


function handleTagChange(value: string) {
  console.log(`832 ${value}`);
}

// 取消
function handleCancel() {
  visible.value = false;
}

async function getList() {
  const res = await getSchemaList({
    "serveId": props.serveId,
    "page": 1,
    "pageSize": 20
  });
  if (res.code === 0) {
    dataSource.value = res.data.result;
  }
}


async function onDelete(record: any) {
  const res = await deleteSchema(record.id);
  if (res.code === 0) {
    message.success('删除成功');
    await getList();
  } else {
    message.error('删除失败');
  }
}


async function onCopy(record: any) {
  const res = await copySchema(record.id);
  if (res.code === 0) {
    message.success('复制服务成功');
    await getList();
  } else {
    message.error('复制服务失败');
  }
}


watch(() => {
  return props.serveId
},async () => {
  await getList();
},{
  immediate:true
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

.editModal-content {
  min-height: 200px;
  position: relative;

  .btns {
    position: absolute;
    top: 8px;
    right: 8px;
  }

}
</style>
