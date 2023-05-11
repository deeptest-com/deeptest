<template>
  <a-card
      class="card-baseInfo"
      :bordered="false"
      :headStyle="{}"
      :bodyStyle="{padding:'12px 24px 12px 24px'}"
      :title="null">
    <div style="margin-bottom: 8px;">
      <ConBoxTitle :backgroundStyle="'background: #FBFBFB;'" :title="'基本信息'"/>
    </div>

    <a-descriptions :size="'small'" :title="null">
      <a-descriptions-item label="创建人">{{ detailResult?.createUserName }}</a-descriptions-item>
      <a-descriptions-item label="状态">
        <EditAndShowSelect :label="scenarioStatus.get(detailResult?.status || 0 )"
                           :value="detailResult?.status"
                           :options="endpointStatusOpts"
                           @update="(val) => {
                            handleChange('status',val)
                           }"/>
      </a-descriptions-item>
      <a-descriptions-item label="优先级">
        <EditAndShowSelect :label="scenarioPriority.get(detailResult?.priority || '0' )"
                           :value="detailResult?.status"
                           :options="endpointStatusOpts"
                           @update="(val) => {
                            handleChange('priority',val)
                           }"/>
      </a-descriptions-item>
      <a-descriptions-item label="描述">
        <EditAndShowField :placeholder="'请输入描述'" :value="detailResult?.desc || '暂无'"
                          @update="updateDescription"/>
      </a-descriptions-item>
      <a-descriptions-item label="分类">
        <EditAndShowTreeSelect
            :label="categoryLabel"
            :value="detailResult?.categoryId"
            :treeData="treeData"
            @update="handleChangeCategory"/>
      </a-descriptions-item>
      <a-descriptions-item label="创建时间">{{ momentUtc(detailResult?.createdAt) }}</a-descriptions-item>
      <a-descriptions-item label="最近更新">{{ momentUtc(detailResult?.updatedAt) }}</a-descriptions-item>
    </a-descriptions>
  </a-card>
</template>
<script lang="ts" setup>

import {
  defineProps,
  defineEmits, computed,
} from 'vue';
import {endpointStatusOpts, scenarioStatus,scenarioPriority} from '@/config/constant';
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
  return treeDataCategory.value?.[0]?.children || [];
});
const categoryLabel = computed(() => {
  if(!detailResult.value?.categoryId){
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

const emit = defineEmits(['change', 'changeDescription','changeCategory']);

function handleChange(type,val) {
  emit('change', type,val);
}

function handleChangeCategory(val) {
  emit('changeCategory', val);
}

function updateDescription(val: string) {
  emit('changeDescription', val);
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
