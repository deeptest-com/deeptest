<template>
  <ContentPane 
    :showExpand="true" 
    :containerStyle="{ padding:0, margin:0, overflow: 'hidden'}">
    <template #left>
      <ScenarioTree/>
    </template>
    <template #right>
      <Designer/>
    </template>
  </ContentPane>
</template>

<script setup lang="ts">
import {computed, defineProps, onMounted, onUnmounted, ref, watch} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";

import {StateType as GlobalStateType} from "@/store/global";
import {StateType as ScenarioStateType} from "../../store";
// import ScenarioTree from "./Tree.vue"
import ScenarioTree from "./ScenarioTree.vue"
import Designer from "./Designer.vue"
import ContentPane from '@/views/component/ContentPane/index.vue';

const router = useRouter();
const store = useStore<{ Scenario: ScenarioStateType; Global: GlobalStateType; }>();
const selectedNode = computed<any>(()=> store.state.Scenario.nodeData);

onMounted(() => {
  store.dispatch('Scenario/getNode', null)
})

watch(() => {
  return selectedNode.value;
}, val => {
  console.log(val);
}, {
  immediate: true
})
</script>
