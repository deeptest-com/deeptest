<template>
  <div class="interface-list-main">
    <a-table
        :row-selection="{
          selectedRowKeys: selectedRowKeys,
          onChange: onSelectChange
        }"
        :pagination="{
            ...pagination,
            onChange: (page) => {
              loadList(page, pagination.pageSize);
            },
        }"
        :columns="columns"
        :data-source="interfaces"
        rowKey="id">

      <template #colName="{text}">
        {{text}}
      </template>

      <template #colMethod="{text}">
        {{text}}
      </template>

      <template #colUrl="{text}">
        <div class="customPathColRender">
          <a-tag>{{ text }}</a-tag>
        </div>
      </template>
    </a-table>
  </div>
</template>
<script setup lang="ts">
import {
  computed, reactive, toRefs, ref, onMounted,
  watch, defineProps
} from 'vue';
import { useRouter } from 'vue-router';
import debounce from "lodash.debounce";
import EndpointTree from './list/tree.vue';
import {ColumnProps} from 'ant-design-vue/es/table/interface';
import {MoreOutlined} from '@ant-design/icons-vue';
import {endpointStatusOpts} from '@/config/constant';
import EditAndShowField from '@/components/EditAndShow/index.vue';
import CreateEndpointModal from './components/CreateEndpointModal.vue';
import TableFilter from './components/TableFilter.vue';
import Drawer from './components/Drawer/index.vue'
import {useStore} from "vuex";
import {Endpoint, PaginationConfig} from "@/views/endpoint/data";

import {StateType as ServeStateType} from "@/store/serve";
import {StateType as Debug} from "@/views/component/debug/store";
import { message, Modal } from 'ant-design-vue';
import {listEndpointInterface} from "@/views/endpoint/service";

const store = useStore<{ Endpoint, ProjectGlobal, Debug: Debug, ServeGlobal: ServeStateType }>();

const interfaces = ref<any[]>([])
type Key = ColumnProps['key'];

const props = defineProps({
  categoryId: {
    type: Number,
    required: true,
  },
})

const pagination = ref<PaginationConfig>( {
  total: 0,
  current: 1,
  pageSize: 10,
  showSizeChanger: true,
  showQuickJumper: true,
})

watch(props, () => {
  console.log('watch props.categoryId for reload', props.categoryId)
  selectCategory()
}, {deep: true})

const selectCategory = async () => {
  console.log('selectCategory', props.categoryId)

  if (props.categoryId === 0) {
    interfaces.value = []
    return
  }

  const result = await listEndpointInterface(props.categoryId, pagination)
  console.log('listInterface', result)
  interfaces.value = result?.list
  pagination.value.total = result?.total
}

const columns = [
  {
    title: '序号',
    dataIndex: 'index',
    width: 80,
    customRender: ({text, index}: { text: any; index: number }) => (pagination.value.current - 1) * pagination.value.pageSize + index + 1,
  },
  {
    title: '名称',
    dataIndex: 'name',
    slots: {customRender: 'colName'},
  },
  {
    title: '请求方法',
    dataIndex: 'method',
    slots: {customRender: 'colMethod'},
  },
  {
    title: '路径',
    dataIndex: 'url',
    slots: {customRender: 'colUrl'},
  }
];
const selectedRowKeys = ref<Key[]>([]);
const loading = false;
// 抽屉是否打开
const drawerVisible = ref<boolean>(false);
const selectedCategoryId = ref<string>('');
const onSelectChange = (keys: Key[], rows: any) => {
  selectedRowKeys.value = [...keys];
};

async function handleChangeStatus(value: any, record: any,) {
  await store.dispatch('Endpoint/updateStatus', {
    id:record.id,
    status: value
  });
}

async function copy(record: any) {
  await store.dispatch('Endpoint/copy', record);
}

async function disabled(record: any) {
  await store.dispatch('Endpoint/disabled', record);
}

async function del(record: any) {
  await store.dispatch('Endpoint/del', record);
}

const loadList = debounce(async (page, size, opts?: any) => {
  console.log('1')
  // await store.dispatch('Endpoint/loadList', {
  //   "projectId": currProject.value.id,
  //   "page": page,
  //   "pageSize": size,
  //   opts,
  // });
}, 300)

async function refreshList() {
  await loadList(pagination.value.current, pagination.value.pageSize);
}

</script>

<style scoped lang="less">
.interface-list-main {
  margin: 0;
}

</style>
