<template>
  <div class="scenario-basicinfo">
    <div class="scenario-name">{{ record.name }}</div>
    <div class="scenario-priority">{{ record.priority || 'P1' }}</div>
    <div :class="['scenario-status', record.resultStatus]">{{ statusMap.get(record.resultStatus) }}
    </div>
    <div class="scenario-rate">
      <a-progress :percent="100" :success-percent="progressValue" :showInfo="true" status="exception"/>
      <div style="margin-left: 24px;">通过率 {{ progressValue }}</div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import {defineProps, ref, computed, watch} from 'vue';
import {getPercent} from '@/utils/number';

const props = defineProps(['record', 'showScenarioInfo', 'expandActive']);
const statusMap = new Map([['pass', '通过'], ['fail', '失败']]);

const logInfo = computed(() => {
  return props.record?.logs?.[0] || {};
})
const progressValue = computed(() => getPercent(logInfo.value?.passAssertionNum || 0, logInfo?.value.totalAssertionNum || 0));


</script>

<style scoped lang="less">

.scenario-basicinfo {
  display: flex;
  align-items: center;
  justify-content: space-between;

  div {
    overflow-wrap: break-word;
  }
}

.scenario-priority {
  font-weight: bold;
}

.scenario-status {
  display: flex;
  align-items: center;

  &:before {
    content: '';
    display: block;
    width: 6px;
    height: 6px;
    border-radius: 50%;
    background-color: #04C495;
    margin-right: 10px;
  }

  &.pass {
    &:before {
      background-color: #04C495;
    }
  }

  &.fail {
    &:before {
      background-color: #FF6963;
    }
  }
}

.scenario-name {
  width: 333px;
  font-weight: bold;
}

.scenario-rate {
  // width: 292px;
  width: 240px;
  margin-right: 24px;
  display: flex;
  align-items: center;
  font-family: 'PingFang SC';
  font-style: normal;
  font-weight: 400;
  font-size: 14px;
  color: rgba(0, 0, 0, 0.85);
  padding: 0;

}

.scenario-action {
  width: 54px;
}
</style>

