<template>
  <div class="scenario-node-edit-main">
    <ProcessorSimpleDefault v-if="selectedNode.entityType === 'processor_simple_default'" />

  </div>
</template>

<script setup lang="ts">
import {computed} from "vue";
import {useRouter} from "vue-router";

import {useStore} from "vuex";

import {StateType as ScenarioStateType} from "../../store";
import ProcessorSimpleDefault from "./proccessors/simple/default.vue";

const router = useRouter();
const store = useStore<{ Scenario: ScenarioStateType; }>();
const selectedNode = computed<boolean>(()=> store.state.Scenario.nodeData);

</script>

<style lang="less" scoped>
.scenario-node-edit-main {
  padding: 16px;
}

</style>