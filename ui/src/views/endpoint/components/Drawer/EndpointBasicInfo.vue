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
        <a-descriptions-item label="描述">
          <EditAndShowField :placeholder="'请输入描述'" :value="endpointDetail?.description || '暂无'" @update="updateDescription"/>
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
import {endpointStatusOpts} from '@/config/constant';
import {useStore} from "vuex";
import {Endpoint} from "@/views/endpoint/data";
import EditAndShowField from '@/components/EditAndShow/index.vue';
const props = defineProps({
})

const store = useStore<{ Endpoint }>();
const endpointDetail = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);

const emit = defineEmits(['changeStatus', 'changeDescription']);
function handleChangeStatus(val) {
  emit('changeStatus',val);
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
