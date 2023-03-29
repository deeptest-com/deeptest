<template>
  <div class="content">
    <div class="header">
      <a-form layout="inline" :model="formState">
        <a-form-item>
          <a-input v-model:value="formState.name" placeholder="请输入名称"/>
        </a-form-item>
        <a-form-item>
          <a-select
              v-model:value="formState.type"
              show-search
              placeholder="请选择类型"
              style="width: 200px"
              :options="securityTypeOptions"
              @focus="selectUserFocus"
              @blur="handleBlur"
              @change="handleChange"
          >
          </a-select>
        </a-form-item>
       
        <a-button
            class="editable-add-btn"
            @click="handleSave"
            type="primary"
            style="margin-bottom: 8px">
          添加Security
        </a-button>
      </a-form>
    </div>
    <a-table bordered :data-source="dataSource" :columns="securityColumns">
      <template #name="{ text, record }">
        <div class="editable-cell">
          <div class="editable-cell-text-wrapper">
            <a href="javascript:void (0)" @click="edit(record)">{{ text }}</a>
          </div>
        </div>
      </template>
      <template #operation="{ record }">
        <a-space>
        
          <a href="javascript:void (0)" @click="onDelete(record)">删除</a>
        </a-space>
      </template>
    </a-table>

    <!-- ::::编辑security组件 -->
    <a-modal v-model:visible="securityVisible"
             @cancel="handleCancel"
             width="882px"
             :closable="false"
             :key="securityVisibleKey"
             @ok="handleSave">
      <div class="editModal-content">
        <div class="modal-header">
          <div class="header-desc">
            <div class="name">
              编辑Security
            </div>
          </div>
        </div>
        <div>
            <a-form :model="formState" :label-col=" { span: 6 }" :wrapper-col="{ span: 14 }">
              <a-form-item label="Security名称" >       
                 <span>{{formState.name}}</span>        
                 <span class="default">
                   <a-checkbox :checked="formState.default" @change="changeDefault" />
                    设为默认
                 </span>  
               
              </a-form-item>
              <a-form-item label="Security类型" >
                 <a-select
                    v-model:value="formState.type"
                    show-search
                    placeholder="请选择类型"
                    :options="securityTypeOptions"
              
                >
                </a-select>
              </a-form-item>
              <div v-if="formState.type=='apiKey'">
                <a-form-item label="添加位置" >
                  <a-select
                      v-model:value="formState.in"
                      show-search
                      placeholder="请选择添加位置"
                      :options="addPositionOptions"
                  >
                  </a-select>
                </a-form-item>
                <a-form-item label="Key" >
                  <a-input v-model:value="formState.key"/>
                </a-form-item>
                <a-form-item label="Value" >
                  <a-input v-model:value="formState.value"/>
                </a-form-item>
              </div>
              <div v-if="formState.type=='bearerToken'">
                <a-form-item label="Token" >
                  <a-input v-model:value="formState.token"/>
                </a-form-item>
              </div>
               <div v-if="formState.type=='basicAuth'">
                <a-form-item label="Username" >
                  <a-input v-model:value="formState.username"/>
                </a-form-item>
                <a-form-item label="Password" >
                  <a-input v-model:value="formState.password"/>
                </a-form-item>
              </div>
            </a-form>
          </div>
  
      </div>
    </a-modal>

  </div>
</template>
<script setup lang="ts">

import {
  computed,
  defineEmits,
  defineProps,
  reactive,
  Ref,
  ref,
  UnwrapRef,
  watch
} from 'vue';
import {schema2yaml,saveSecurityInfo} from '../../service';
import SchemaEditor from '@/components/SchemaEditor/index.vue';
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import Filter from '../commom/Filter.vue';
import {MonacoOptions} from '@/utils/const';
import { securityColumns } from '../../config';
import { useStore } from 'vuex';
import {StateType as ProjectStateType} from "@/store/project";
import {StateType as ProjectSettingStateType} from '../../store';
import { Schema } from '../../data';
import { message } from 'ant-design-vue';
import {SelectTypes} from 'ant-design-vue/es/select';


const props = defineProps({
  serveId: {
    type: String,
    required: true
  },
})
const emit = defineEmits(['ok', 'close', 'refreshList']);

interface FormState {
  name: string;
  type: string,
  description?:string,
  in?:string,
  key?:string,
  value?:string,
  token?:string,
  username?:string,
  password?:string,
  default?:boolean
}

interface DataItem {
  key: string;
  name: string;
  age: number;
  address: string;
}


const formState: UnwrapRef<FormState> = reactive({
  name: '',
  type: '',
  description:"",
  in:"",
  key:"",
  value:"",
  token:"",
  username:"",
  password:"",
  default:false
});

const securityTypeOptions:SelectTypes['options']=[
  {
      label: "apiKey",
      value: "apiKey",
  },
  {
      label: "bearerToken",
      value: "bearerToken",
  },
  {
      label: "basicAuth",
      value: "basicAuth",
  }
]
const addPositionOptions:SelectTypes['options']=[
  {
      label: "Header",
      value: "header",
  },
  {
      label: "Query Params",
      value: "quseryParams",
  },
 
]




const securityVisible = ref(false);
const contentStr = ref('');
const schemaType = ref('object');
const exampleStr = ref('');
const keyword = ref('');
const isCheckAll = ref(false)
const store = useStore<{ ProjectGlobal: ProjectStateType, ProjectSetting: ProjectSettingStateType }>();
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const dataSource = computed<any>(() => store.state.ProjectSetting.securityList);

const securityVisibleKey = ref(0);
const edit = (record: any) => {
  securityVisible.value = true;
  formState.name=record.name;
  formState.type=record.type;
  formState.default=record.default;

  switch (record.type) {
    case 'apiKey': 
      formState.in=record.in;
      formState.key=record.key;
      formState.value=record.value;
      break;
    case 'bearerToken': 
      formState.token=record.token;
      break;
    case 'basicAuth': 
      formState.username=record.username;
      formState.password=record.password;
      break;
    default: 
      break;
  }

};

// 保存security
async function handleSave() {
   if (!formState.name) {
    message.error('security名称不能为空');
    return;
  }
   if (!formState.type) {
    message.error('security类型不能为空');
    return;
  }
  const res = await saveSecurityInfo({
    "projectId": currProject.value.id,
    "serveId": props.serveId,
    ...formState
   
  });
  if (res.code === 0) {
    message.success('保存security成功');
    // 清空表单中的数据
    formState.name = '';
    formState.type = '';
    securityVisible.value = false
    await getList();
  } else {
    message.error('保存失败：'+res.msg);
  }
}


function handleCancel() {
  formState.name = '';
  formState.type = '';
  securityVisible.value = false;
  
}

async function getList() {
  await store.dispatch('ProjectSetting/getSecurityList', {
    serveId: props.serveId,
   
  })
}


async function onDelete(record: any) {
  await store.dispatch('ProjectSetting/deleteSecurity', {
    id: record.id,
    serveId: props.serveId,
  })
}

function  changeDefault() {
 formState.default=!formState.default
}





watch(() => {
  return props.serveId
}, async () => {
  await getList();
}, {
  immediate: true
})
watch(() => {
  return securityVisible.value
}, () => {
  securityVisibleKey.value++
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
  .default{
    float: right;
  }
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
