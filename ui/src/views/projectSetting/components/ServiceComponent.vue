<template>
  <div class="content">
    <div class="header">
      <a-button class="editable-add-btn"
                @click="handleAdd"
                type="primary"
                style="margin-bottom: 8px">新建组件
      </a-button>
      <a-input-search
          v-model:value="keyword"
          placeholder="输入组件名称搜索"
          style="width: 300px"
          @search="onSearch"/>
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
             :bodyStyle="{position:'relative'}"
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
              @change="handleTagChange">
          </a-select>
        </a-form-item>
      </a-form>
    </a-modal>

    <!-- ::::编辑scheme组件 -->
    <a-modal v-model:visible="schemeVisible"
             @cancel="handleSchemeCancel"
             width="882px"
             :closable="false"
             :key="schemeVisibleKey"
             @ok="handleEdit">
      <div class="editModal-content">
        <div class="modal-header">
          <div class="header-desc">
            <div class="name" v-if="showMode === 'form'">
              <a-input
                  @focusout="updateModelInfo"
                  @pressEnter="updateModelInfo"
                  @change="(e) => {
                    changeModelInfo('name',e)
                  }"
                  :value="activeSchema?.name"
                  placeholder="请输入内容"/>
            </div>
            <div class="desc" v-if="showMode === 'form'">
              <a-input
                  @focusout="updateModelInfo"
                  @pressEnter="updateModelInfo"
                  @change="(e) => {
                    changeModelInfo('desc',e)
                  }"
                  :value="activeSchema?.description"
                  placeholder="请输入内容"/>
            </div>
          </div>
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
        <!-- ::::表单模式 -->
        <div class="content-form" v-if="showMode === 'form'">
          <SchemaEditor
              :schemeVisibleKey="schemeVisibleKey"
              @generateFromJSON="generateFromJSON"
              @exampleChange="handleExampleChange"
              @generateExample="handleGenerateExample"
              @schemaTypeChange="handleSchemaTypeChange"
              @contentChange="handleContentChange"
              :value="activeSchema"/>
        </div>
        <!-- ::::代码模式 -->
        <div class="content-code" v-if="showMode === 'code'">
          <div style="border: 1px solid #f0f0f0; padding: 8px 0;">
            <MonacoEditor
                class="editor"
                :value="yamlCode"
                :language="'yaml'"
                :height="400"
                theme="vs"
                :options="{...MonacoOptions}"
                @change="handleCodeChange"/>
          </div>
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
// import {getYaml} from "@/views/interfaceV2/service";
import {
  saveSchema,
  deleteSchema,
  disableSchema,
  copySchema,
  getSchemaList,
  saveServe,
  example2schema,
  schema2example,
  schema2yaml,
} from '../service';
import SchemaEditor from '@/components/SchemaEditor/index.vue';
import {message} from "ant-design-vue";
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import {MonacoOptions} from '@/utils/const';

