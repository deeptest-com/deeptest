<template>
  <div class="response-xml-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="1">
          <span style="margin-left:5px;">XML</span>
        </a-col>

        <a-col flex="100px" class="dp-right">
<!--          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>格式化</template>
            <ClearOutlined @click="format" class="dp-icon-btn dp-trans-80" />
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
        ref="monacoEditor"
        customId="xml-lens-main"
        class="editor"
        :value="content"
        :language="responseData.contentLang"
        theme="vs"
        :options="editorOptions"
        :onExtractor="responseExtractor"
      />
    </div>

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
</template>

<script setup lang="ts">
import {computed, inject, onMounted, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { DownloadOutlined, CopyOutlined, ClearOutlined } from '@ant-design/icons-vue';
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import {MonacoOptions} from "@/utils/const";
import {formatXml} from "@/utils/dom";
import {parseXml, testExpr} from "@/views/component/debug/service";
import {ExtractorSrc, ExtractorType, UsedBy} from "@/utils/enum";
import {StateType as Debug} from "@/views/component/debug/store";
import ResponseExtractor from "@/components/Editor/ResponseExtractor.vue";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";

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
const content = ref(formatXml(responseData.value.content))

const responseExtractorVisible = ref(false)
const expr = ref('')
const exprType = ref('')
const result = ref('')
const monacoEditor = ref();

const responseExtractor = (data) => {
  result.value = ''

  parseXml({
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

const testParse = (expr1, exprType1) => {
  console.log('testParse')
  testExpr({
    content: responseData.value.content,
    type: responseData.value.contentLang,
    expr: expr1,
    exprType: exprType1,
  }).then((json) => {
    console.log('json', json)
    result.value = json.data.result
  })
}

const responseExtractorFinish = (conf) => {
  console.log('responseExtractorFinish')

  conf.type = conf.expressionType === 'regx' ? ExtractorType.regx : ExtractorType.xmlquery
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

watch(responseData, () => {
  console.log('watch responseData')
  content.value = formatXml(responseData.value.content)
}, {deep: true})

const format = (item) => {
  console.log('format', item)
  bus.emit(settings.eventEditorAction, {act: settings.eventTypeFormat})
}

onMounted(() => {
  bus.on(settings.paneResizeTop, () => {
    monacoEditor.value?.resizeIt({
      act: settings.eventTypeContainerHeightChanged,
      container: 'response-xml-main',
      id: 'xml-lens-main',
    });
  })
});

</script>

<style lang="less">
.response-xml-main {
  .xmleditor-vue {
    height: 100%;
    .xmleditor-menu {
      display: none;
    }
    .xmleditor-outer {
      margin: 0;
      padding: 0;
      height: 100%;
      .ace-xmleditor {
        height: 100%;
      }
    }
  }
}
</style>

<style lang="less" scoped>
.response-xml-main {
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
