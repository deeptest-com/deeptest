<template>
  <div class="endpoint-expand">
    <div class="endpoint-expand-content">
                <span v-if="endpointData.resultStatus === 'fail'">
                    <exclamation-circle-outlined style="color: #f5222d" /> &nbsp;
                    {{ resContent.statusContent || '' }}</span>
      <span v-if="endpointData.resultStatus === 'pass'">
                    <check-circle-outlined style="color: #04C495" /> &nbsp;
                    返回数据结构校验通过
                </span>
    </div>
    <div class="endpoint-expand-btn" @click="handleQueryDetail">
      更多详情 &nbsp;&nbsp;
      <RightOutlined />
    </div>
    <ResponseDrawer
        :data="currRespDetail"
        :response-drawer-visible="logResponseDetailVisible"
        @onClose="logResponseDetailVisible = false" />
  </div>
</template>
<script setup lang="ts">
import { defineProps, h, defineEmits, computed, toRefs,ref } from 'vue';
import ResponseDrawer from '@/views/component/Report/Response/index.vue';
import { RightOutlined, LoadingOutlined, ExclamationCircleOutlined, CheckCircleOutlined } from '@ant-design/icons-vue';
import { responseCodes } from '@/config/constant';
import { formatWithSeconds } from '@/utils/datetime';

const props = defineProps({
    collapseKey: {
        required: false
    },
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

const currRespDetail = ref({ resContent: {}, reqContent: {} });
const logResponseDetailVisible = ref(false);

function handleQueryDetail() {
  currRespDetail.value = { resContent: resContent.value, reqContent: reqContent.value  };
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
            background: #E6FFF4;
            color: #04C495;

        }

        &.endpoint-error {
            background: #FFF2F0;
            ;
            color: #F63838;
        }

        &.endpoint-expires {
            background: #FFF2F0;
            ;
            color: #F63838;
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
            color: #04C495;
        }

        &.endpoint-error,
        &.endpoint-expires {
            color: #F63838;
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
            color: #04C495;
        }
    }

    .endpoint-time {
        width: 150px;
      text-align: right;
    }

    .endpoint-type {
        font-size: 14px;
        line-height: 22px;
        color: #447DFD;
        cursor: pointer;
    }

    .endpoint-response {
        width: 272px;
        display: flex;
        align-items: center;
    }
}

.enpoint-expand {}
</style>
