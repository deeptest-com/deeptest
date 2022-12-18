<template>
  <div class="response-html-main">
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

    <iframe id="iframe" height="0"></iframe>

    <div class="body">
      <MonacoEditor
          class="editor"
          :interfaceId="interfaceData.id"
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
        :content="responseData.content"
        :selection="selectionRef"
        :onFinish="responseExtractorFinish"
        :onCancel="responseExtractorCancel"
    />
  </div>
</template>

<script setup lang="ts">
import {computed, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { DownloadOutlined, CopyOutlined, ClearOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import {MonacoOptions} from "@/utils/const";
import {Interface, Response} from "@/views/interface/data";
import ResponseExtractor from "@/components/Editor/ResponseExtractor.vue";
import {initIFrame} from "@/services/parser-html";

const {t} = useI18n();
const store = useStore<{ Interface: StateType }>();
const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);
const responseData = computed<Response>(() => store.state.Interface.responseData);

const responseExtractorVisible = ref(false)
const requestReplaceVisible = ref(false)

const editorOptions = ref(Object.assign({usedWith: 'response'}, MonacoOptions) )

const selectionRef = ref({} as any)

let frameElem: HTMLIFrameElement
let frameDoc: Document

const responseExtractor = (data) => {
  console.log('responseExtractor', data)
  // responseExtractorVisible.value = true

  frameDoc = initIFrame(data.html)

  frameDoc.getElementById('abc')?.setAttribute('data-com-deeptest-selection', '123456')
  const elem = frameDoc.querySelector('*[data-com-deeptest-selection="123456"]')

  console.log('==div===', elem)
}

const requestReplace = (data) => {
  console.log('requestReplace', data)
  requestReplaceVisible.value = true
}

const responseExtractorFinish = () => {
  console.log('responseExtractorFinish')
  responseExtractorVisible.value = false
}
const responseExtractorCancel = () => {
  console.log('responseExtractorCancel')
  responseExtractorVisible.value = false
}

</script>

<style lang="less">
.response-html-main {
  .htmleditor-vue {
    height: 100%;
    .htmleditor-menu {
      display: none;
    }
    .htmleditor-outer {
      margin: 0;
      padding: 0;
      height: 100%;
      .ace-htmleditor {
        height: 100%;
      }
    }
  }
}
</style>

<style lang="less" scoped>
.response-html-main {
  height: 100%;
  .head {
    padding: 2px 3px;
    border-bottom: 1px solid #d9d9d9;
  }
  .body {
    height: calc(100% - 30px);
    overflow-y: hidden;
    &>div.monaco-editor-vue3 {
      height: 100%;
    }
  }
}
</style>