<template>
  <div class="progress-container">
    <a-progress :status="execStatusMap.get(execStatus)"
                :strokeColor="execStatusColorMap.get(execStatus)"
                :percent="percentVal"
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
import {defineProps, defineEmits, computed, onUnmounted, ref, watch, onMounted} from 'vue';


const props = defineProps(['percent', 'execStatus']);

const emits = defineEmits(['execCancel']);
const execStatusMap = new Map([['in_progress', 'active'], ['end', 'success'], ['failed', 'exception'], ['cancel', 'exception']]);
const execStatusTextMap = new Map([['in_progress', '停止执行'], ['end', '执行完成'], ['failed', '执行失败'], ['cancel', '已停止']]);
const execStatusColorMap = new Map([['in_progress', '#1890ff'], ['end', '#04C495'], ['failed', '#F63838'], ['cancel', '#F63838']]);
const handleExecCancel = () => {
  emits('execCancel');
};

/*
* TODO 处理进度条效果，未来需要实时计算进度条的值
* */
let timeId: any = ref(null);
const percentVal: any = ref(10);
onMounted(() => {
  if (props.execStatus === "end") {
    percentVal.value = 100;
  }
  timeId.value = window?.setInterval(() => {
    if (percentVal.value >= 80) {
      timeId.value && window.clearInterval(timeId.value);
      return;
    }
    percentVal.value += 1;
  }, 1000);
});

watch(() => {
  return props.execStatus
}, (newVal) => {
  if (newVal === "end") {
    percentVal.value = 100;
    timeId.value && window.clearInterval(timeId.value);
  }
})
onUnmounted(() => {
  clearInterval(timeId.value);
});
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
