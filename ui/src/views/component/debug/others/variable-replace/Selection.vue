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
        <a-col flex="100px" class="dp-border">环境变量</a-col>
      </a-row>
      <a-row v-for="(item, idx) in environmentData?.vars" :key="idx" type="flex">
        <a-col flex="100px">{{item.name}}</a-col>
        <a-col :flex="3">{{item.rightValue}}</a-col>

        <a-col flex="100px">
          <span @click="select(item)" class="dp-link-primary">选择</span>
        </a-col>
      </a-row>

      <br/>

      <a-row>
        <a-col flex="100px" class="dp-border">共享变量</a-col>
      </a-row>
      <a-row v-for="(item, idx) in validExtractorVariablesData" :key="idx" type="flex">
        <a-col flex="100px">{{item.name}}</a-col>
        <a-col :flex="3">{{item.rightValue==='extractor_err'? t(item.rightValue+'_short') : item.value}}</a-col>

        <a-col flex="100px">
          <span @click="select(item)" class="dp-link-primary">选择</span>
        </a-col>
      </a-row>

      <br/>

      <a-row v-for="(item, idx) in validExtractorVariablesData" :key="idx" type="flex">
        <a-col flex="100px"></a-col>
        <a-col :flex="3" class="dp-center">
          <a-button @click="() => onCancel()" type="primary">关闭</a-button>
        </a-col>
        <a-col flex="100px"></a-col>
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
const validExtractorVariablesData = computed<any>(() => store.state.Debug?.validExtractorVariablesData);
const environmentData = computed<any>(() => store.state.EnvironmentState?.environmentData);

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

const select = async (item) => {
  console.log('select', item, debugData.value)

  if (variableSelectionData.value.src === 'body') {
    const body = getEditorContent(item.name)

    await store.dispatch('Interface/updateBody', body)

  } else if (variableSelectionData.value.src === 'url') {
    let url = debugData.value.url

    url = getInputContent(item.name, url,
        variableSelectionData.value.data.selectionStart, variableSelectionData.value.data.selectionEnd)

    store.dispatch('Interface/updateUrl', url).then((res) => {
      console.log('res', res)
    })

  } else if (variableSelectionData.value.src === 'param') {
    let param = debugData.value.params[variableSelectionData.value.index].value

    param = getInputContent(item.name, param,
        variableSelectionData.value.data.selectionStart, variableSelectionData.value.data.selectionEnd)

    await store.dispatch('Interface/updateParam', {
      value: param,
      index: variableSelectionData.value.index,
    })

  } else if (variableSelectionData.value.src === 'header') {
    let header = debugData.value.headers[variableSelectionData.value.index].value

    header = getInputContent(item.name, header,
            variableSelectionData.value.data.selectionStart, variableSelectionData.value.data.selectionEnd)

    await store.dispatch('Interface/updateHeader', {
      value: header,
      index: variableSelectionData.value.index,
    })
  }

  requestVariableVisible.value = false
}

const getInputContent = (variName, val, start, end) => {
  const ret = val.substring(0, variableSelectionData.value.data.selectionStart) +
      '${' + variName + '}' +
      val.substring(variableSelectionData.value.data.selectionEnd)

  return ret
}

const getEditorContent = (variName) => {
  console.log('getEditorContent', variName)

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