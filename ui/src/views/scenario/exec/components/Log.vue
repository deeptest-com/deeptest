<template>
  <div class="scenario-exec-log-main">
    <div v-for="(item, index) in logs" :key="index" class="log">
      <div>
        <div v-if="!item.respContent">
          {{ item.name }}&nbsp; {{ joinArr(item.summary) }}
        </div>

        <a-collapse v-if="item.respContent" expand-icon-position="right">
          <a-collapse-panel key="index" :header="getHeader(item)">
            <template #extra>
              <span :class="getResultCls(item.resultStatus)">{{ t(item.resultStatus) }}</span>
            </template>

            <div class="resp-content">
              <a-row class="url">
                <a-col flex="150px">{{ getReq(item).method }}</a-col>
                <a-col flex="200px">{{ getResp(item).statusCode }}</a-col>

                <a-col flex="1">{{ getReq(item).url }}</a-col>
              </a-row>

              <div class="extractor">
                <div class="title">提取器</div>
                <a-row v-for="(extractor, idx) in item.interfaceExtractorsResult" :key="idx" type="flex" class="item">
                  <a-col flex="50px">{{idx + 1}}</a-col>
                  <a-col flex="100px">{{ t(extractor.src) }}</a-col>
                  <a-col flex="100px">{{ extractor.type ? t('processor_extractor_'+extractor.type) : '' }}</a-col>
                  <a-col flex="100px">
                    <span v-if="item.src === ExtractorSrc.header">
                      {{ extractor.key }}
                    </span>
                    <span v-if="extractor.src === ExtractorSrc.body">
                      {{ item.type === ExtractorType.boundary ?
                                `${extractor.boundaryStart}-${extractor.boundaryEnd}[${extractor.boundaryIndex}] ${extractor.boundaryIncluded}` :
                                extractor.expression }}
                    </span>
                  </a-col>
                  <a-col flex="100px">{{ extractor.variable }}</a-col>
                  <a-col flex="1">{{extractor.result}}</a-col>
                </a-row>
              </div>

              <div class="checkpoint">
                <div class="title">检查点</div>
                <a-row v-for="(checkpoint, idx) in item.interfaceCheckpointsResult" :key="idx" type="flex">
                  <a-col flex="50px">{{idx + 1}}</a-col>
                  <a-col flex="100px">{{t(checkpoint.type)}}</a-col>
                  <a-col flex="100px">{{ checkpoint.type === CheckpointType.extractor ? checkpoint.extractorVariable : checkpoint.expression }} </a-col>
                  <a-col flex="100px">{{ t(checkpoint.operator) }}</a-col>
                  <a-col flex="100px">{{ checkpoint.value }}</a-col>
                  <a-col flex="1">
                    {{ checkpoint.actualResult }}
                  </a-col>
                  <a-col flex="100px">
                    <span :class="getResultCls(checkpoint.resultStatus)">{{ t(checkpoint.resultStatus) }}</span>
                  </a-col>
                </a-row>
              </div>

              <div class="header">
                <div class="title">响应头</div>
                <a-row v-for="(header, idx) in getResp(item).headers" :key="idx" type="flex" class="item">
                  <a-col flex="50px">{{idx + 1}}</a-col>
                  <a-col flex="300px">{{ header.name }}</a-col>
                  <a-col flex="1">{{ header.value }}</a-col>
                </a-row>
              </div>

              <div class="resp">
                <a-link @click="showModal(item)" to="">显示响应内容</a-link>
              </div>
            </div>

          </a-collapse-panel>
        </a-collapse>
      </div>

      <Log v-if="item.logs" :logs="item.logs"></Log>
    </div>

    <a-modal
        v-model:visible="visible"
        title="响应内容"
        width="100%"
        wrapClassName="dp-full-modal"
    >
      <p>{{ respContent }}</p>
      <template #footer>
        <a-button key="back" @click="handleCancel">关闭</a-button>
      </template>
    </a-modal>

  </div>
</template>

<script setup lang="ts">
import {defineProps, ref} from "vue";
import {Collapse, CollapsePanel} from 'ant-design-vue';
import ALink from "@/components/ALink/index.vue";
import {ExtractorSrc, ExtractorType, ComparisonOperator, CheckpointType} from "@/utils/enum";
import Log from "./Log.vue"
import {getResultCls} from "@/utils/dom"
import {useI18n} from "vue-i18n";

defineProps<{
  logs: []
}>()

const { t } = useI18n();
const respContent = ref('')
const visible = ref<boolean>(false);

const showModal = (item) => {
  respContent.value = getResp(item).content
  visible.value = true;
};

const handleCancel = (e: MouseEvent) => {
  visible.value = false;
};

const getHeader = (item) => {
  if (item.processorType === 'processor_interface_default') {
    let ret = item.name
    if (item.summary && item.summary.length > 0) ret += '&nbsp;' + joinArr(item.summary)
    return ret

  } else {
    return item.name + '&nbsp;' + joinArr(item.summary)
  }
}
const getReq = (item) => {
  if (!item.reqContent) return {}

  return JSON.parse(item.reqContent)
}
const getResp = (item) => {
  if (!item.reqContent) return {}

  return JSON.parse(item.respContent)
}

const joinArr = (arr : string[]) => {
  if (!arr) return ''

  if (Array.isArray(arr)) {
    return arr.join(' ')
  } else {
    return arr + ''
  }
}

</script>

<style lang="less" scoped>
.scenario-exec-log-main {
  height: 100%;
  padding: 10px;
  .title {
    font-weight: bolder;
  }
  .url {
    margin-bottom: 10px;
  }
  .extractor {
    margin-bottom: 10px;
  }
  .checkpoint {
    margin-bottom: 10px;
  }
  .header {
    margin-bottom: 10px;
  }
  .resp {
    margin-top: 10px;
  }

  .log {
    .resp-content {
      width: 100%;
      word-break:break-word;
      overflow: auto;
    }
  }
}

</style>