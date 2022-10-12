<template>
  <div class="scenario-exec-log-main">
    <div v-for="(item, index) in logs" :key="index" class="log">
      <div>
        <div v-if="item.processorCategory !== 'processor_interface'">
          {{ item.name }}&nbsp; {{ joinArr(item.summary) }}
        </div>

        <a-collapse v-if="item.processorCategory === 'processor_interface'" expand-icon-position="right">
          <a-collapse-panel key="index" :header="getHeader(item)">
            <template #extra>
              <span :class="getResultCls(item.resultStatus)">{{ t(item.resultStatus) }}</span>
            </template>

            <div class="resp-content">
              <div class="section request">
                <div class="title">请求</div>
                <div class="content">
                  <a-row class="url item">
                    <a-col flex="200px">{{ getResp(item).statusCode }}</a-col>
                    <a-col flex="150px">{{ getReq(item).method }}</a-col>
                    <a-col flex="1">{{ getReq(item).url }}</a-col>
                  </a-row>
                </div>
              </div>

              <div v-if="getReq(item).method === 'POST'" class="show-detail">
                <a-link @click="showModal(item, 'req')" to="">显示请求体</a-link>
              </div>

              <div v-if="item.extractorsResult" class="section extractor">
                <div class="title">提取器</div>
                <div class="content">
                  <a-row v-for="(extractor, idx) in item.extractorsResult" :key="idx" type="flex" class="item">
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
              </div>

              <div v-if="item.checkpointsResult" class="section checkpoint">
                <div class="title">检查点</div>

                <div class="content">
                  <a-row v-for="(checkpoint, idx) in item.checkpointsResult" :key="idx" type="flex" class="item">
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
              </div>

              <div v-if="getResp(item).headers" class="section header">
                <div class="title">响应头</div>

                <div class="content">
                  <a-row v-for="(header, idx) in getResp(item).headers" :key="idx" type="flex" class="item">
                    <a-col flex="50px">{{idx + 1}}</a-col>
                    <a-col flex="300px">{{ header.name }}</a-col>
                    <a-col flex="1">{{ header.value }}</a-col>
                  </a-row>
                </div>
              </div>

              <div class="show-detail">
                <a-link @click="showModal(item, 'resp')" to="">显示响应内容</a-link>
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
      <div class="editor-wrapper">
        <MonacoEditor
            class="editor"
            :value="req.body ? req.body : resp.content"
            :language="req.bodyLang ? req.bodyLang : resp.contentLang"
            theme="vs"
            :options="editorOptions"
        />
      </div>
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
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import {MonacoOptions} from "@/utils/const";

defineProps<{
  logs: []
}>()

const { t } = useI18n();
const editorOptions = ref(MonacoOptions)

const req = ref({})
const resp = ref({})
const visible = ref<boolean>(false);

const showModal = (item, type) => {
  if (type === 'req') {
    req.value = getReq(item)
    resp.value = {}
  }
  else if (type === 'resp') {
    req.value = {}
    resp.value = getResp(item)
  }

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

<style lang="less">
.scenario-exec-log-main {
  height: 100%;
  .ant-collapse-content-box {
    padding-top: 8px !important;
  }
}
</style>

<style lang="less" scoped>
.scenario-exec-log-main {
  height: 100%;
  padding: 5px 10px 0px 10px;

  .log {
    padding-top: 8px;
  }

  .editor-wrapper {
    height: calc(100% - 30px);
    overflow-x: hidden;
    overflow-y: hidden;
    &>div {
      height: 100%;
    }
  }

  .resp-content {
    .section {
      margin: 10px 0;
      border: 1px solid #dedfe1;

      .title {
        padding: 6px 10px;
        border-bottom: 1px solid #dedfe1;
        font-weight: bolder;
      }
      .content {
        padding: 6px 0;
        .item {
          padding: 3px 10px;
          line-height: 20px;
          &:nth-child(even) {
            background-color: #f5f6f8 !important;
          }
        }
      }
      .show-detail {
        margin: 5px 0;
      }
    }
  }
}

</style>