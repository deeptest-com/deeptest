<template>
  <div class="content">
    <div class="header">
      <CustomForm
        :form-config="formConfig"
        :rules="rules"
        :search-placeholder="'输入组件名称搜索'"
        :show-search="true"
        @handle-ok="handleAdd"
        @handle-search="onSearch"
      />
    </div>
    <a-table bordered :data-source="dataSource" :columns="schemaColumns">
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
              <EditAndShowField :placeholder="'请输入内容'" :value="activeSchema?.name" @update="(e: string) => changeModelInfo('name', e)" />
            </div>
            <div class="desc" v-if="showMode === 'form'">
              <EditAndShowField :placeholder="'请输入内容'" :value="activeSchema?.description" @update="(e: string) => changeModelInfo('desc', e)" />
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
  defineProps,
  ref,
  watch,
  createVNode
} from 'vue';
import { useStore } from 'vuex';
import { Modal } from 'ant-design-vue';
import { ExclamationCircleOutlined,CodeOutlined,BarsOutlined } from '@ant-design/icons-vue';
import {schema2yaml} from '../../service';
import SchemaEditor from '@/components/SchemaEditor/index.vue';
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import CustomForm from '../common/CustomForm.vue';
import EditAndShowField from '@/components/EditAndShow/index.vue';
import {MonacoOptions} from '@/utils/const';
import { schemaColumns } from '../../config';
import {StateType as ProjectSettingStateType} from '../../store';

const props = defineProps({
  serveId: {
    type: String,
    required: true
  },
})

const rules = {
  name: [
    {
      required: true,
      message: '组件名称不能为空'
    }
  ]
}

const formConfig = [
  {
    type: 'input',
    modelName: 'name',
    placeholder: '请输入组件名称',
    valueType: 'string'
  },
  {
    type: 'select',
    modelName: 'tags',
    placeholder: '标签',
    options: [],
    valueType: 'array',
    mode: 'tags',
  },
  {
    type: 'button',
    text: '新建组件',
    modelName: ''
  }
]

const schemeVisible = ref(false);
const activeSchema: any = ref(null);
const contentStr = ref('');
const schemaType = ref('object');
const exampleStr = ref('');
const keyword = ref('');


const store = useStore<{ ProjectSetting: ProjectSettingStateType }>();
const dataSource = computed<any>(() => store.state.ProjectSetting.schemaList);

async function onSearch(e: any) {
  keyword.value = e;
  await getList();
}

async function changeModelInfo(type, e) {
  if (type === 'desc') {
    activeSchema.value.description = e;
  }
  if (type === 'name') {
    activeSchema.value.name = e;
  }
}

const showMode = ref('form');
const yamlCode = ref('');

async function switchMode(val: any) {
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

// 保存组件
async function handleAdd(formState: any) {
  if(formState.name) {
    formState.name = formState.name.trim();
  }
  // 判断组件名称必须是英文，复合 URL 规则
  // 匹配由数字、26个英文字母或者下划线或者 - 组成的字符串
  const reg = /^[a-zA-Z][\w-]*$/;
  if (!reg.test(formState.name)) {
    Modal.error({
      title: '组件名称必须是英文、下划线或 -',
      icon: createVNode(ExclamationCircleOutlined)
    });
    return;
  }

  const result = await store.dispatch('ProjectSetting/saveSchema', {
    schemaInfo: {
      "name": formState.name,
      "serveId": props.serveId,
      "tags": formState.tags.join(','),
    },
    action: 'create'
  })
  if (result) {
    await getList();
  }
}

async function handleEdit() {
  const result = await store.dispatch('ProjectSetting/saveSchema', {
    schemaInfo: {
      "name": activeSchema.value.name,
      "id": activeSchema.value.id,
      "serveId": props.serveId,
      "tags": activeSchema.value.tabs,
      "content": contentStr.value,
      "examples": exampleStr.value,
      "type": schemaType.value,
      "description": activeSchema.value.description
    },
    action: 'update'
  })
  if (result) {
    schemeVisible.value = false;
    await getList();
  }
}

function handleSchemeCancel() {
  schemeVisible.value = false;
}

async function getList() {
  await store.dispatch('ProjectSetting/getSchemaList', {
    serveId: props.serveId,
    name: keyword.value
  })
}


async function onDelete(record: any) {
  Modal.confirm({
    title: '确认要删除该组件吗',
    icon: createVNode(ExclamationCircleOutlined),
    onOk() {
      store.dispatch('ProjectSetting/deleteSchema', {
        id: record.id,
        serveId: props.serveId,
        name: keyword.value
      })
    }
  })
}


async function onCopy(record: any) {
  await store.dispatch('ProjectSetting/copySchema', {
    id: record.id,
    serveId: props.serveId,
    name: keyword.value
  })
}


function handleCodeChange() {
  console.log('代码改变');
}

async function generateFromJSON(JSONStr: string) {
  const result = await store.dispatch('ProjectSetting/generateSchema', {
    data: JSONStr
  })
  if (result) {
    activeSchema.value.content = result;
    contentStr.value = JSON.stringify(result);
    schemaType.value = result.type;
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
  const result = await store.dispatch('ProjectSetting/generateExample', {
    data: contentStr.value
  })
  if (result) {
    const example = {
      name: `Example ${examples.length + 1}`,
      content: JSON.stringify(result),
    };
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
