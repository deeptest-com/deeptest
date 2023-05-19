<template>
  <div class="interface-list-main">
    <div class="table-toolbar">
      <div class="filters">
        <TreeSelect/>
        <a-input-search
            style="display: flex;justify-content: end;width: 300px;margin-bottom: 16px; "
            placeholder="请输入关键词"
            enter-button
            v-model:value="filters.keywords"
            @change="onKeywordsChanged"
            @search="selectCategory"/>
      </div>
    </div>

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
            onShowSizeChange: (page, size) => {
              loadList(page,size);
            },
        }"
        :columns="columns"
        :data-source="interfaces"
        rowKey="id">

      <template #colName="{text}">
        {{ text }}
      </template>

      <template #colMethod="{text}">
        {{ text }}
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
import {useRouter} from 'vue-router';
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
import {message, Modal} from 'ant-design-vue';
import {listEndpointInterface} from "@/views/endpoint/service";
import TreeSelect from "./TreeSelect.vue";

const store = useStore<{ Endpoint, ProjectGlobal, Debug: Debug, ServeGlobal: ServeStateType }>();

const interfaces = ref<any[]>([])
const filters = ref<any>({})
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);

const props = defineProps({
  // categoryId: {
  //   type: Number,
  //   required: true,
  // },
  selectInterface: {
    type: Function,
    required: true,
  },
})

const pagination = ref<any>({
  total: 0,
  page: 1,
  pageSize: 10,
  showSizeChanger: true,
  showQuickJumper: true,
})

// watch(props, () => {
//   console.log('watch props.categoryId for reload', props.categoryId)
//   selectCategory()
// }, {deep: true})
//
const selectCategory = async () => {
  // console.log('selectCategory', props.categoryId)
  //
  // if (props.categoryId === 0) {
  //   interfaces.value = []
  //   return
  // }
  loadList(pagination.value.page, pagination.value.pageSize)
}

const onKeywordsChanged = debounce(async () => {
  console.log('onKeywordsChanged')
  await selectCategory()
}, 600)

const loadList = debounce(async (page, pageSize) => {
  pagination.value.page = page
  pagination.value.pageSize = pageSize
  const data = {
    // categoryId: props.categoryId,
    keywords: filters.value.keywords,
    projectId: currProject.value.id,
  }
  const result = await listEndpointInterface(data, pagination.value)
  console.log('listInterface', result)
  interfaces.value = result?.list
  pagination.value.total = result?.total
}, 300)

type Key = ColumnProps['key'];
const selectedRowKeys = ref<Key[]>([]);

const onSelectChange = (keys: Key[], rows: any) => {
  selectedRowKeys.value = [...keys];
  props.selectInterface(selectedRowKeys.value)
};

const columns = [
  {
    title: '序号',
    dataIndex: 'index',
    width: 80,
    customRender: ({
                     text,
                     index
                   }: { text: any; index: number }) => (pagination.value.page - 1) * pagination.value.pageSize + index + 1,
  },
  {
    title: '端点名称',
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

onMounted(async () => {
  await loadList(pagination.value.page, pagination.value.pageSize)
})
</script>

<style scoped lang="less">
.interface-list-main {
  margin: 0;

  .table-toolbar {
    display: flex;

    .actions {
      width: 200px;
      line-height: 36px;
    }

    .filters {
      flex: 1;
      display: flex;
      justify-content: space-between;
    }
  }
}

</style>
