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
import {computed, ref, watch} from "vue";
import { useStore } from "vuex";
import {StateType as Debug} from "@/views/component/debug/store";
import debounce from "lodash.debounce";
import {StateType as TestInterfaceStateType} from "@/views/debugger/store";

const store = useStore<{TestInterface: TestInterfaceStateType, Debug: Debug, Endpoint}>();

const debugData = computed<any>(() => store.state.Debug.debugData);
const serveServers: any = computed(() => store.state.TestInterface.serveServers);

const serverId = computed(() => {
  return debugData?.value?.serverId || serveServers?.value[0]?.value || ''
});

function changeServer(id) {
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
