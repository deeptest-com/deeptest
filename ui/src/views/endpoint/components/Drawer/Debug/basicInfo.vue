<template>
    <div class="path-param-header">
        <a-input class="path-param-header-input" placeholder="请输入路径"
                 :value="url"
                 @change="changeUrl">
            <template #addonBefore>
                <a-select :options="serveServers" :value="serverId || null" @change="changeServer"
                    placeholder="请选择环境" class="select-env">
                </a-select>
                <span v-if="envURL" class="current-env-url">{{ envURL || '---' }}</span>
            </template>
        </a-input>
    </div>
</template>
<script setup lang="ts">
import { computed, ref } from "vue";
import { useStore } from "vuex";
import {Endpoint} from "@/views/endpoint/data";
import {StateType as Debug} from "@/views/component/debug/store";
import debounce from "lodash.debounce";
const store = useStore<{  Debug: Debug, Endpoint }>();

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
  return debugData?.value.url || endpointDetail.value.path
});

function changeServer(id) {
  store.dispatch('Debug/changeServer', id)
}

const changeUrl = (e) => {
  store.commit('Debug/setUrl', e.target.value.trim())
}

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
