<template>
  <div class="progress-container">
    <a-progress :status="execStatusMap.get(execStatus)"
                :strokeColor="execStatusColorMap.get(execStatus)"
                :percent="percent || 10"
                :showInfo="false"/>
    <div>
      <a-button v-if="execStatus === WsMsgCategory.InProgress"
                class="exec-cancel"
                type="default"
                @click="handleExecCancel">{{ execStatusTextMap.get(execStatus) }}
      </a-button>
      <div v-else
           class="exec-status"
           :style="{ color: `${execStatusColorMap.get(execStatus)}` }">
        {{ execStatusTextMap.get(execStatus) }}
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import {WsMsgCategory} from '@/utils/enum';
import {defineProps, defineEmits} from 'vue';

defineProps<{
  percent: number
  execStatus: string
}>();

const emits = defineEmits(['execCancel']);
const execStatusMap = new Map([['in_progress', 'active'], ['end', 'success'], ['failed', 'exception'], ['cancel', 'exception']]);
const execStatusTextMap = new Map([['in_progress', '停止执行'], ['end', '执行完成'], ['failed', '执行失败'], ['cancel', '已停止']]);
const execStatusColorMap = new Map([['in_progress', '#1890ff'], ['end', '#04C495'], ['failed', '#F63838'], ['cancel', '#F63838']]);
const handleExecCancel = () => {
  emits('execCancel');
};
</script>

<style scoped lang="less">
.progress-container {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;

  :deep(.ant-progress-bg) {
    height: 22px !important;
  }

  .exec-cancel {
    margin-left: 40px;
  }

  .exec-status {
    width: 84px;
    height: 32px;
    line-height: 32px;
    text-align: center;
    margin-left: 40px;
    display: block;
  }
}
</style>
