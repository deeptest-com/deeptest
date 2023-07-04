<template>
  <div class="endpoint-header">
    <div :class="['endpoint-status', ClassMap[endpointData.resultStatus]]">
      <span v-if="endpointData.resultStatus !== 'loading'">{{ StatusMap[endpointData.resultStatus] }}</span>
      <span v-else><a-spin :indicator="indicator"/></span>
    </div>
    <div :class="['endpoint-method', ClassMap[endpointData.resultStatus]]">
      {{ reqContent.method }}
    </div>
    <a-tooltip  :title="reqContent.url">
      <div class="endpoint-url">
        {{ reqContent.url }}
      </div>
    </a-tooltip>

    <a-tooltip  :title="`${endpointData.name}`">
      <div class="endpoint-name">
        {{ endpointData.name || '' }}
      </div>
    </a-tooltip>


    <div class="endpoint-response">
      <div class="endpoint-code" v-if="endpointData.resultStatus !== 'loading'">
        状态码: &nbsp; <span :style="{ color: `${responseCodeColorMap[resContent.statusCode]}` }">{{
          resContent.statusCode
        }}</span>
      </div>

      <div :class="['endpoint-time', ClassMap[endpointData.resultStatus]]"
           v-if="endpointData.resultStatus !== 'loading'">
        耗时:
        <a-tooltip  :title="`${resContent.time} ms`">&nbsp;<span
            v-html="formatWithSeconds(resContent.time)"></span>
        </a-tooltip>
      </div>


    </div>

  </div>
</template>
<script setup lang="ts">
import {defineProps, h, defineEmits, computed, toRefs} from 'vue';
import {RightOutlined, LoadingOutlined, ExclamationCircleOutlined, CheckCircleOutlined} from '@ant-design/icons-vue';
import {responseCodes} from '@/config/constant';
import {formatWithSeconds} from '@/utils/datetime';

enum StatusMap {
  'pass' = '通过',
  'expires' = '过期',
  'fail' = '失败'
}

enum ClassMap {
  'pass' = 'endpoint-success',
  'expires' = 'endpoint-expires',
  'fail' = 'endpoint-error',
  'loading' = 'endpoint-loading'
}

const props = defineProps({
  endpointData: {
    type: Object,
    required: true
  }
});

const emits = defineEmits(['queryDetail']);
const reqContent = computed(() => props.endpointData.reqContent ? JSON.parse(props.endpointData.reqContent) : {});
const resContent = computed(() => props.endpointData.respContent ? JSON.parse(props.endpointData.respContent) : {});
const responseCodeColorMap = {};

responseCodes.forEach(e => {
  responseCodeColorMap[e.value] = e.color;
})

const indicator = h(LoadingOutlined, {
  style: {
    fontSize: '16px',
    color: '#b0b0b0'
  },
  spin: true,
});

function handleQueryDetail() {
  emits('queryDetail', {resContent: resContent.value, reqContent: reqContent.value});
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
    padding: 16px;
  }

  .endpoint-expand-btn {
    margin-top: 15px;
    text-align: center;
    cursor: pointer;
    line-height: 22px;
    font-size: 14px;
  }

}

.endpoint-header {
  display: flex;
  align-items: center;
  overflow: hidden;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  text-align: right;

  .endpoint-status {
    width: 36px;
    min-width: 36px;
    height: 22px;
    font-size: 12px;
    border-radius: 2px;
    text-align: center;
    line-height: 22px;
    margin-right: 16px;

    &.endpoint-success {
      background: #E6FFF4;
      color: #04C495;

    }

    &.endpoint-error {
      background: #FFF2F0;;
      color: #F63838;
    }

    &.endpoint-expires {
      background: #FFF2F0;;
      color: #F63838;
    }
  }

  .endpoint-method {
    font-weight: bold;
    font-size: 14px;
    min-width: 54px;
    line-height: 22px;
    text-align: center;
    margin-right: 16px;

    &.endpoint-success {
      color: #04C495;
    }

    &.endpoint-error,
    &.endpoint-expires {
      color: #F63838;
    }
  }

  .endpoint-url {
    width: calc(100% - 380px);
    margin-right: 16px;
    flex:1;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    text-align: left;
  }

  .endpoint-url,
  .endpoint-name {
    line-height: 22px;
  }

  .endpoint-name {
    width: 120px;
    margin-right: 16px;
    min-width: 120px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .endpoint-code {
    margin-right: 12px;
  }

  .endpoint-time,
  .endpoint-code {
    width: 80px;
    text-align: center;
    font-size: 14px;
    line-height: 22px;

    //margin-right: 29px;
    color: rgba(0, 0, 0, 0.85);

    span {
      color: #04C495;
    }
  }

  .endpoint-time {
    width: 120px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    text-align: right;
  }

  .endpoint-type {
    font-size: 14px;
    line-height: 22px;
    color: #447DFD;
    cursor: pointer;
  }

  .endpoint-response {
    display: flex;
    align-items: center;
    justify-content: flex-end;
  }
}

.enpoint-expand {
}
</style>
