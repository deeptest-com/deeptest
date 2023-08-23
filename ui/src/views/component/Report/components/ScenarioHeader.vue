<template>
  <div class="scenario-basicinfo">
    <div class="scenario-name">{{ record.name }}</div>
    <div class="scenario-priority">{{ record.priority}}</div>
    <div :class="['scenario-status', logInfo.resultStatus]">{{ statusMap.get(status) }}
    </div>
    <div class="scenario-rate">
      <div class="scenario-rate-progress">
        <a-progress :percent="progressInfo.progressValue" :status="progressInfo.status" :show-info="false"/>
      </div>
      <div class="scenario-rate-info">通过率 {{ `${progressInfo.progressValue}%` }}</div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import {defineProps, ref, computed, watch} from 'vue';
import {getPercent, getPercentStr, num2Percent} from '@/utils/number';

const props = defineProps(['record', 'showScenarioInfo', 'expandActive']);
const statusMap = new Map([['pass', '通过'], ['fail', '失败'],['exception','失败'], ['in-progress', '进行中']]);

const logInfo = computed(() => {
  return props.record?.logs?.[0] || {};
})

const status = computed(() => {
  return props.record?.resultStatus || 'in-progress'
})

const progressInfo = ref({
  status: 'active',
  progressValue: 20
})

watch(() => props.record?.resultStatus, (newVal) => {
  console.log(8322222, newVal)
  if (newVal) {
    if (newVal === 'fail') {
      progressInfo.value = {
        status: 'exception',
        progressValue: 50
      }
    }
    if (newVal === 'pass') {
      progressInfo.value = {
        status: 'success',
        progressValue: 100
      }
    }
  }
}, {
  immediate: false
});

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

  &.in-progress {
    &:before {
      background-color: #FFC107;
    }
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
  &.exception {
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

.scenario-rate-progress {
  width: 170px;
}

.scenario-rate-info {
  margin-left: 40px;
  font-size: 12px;
  color: rgba(0, 0, 0, 0.45);
}
</style>

