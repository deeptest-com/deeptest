<template>
  <div class="endpoint-expand">
    <!-- 请求体检查信息 -->
    <template v-if="invokeDetail.responseDefine">
      <div
        class="endpoint-expand-content"
      >
      <span :style="{color: resultMap[invokeDetail.responseDefine.resultStatus] }">
        <template v-if="invokeDetail.responseDefine.resultStatus === 'fail'">
          <exclamation-circle-outlined /> &nbsp;
        </template>
        <template v-else> <check-circle-outlined /> &nbsp; </template>
        {{ invokeDetail.responseDefine.resultMsg || '' }}
      </span>
      </div>
    </template>
    <!-- 请求断言检查信息 -->
    <template v-if="invokeDetail.checkpoint && invokeDetail.checkpoint.length">
      <div class="endpoint-expand-content">断言结果</div>
      <div
        class="endpoint-expand-content"
        v-for="(item, index) in invokeDetail.checkpoint"
        :key="index"
      >
        <span :style="{color: resultMap[item.resultStatus] }">
          <template v-if="item.resultStatus === 'fail'">
            <exclamation-circle-outlined /> &nbsp;
          </template>
          <template v-else> <check-circle-outlined /> &nbsp; </template>
          {{ item.resultMsg || "" }}
        </span>
      </div>
    </template>
    <!-- 请求未通过其他异常信息 -->
    <template v-if="invokeDetail.result">
      <div class="endpoint-expand-content">{{ invokeDetail.result || '' }}</div>
    </template>
  </div>
</template>
<script setup lang="ts">
import { defineProps, computed, watch } from "vue";
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

const invokeDetail = computed(() => {
  return props.endpointData.respContent
    ? JSON.parse(props.endpointData.detail)
    : { };
});

const resultMap = {
  fail: '#f5222d',
  pass: '#14945a'
}

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
      color: #14945a;
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
      color: #14945a;
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
