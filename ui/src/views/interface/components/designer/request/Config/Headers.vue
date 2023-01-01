<template>
  <div class="headers-main">
    <div class="dp-param-grid">
      <div class="head">
        <a-row type="flex">
          <a-col flex="1">请求头列表</a-col>
          <a-col flex="80px" class="dp-right">
            <a-tooltip overlayClassName="dp-tip-small">
              <template #title>帮助</template>
              <QuestionCircleOutlined class="dp-icon-btn dp-trans-80"/>
            </a-tooltip>

            <a-tooltip @click="removeAll" overlayClassName="dp-tip-small">
              <template #title>全部清除</template>
              <DeleteOutlined class="dp-icon-btn dp-trans-80"/>
            </a-tooltip>

            <a-tooltip @click="add" overlayClassName="dp-tip-small">
              <template #title>新增</template>
              <PlusOutlined class="dp-icon-btn dp-trans-80"/>
            </a-tooltip>
          </a-col>
        </a-row>
      </div>
      <div class="params">
        <a-row v-for="(item, idx) in interfaceData.headers" :key="idx" type="flex" class="param">
          <a-col flex="1">
            <a-input v-model:value="item.name" @change="onParamChange(idx)" class="dp-bg-input-transparent" />
          </a-col>
          <a-col flex="1">
            <a-input
                v-model:value="item.value"
                @change="onParamChange(idx)"
                v-contextmenu="e => onContextMenuShow(idx, e)"
                class="dp-bg-input-transparent" />
          </a-col>
          <a-col flex="80px" class="dp-right dp-icon-btn-container">
            <a-tooltip v-if="!item.disabled" @click="disable(idx)" overlayClassName="dp-tip-small">
              <template #title>禁用</template>
              <CheckCircleOutlined class="dp-icon-btn dp-trans-80" />
            </a-tooltip>

            <a-tooltip v-if="item.disabled" @click="disable(idx)" overlayClassName="dp-tip-small">
              <template #title>启用</template>
              <CloseCircleOutlined class="dp-icon-btn dp-trans-80 dp-light" />
            </a-tooltip>

            <a-tooltip @click="remove(idx)" overlayClassName="dp-tip-small">
              <template #title>移除</template>
              <DeleteOutlined class="dp-icon-btn dp-trans-80"/>
            </a-tooltip>

            <a-tooltip @click="insert(idx)" overlayClassName="dp-tip-small">
              <template #title>插入</template>
              <PlusOutlined class="dp-icon-btn dp-trans-80"/>
            </a-tooltip>
          </a-col>
        </a-row>
      </div>
    </div>

    <div v-if="showContextMenu" :style="contextMenuStyle" class="context-menu">
      <div @click="onMenuClick('replaceVari')" class="item">使用变量</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, ComputedRef, defineComponent, PropType, Ref, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined, DeleteOutlined, PlusOutlined, CheckCircleOutlined, CloseCircleOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface/store";
import {Header, Interface} from "@/views/interface/data";
import {getContextMenuStyle} from "@/views/interface/service";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";

const {t} = useI18n();
const store = useStore<{ Interface: StateType }>();
const interfaceData = computed<Interface>(() => store.state.Interface.interfaceData);

const onParamChange = (idx) => {
  console.log('onParamChange', idx)
  if (interfaceData.value.headers.length <= idx + 1
      && (interfaceData.value.headers[idx].name !== '' || interfaceData.value.headers[idx].value !== '')) {
    interfaceData.value.headers.push({} as Header)
  }
};

const add = () => {
  console.log('add')
  interfaceData.value.params.push({} as Header)
}
const removeAll = () => {
  console.log('removeAll', interfaceData.value.headers)
  interfaceData.value.headers = [{} as Header]
}

const disable = (idx) => {
  console.log('enable', idx)
  interfaceData.value.headers[idx].disabled = !interfaceData.value.headers[idx].disabled
}
const remove = (idx) => {
  console.log('remove')
  interfaceData.value.headers.splice(idx, 1)
  add()
}
const insert = (idx) => {
  console.log('insert')
  interfaceData.value.headers.splice(idx+1, 0, {} as Header)
}


const showContextMenu = ref(false)
const headerIndex = ref(-1)
let contextTarget = {} as any
const contextMenuStyle = ref({} as any)

const onContextMenuShow = (idx, e) => {
  console.log('onContextMenuShow', idx, e)
  if (!e) return

  contextMenuStyle.value = getContextMenuStyle(e)
  contextTarget = e.target
  headerIndex.value = idx

  showContextMenu.value = true
}

const onMenuClick = (item) => {
  console.log('onMenuClick', item)
  showContextMenu.value = false

  bus.emit(settings.eventVariableSelectionStatus, {src: 'header', index: headerIndex.value, data: contextTarget});
}

</script>

<style lang="less" scoped>
.headers-main {
  height: 100%;
}

</style>