<template>
  <div class="endpoint-header">
    <div class="endpoint-method">
      <IconSvg :type="DESIGN_TYPE_ICON_MAP[endpointData.processorType]" class="processor-icon-svg"/>
      <span :style="{ color: getMethodColor(reqContent.method) }">
              {{ reqContent.method }}
      </span>
    </div>

    <div class="summary">
      <a-tooltip :title="reqContent.url">
        <a class="endpoint-url" href="javascript:void (0)">
          {{ reqContent.url }}
        </a>
      </a-tooltip>

      <a-tooltip :title="`${endpointData.name}`">
        <span class="endpoint-name">
          {{ endpointData.name || '' }}
        </span>
      </a-tooltip>

      <span class="endpoint-code" v-if="endpointData.resultStatus !== 'loading'">
          状态码: &nbsp; <span :style="{ color: `${responseCodeColorMap[resContent.statusCode]}` }">{{
          resContent.statusCode
        }}</span>
        </span>

      <span :class="['endpoint-time', ClassMap[endpointData.resultStatus]]"
            v-if="endpointData.resultStatus !== 'loading'">
          耗时:
          <a-tooltip :title="`${resContent.time} ms`">&nbsp;<span
              v-html="formatWithSeconds(resContent.time)"></span>
          </a-tooltip>
        </span>


    </div>

    <div class="status endpoint-status">
      <span :class="[ ClassMap[endpointData.resultStatus]]"
            v-if="endpointData.resultStatus !== 'loading'">{{ StatusMap[endpointData.resultStatus] }}</span>
      <span v-else><a-spin :indicator="indicator"/></span>
    </div>

    <div class="endpoint-expand-btn" @click.stop="handleQueryDetail">
      详情
      <RightOutlined/>
    </div>

    <ResponseDrawer
      :data="currRespDetail"
      :response-drawer-visible="logResponseDetailVisible"
      @onClose="logResponseDetailVisible = false" />
  </div>
</template>
<script setup lang="ts">
import {defineProps, h, defineEmits, computed, ref, reactive} from 'vue';
import {RightOutlined, LoadingOutlined, ExclamationCircleOutlined, CheckCircleOutlined} from '@ant-design/icons-vue';
import {responseCodes} from '@/config/constant';
import IconSvg from "@/components/IconSvg";
import {formatWithSeconds} from '@/utils/datetime';
import {
  scenarioTypeMapToText,
  showArrowScenarioType,
  DESIGN_TYPE_ICON_MAP,
  showScenarioExecStatus,
} from "@/views/scenario/components/Design/config";
import ResponseDrawer from '@/views/component/Report/Response/index.vue';
import { getMethodColor } from '@/utils/interface';


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
const logResponseDetailVisible = ref(false);
const currRespDetail = reactive({ reqContent: {}, resContent: {} });

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
  Object.assign(currRespDetail, {
    reqContent: reqContent.value,
    resContent: resContent.value,
  });
  logResponseDetailVisible.value = true;
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
  justify-content: space-between;
  overflow: hidden;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  text-align: right;

  .endpoint-status {
    width: 60px;
    height: 22px;
    font-size: 14px;
    border-radius: 2px;
    text-align: left;
    line-height: 22px;
    margin-right: 16px;

    .endpoint-success {
      background: #E6FFF4;
      color: #04C495;

    }

    .endpoint-error {
      background: #FFF2F0;;
      color: #F63838;
    }

    .endpoint-expires {
      background: #FFF2F0;;
      color: #F63838;
    }
  }

  .endpoint-method {
    font-weight: bold;
    font-size: 14px;
    line-height: 22px;
    text-align: left;
    margin-right: 20px;
  }

  .endpoint-url {
    max-width: 280px;
    margin-right: 20px;
    display: inline-block;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    text-align: left;
  }

  .endpoint-url,
  .endpoint-name {
    line-height: 22px;
    margin-right: 16px;
  }

  .endpoint-name {
    margin-right: 16px;
    max-width: 150px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .endpoint-code {
    margin-right: 12px;
  }

  .endpoint-time,
  .endpoint-code {
    max-width: 80px;
    text-align: center;
    font-size: 14px;
    line-height: 22px;
    margin-right: 16px;

    //margin-right: 29px;
    color: rgba(0, 0, 0, 0.85);

    span {
      color: #04C495;
    }
  }

  .endpoint-time {
    max-width: 120px;
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
    //display: flex;
    //align-items: center;
    //justify-content: flex-end;
  }
}

.enpoint-expand {
}

.status {
  width: 60px;
}

.processor-icon-svg {
  display: inline-block;
  margin-right: 4px;
}

.summary {
  flex: 1;
  margin-right: 16px;
  display: flex;
  align-items: center;
}
</style>
