<template>
  <div class="path-param-header">
    <!-- debug-url-1 -->
    <a-input class="path-param-header-input" placeholder="请输入路径"
             :value="url"
             @change="changeUrl"
             v-contextmenu="e => onContextMenuShow(0, e)">

      <template #addonBefore>
        <a-select :options="serveServers" :value="serverId || null" @change="changeServer"
                  placeholder="请选择环境" class="select-env">
        </a-select>
        <span v-if="envURL" class="current-env-url">{{ envURL || '---' }}</span>
      </template>
    </a-input>

    <ContextMenu
        :isShow="showContextMenu"
        :style="contextMenuStyle"
        :menu-click="onMenuClick">
    </ContextMenu>
  </div>
</template>

<script setup lang="ts">
import {computed, ref,watch} from "vue";
import {useStore} from "vuex";
import {Endpoint} from "@/views/endpoint/data";
import {StateType as Debug} from "@/views/component/debug/store";
import {getContextMenuStyle2} from "@/utils/dom";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import ContextMenu from "@/views/component/debug/others/variable-replace/ContextMenu.vue"
import useVariableReplace from "@/hooks/variable-replace";

const store = useStore<{ Debug: Debug, Endpoint }>();

const debugData = computed<any>(() => store.state.Debug.debugData);

const endpointDetail: any = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);
const serveServers: any = computed<Endpoint>(() => store.state.Endpoint.serveServers);

const serverId = computed(() => {
  return debugData?.value?.serverId || endpointDetail?.value?.serverId || serveServers?.value[0]?.value || ''
});
const envURL = computed(() => {
  return serveServers.value?.find((item) => {
    return serverId.value === item.id;
  })?.url
});

const url = computed(() => {
  console.log('computed url')
  const u = debugData?.value.url || endpointDetail.value.path
  return u
});
const changeUrl = (e) => {
  store.commit('Debug/setUrl', e.target.value.trim())
}

function changeServer(id) {
  console.log('Debug/changeServer',id)
  store.dispatch('Debug/changeServer', id)
}

const { showContextMenu, contextMenuStyle, onContextMenuShow, onMenuClick } = useVariableReplace('endpointInterfaceUrl')

watch(() => {
  return envURL.value
}, (newVal) => {
  console.log('Debug/updateBaseUrl',newVal)
  store.dispatch("Debug/updateBaseUrl",newVal)
}, {
  immediate: true
})
</script>

<style scoped lang="less">
.path-param-header {
  display: inline-block;
  overflow: hidden;
  width: 100%;
}

.select-env {
  min-width: 100px;
  text-align: left;
  border-right: 1px solid #d9d9d9;
}

.current-env-url {
  min-width: 120px;
  padding-left: 16px;
  display: inline-block
}
</style>
