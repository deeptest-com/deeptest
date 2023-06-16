<template>
  <div class="url-input-main">
    <div class="url">
      <a-input-group>
        <a-row type="flex" :gutter="0">
          <a-col flex="80px">
            <a-select class="select-env"
                      :options="methods"
                      :value="method"
                      @change="changeMethod">
            </a-select>
          </a-col>

          <a-col flex="3">
            <a-input placeholder="站点地址"
                     :value="currentEnvUrl">
            </a-input>
          </a-col>
          <a-col flex="3">
            <a-input class="uri" placeholder="请求路径"
                     :value="url">
            </a-input>
          </a-col>
        </a-row>
      </a-input-group>

    </div>
  </div>
</template>
<script setup lang="ts">
import {computed, ref} from "vue";
import {useStore} from "vuex";
import {StateType as DebugStateType} from "@/views/component/debug/store";
import {StateType as TestInterfaceStateType} from "@/views/debugger/store";
import {StateType as EndpointStateType} from "@/views/endpoint/store";
import {Methods} from "@/utils/enum";
import {getArrSelectItems} from "@/utils/comm";

const store = useStore<{ TestInterface: TestInterfaceStateType, Debug: DebugStateType, Endpoint: EndpointStateType }>();

const debugData = computed<any>(() => store.state.Debug.debugData);
const serveServers: any = computed(() => store.state.TestInterface.serveServers);

const method = ref('GET')
const methods = getArrSelectItems(Methods)

const url = computed(() => {
  return debugData?.value.url
});

const currentServerId = ref(debugData.value.serverId || null);
const currentEnvUrl = computed(() => {
  console.log('computed currentEnvUrl', currentServerId.value, serveServers.value)

  return serveServers.value?.find((item) => {
    return currentServerId.value === item.id;
  })?.url
});

const changeMethod = (item) => {
  console.log('changeMethod', item)
}

</script>

<style scoped lang="less">
.url-input-main {
  display: inline-block;
  overflow: hidden;
  width: 100%;

  .url {
    .select-env {
      width: 100%;
    }
  }
}
</style>
