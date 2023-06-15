<template>
    <div class="url-input-main">
      <a-input class="path-param-header-input" placeholder="请求站点"
               :value="currentEnvURL">
      </a-input>

      <a-input class="path-param-header-input" placeholder="请求路径"
               :value="url">
      </a-input>
    </div>
</template>
<script setup lang="ts">
import { computed, ref } from "vue";
import { useStore } from "vuex";
import {Endpoint} from "@/views/endpoint/data";
import {StateType as DebugStateType, StateType as Debug} from "@/views/component/debug/store";
import debounce from "lodash.debounce";
import {StateType as TestInterfaceStateType} from "@/views/debugger/store";
import {StateType as EndpointStateType} from "@/views/endpoint/store";

const store = useStore<{TestInterface: TestInterfaceStateType, Debug: DebugStateType, Endpoint: EndpointStateType}>();

const debugData = computed<any>(() => store.state.Debug.debugData);
const serveServers: any = computed(() => store.state.TestInterface.serveServers);

const serverId = computed(() => {
  return debugData?.value?.serverId || serveServers?.value[0]?.value || ''
});

const url = computed(() => {
  return debugData?.value.url
});

const currentServerId = ref(debugData.value.serverId || null);
const currentEnvURL = computed(() => {
  console.log('computed currentEnvURL', currentServerId.value, serveServers.value)

  return serveServers.value?.find((item) => {
    return currentServerId.value === item.id;
  })?.url
});

</script>

<style scoped lang="less">
.url-input-main {
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
