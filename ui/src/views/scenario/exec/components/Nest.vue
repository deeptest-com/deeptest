<template>
  <div class="scenario-exec-log-main">
    <div v-for="(item, index) in logs" :key="index">
      <div>{{item.id}} - {{ item.name }}</div>
      <LogComp v-if="item.logs" :logs="item.logs"></LogComp>
    </div>
  </div>
</template>

<script setup lang="ts" name="LogComp">
import {defineProps} from "vue";
import LogComp from "./Nest.vue"

defineProps<{
  logs: []
}>()

</script>

<style lang="less" scoped>
.scenario-exec-log-main {
  height: 100%;
  padding: 10px;
}

</style>