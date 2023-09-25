<template>
  <div class="request-body-main">
    <div class="head">
      <a-row type="flex">
        <a-col flex="1">
          <span>
            原始请求体
          </span>
          <a-select
              ref="bodyType"
              :options="bodyTypes"
              v-model:value="debugData.bodyType"
              size="small"
              :dropdownMatchSelectWidth="false"
              :bordered="false"
          >
          </a-select>
        </a-col>

        <a-col flex="100px" class="dp-right">
          <a-tooltip overlayClassName="dp-tip-small dp-tip-white">
            <template #title>
              <div class="tips">
                <div>发送到服务端的请求Body内容。</div>
                <div @click="openHelp('why_interface')" class="dp-link-primary">
                  了解更多 <ArrowRightOutlined/>
                </div>
              </div>
            </template>
            <QuestionCircleOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>

<!--      <a-tooltip overlayClassName="dp-tip-small">
            <template #title>全部清除</template>
            <DeleteOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>格式化</template>
            <ClearOutlined class="dp-icon-btn dp-trans-80" />
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>导入</template>
            <ImportOutlined class="dp-icon-btn dp-trans-80" />
          </a-tooltip>-->
        </a-col>
      </a-row>
    </div>

    <div class="body">
      <div v-if="debugData.bodyType === 'multipart/form-data'">
        <BodyFormData></BodyFormData>
      </div>
      <div v-if="debugData.bodyType === 'application/x-www-form-urlencoded'">
        <BodyFormUrlencoded></BodyFormUrlencoded>
      </div>

      <div v-else class="editor-container">
        <MonacoEditor
          ref="monacoEditor"
          customId="request-body-main"
          class="editor"
          v-model:value="debugData.body"
          :language="codeLang"
          theme="vs"
          :options="editorOptions"
          @change="editorChange"
          :onReplace="replaceRequest"
          :timestamp="timestamp" />
      </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import {computed, inject, nextTick, onMounted, onUnmounted, ref, watch} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined, ArrowRightOutlined } from '@ant-design/icons-vue';
import {MonacoOptions} from "@/utils/const";
import {openHelp} from "@/utils/link";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import {getCodeLangByType} from "@/views/component/debug/service";
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import BodyFormUrlencoded from "./Body-FormUrlencoded.vue";
import BodyFormData from "./Body-FormData.vue";
import {getRequestBodyTypes} from "@/views/scenario/service";
import {UsedBy} from "@/utils/enum";
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();

import {StateType as Debug} from "@/views/component/debug/store";
const store = useStore<{  Debug: Debug }>();

const debugData = computed<any>(() => store.state.Debug.debugData);

const codeLang = computed(() => {
  return getCodeLang()
})
const editorOptions = ref(Object.assign({usedWith: 'request'}, MonacoOptions))
const bodyTypes = ref(getRequestBodyTypes())
const timestamp = ref('')
const monacoEditor = ref();

watch(debugData, (newVal) => {
  timestamp.value = Date.now() + ''
}, {immediate: true, deep: true})

const getCodeLang = () => {
  console.log('debugData.value.bodyType', debugData.value.bodyType)
  return getCodeLangByType(debugData.value.bodyType)
}

const editorChange = (newScriptCode) => {
  debugData.value.body = newScriptCode;
}

const replaceRequest = (data) => {
  console.log('replaceRequest', data)
  bus.emit(settings.eventVariableSelectionStatus, {src: 'body', index: 0, data: data});
}

onMounted(() => {
  bus.on(settings.paneResizeTop, async () => {
    monacoEditor.value?.resizeIt({
      act: settings.eventTypeContainerHeightChanged,
      container: 'request-body-main',
      id: 'request-body-main'
    });
  })
})

</script>

<style lang="less">
.request-body-main {
  .editor-container {
    height: 100%;
    min-height: 160px;
    .editor {
      height: 100%;
      min-height: 160px;
    }

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
}
</style>

<style lang="less" scoped>
.request-body-main {
  height: 100%;
  overflow-y: scroll;
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
