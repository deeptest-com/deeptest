<template>
  <ContentPane>
    <template #left>
      <PlanTree/>
    </template>
    <template #right>
      <PlanList/>
    </template>
  </ContentPane>
</template>

<script setup lang="ts">
import PlanTree from './list/tree.vue';
import PlanList from './list/list.vue';
import ContentPane from '@/views/component/ContentPane/index.vue';
</script>
