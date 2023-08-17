<template>
  <div class="endpoint-expand">
    <div class="endpoint-expand-content">
      <span v-if="endpointData.resultStatus === 'fail'">
        <exclamation-circle-outlined style="color: #f5222d" /> &nbsp;
        {{ resContent.statusContent || "" }}</span
      >
      <span v-if="endpointData.resultStatus === 'pass'">
        <check-circle-outlined style="color: #04c495" /> &nbsp;
        返回数据结构校验通过
      </span>
    </div>
  </div>
</template>
<script setup lang="ts">
import { defineProps, computed } from "vue";
import {
  ExclamationCircleOutlined,
  CheckCircleOutlined,
} from "@ant-design/icons-vue";

const props = defineProps({
  collapseKey: {
    required: false,
  },
  endpointData: {
    type: Object,
    required: true,
  },
});

const resContent = computed(() =>
  props.endpointData.respContent
    ? JSON.parse(props.endpointData.respContent)
    : {}
);
</script>

<style scoped lang="less">
.endpoint-collapse-item {
  :deep(.ant-collapse-content.ant-collapse-content-active) {
    background-color: #fbfbfb;

    .ant-collapse-content-box {
      padding: 9px 16px;
    }
  }
}

.endpoint-expand {
  .endpoint-expand-content {
    background-color: #fff;
    padding: 8px 16px;
  }

  .endpoint-expand-btn {
    margin-top: 8px;
    text-align: center;
    cursor: pointer;
    line-height: 22px;
    font-size: 14px;
  }
}

.endpoint-header {
  display: flex;
  align-items: center;
  padding-right: 10px;

  .endpoint-status {
    width: 36px;
    height: 22px;
    font-size: 12px;
    border-radius: 2px;
    text-align: center;
    line-height: 22px;
    margin-right: 16px;

    &.endpoint-success {
      background: #e6fff4;
      color: #04c495;
    }

    &.endpoint-error {
      background: #fff2f0;
      color: #f63838;
    }

    &.endpoint-expires {
      background: #fff2f0;
      color: #f63838;
    }
  }

  .endpoint-method {
    font-weight: bold;
    font-size: 14px;
    width: 54px;
    line-height: 22px;
    text-align: center;
    margin-right: 36px;

    &.endpoint-success {
      color: #04c495;
    }

    &.endpoint-error,
    &.endpoint-expires {
      color: #f63838;
    }
  }

  .endpoint-url,
  .endpoint-data {
    flex: 1;
    line-height: 22px;
  }

  .endpoint-data {
    flex-grow: 2;
  }

  .endpoint-time,
  .endpoint-code {
    width: 84px;
    text-align: center;
    font-size: 14px;
    line-height: 22px;
    margin-right: 29px;
    color: rgba(0, 0, 0, 0.85);

    span {
      color: #04c495;
    }
  }

  .endpoint-time {
    width: 150px;
    text-align: right;
  }

  .endpoint-type {
    font-size: 14px;
    line-height: 22px;
    color: #447dfd;
    cursor: pointer;
  }

  .endpoint-response {
    width: 272px;
    display: flex;
    align-items: center;
  }
}

.enpoint-expand {
}
</style>
