<template>
    <a-card
        class="card-baseInfo"
        :bordered="false"
        title="基本信息">
      <a-descriptions :title="null">
        <a-descriptions-item label="创建人">{{ interfaceDetail?.createUser }}</a-descriptions-item>
        <a-descriptions-item label="状态">
          <a-select
              :value="interfaceDetail?.status"
              style="width: 100px"
              :size="'small'"
              placeholder="请修改接口状态"
              :options="interfaceStatusOpts"
              @change="handleChangeStatus"
          >
          </a-select>
        </a-descriptions-item>
        <a-descriptions-item label="描述">{{ interfaceDetail?.description || '暂无' }}</a-descriptions-item>
        <a-descriptions-item label="创建时间">{{ interfaceDetail?.createdAt }}</a-descriptions-item>
        <a-descriptions-item label="最近更新">{{ interfaceDetail?.updatedAt }}</a-descriptions-item>
      </a-descriptions>
    </a-card>
</template>
<script lang="ts" setup>

import {
  defineProps,
  defineEmits, computed,
} from 'vue';
import {interfaceStatusOpts} from '@/config/constant';
import {useStore} from "vuex";
import {Interface} from "@/views/interface/data";
const props = defineProps({
})

const store = useStore<{ Interface }>();
const interfaceDetail = computed<Interface>(() => store.state.Interface.interfaceDetail);

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
