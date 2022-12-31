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
              v-model:value="interfaceData.bodyType"
              size="small"
              :dropdownMatchSelectWidth="false"
              :bordered="false"
          >
          </a-select>
        </a-col>

        <a-col flex="100px" class="dp-right">
          <a-tooltip overlayClassName="dp-tip-small">
            <template #title>帮助</template>
            <QuestionCircleOutlined class="dp-icon-btn dp-trans-80"/>
          </a-tooltip>

          <a-tooltip overlayClassName="dp-tip-small">
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
          </a-tooltip>
        </a-col>
      </a-row>
    </div>

    <div class="body">
      <div v-if="interfaceData.bodyType === 'multipart/form-data'">
        <BodyFormData></BodyFormData>
      </div>
      <div v-if="interfaceData.bodyType === 'application/x-www-form-urlencoded'">
        <BodyFormUrlencoded></BodyFormUrlencoded>
      </div>

      <div v-else>
        <MonacoEditor
            class="editor"
            v-model:value="interfaceData.body"
            :language="codeLang"
            theme="vs"
            :options="editorOptions"
            @change="editorChange"
            :onReplace="replaceRequest"
        />
      </div>
    </div>

    <RequestVariable
        v-if="requestVariableVisible"
        :interfaceId="interfaceData.id"
        :onFinish="requestVariableSelectFinish"
        :onCancel="requestVariableSelectCancel"
    />
  </div>
</template>

<script setup lang="ts">
import {computed, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined, DeleteOutlined, ClearOutlined, ImportOutlined, CheckCircleOutlined, CloseCircleOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import {MonacoOptions} from "@/utils/const";
import {Interface} from "@/views/interface/data";
import {getCodeLangByType} from "@/views/interface/service";
import MonacoEditor from "@/components/Editor/MonacoEditor.vue";
import BodyFormUrlencoded from "./Body-FormUrlencoded.vue";
import BodyFormData from "./Body-FormData.vue";
import {getRequestBodyTypes} from "@/views/scenario/service";
import RequestVariable from "@/components/Editor/RequestVariable.vue";

const {t} = useI18n();
const store = useStore<{ Interface: StateType }>();
const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);
const codeLang = computed(() => {
  return getCodeLang()
})
const editorOptions = ref(Object.assign({usedWith: 'request'}, MonacoOptions))

const bodyTypes = ref(getRequestBodyTypes())

const getCodeLang = () => {
  console.log('interfaceData.value.bodyType', interfaceData.value.bodyType)
  return getCodeLangByType(interfaceData.value.bodyType)
}

const editorChange = (newScriptCode) => {
  interfaceData.value.body = newScriptCode;
}

const replaceRequest = (data) => {
  console.log('replaceRequest', data)
  requestVariableVisible.value = true
}

const requestVariableVisible = ref(false)

const requestVariableSelectFinish = (data) => {
  console.log('requestVariableSelectFinish', data)

  data.interfaceId = interfaceData.value.id
  data.projectId = interfaceData.value.projectId
  // store.dispatch('Interface/createExtractorOrUpdateResult', data).then((result) => {
  //   if (result) {
  //     requestVariableVisible.value = false
  //   }
  // })
}
const requestVariableSelectCancel = () => {
  console.log('requestVariableSelectCancel')
  requestVariableVisible.value = false
}

</script>

<style lang="less">
.request-body-main {
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
.request-body-main {
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