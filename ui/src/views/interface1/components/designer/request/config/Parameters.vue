<template>
  <div class="parameters-main">
    <div class="dp-param-grid">
      <div class="head">
        <a-row type="flex">
          <a-col flex="1">查询参数</a-col>
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
        <a-row v-for="(item, idx) in interfaceData.params" :key="idx" type="flex" class="param">
          <a-col flex="1">
            <a-input v-model:value="item.name" @change="onParamChange(idx)" class="dp-bg-input-transparent" />
          </a-col>
          <a-col flex="1">
            <a-input v-model:value="item.value"
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

    <ContextMenu
        :isShow="showContextMenu"
        :style="contextMenuStyle"
        :menu-click="onMenuClick">
    </ContextMenu>

  </div>
</template>

<script setup lang="ts">
import {computed, inject, ref} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { QuestionCircleOutlined, DeleteOutlined, PlusOutlined, CheckCircleOutlined, CloseCircleOutlined } from '@ant-design/icons-vue';
import {StateType} from "@/views/interface1/store";
import {Param, Interface} from "@/views/interface1/data";
import {getContextMenuStyle} from "@/views/interface1/service";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import ContextMenu from "@/components/Editor/ContextMenu.vue"
import {UsedBy} from "@/utils/enum";
import {StateType as ScenarioStateType} from "@/views/scenario/store";
const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();
const store = useStore<{ Interface1: StateType, Scenario: ScenarioStateType }>();
const interfaceData = computed<Interface>(
    () => usedBy === UsedBy.InterfaceDebug ? store.state.Interface1.interfaceData : store.state.Scenario.interfaceData);

const onParamChange = (idx) => {
  console.log('onParamChange', idx)
  if (interfaceData.value.params.length <= idx + 1
        && (interfaceData.value.params[idx].name !== '' || interfaceData.value.params[idx].value !== '')) {
    interfaceData.value.params.push({} as Param)
  }
};

const add = () => {
  console.log('add')
  interfaceData.value.params.push({} as Param)
}
const removeAll = () => {
  console.log('removeAll', interfaceData.value.params)
  interfaceData.value.params = [{} as Param]
}

const disable = (idx) => {
  console.log('enable', idx)
  interfaceData.value.params[idx].disabled = !interfaceData.value.params[idx].disabled
}
const remove = (idx) => {
  console.log('remove')
  interfaceData.value.params.splice(idx, 1)
  add()
}
const insert = (idx) => {
  console.log('insert')
  interfaceData.value.params.splice(idx+1, 0, {} as Param)
}

const showContextMenu = ref(false)
const paramIndex = ref(-1)
let contextTarget = {} as any
const contextMenuStyle = ref({} as any)

const onContextMenuShow = (idx, e) => {
  console.log('onContextMenuShow', idx, e.target)
  if (!e) return

  contextMenuStyle.value = getContextMenuStyle(e)
  contextTarget = e.target
  paramIndex.value = idx

  showContextMenu.value = true
}

const onMenuClick = (key) => {
  console.log('onMenuClick', key)

  if (key === 'use-variable') {
    bus.emit(settings.eventVariableSelectionStatus, {src: 'param', index: paramIndex.value, data: contextTarget});
  }

  showContextMenu.value = false
}

</script>

<style lang="less" scoped>
.parameters-main {
  height: 100%;
}

</style>