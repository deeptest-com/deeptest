<template>
  <div class="interface-debug">
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
import {StateType as Interface} from "../../store";
import {StateType as Debug} from "@/store/debug";
import {StateType as ProjectGlobal} from "@/store/project";

const props = defineProps({});
const store = useStore<{ Debug, ProjectGlobal, Interface  }>();

console.log(store.state.ProjectGlobal)

const currProject = computed<any>(() => store.state.ProjectGlobal.currProject);
const interfaceDetail = computed<any>(() => store.state.Interface.interfaceDetail);
const debugData = computed<any>(() => store.state.Debug.debugData);

store.dispatch('Debug/loadDebugData', {endpointId: interfaceDetail.value.id});

</script>

<style lang="less" scoped>
.interface-debug {

}
</style>
