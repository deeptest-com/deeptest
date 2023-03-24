<template>
    <a-card
        class="card-baseInfo"
        :bordered="false"
        title="基本信息">
      <a-descriptions :title="null">
        <a-descriptions-item label="创建人">{{ interfaceDetail?.createUser }}</a-descriptions-item>
        <a-descriptions-item label="状态">
          <a-tag :color="interfaceStatusColor.get(interfaceDetail?.status || 0)">{{ interfaceStatus.get(interfaceDetail?.status || 0) }}</a-tag>
        </a-descriptions-item>
        <a-descriptions-item label="描述">{{ interfaceDetail?.description || '暂无' }}</a-descriptions-item>
        <a-descriptions-item label="服务版本">{{ 'V1.0' }}</a-descriptions-item>
        <a-descriptions-item label="创建时间">{{ interfaceDetail?.createdAt }}</a-descriptions-item>
        <a-descriptions-item label="最近更新">{{ interfaceDetail?.updatedAt }}</a-descriptions-item>
      </a-descriptions>
    </a-card>
</template>
<script lang="ts" setup>

import {
  defineProps,
  defineEmits,
} from 'vue';
import {interfaceStatus, interfaceStatusColor} from '@/config/constant';

const props = defineProps({
  interfaceDetail: {
    required: true,
    type: Object,
  }
})

const emit = defineEmits([]);

</script>

<style lang="less" scoped>
.card-baseInfo {
  width: 100%;

  :deep(.ant-card-body) {
    padding: 12px 24px;
  }
}

</style>
