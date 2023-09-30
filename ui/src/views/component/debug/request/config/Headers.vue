<template>
  <div class="headers-main">
    <div class="dp-param-grid">
      <div class="head">
        <a-row type="flex">
          <a-col flex="1" class="title">参数名</a-col>
          <a-col flex="1" class="title">参数值</a-col>

          <a-col flex="80px" class="dp-right">
            <Tips section="i" title="帮助" />

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
        <a-row v-for="(item, idx) in debugData.headers" :key="idx" type="flex" class="param">
          <a-col flex="1">
            <a-auto-complete
              class="dp-bg-input-transparent"
              v-model:value="item.name"
              style="width: 100%"
              allowClear
              @change="onParamChange(idx)"
              :options="requestHeaderOptions"
            />
          </a-col>
          <a-col flex="1">
            <a-input :id="'header' + idx"
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

    <ContextMenu
        :isShow="showContextMenu"
        :style="contextMenuStyle"
        :menu-click="onMenuClick">
    </ContextMenu>
  </div>
</template>

<script setup lang="ts">
import {computed, inject} from "vue";
import {useI18n} from "vue-i18n";
import {useStore} from "vuex";
import { DeleteOutlined, PlusOutlined, CheckCircleOutlined, CloseCircleOutlined } from '@ant-design/icons-vue';
import Tips from "@/components/Tips/index.vue";
import {UsedBy} from "@/utils/enum";
import {Header} from "@/views/component/debug/data";
import ContextMenu from "@/views/component/debug/others/variable-replace/ContextMenu.vue"

import {StateType as Debug} from "@/views/component/debug/store";
import useVariableReplace from "@/hooks/variable-replace";
import { requestHeaderOptions } from "@/config/constant";

const usedBy = inject('usedBy') as UsedBy
const {t} = useI18n();
const store = useStore<{  Debug: Debug }>();
const debugData = computed<any>(() => store.state.Debug.debugData);

const onParamChange = (idx) => {
  console.log('onParamChange', idx)
  if (debugData.value.headers.length <= idx + 1
      && (debugData.value.headers[idx].name !== '' || debugData.value.headers[idx].value !== '')) {
    debugData.value.headers.push({} as Header)
  }
};

const add = () => {
  console.log('add')
  debugData.value.headers.push({} as Header)
}
const removeAll = () => {
  console.log('removeAll', debugData.value.headers)
  debugData.value.headers = [{} as Header]
}

const disable = (idx) => {
  console.log('enable', idx)
  debugData.value.headers[idx].disabled = !debugData.value.headers[idx].disabled
}
const remove = (idx) => {
  console.log('remove')
  debugData.value.headers.splice(idx, 1)
  const len = debugData.value.headers.length
  if (len == 0 || !!debugData.value.headers[len-1].name) {
    add()
  }
}
const insert = (idx) => {
  console.log('insert')
  debugData.value.headers.splice(idx+1, 0, {} as Header)
}

const { showContextMenu, contextMenuStyle, onContextMenuShow, onMenuClick } = useVariableReplace('header')

</script>

<style lang="less" scoped>
.headers-main {
  height: 100%;
  //max-height: 180px;
  overflow-y: scroll;
}


</style>
