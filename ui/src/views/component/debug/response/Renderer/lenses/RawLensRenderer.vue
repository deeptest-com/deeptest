<template>
  <div class="response-raw-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="1">
          <span style="margin-left:5px;">RAW</span>
        </a-col>

        <a-col flex="100px" class="dp-right">
<!--          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>格式化</template>
            <ClearOutlined class="dp-icon-btn dp-trans-80" />
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>复制</template>
            <CopyOutlined class="dp-icon-btn dp-trans-80" />
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>下载</template>
            <DownloadOutlined class="dp-icon-btn dp-trans-80" />
          </a-tooltip>-->
        </a-col>
      </a-row>
    </div>

    <div class="body">
      <MonacoEditor
          class="editor"
          :value="responseData.content"
          :timestamp="timestamp"
          :language="responseData.contentLang"
          theme="vs"
          :options="editorOptions"
          :onExtractor="responseExtractor"
      />

      <ResponseExtractor
          v-if="responseExtractorVisible"
          :interfaceId="debugData.endpointInterfaceId"
          :exprType="exprType"
          :expr="expr"
          :result="result"
          :onTest="testParse"
          :onFinish="responseExtractorFinish"
          :onCancel="responseExtractorCancel"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, inject, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { DownloadOutlined, CopyOutlined, ClearOutlined } from '@ant-design/icons-vue';

import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import {MonacoOptions} from "@/utils/const";
import {StateType as Debug} from "@/views/component/debug/store";
import ResponseExtractor from "@/components/Editor/ResponseExtractor.vue";
import {parseText, testExpr} from "@/views/component/debug/service";
import {ExtractorSrc, ExtractorType, UsedBy} from "@/utils/enum";
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();
const store = useStore<{  Debug: Debug }>();

const debugInfo = computed<any>(() => store.state.Debug.debugInfo);
const debugData = computed<any>(() => store.state.Debug.debugData);
const responseData = computed<any>(() => store.state.Debug.responseData);

const timestamp = ref('')
watch(responseData, (newVal) => {
  timestamp.value = Date.now() + ''
}, {immediate: true, deep: true})

const editorOptions = ref(Object.assign({usedWith: 'response',readOnly:false}, MonacoOptions) )

const responseExtractorVisible = ref(false)
const expr = ref('')
const exprType = ref('')
const result = ref('')

const responseExtractor = (data) => {
  result.value = ''

  parseText({
    docContent: data.docContent,
    selectContent: data.selectContent,

    startLine: data.selectionObj.startLineNumber - 1,
    endLine: data.selectionObj.endLineNumber - 1,
    startColumn: data.selectionObj.startColumn - 1,
    endColumn: data.selectionObj.endColumn - 1,
  }).then((json) => {
    console.log('json', json)
    responseExtractorVisible.value = true
    expr.value = json.data.expr
    exprType.value = json.data.exprType
  })
}

const testParse = (expr, exprType) => {
  console.log('testParse')
  testExpr({
    content: responseData.value.content,
    type: responseData.value.contentLang,
    expr: expr,
    exprType: exprType,

  }).then((json) => {
    console.log('json', json)
    result.value = json.data.result
  })
}

const responseExtractorFinish = (conf) => {
  console.log('responseExtractorFinish')

  conf.type = conf.expressionType
  conf.src = ExtractorSrc.body
  conf.result = result.value

  const data = {
    conf,
    info: debugInfo.value,
  } as any

  store.dispatch('Debug/quickCreateExtractor', data).then((result) => {
    if (result) {
      responseExtractorVisible.value = false
    }
  })
}
const responseExtractorCancel = () => {
  console.log('responseExtractorCancel')
  responseExtractorVisible.value = false
}

</script>

<style lang="less">
.response-raw-main {
  .raweditor-vue {
    height: 100%;
    .raweditor-menu {
      display: none;
    }
    .raweditor-outer {
      margin: 0;
      padding: 0;
      height: 100%;
      .ace-raweditor {
        height: 100%;
      }
    }
  }
}
</style>

<style lang="less" scoped>
.response-raw-main {
  height: 100%;
  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .body {
    height: calc(100% - 30px);
    overflow-y: hidden;
    &>div {
      height: 100%;
    }
  }
}
</style>