const props = defineProps({
  serveId: {
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
const activeSchema: any = ref(null);
const contentStr = ref('');
const schemaType = ref('object');
const exampleStr = ref('');

const columns = [
  {
    title: '组件名称',
    dataIndex: 'name',
    width: '30%',
    slots: {customRender: 'name'},
  },
  {
    title: '标签',
    dataIndex: 'tags',
  },
  {
    title: '操作',
    dataIndex: 'operation',
    slots: {customRender: 'operation'},
  },
];

const dataSource: Ref<DataItem[]> = ref([]);
const count = computed(() => dataSource.value.length + 1);

async function onSearch(e) {
  await getList();
  console.log(e.target.value)
}

async function changeModelInfo(type, e) {
  if (type === 'desc') {
    activeSchema.value.description = e.target.value;
  }
  if (type === 'name') {
    activeSchema.value.name = e.target.value;
  }
}

async function updateModelInfo() {
  // isEditModelName.value = false;
  // isEditModelDesc.value = false;
  // if (activeSchema.value.name && activeSchema.value.description) {
  //   const res = await saveSchema({
  //     "projectId": 1,
  //     "name": activeSchema.value.name,
  //     "description": activeSchema.value.description,
  //     "id": props.serveId,
  //   });
  //   if (res.code === 0) {
  //     await getList();
  //   }
  // }
}


const showMode = ref('form');
const yamlCode = ref('');

async function switchMode(val) {
  showMode.value = val;
  // 需求去请求YAML格式
  const content = activeSchema.value.content;
  if (val === 'code') {
    let res = await schema2yaml({
      data: JSON.stringify(content)
    });
    yamlCode.value = res.data;
  }
}

const keyword = ref('');

const activeKey = ref('1');

function onClose() {
  console.log('xxx')
  schemeVisible.value = false;
}

const schemeVisibleKey = ref(0);
const edit = (record: any) => {
  schemeVisible.value = true;
  record.content =  record.content &&  typeof record.content === 'string' ? JSON.parse(record.content) : {type: 'object'};
  record.examples = record.examples && typeof record.examples === 'string' ? JSON.parse(record.examples) : [];
  activeSchema.value = record;
  contentStr.value = JSON.stringify(record?.content || '');
  exampleStr.value = JSON.stringify(record?.examples || '');
  schemaType.value = record?.type || '';
  schemeVisibleKey.value++;
};


const handleAdd = () => {
  visible.value = true;
};

// 保存组件
async function handleOk() {
  const res = await saveSchema({
    "name": formState.name,
    "serveId": props.serveId,
    "tags": formState.tags.join(','),
  });
  if (res.code === 0) {
    message.success('新建组件成功');
    visible.value = false;
    await getList();
  } else {
    message.error('新建组件失败');
  }
}

async function handleEdit() {
  const res = await saveSchema({
    "name": activeSchema.value.name,
    "id": activeSchema.value.id,
    "serveId": props.serveId,
    "tags": activeSchema.value.tabs,
    "content": contentStr.value,
    "examples": exampleStr.value,
    "type": schemaType.value,
    "description": activeSchema.value.description
  });
  if (res.code === 0) {
    schemeVisible.value = false;
    message.success('修改组件成功');
    visible.value = false;
    await getList();
  } else {
    message.error('修改组件失败');
  }
}


function handleTagChange(value: string) {
  console.log(`832 ${value}`);
}

function handleCancel() {
  visible.value = false;
}

function handleSchemeCancel() {
  schemeVisible.value = false;
}

async function getList() {
  const res = await getSchemaList({
    "serveId": props.serveId,
    "page": 1,
    "pageSize": 20,
    name: keyword.value
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


function handleCodeChange() {
  console.log('代码改变');
}

async function generateFromJSON(JSONStr: string) {
  const res = await example2schema({
    data: JSONStr
  });
  if (res.code === 0) {
    activeSchema.value.content = res.data;
    contentStr.value = JSON.stringify(res.data);
    schemaType.value = res.data.type;
  }
}

function handleContentChange(str: string) {
  contentStr.value = str;
}

function handleSchemaTypeChange(str: string) {
  schemaType.value = str;
}

function handleExampleChange(str: string) {
  exampleStr.value = str;
}

async function handleGenerateExample(examples: any) {
  const res = await schema2example({
    data: contentStr.value
  });

  const example = {
    name: `Example ${examples.length + 1}`,
    content: JSON.stringify(res.data),
  };

  if (res.code === 0) {
    activeSchema.value.examples.push(example);
  }
}

watch(() => {
  return props.serveId
}, async () => {
  await getList();
}, {
  immediate: true
})
watch(() => {
  return schemeVisible.value
}, () => {
  schemeVisibleKey.value++
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

.btns {
  display: flex;
  justify-content: flex-end;
}

.modal-header {
  display: flex;
  justify-content: space-between;
}

.editModal-content {
  min-height: 200px;
}

.content-form {
  //margin-top: 32px;
}

.content-code {
  // margin-top: 32px;
}

.header-desc {
  flex: 1;
  margin-right: 36px;

  .name {
    height: 24px;
    line-height: 24px;
    font-weight: bold;
    font-size: 18px;
    margin-bottom: 16px;

    input {
      border: none;
      height: 32px;
      line-height: 32px;
      font-size: 18px;

      &:hover {
        border: 1px solid #1aa391;
      }
    }
  }

  .desc {
    height: 16px;
    line-height: 16px;
    font-weight: bold;
    font-size: 14px;
    margin-bottom: 32px;

    input {
      border: none;
      height: 24px;
      line-height: 24px;
      font-size: 12px;

      &:hover {
        border: 1px solid #1aa391;
      }
    }
  }
}
</style>
