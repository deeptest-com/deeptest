<template>
  <div class="endpoint-debug-cases-list">
    <div class="toolbar">
<!--      <a-button trigger="click" @click="generateCases">
        <span>备选用例</span>
      </a-button>-->

      <a-button type="primary" trigger="click" @click="create">
        <span>新建用例</span>
      </a-button>
    </div>

    <div class="content">
      <a-table
          v-if="caseList.length > 0"
          :data-source="caseList"
          :columns="columns"
          :loading="loading"
          row-key="id"
          class="dp-table">

        <template #name="{ record, text }">
          <EditAndShowField placeholder="名称"
                            :custom-class="'custom-endpoint show-on-hover'"
                            :value="text || ''"
                            @update="(val) => updateName(val, record)"
                            @edit="design(record)"/>
        </template>

        <template #createdAt="{ record }">
          <span>{{ momentUtc(record.createdAt) }}</span>
        </template>

        <template #updatedAt="{ record }">
          <span>{{ momentUtc(record.updatedAt) }}</span>
        </template>

        <template #createUserName="{ record }">
          <span>{{ username(record.createUserName) }}</span>
        </template>

        <template #action="{ record }">
          <a-button type="link" @click="() => copy(record)">
            <CopyOutlined title="复制" />
          </a-button>

          <a-button type="link" @click="() => remove(record)">
            <DeleteOutlined title="删除" />
          </a-button>
        </template>

      </a-table>

      <a-empty class="dp-empty-no-margin"
               v-if="caseList.length === 0"
               :image="simpleImage"/>
    </div>

    <CaseEdit
        v-if="editVisible"
        :visible="editVisible"
        :model="editModel"
        :onFinish="createFinish"
        :onCancel="createCancel"/>

    <GenerateCasePopup
        v-if="generateCasesVisible"
        :visible="generateCasesVisible"
        :model="generateCasesModel"
        :onFinish="generateCasesFinish"
        :onCancel="generateCasesCancel" />
  </div>
</template>

<script lang="ts" setup>
import {computed, defineProps, provide, ref} from "vue";
import {UsedBy} from "@/utils/enum";
import {Empty, notification} from "ant-design-vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import {DeleteOutlined, CopyOutlined} from '@ant-design/icons-vue';
import {momentUtc} from '@/utils/datetime';
import debounce from "lodash.debounce";
import {confirmToDelete} from "@/utils/confirm";

import {StateType as Endpoint} from "@/views/endpoint/store";
import {StateType as Debug} from "@/views/component/debug/store";
import {StateType as Project} from "@/views/project/store";

import EditAndShowField from '@/components/EditAndShow/index.vue';
import CaseEdit from "./edit.vue";
import {notifyError, notifySuccess} from "@/utils/notify";
import GenerateCasePopup from "./generate.vue";

provide('usedBy', UsedBy.InterfaceDebug)
const simpleImage = Empty.PRESENTED_IMAGE_SIMPLE
const {t} = useI18n();

const store = useStore<{ Endpoint: Endpoint, Debug: Debug, Project: Project }>();
const endpoint = computed<any>(() => store.state.Endpoint.endpointDetail);
const caseList = computed<any[]>(() => store.state.Endpoint.caseList);
const userList = computed<any>(() => store.state.Project.userList);

const debugData = computed<any>(() => store.state.Debug.debugData);
const debugInfo = computed<any>(() => store.state.Debug.debugInfo);

const props = defineProps({
  onDesign: {
    type: Function,
    required: true,
  },
})

const loading = ref<boolean>(true);

const list = debounce(async (endpointId: number): Promise<void> => {
  console.log('getList')

  loading.value = true;
  await store.dispatch('Endpoint/listCase', endpointId);
  loading.value = false
}, 300)
list(endpoint.value.id)

const editVisible = ref(false)
const editModel = ref({} as any)
const create = () => {
  console.log('create')
  editVisible.value = true
  editModel.value = {title: ''}
}
const createFinish = (data) => {
  console.log('createFinish', data)

  data.endpointId = endpoint.value.id

  store.dispatch('Endpoint/saveCase', data).then((result) => {
    console.log('saveCase', result)

    editVisible.value = false
  })
}
const createCancel = () => {
  console.log('createVisible')
  editVisible.value = false
}
const remove = (record) => {
  console.log('remove', record)

  const title = '确定删除该用例吗？'
  confirmToDelete(title, '', () => {
    store.dispatch('Endpoint/removeCase', record);
  })
}
const copy  = (record) => {
  console.log('copy', record)
  store.dispatch('Endpoint/copyCase', record.id).then((po) => {
    if (po.id > 0) {
      notifySuccess(`复制成功`);
      design(po)
    } else {
      notifyError(`复制失败`);
    }
  })
}

const design = async (record: any) => {
  props.onDesign(record)
}
const updateName = async (value: string, record: any) => {
  await store.dispatch('Endpoint/updateCaseName', {
    id: record.id,
    name: value,
    endpointId: endpoint.value.id,
  });
  list(endpoint.value.id)
}

const username = (user:string)=>{
  let result = userList.value.find(arrItem => arrItem.value == user);
  return result?.label || '-'
}


const generateCasesVisible = ref(false)
const generateCasesModel = ref({} as any)
const generateCases = () => {
  console.log('generateCases')
  generateCasesVisible.value = true
  generateCasesModel.value = {}
}
const generateCasesFinish = async (model) => {
  console.log('generateCasesFinish', model, debugData.value.url)

  const data = Object.assign({...model}, debugInfo.value)

  store.commit("Global/setSpinning",true)
  const res = await store.dispatch('Debug/generateCases', data)
  store.commit("Global/setSpinning",false)

  if (res === true) {
    generateCasesVisible.value = false

    notifySuccess(`自动生成用例成功`);
  } else {
    notifyError(`自动生成用例保存失败`);
  }
}
const generateCasesCancel = () => {
  console.log('generateCasesCancel')
  generateCasesVisible.value = false
}

const columns = [
  {
    title: '编号',
    dataIndex: 'serialNumber',
    width: 120,
  },
  {
    title: '名称',
    dataIndex: 'name',
    slots: {customRender: 'name'},
  },
  {
    title: '创建人',
    dataIndex: 'createUserName',
    slots: {customRender: 'createUserName'},
    ellipsis: true,
    width: '100px',
  },
  {
    title: '创建时间',
    dataIndex: 'createdAt',
    slots: {customRender: 'createdAt'},
    ellipsis: true,
    width: '190px',
  },
  {
    title: '更新时间',
    dataIndex: 'updatedAt',
    slots: {customRender: 'updatedAt'},
    ellipsis: true,
    width: '190px',
  },
  {
    title: '操作',
    key: 'action',
    width: 200,
    slots: {customRender: 'action'},
  },
];

</script>

<style lang="less" scoped>
.endpoint-debug-cases-list {
  position: relative;

  height: 100%;

  .toolbar {
    position: absolute;
    z-index: 999999;
    top: -42px;
    right: 0;
    width: 200px;
    .ant-btn {
      margin-left: 10px;
    }
  }

  .content {
    height: 100%;
  }
}

</style>
