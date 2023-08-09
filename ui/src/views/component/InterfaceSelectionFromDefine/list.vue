<template>
  <div class="interface-list-main">
    <div class="table-toolbar">
      <div class="filters">
        <a-input-search
            style="display: flex;justify-content: end;width: 300px;margin-bottom: 16px; "
            placeholder="请输入接口名称或路径"
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
        :scroll="{ y: 310 }"
        rowKey="id">

      <template #colName="{text}">
        {{ text }}
      </template>

      <template #colMethod="{text}">
        {{ text }}
      </template>

      <template #colUrl="{text}">
        <div class="customPathColRender">
          <a-tooltip placement="topLeft">
            <template #title>
              <span>{{ text }}</span>
            </template>
            <a-tag>{{ text }}</a-tag>
          </a-tooltip>
        </div>
      </template>
    </a-table>
  </div>
</template>
<script setup lang="ts">
import {
  computed, reactive, toRefs, ref, onMounted,
  watch, defineProps, nextTick
} from 'vue';
import debounce from "lodash.debounce";
import {ColumnProps} from 'ant-design-vue/es/table/interface';
import {useStore} from "vuex";

import {StateType as ServeStateType} from "@/store/serve";
import {StateType as Debug} from "@/views/component/debug/store";
import {listEndpointInterface} from "@/views/endpoint/service";

const store = useStore<{ Endpoint, ProjectGlobal, Debug: Debug, ServeGlobal: ServeStateType }>();

const interfaces = ref<any[]>([])
const filters = ref<any>({})
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);

const props = defineProps({
  serveId: {
    type: Number,
    required: true,
  },
  categoryId: {
    type: Number,
    required: true,
  },
  selectInterfaces: {
    type: Function,
    required: true,
  },
})

const pagination = ref<any>({
  total: 0,
  current: 1,
  pageSize: 10,
  showSizeChanger: true,
  showQuickJumper: true,
})

watch(() => {return props.categoryId}, () => {
  console.log('watch props.categoryId for reload', props.categoryId)
  selectCategory()
}, {deep: false})

const loadList = debounce(async (page, pageSize) => {
  pagination.value.current = page
  pagination.value.pageSize = pageSize
  const data = {
    serveId: props.serveId,
    categoryId: props.categoryId,
    projectId: currProject.value.id,
    keywords: filters.value.keywords,
  }
  const result = await listEndpointInterface(data, { ...pagination.value, page: pagination.value.current })
  console.log('listEndpointInterface', result)
  interfaces.value = result?.list
  pagination.value.total = result?.total
}, 300)

const selectCategory = async () => {
  console.log('selectCategory', props.categoryId)
  pagination.value.current = 1;
  await nextTick();
  loadList(pagination.value.current, pagination.value.pageSize)
}

watch(props, () => {
  console.log('watch props for reload', props.serveId, props.categoryId)
  selectCategory()
}, {deep: true, immediate: true})

const onKeywordsChanged = debounce(async () => {
  console.log('onKeywordsChanged')
  await selectCategory()
}, 600)

type Key = ColumnProps['key'];
const selectedRowKeys = ref<Key[]>([]);

const onSelectChange = (keys: Key[], rows: any) => {
  selectedRowKeys.value = [...keys];
  props.selectInterfaces(selectedRowKeys.value)
};

const columns = [
  {
    title: '序号',
    dataIndex: 'index',
    width: 80,
    customRender: ({
                     text,
                     index
                   }: { text: any; index: number }) => (pagination.value.current - 1) * pagination.value.pageSize + index + 1,
  },
  {
    title: '端点名称',
    dataIndex: 'name',
    slots: {customRender: 'colName'},
  },
  {
    title: '请求方法',
    dataIndex: 'method',
    width: 100,
    slots: {customRender: 'colMethod'},
  },
  {
    title: '路径',
    width: 200,
    dataIndex: 'url',
    slots: {customRender: 'colUrl'},
  }
];

onMounted(async () => {
  await loadList(pagination.value.current, pagination.value.pageSize)
})
</script>

<style scoped lang="less">
.interface-list-main {
  margin: 0;
  padding-top: 20px;

  :deep(.ant-table) {
    .ant-table-body {
      overflow-x: hidden;
    }
  }

  :deep(.ant-tag) {
    max-width: 100%;
    text-overflow: ellipsis;
    overflow: hidden;
    cursor: pointer;
  }

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
