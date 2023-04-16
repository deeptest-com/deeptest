<template>
  <div class="endpoint-debug">
    <DebugComp></DebugComp>
  </div>
</template>

<script setup lang="ts">
import DebugComp from '@/views/component/debug/index.vue';
import {provide} from "vue";
import {UsedBy} from "@/utils/enum";

provide('usedBy', UsedBy.InterfaceDebug)
</script>

<style lang="less" scoped>
.endpoint-debug {
}
</style>
