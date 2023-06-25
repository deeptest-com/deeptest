<template>
  <a-modal
      title="使用变量"
      :destroy-on-close="true"
      :mask-closable="false"
      :visible="requestVariableVisible"
      :footer="null"
      :onCancel="onCancel"
      width="800px"
      height="600px">
    <div>
      <a-row>
        <a-col flex="100px" class="dp-border">共享变量</a-col>
      </a-row>
      <a-row v-for="(item, idx) in debugData.shareVars" :key="idx" type="flex">
        <a-col flex="100px">{{item.name}}</a-col>
        <a-col :flex="3">{{item.rightValue==='extractor_err'? t(item.rightValue+'_short') : item.value}}</a-col>

        <a-col flex="100px">
          <span @click="select(item)" class="dp-link-primary">选择</span>
        </a-col>
      </a-row>
      <br/>

      <a-row>
        <a-col flex="100px" class="dp-border">环境变量</a-col>
      </a-row>
      <a-row v-for="(item, idx) in debugData.envVars" :key="idx" type="flex">
        <a-col flex="100px">{{item.name}}</a-col>
        <a-col :flex="3">{{item.rightValue}}</a-col>

        <a-col flex="100px">
          <span @click="select(item)" class="dp-link-primary">选择</span>
        </a-col>
      </a-row>
      <br/>

      <a-row>
        <a-col flex="100px" class="dp-border">全局变量</a-col>
      </a-row>
      <a-row v-for="(item, idx) in debugData.globalVars" :key="idx" type="flex">
        <a-col flex="100px">{{item.name}}</a-col>
        <a-col :flex="3">{{item.rightValue}}</a-col>

        <a-col flex="100px">
          <span @click="select(item)" class="dp-link-primary">选择</span>
        </a-col>
      </a-row>

    </div>

  </a-modal>
</template>

<script setup lang="ts">
import {onMounted, ref, computed, onUnmounted} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";

import {Interface} from "@/views/component/debug/data";
import {StateType as DebugStateType} from "@/views/component/debug/store";
import {StateType as EnvironmentStateType} from "@/store/environment";

const { t } = useI18n();

const store = useStore<{ Debug: DebugStateType, EnvironmentState: EnvironmentStateType }>();
const debugData = computed<Interface>(() => store.state.Debug?.debugData);

const requestVariableVisible = ref(false)

onMounted(()=> {
  console.log('onMounted')
  bus.on(settings.eventVariableSelectionStatus, onVariableSelectionStatus);
})

onUnmounted(() => {
  console.log('onUnmounted')
  bus.off(settings.eventVariableSelectionStatus, onVariableSelectionStatus);
})

const variableSelectionData = ref({}as any)
const onVariableSelectionStatus = (data) => {
  console.log('onVariableSelectionStatus', data)

  variableSelectionData.value = data
  requestVariableVisible.value = true
}

const selectMenuItem = async (item) => {
  console.log('select', item, variableSelectionData.value, debugData.value)
  const targetElemId = '' + variableSelectionData.value.src + variableSelectionData.value.index

  if (variableSelectionData.value.src.indexOf('InterfaceUrl') > -1) {
    let url = debugData.value.url
    url = getInputNewContent(item.name, url,
        variableSelectionData.value.data.selectionStart, variableSelectionData.value.data.selectionEnd)

    store.dispatch("Debug/updateUrl", url);

  } else if (variableSelectionData.value.src === 'body') {
    const body = getEditorNewContent(item.name)
    store.dispatch("Debug/updateBody", body);

  } else if (variableSelectionData.value.src === 'queryParam') {
    let param = debugData.value.queryParams[variableSelectionData.value.index].value
    param = getInputNewContent(item.name, param,
        variableSelectionData.value.data.selectionStart, variableSelectionData.value.data.selectionEnd)

    updateInput(targetElemId, param)

  }  else if (variableSelectionData.value.src === 'pathParam') {
    let param = debugData.value.pathParams[variableSelectionData.value.index].value
    param = getInputNewContent(item.name, param,
        variableSelectionData.value.data.selectionStart, variableSelectionData.value.data.selectionEnd)

    updateInput(targetElemId, param)

  } else if (variableSelectionData.value.src === 'header') {
    let header = debugData.value.headers[variableSelectionData.value.index].value
    header = getInputNewContent(item.name, header,
            variableSelectionData.value.data.selectionStart, variableSelectionData.value.data.selectionEnd)

    updateInput(targetElemId, header)

  }

  requestVariableVisible.value = false
}

const updateInput = (id, val) => {
  const targetElem = document.getElementById(id)
  if (targetElem) {
    targetElem.value = val
    targetElem.dispatchEvent(new Event('input'));
  }
}

const getInputNewContent = (variName, val, start, end) => {
  const ret = val.substring(0, variableSelectionData.value.data.selectionStart) +
      '${' + variName + '}' +
      val.substring(variableSelectionData.value.data.selectionEnd)

  return ret
}

const getEditorNewContent = (variName) => {
  console.log('getEditorNewContent', variName)

  const docContent = variableSelectionData.value.data.docContent
  const selectContent = variableSelectionData.value.data.selectContent
  const selectionObj = variableSelectionData.value.data.selectionObj

  const startLine = selectionObj.startLineNumber - 1
  const startColumn = selectionObj.startColumn - 1
  const endLine = selectionObj.endLineNumber - 1
  const endColumn = selectionObj.endColumn - 1

  const lines = docContent.split('\n')
  if (startLine === endLine) {
    let line = lines[startLine]
    lines[startLine] = line.substring(0, startColumn) + '${' + variName + '}' + line.substring(endColumn)

    return lines.join('\n')

  } else if (startLine < endLine) {
    let ret = [] as string[]

    for (let i = startLine + 1; i++; i < endLine) {
      if (i === startLine) {
        ret.push(lines[i].substring(0, startColumn) + '${' + variName + '}')

      } else if (i === endLine) {
        ret.push(lines[i].substring(endColumn))

      } else if (i < startLine || i > endLine) {
        ret.push(lines[i])
      } else if (i > startLine && i < endLine) {
        // ignore
      }
    }

    return ret.join('\n')
  }


}

const onCancel = () => {
  requestVariableVisible.value = false
}

const labelCol = { span: 6 }
const wrapperCol = { span: 16 }

</script>

<style lang="less">
.request-variable-main {

}
</style>