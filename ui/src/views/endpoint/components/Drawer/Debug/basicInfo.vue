<template>
    <div class="path-param-header">
        <a-input class="path-param-header-input" :value="endpointDetail.path" placeholder="请输入路径">
            <template #addonBefore>
                <a-select :options="serveServers" :value="currentServerId || null" @change="changeServer"
                    placeholder="请选择环境" class="select-env">
                </a-select>
                <span v-if="currentEnvURL" class="current-env-url">{{ currentEnvURL || '---' }}</span>
            </template>
        </a-input>
    </div>
</template>
<script setup lang="ts">
import { computed, ref } from "vue";
import { useStore } from "vuex";
import {Endpoint} from "@/views/endpoint/data";
import {StateType as Debug} from "@/views/component/debug/store";
const store = useStore<{  Debug: Debug, Endpoint }>();

const endpointDetail: any = computed<Endpoint>(() => store.state.Endpoint.endpointDetail);
const serveServers: any = computed<Endpoint>(() => store.state.Endpoint.serveServers);

const currentServerId = ref(endpointDetail?.value?.serverId || serveServers?.value[0]?.value || '');
const currentEnvURL = computed(() => {
  return serveServers.value?.find((item) => {
    return currentServerId.value === item.id;
  })?.url
});

function changeServer(val) {
  currentServerId.value = val;
  endpointDetail.value.serverId = val;
  store.commit('Endpoint/setEndpointDetail', {
    ...endpointDetail.value,
  })
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
