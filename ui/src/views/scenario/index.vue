<template>
  <ContentPane>
    <template #left>
      <Tree/>
    </template>
    <template #right>
      <ScenarioList/>
    </template>
  </ContentPane>
</template>

<script setup lang="ts">
import Tree from './components/Tree/index.vue';
import ScenarioList from './components/List/index.vue';
import ContentPane from '@/views/component/ContentPane/index.vue';
</script>
