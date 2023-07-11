<template>
    <div class="env-selection-main">
      <a-select class="select-env"
                :options="serveServers"
                :value="serverId || null"
                @change="changeServer"
                placeholder="请选择环境">
      </a-select>
    </div>
</template>
<script setup lang="ts">
import {computed} from "vue";
import { useStore } from "vuex";
import {StateType as DebugStateType} from "@/views/component/debug/store";
import {StateType as EndpointStateType} from "@/views/endpoint/store";
import {StateType as DiagnoseInterfaceStateType} from "@/views/diagnose/store";

const store = useStore<{DiagnoseInterface: DiagnoseInterfaceStateType, Debug: DebugStateType, Endpoint: EndpointStateType}>();

const debugData = computed<any>(() => store.state.Debug.debugData);
const serveServers: any = computed(() => store.state.DiagnoseInterface.serveServers);

const serverId = computed(() => {
  return debugData?.value?.serverId || serveServers?.value[0]?.value || ''
});

function changeServer(id) {
  console.log('changeServer', id)
  store.dispatch('Debug/changeServer', id)
}

</script>

<style scoped lang="less">
.env-selection-main {
  .select-env {
    min-width: 100px;
    text-align: left;
    border-right: 1px solid #d9d9d9;
  }
}

</style>
