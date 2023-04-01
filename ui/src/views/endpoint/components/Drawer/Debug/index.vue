<template>
  <div class="endpoint-debug">
    {{debugData}}
  </div>
</template>

<script setup lang="ts">
import {
  ref,
  defineProps,
  computed,
} from 'vue';
import {useStore} from "vuex";
import {StateType as Endpoint} from "../../store";
import {StateType as Debug} from "@/store/debug";
import {StateType as ProjectGlobal} from "@/store/project";

const props = defineProps({});
const store = useStore<{ Debug, ProjectGlobal, Endpoint  }>();

console.log(store.state.ProjectGlobal)

const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const endpointDetail = computed<any>(() => store.state.Endpoint.endpointDetail);
const debugData = computed<any>(() => store.state.Debug.debugData);

store.dispatch('Debug/loadDebugData', {endpointId: endpointDetail.value.id});

</script>

<style lang="less" scoped>
.endpoint-debug {

}
</style>
