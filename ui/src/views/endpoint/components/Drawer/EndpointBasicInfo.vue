<template>
    <a-card
        class="card-baseInfo"
        :bordered="false"
        title="基本信息">
      <a-descriptions :title="null">
        <a-descriptions-item label="创建人">{{ endpointDetail?.createUser }}</a-descriptions-item>
        <a-descriptions-item label="状态">
          <a-select
              :value="endpointDetail?.status"
              style="width: 100px"
              :size="'small'"
              placeholder="请修改接口状态"
              :options="endpointStatusOpts"
              @change="handleChangeStatus"
          >
          </a-select>
        </a-descriptions-item>
        <a-descriptions-item label="描述">{{ endpointDetail?.description || '暂无' }}</a-descriptions-item>
        <a-descriptions-item label="创建时间">{{ endpointDetail?.createdAt }}</a-descriptions-item>
        <a-descriptions-item label="最近更新">{{ endpointDetail?.updatedAt }}</a-descriptions-item>
      </a-descriptions>
    </a-card>
</template>
<script lang="ts" setup>

import {
  defineProps,
  defineEmits, computed,
} from 'vue';
import {endpointStatusOpts} from '@/config/constant';
import {useStore} from "vuex";
import {Endpoint} from "@/views/endpoint/data";
const props = defineProps({
})

const store = useStore<{ Endpoint }>();
const endpointDetail = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);

const emit = defineEmits(['changeStatus']);
function handleChangeStatus(val) {
  emit('changeStatus',val);
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
