<template>
  <div class="scenario-node-edit-main">
    <div>id: {{selectedNode.id}}</div>
    <div>name: {{selectedNode.name}}</div>

    <div>entityCategory: {{selectedNode.entityCategory}}</div>
    <div>entityType: {{selectedNode.entityType}}</div>
    <div>entityId: {{selectedNode.entityId}}</div>
    <div>interfaceId: {{selectedNode.interfaceId}}</div>
  </div>
</template>

<script setup lang="ts">
import {computed} from "vue";
import {useRouter} from "vue-router";

import {useStore} from "vuex";

import {StateType as ScenarioStateType} from "../../store";

const router = useRouter();
const store = useStore<{ Scenario: ScenarioStateType; }>();
const selectedNode = computed<boolean>(()=> store.state.Scenario.nodeData);

</script>

<style lang="less" scoped>
.scenario-node-edit-main {
  padding: 16px;
}

</style>