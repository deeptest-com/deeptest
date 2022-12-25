<template>
  <div class="response-json-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="1">
          <span>响应体</span>
        </a-col>

        <a-col flex="100px" class="dp-right">
          <a-tooltip overlayClassName="dp-tip-small">
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
          </a-tooltip>
        </a-col>
      </a-row>
    </div>

    <div class="body">
      <MonacoEditor
          class="editor"
          :value="responseData.content"
          :language="responseData.contentLang"
          theme="vs"
          :options="editorOptions"

          :onExtractor="responseExtractor"
          :onReplace="responseExtractor"
      />
    </div>

    <ResponseExtractor
        v-if="responseExtractorVisible"
        :interfaceId="interfaceData.id"
        type="json"
        :xpath="xpath"
        :result="result"
        :onTest="testParse"
        :onFinish="responseExtractorFinish"
        :onCancel="responseExtractorCancel"
    />
  </div>
</template>

<script setup lang="ts">
import {computed, ComputedRef, defineComponent, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { DownloadOutlined, CopyOutlined, ClearOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import {isInArray} from "@/utils/array";
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import {MonacoOptions} from "@/utils/const";
import {Interface, Response} from "@/views/interface/data";

import {parseHtml, parseJson, testXPath} from "@/views/interface/service";
import {ExtractorSrc, ExtractorType} from "@/utils/enum";
import ResponseExtractor from "@/components/Editor/ResponseExtractor.vue";

const {t} = useI18n();
const store = useStore<{ Interface: StateType }>();
const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);
const responseData = computed<Response>(() => store.state.Interface.responseData);
const editorOptions = ref(Object.assign({usedWith: 'response'}, MonacoOptions) )

const responseExtractorVisible = ref(false)
const xpath = ref('')
const result = ref('')

const responseExtractor = (data) => {
  // console.log('responseExtractor', data)
  result.value = ''

  parseJson({
    docHtml: data.docHtml,
    selectContent: data.selectContent,

    startLine: data.selectionObj.startLineNumber - 1,
    endLine: data.selectionObj.endLineNumber - 1,
    startColumn: data.selectionObj.startColumn - 1,
    endColumn: data.selectionObj.endColumn - 1,
  }).then((json) => {
    console.log('json', json)
    responseExtractorVisible.value = true
    xpath.value = json.data.xpath
  })
}

const testParse = (xpath) => {
  console.log('testXPath')
  testXPath({
    content: responseData.value.content,
    type: responseData.value.contentLang,
    xpath: xpath,
  }).then((json) => {
    console.log('json', json)
    result.value = json.data.result
  })
}

const responseExtractorFinish = (data) => {
  console.log('responseExtractorFinish')
  data.type = ExtractorType.htmlquery
  data.src = ExtractorSrc.body
  data.result = result.value

  data.interfaceId = interfaceData.value.id
  data.projectId = interfaceData.value.projectId
  store.dispatch('Interface/createExtractorOrUpdateResult', data).then((result) => {
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
.response-json-main {
  .jsoneditor-vue {
    height: 100%;
    .jsoneditor-menu {
      display: none;
    }
    .jsoneditor-outer {
      margin: 0;
      padding: 0;
      height: 100%;
      .ace-jsoneditor {
        height: 100%;
      }
    }
  }
}
</style>

<style lang="less" scoped>
.response-json-main {
  height: 100%;
  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .body {
    height: calc(100% - 30px);
    overflow-x: hidden;
    overflow-y: hidden;
    &>div {
      height: 100%;
    }
  }
}
</style>