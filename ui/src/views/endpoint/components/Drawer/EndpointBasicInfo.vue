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
      <a-descriptions-item label="创建人">{{ endpointDetail?.createUser }}</a-descriptions-item>
      <a-descriptions-item label="状态">
        <EditAndShowSelect :label="endpointStatus.get(endpointDetail?.status || 0 )"
                           :value="endpointDetail?.status"
                           :options="endpointStatusOpts"
                           @update="handleChangeStatus"/>
      </a-descriptions-item>
      <a-descriptions-item label="描述">
        <EditAndShowField :placeholder="'请输入描述'" :value="endpointDetail?.description || '暂无'"
                          @update="updateDescription"/>
      </a-descriptions-item>
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
import {endpointStatusOpts, endpointStatus} from '@/config/constant';
import {useStore} from "vuex";
import {Endpoint} from "@/views/endpoint/data";
import EditAndShowField from '@/components/EditAndShow/index.vue';
import EditAndShowSelect from '@/components/EditAndShowSelect/index.vue';
import ConBoxTitle from '@/components/ConBoxTitle/index.vue';

const props = defineProps({})

const store = useStore<{ Endpoint }>();
const endpointDetail = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);

const emit = defineEmits(['changeStatus', 'changeDescription']);

function handleChangeStatus(val) {
  emit('changeStatus', val);
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
