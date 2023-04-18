<template>
  <div class="workbench">
      <StatisticHeader/>
      <List style="margin-top:16px"/>
  </div>
</template>


<script setup lang="ts">
import { computed, onMounted, reactive, ref } from "vue";
import StatisticHeader from "@/components/StatisticHeader/index.vue";
import List from "./component/List/index.vue";




</script>

<style lang="less" scoped>
.workbench {
  padding: 16px;
}
</style>