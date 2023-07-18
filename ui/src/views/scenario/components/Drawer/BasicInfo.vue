<template>
  <a-descriptions :size="'small'" :title="null" :column="4">
    <a-descriptions-item label="创建人">{{ detailResult?.createUserName }}</a-descriptions-item>
    <a-descriptions-item label="状态">
      <EditAndShowSelect :label="scenarioStatus.get(detailResult?.status) || '未设置'"
                         :value="detailResult?.status || null"
                         :options="scenarioStatusOptions"
                         @update="(val) => {
                            handleChange('status',val)
                           }"/>
    </a-descriptions-item>
    <a-descriptions-item label="优先级">
      <EditAndShowSelect :label="scenarioPriority.get(detailResult?.priority) || '未设置'"
                         :value="detailResult?.priority || null"
                         :options="priorityOptions"
                         @update="(val) => {
                            handleChange('priority',val)
                           }"/>
    </a-descriptions-item>
    <a-descriptions-item label="描述">
      <EditAndShowField :placeholder="'请输入描述'" :value="detailResult?.desc || ''"
                        @update="(val) => {
                            handleChange('desc',val)
                           }"/>
    </a-descriptions-item>
    <a-descriptions-item label="分类">
      <EditAndShowTreeSelect
          :label="categoryLabel"
          :value="detailResult?.categoryId || -1"
          :treeData="treeData"
          @update="(val) => {
               handleChange('categoryId',val)
            }"/>
    </a-descriptions-item>
    <a-descriptions-item label="测试类型">
      <EditAndShowSelect :label="testTypeMap.get(detailResult?.type) || '未设置'"
                         :value="detailResult?.type || null"
                         :options="testTypeOptions"
                         @update="(val) => {
                            handleChange('type',val)
                           }"/>
    </a-descriptions-item>
    <a-descriptions-item label="创建时间">{{ momentUtc(detailResult?.createdAt) }}</a-descriptions-item>
    <a-descriptions-item label="最近更新">{{ momentUtc(detailResult?.updatedAt) }}</a-descriptions-item>
  </a-descriptions>
</template>
<script lang="ts" setup>

import {
  defineProps,
  ref,
  defineEmits, computed,
} from 'vue';
import {
  endpointStatusOpts,
  scenarioStatus,
  scenarioPriority,
  scenarioStatusOptions,
  priorityOptions,
  testTypeMap,
  testTypeOptions
} from '@/config/constant';
import {useStore} from "vuex";
import {Scenario} from "@/views/Scenario/data";
import EditAndShowField from '@/components/EditAndShow/index.vue';
import EditAndShowSelect from '@/components/EditAndShowSelect/index.vue';
import EditAndShowTreeSelect from '@/components/EditAndShowTreeSelect/index.vue';
import ConBoxTitle from '@/components/ConBoxTitle/index.vue';
import {momentUtc} from '@/utils/datetime';

const props = defineProps({})

const store = useStore<{ Scenario }>();
const detailResult: any = computed<Scenario>(() => store.state.Scenario.detailResult);
const treeDataCategory = computed<any>(() => store.state.Scenario.treeDataCategory);
const treeData: any = computed(() => {
  return treeDataCategory.value?.[0]?.children || [{
    label:'未分类',
    value:-1,
  }];
});
const categoryLabel = computed(() => {
  if (!detailResult.value?.categoryId) {
    return '未分类'
  }
  const data = treeDataCategory.value?.[0]?.children || [];
  let label = "";
  let hasFind = false;
  // 递归查找目录树的文案
  function fn(arr: any) {
    if (!Array.isArray(arr)) {
      return;
    }
    for (let i = 0; i < arr.length; i++) {
      const item = arr[i];
      if (item.id === detailResult.value?.categoryId) {
        label = item.name;
        hasFind = true;
      }
      if (Array.isArray(item.children) && !hasFind) {
        fn(item.children)
      }
    }
  }
  fn(data);
  return label;
});

const emit = defineEmits(['change']);

function handleChange(type, val) {
  emit('change', type, val);
}



</script>

<style lang="less" scoped>
.card-baseInfo {
  width: 100%;

  :deep(.ant-card-body) {
    padding: 12px 24px;
  }
}

</style>
