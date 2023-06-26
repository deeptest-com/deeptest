<template>
  <div class="path-param-header">
    <!-- debug-url-2 -->
    <a-input-group>
      <a-row type="flex" :gutter="0">
        <a-col flex="80px">
          <a-select class="select-env"
                    :options="methods"
                    v-model:value="debugData.method">
          </a-select>
        </a-col>
        <a-col flex="3">

          <a-input class="path-param-header-input" placeholder="请输入路径"
                   :value="debugData.url"
                   @change="changeUrl"
                   v-contextmenu="e => onContextMenuShow(0, e)">

            <template #addonBefore>
              <a-select :options="servers" :value="serverId || null" @change="changeServer"
                        placeholder="请选择环境" class="select-env">
              </a-select>
              <span v-if="envURL" class="current-env-url">{{ envURL || '---' }}</span>
            </template>
          </a-input>
        </a-col>
      </a-row>
    </a-input-group>

    <ContextMenu
        :isShow="showContextMenu"
        :style="contextMenuStyle"
        :menu-click="onMenuClick">
    </ContextMenu>
  </div>
</template>
<script setup lang="ts">
import {computed, ref, watch} from "vue";
import {useStore} from "vuex";
import {StateType as Debug} from "@/views/component/debug/store";
import {serverList} from "@/views/project-settings/service";
import ContextMenu from "@/views/component/debug/others/variable-replace/ContextMenu.vue"
import useVariableReplace from "@/hooks/variable-replace";
import {getArrSelectItems} from "@/utils/comm";
import {Methods} from "@/utils/enum";

const store = useStore<{ Debug: Debug, Endpoint }>();

const debugData = computed<any>(() => store.state.Debug.debugData);
const servers = ref([] as any);

const methods = getArrSelectItems(Methods)

const serverId = computed(() => {
  return debugData?.value?.serverId || servers?.value[0]?.value || ''
});
const envURL = computed(() => {
  const server = servers.value?.find((item) => {
    return serverId.value === item.value;
  })
  return server?.url
});
// const url = computed(() => {
//   return debugData?.value.url
// });

const listServer = async (serverId) => {
  servers.value = []
  if (!serverId) {
    return
  }

  const res = await serverList({
    serverId: serverId
  });
  if (res.code === 0) {
    res.data.forEach((item: any) => {
      servers.value.push({
        label: item.description,
        value: item.id,
        url: item.url,
      })
    })
  }
}

watch(debugData, (val) => {
  console.log('-----', debugData.value.serverId)
  listServer(debugData.value.serverId)
}, {immediate: true, deep: true});

function changeServer(id) {
  store.dispatch('Debug/changeServer', id)
}

const changeUrl = (e) => {
  store.commit('Debug/setUrl', e.target.value.trim())
}

const {showContextMenu, contextMenuStyle, onContextMenuShow, onMenuClick} = useVariableReplace('scenarioInterfaceUrl')

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
