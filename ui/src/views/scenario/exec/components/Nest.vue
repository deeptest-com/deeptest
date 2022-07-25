<template>
  <div class="scenario-exec-log-main">
    <div v-for="(item, index) in logs" :key="index">
      <div>{{index}} - {{ item.name }} - {{ item.logs }}</div>
      - <LogComp :logs="item.logs"></LogComp>
    </div>
  </div>
</template>

<script setup lang="ts" name="logComp">
import {defineProps} from "vue";

defineProps<{
  logs: []
}>()

</script>

<style lang="less" scoped>
.scenario-exec-log-main {
  height: 100%;
  padding: 6px;
}

</style>