<template>
  <div class="endpoint-debug-cases-list">
    <div class="toolbar">
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

import EditAndShowField from '@/components/EditAndShow/index.vue';
import CaseEdit from "./edit.vue";
import {NotificationKeyCommon} from "@/utils/const";

provide('usedBy', UsedBy.InterfaceDebug)
const simpleImage = Empty.PRESENTED_IMAGE_SIMPLE
const {t} = useI18n();

const store = useStore<{ Endpoint,Project }>();
const endpoint = computed<any>(() => store.state.Endpoint.endpointDetail);
const caseList = computed<any[]>(() => store.state.Endpoint.caseList);
const userList = computed<any>(() => store.state.Project.userList);
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
      notification.success({
        key: NotificationKeyCommon,
        message: `复制成功`,
      });
      design(po)
    } else {
      notification.error({
        key: NotificationKeyCommon,
        message: `复制失败`,
      });
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
    width: 100px;
  }

  .content {
    height: 100%;
  }
}

</style>
