<template>
  <div class="table-filter">
    <div class="left" v-if="linked">
      <a-button type="primary" @click="visible = true">
        <template #icon>
          <plus-outlined/>
        </template>
        关联测试计划
      </a-button>
      <a-button type="default" @click="handleRemove" :disabled="!selectedRowKeys?.length">批量移除</a-button>
    </div>
    <div class="right">
      <a-form-item label="状态">
        <a-select allowClear ref="select"
                  v-model:value="formState.status"
                  style="width: 140px"
                  :options="scenarioStatusOptions"
                  @change="handleChange" placeholder="请选择状态"></a-select>
      </a-form-item>
      <a-form-item label="负责人">
        <a-select allowClear ref="select" v-model:value="formState.adminId" style="width: 140px" :options="members"
                  @change="handleChange" placeholder="请选择创建人"></a-select>
      </a-form-item>
      <a-form-item>
        <a-input-search allowClear v-model:value="formState.keywords" placeholder="请输入测试计划名称" @search="handleChange"
                        style="width: 220px"/>
      </a-form-item>
    </div>
  </div>
  <a-table
      :row-selection="{
            selectedRowKeys: selectedRowKeys,
            onChange: onSelectChange
        }"
      :pagination="{
            ...pagination,
            showSizeChanger: false,
            onChange: (page) => {
              pagination.current = page;
              getPlans({ page });
            },
        }"
      row-key="id"
      :loading="loading"
      :columns="linked ? columns : columnsForSelect"
      :data-source="linked ? linkedPlans : notLinkedPlans">
    <template #status="{ record }">
      {{ scenarioStatus.get(record.status) || '未知' }}
    </template>
    <template #updatedAt="{ record }">
      <span>{{ momentUtc(record.updatedAt) }}</span>
    </template>
    <template #testPassRate="{ record }">
      <span>{{ record.testPassRate }}</span>
    </template>
    <template #operation="{ record }">
      <a-button type="primary" @click="handleRemove(record)">
        移除
      </a-button>
    </template>
  </a-table>

  <!-- 使用v-if 来保证组件重新渲染 -->
  <div v-if="visible">
    <SelectPlan
        :visible="visible"
        @cancel="visible = false"
        @ok="handleSelect"/>
  </div>

</template>
<script lang="ts" setup>
import {ref, reactive, defineProps, defineEmits, PropType, computed, onMounted, watch} from 'vue';
import {useStore} from 'vuex';
import {PlusOutlined} from '@ant-design/icons-vue';
import SelectPlan from './SelectPlan.vue';
import {Modal} from 'ant-design-vue';
import {momentUtc} from "@/utils/datetime";
import {Scenario} from "@/views/scenario/data";

import {priorityOptions, scenarioStatusOptions, scenarioStatus} from "@/config/constant";

const props = defineProps({
  linked: {
    type: Boolean,
    required: true
  }
})

const emits = defineEmits(['select']);
const store = useStore<{Report,Scenario, ProjectGlobal, ServeGlobal }>();
const detailResult: any = computed<Scenario>(() => store.state.Scenario.detailResult);
const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const linkedPlans = computed(() => store.state.Scenario.linkedPlans);
let pagination = computed(() => store.state.Scenario.linkedPlansPagination);
const notLinkedPlans = computed(() => store.state.Scenario.notLinkedPlans);
const members = computed(() => store.state.Report.members);
const visible = ref(false);
const selectedRowKeys = ref<any[]>([]); // Check here to configure the default column

const onSelectChange = (keys: string[], rows: any) => {
  selectedRowKeys.value = keys;
  emits('select', keys);
};
const formState = reactive({status: null, adminId: null, keywords: ''});
const columns = [
  {
    title: '编号',
    dataIndex: 'serialNumber',
  },
  {
    title: '名称',
    dataIndex: 'name',
    slots: {customRender: 'name'}
  },
  {
    title: '状态',
    dataIndex: 'status',
    slots: {customRender: 'status'}
  },
  {
    title: '测试通过率',
    dataIndex: 'testPassRate',
    slots: {customRender: 'testPassRate'}
  },
  {
    title: '负责人',
    dataIndex: 'adminName',
  },
  {
    title: '最近更新',
    dataIndex: 'updatedAt',
    slots: {customRender: 'updatedAt'}
  },
  {
    title: '操作',
    dataIndex: 'action',
    slots: {customRender: 'operation'}
  },
];
const columnsForSelect = [
  {
    title: '编号',
    dataIndex: 'serialNumber',
  },
  {
    title: '名称',
    dataIndex: 'name',
    slots: {customRender: 'name'}
  },
  {
    title: '状态',
    dataIndex: 'status',
    slots: {customRender: 'status'}
  },
  {
    title: '测试通过率',
    dataIndex: 'testPassRate',
    slots: {customRender: 'testPassRate'}
  },
  {
    title: '负责人',
    dataIndex: 'adminName',
  },
  {
    title: '最近更新',
    dataIndex: 'updatedAt',
    slots: {customRender: 'updatedAt'}
  },
];
const loading = ref(false);

async function getPlans(opts = {}) {
  loading.value = true;
  await store.dispatch('Scenario/getPlans', {
    currProjectId: currProject.value.id,
    id: detailResult.value.id,
    data: {
      "ref": props.linked,
      ...formState,
      ...pagination.value,
      page: pagination.value.current,
      ...opts,
    }
  });
  loading.value = false;
}

watch(() => {
  return detailResult.value.id
}, async (newVal) => {
  if(newVal) {
    await getPlans({ page: 1 });
  }
}, {
  immediate: true
})

const getMember = async (): Promise<void> => {
  await store.dispatch('Report/getMembers', currProject.value.id)
}
watch(() => {
    return currProject.value;
}, (val: any) => {
    if (val.id) {
        getMember();
    }
}, {
    immediate: true
})

// onMounted(async () => {
//   await getPlans();
// })

const handleChange = () => {
  getPlans({ page: 1 });
};

async function handleAdd(keys: any[]) {
  await store.dispatch('Scenario/addPlans', {
    currProjectId: currProject.value.id,
    id: detailResult.value.id,
    data: keys,
  });
  await getPlans();
}

async function handleRemove(record?: any) {
  Modal.confirm({
    title: '确认要解除与该测试计划的关联吗?',
    onOk: async () => {
      let keys: any[] = [];
      if (record && record.id) {
        keys.push(record.id);
      } else {
        keys = selectedRowKeys.value;
      }
      await store.dispatch('Scenario/removePlans', {
        currProjectId: currProject.value.id,
        id: detailResult.value.id,
        data: keys
      });
      await getPlans(!record 
        ? { page: 1 } 
        : { 
          page: linkedPlans.value.length === 1 && pagination.value.current > 1 
          ? pagination.value.current - 1 
          : pagination.value.current }
        );
    }
  })
}

async function handleSelect(keys) {
  visible.value = false;
  await handleAdd(keys);
}
</script>
<style scoped lang="less">
.table-filter {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-bottom: 20px;

  .left, .right {
    display: flex;
    align-items: center;

    :deep(.ant-row.ant-form-item), :deep(.ant-btn) {
      margin-right: 20px;
      margin-bottom: 0;

      &:last-child {
        margin: 0;
      }
    }
  }
}
</style>

